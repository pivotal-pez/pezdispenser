package pezdispenser

import (
	"github.com/pivotal-pez/pezdispenser/service/_integrations"
	"labix.org/v2/mgo/bson"
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

//FindLockFirstCallerName - find and lock the first matching task, then return
//it
func (s *TaskManager) FindLockFirstCallerName(callerName string) (t *Task, err error) {
	return
}

//UnLockTask - this will unlock a task with given id
func (s *TaskManager) UnLockTask(id string) (t *Task, err error) {
	return
}

//FindTask - this will find and return a task with a given ID
func (s *TaskManager) FindTask(id string) (t *Task, err error) {
	return
}
