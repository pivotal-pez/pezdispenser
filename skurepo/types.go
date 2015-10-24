package skurepo

import "github.com/pivotal-pez/pezdispenser/taskmanager"

type (
	//Sku - interface for a sku object
	Sku interface {
		Procurement() *taskmanager.Task
		ReStock() *taskmanager.Task
		PollForTasks()
		New(tm TaskManager, procurementMeta map[string]interface{}) Sku
	}

	//TaskManager - an interface representing a taskmanager object
	TaskManager interface {
		SaveTask(t *taskmanager.Task) (*taskmanager.Task, error)
		FindAndStallTaskForCaller(callerName string) (t *taskmanager.Task, err error)
		FindTask(id string) (t *taskmanager.Task, err error)
		NewTask(n string, p taskmanager.ProfileType, s string) (t *taskmanager.Task)
	}
)
