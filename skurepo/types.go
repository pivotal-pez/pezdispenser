package skurepo

import (
	"time"

	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

type (
	//SkuBuilder - a object that can build skus
	SkuBuilder interface {
		New(tm TaskManager, meta map[string]interface{}) Sku
	}

	//Sku - interface for a sku object
	Sku interface {
		Procurement() *taskmanager.Task
		ReStock() *taskmanager.Task
		PollForTasks()
	}

	//TaskManager - an interface representing a taskmanager object
	TaskManager interface {
		SaveTask(t *taskmanager.Task) (*taskmanager.Task, error)
		FindAndStallTaskForCaller(callerName string) (t *taskmanager.Task, err error)
		FindTask(id string) (t *taskmanager.Task, err error)
		NewTask(n string, p taskmanager.ProfileType, s string) (t *taskmanager.Task)
		ScheduleTask(t *taskmanager.Task, expireTime time.Time)
	}
)
