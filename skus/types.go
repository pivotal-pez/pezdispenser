package skus

import (
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

type (
	//Sku - interface for a sku object
	Sku interface {
		Procurement() (status string, taskMeta map[string]interface{})
		ReStock() (status string, taskMeta map[string]interface{})
		PollForTasks()
		New(tm TaskManager, procurementMeta map[string]interface{}) Sku
	}
	//Sku2CSmall - a object representing a 2csmall sku
	Sku2CSmall struct {
		Client          vcdClient
		TaskManager     TaskManager
		name            string
		ProcurementMeta map[string]interface{}
	}

	vcdClient interface {
		UnDeployVApp(vappID string) (task *vcloudclient.TaskElem, err error)
		DeployVApp(templateName, templateHref, vcdHref string) (vapp *vcloudclient.VApp, err error)
		Auth(username, password string) (err error)
		QueryTemplate(templateName string) (vappTemplate *vcloudclient.VAppTemplateRecord, err error)
	}

	//TaskManager - an interface representing a taskmanager object
	TaskManager interface {
		SaveTask(t *taskmanager.Task) (*taskmanager.Task, error)
		FindLockFirstCallerName(callerName string) (t *taskmanager.Task, err error)
		UnLockTask(id string) (t *taskmanager.Task, err error)
		FindTask(id string) (t *taskmanager.Task, err error)
		NewTask(n string, p taskmanager.ProfileType, s string) (t *taskmanager.Task)
	}
)
