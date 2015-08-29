package skus

import (
	"fmt"

	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

//New - create a new instance of the given object type, initialized with some vars
func (s *Sku2CSmall) New(tm TaskManager, procurementMeta map[string]interface{}) Sku {
	httpClient := vcloudclient.DefaultClient()
	baseURI := fmt.Sprintf("%s", procurementMeta[VCDBaseURIField])

	return &Sku2CSmall{
		Client:          vcloudclient.NewVCDClient(httpClient, baseURI),
		ProcurementMeta: procurementMeta,
		TaskManager:     tm,
	}
}

//Procurement - this method will walk the procurement flow for the 2csmall
//object
func (s *Sku2CSmall) Procurement() (status string, taskMeta map[string]interface{}) {
	status = StatusComplete
	return
}

//ReStock - this method will walk the restock flow for the 2csmall object
func (s *Sku2CSmall) ReStock() (status string, taskMeta map[string]interface{}) {
	taskMeta = make(map[string]interface{})
	user := fmt.Sprintf("%s", s.ProcurementMeta[VCDUsernameField])
	pass := fmt.Sprintf("%s", s.ProcurementMeta[VCDPasswordField])
	vAppID := fmt.Sprintf("%s", s.ProcurementMeta[VCDAppIDField])
	s.Client.Auth(user, pass)

	if vcdResponseTaskElement, err := s.Client.UnDeployVApp(vAppID); err == nil {
		status = StatusOutsourced
		task := s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusProcessing)
		task.MetaData = s.ProcurementMeta
		task.MetaData[VCDTaskElementHrefMetaName] = vcdResponseTaskElement.Href
		task.MetaData[taskmanager.TaskActionMetaName] = TaskActionUnDeploy
		taskMeta[VCDTaskElementHrefMetaName] = vcdResponseTaskElement.Href
		taskMeta[taskmanager.TaskActionMetaName] = TaskActionUnDeploy
		taskMeta[SubTaskIDField] = task.ID.Hex()
		s.TaskManager.SaveTask(task)

	} else {
		status = StatusFailed
	}
	return
}

//PollForTasks - this is a method for polling the current long poll task queue and acting on it
func (s *Sku2CSmall) PollForTasks() {
	var (
		err            error
		task           *taskmanager.Task
		vcdTaskElement *vcloudclient.TaskElem
	)
	task, err = s.TaskManager.FindAndStallTaskForCaller(SkuName2CSmall)

	if vcdTaskURI := fmt.Sprintf("%s", task.MetaData[VCDTaskElementHrefMetaName]); vcdTaskURI != "" {

		if s.Client == nil {
			httpClient := vcloudclient.DefaultClient()
			s.Client = vcloudclient.NewVCDClient(httpClient, fmt.Sprintf("%s", task.MetaData[VCDBaseURIField]))
		}

		if vcdTaskElement, err = s.Client.PollTaskURL(vcdTaskURI); err == nil {
			s.evaluateStatus(vcdTaskElement.Status, task)
		}
	}
	s.TaskManager.SaveTask(task)
}

func (s *Sku2CSmall) evaluateStatus(status string, task *taskmanager.Task) {
	task.Status = status

	switch status {
	case vcloudclient.TaskStatusSuccess:
		s.expireLongRunningTask(task)

		if task.MetaData[taskmanager.TaskActionMetaName] == TaskActionUnDeploy {
			s.deployNew2CSmall(task)
		}

	case vcloudclient.TaskStatusError, vcloudclient.TaskStatusAborted, vcloudclient.TaskStatusCanceled:
		s.expireLongRunningTask(task)
	}
}

func (s *Sku2CSmall) deployNew2CSmall(task *taskmanager.Task) {
	var (
		err          error
		username     = fmt.Sprintf("%s", task.MetaData[VCDUsernameField])
		password     = fmt.Sprintf("%s", task.MetaData[VCDPasswordField])
		templatename = fmt.Sprintf("%s", task.MetaData[VCDTemplateNameField])
		vapp         *vcloudclient.VApp
		vappTemplate *vcloudclient.VAppTemplateRecord
		newTask      *taskmanager.Task
	)
	s.Client.Auth(username, password)

	if vappTemplate, err = s.Client.QueryTemplate(templatename); err == nil {

		if vapp, err = s.Client.DeployVApp(templatename, vappTemplate.Href, vappTemplate.Vdc); err != nil {
			newTask = s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusFailed)

		} else {
			newTask = s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusOutsourced)
			newTask.MetaData[VCDTaskElementHrefMetaName] = vapp.Tasks.Task.Href
		}
		newTask, err = s.TaskManager.SaveTask(newTask)
		fmt.Println("this is what was returned: ", newTask)
	}
}

func (s *Sku2CSmall) expireLongRunningTask(task *taskmanager.Task) {
	task.Expires = taskmanager.ExpiredTask
	s.TaskManager.SaveTask(task)
}
