package taskmanager

import (
	"sync"
	"time"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"labix.org/v2/mgo/bson"
)

type (
	//RedactedTask - a task object without sensitive information
	RedactedTask struct {
		ID         bson.ObjectId          `bson:"_id"`
		Timestamp  int64                  `bson:"timestamp"`
		Expires    int64                  `bson:"expires"`
		Status     string                 `bson:"status"`
		Profile    ProfileType            `bson:"profile"`
		CallerName string                 `bson:"caller_name"`
		MetaData   map[string]interface{} `bson:"metadata"`
	}

	//Task - a task object
	Task struct {
		ID              bson.ObjectId          `bson:"_id"`
		Timestamp       int64                  `bson:"timestamp"`
		Expires         int64                  `bson:"expires"`
		Status          string                 `bson:"status"`
		Profile         ProfileType            `bson:"profile"`
		CallerName      string                 `bson:"caller_name"`
		MetaData        map[string]interface{} `bson:"metadata"`
		PrivateMetaData map[string]interface{} `bson:"private_metadata"`
		mutex           sync.RWMutex
		taskManager     TaskManagerInterface
	}

	//TaskManagerInterface ---
	TaskManagerInterface interface {
		NewTask(callerName string, profile ProfileType, status string) (t *Task)
		FindTask(id string) (t *Task, err error)
		FindAndStallTaskForCaller(callerName string) (task *Task, err error)
		SaveTask(t *Task) (*Task, error)
		ScheduleTask(t *Task, expireTime time.Time)
	}

	//TaskManager - manages task interactions crud stuff
	TaskManager struct {
		taskCollection integrations.Collection
	}

	//Agent an object which knows how to handle long running tasks. polling,
	//timeouts etc
	Agent struct {
		killTaskPoller  chan bool
		processComplete chan bool
		taskPollEmitter chan bool
		statusEmitter   chan string
		task            *Task
		taskManager     TaskManagerInterface
	}

	//ProfileType - indicator of the purpose of the task to be performed
	ProfileType string
)
