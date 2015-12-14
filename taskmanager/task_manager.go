package taskmanager

import (
	"time"
	"sync"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/xchapter7x/lo"
)

//NewTaskManager - this creates a new task manager object and returns it
func NewTaskManager(taskCollection integrations.Collection) (tm *TaskManager) {
	tm = &TaskManager{
		taskCollection: taskCollection,
	}
	return
}

//SaveTask - saves the given task
func (s *TaskManager) SaveTask(t *Task) (*Task, error) {

	if t.ID.Hex() == "" {
		t.ID = bson.NewObjectId()
	}
	_, err := s.taskCollection.UpsertID(t.ID, t)
	return t, err
}

//FindAndStallTaskForCaller - find and lock the first matching task, then return
//it
func (s *TaskManager) FindAndStallTaskForCaller(callerName string) (task *Task, err error) {
	nowEpoch := time.Now().UnixNano()
	task = new(Task)
	s.taskCollection.FindAndModify(
		bson.M{
			"caller_name": callerName,
			"profile":     TaskLongPollQueue,
			"expires": bson.M{
				"$lte": nowEpoch,
				"$ne":  ExpiredTask,
			},
		},
		bson.M{
			"expires": time.Now().Add(5 * time.Minute).UnixNano(),
		},
		task,
	)

	if s.isDefaultTask(*task) {
		task = nil
		err = ErrNoResults
	}
	return
}

func (s *TaskManager) isDefaultTask(task Task) bool {
	defaultTask := Task{}
	return (task.ID == defaultTask.ID &&
		task.Timestamp == defaultTask.Timestamp &&
		task.Expires == defaultTask.Expires &&
		task.Status == defaultTask.Status &&
		task.Profile == defaultTask.Profile &&
		task.CallerName == defaultTask.CallerName)
}

//FindTask - this will find and return a task with a given ID
func (s *TaskManager) FindTask(id string) (t *Task, err error) {
	t = new(Task)
	err = s.taskCollection.FindOne(id, t)
	return
}

//SubscribeToSchedule - subscribe to a schedule and get a channel to listen on
//for a task when it hits its scheduled time
func (s *TaskManager) SubscribeToSchedule(callerName string) (subscription chan *Task) {
	subscription = make(chan *Task, 1)
	go func() {
		for {
			task := new(Task)
			s.taskCollection.FindAndModify(
				bson.M{
					"caller_name": callerName,
					"profile":     TaskAgentScheduledTask,
					"expires": bson.M{
						"$lte": time.Now().UnixNano(),
						"$ne":  ExpiredTask,
					},
				},
				bson.M{
					"expires": time.Now().Add(AgentTaskPollerTimeout).UnixNano(),
					"profile": TaskAgentLongRunning,
					"status":  AgentTaskStatusInitializing,
				},
				task,
			)

			if !s.isDefaultTask(*task) {
				go func() { subscription <- task }()
			}
			time.Sleep(AgentTaskPollerInterval)
		}
	}()
	return
}

func (s *TaskManager) ScheduleTask(t *Task, expireTime time.Time) {
	t.Expires = expireTime.UnixNano()
	t.Profile = TaskAgentScheduledTask
	t.Status = AgentTaskStatusScheduled
	return
}

func (s *TaskManager) GarbageCollectExpiredAgents(callerName string) (changeInfo *mgo.ChangeInfo, err error) {
	nowEpoch := time.Now().UnixNano()
	changeInfo, err = s.taskCollection.FindAndModify(
		bson.M{
			"caller_name": callerName,
			"profile":     TaskAgentLongRunning,
			"expires": bson.M{
				"$lte": nowEpoch,
				"$ne":  ExpiredTask,
			},
		},
		bson.M{
			"expires": ExpiredTask,
			"status":  AgentTaskStatusFailed,
		},
		nil,
	)
	lo.G.Debug("change info on agent gc: ", changeInfo)
	return
}

//NewTask - get us a new empty task
func (s *TaskManager) NewTask(callerName string, profile ProfileType, status string) (t *Task) {
	t = new(Task)
	t.CallerName = callerName
	t.Profile = profile
	t.Status = status
	t.ID = bson.NewObjectId()
	t.Timestamp = time.Now().UnixNano()
	t.MetaData = make(map[string]interface{})
	t.PrivateMetaData = make(map[string]interface{})
	t.mutex = sync.RWMutex{}
	t.taskManager = s
	return
}
