package skus

import (
	"fmt"
	"log"

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
func (s *Sku2CSmall) Procurement() (task *taskmanager.Task) {
	task = s.createResponseTaskWithConsumerMetaData()
	s.createSelfDestructTask()
	return
}

func (s *Sku2CSmall) createResponseTaskWithConsumerMetaData() (task *taskmanager.Task) {
	task = s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusComplete)
	s.TaskManager.SaveTask(task)
	return
}

func (s *Sku2CSmall) createSelfDestructTask() {
	task := s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusProcessing)
	task.PrivateMetaData = s.ProcurementMeta
	task.SetPrivateMeta(taskmanager.TaskActionMetaName, TaskActionSelfDestruct)
	task.Expires = task.GetPrivateMeta(LeaseExpiresFieldName).(int64)
	s.TaskManager.SaveTask(task)
}

//ReStock - this method will walk the restock flow for the 2csmall object
func (s *Sku2CSmall) ReStock() (task *taskmanager.Task) {

	if vcdResponseTaskElement, err := s.undeployVapp(); err == nil {
		task = s.createUndeployPollingTask(vcdResponseTaskElement)

	} else {
		task = new(taskmanager.Task)
		task.Status = StatusFailed
	}
	s.TaskManager.SaveTask(task)
	return
}

func (s *Sku2CSmall) undeployVapp() (*vcloudclient.TaskElem, error) {
	user := fmt.Sprintf("%s", s.ProcurementMeta[VCDUsernameField])
	pass := fmt.Sprintf("%s", s.ProcurementMeta[VCDPasswordField])
	vAppID := fmt.Sprintf("%s", s.ProcurementMeta[VCDAppIDField])
	s.Client.Auth(user, pass)
	return s.Client.UnDeployVApp(vAppID)
}

func (s *Sku2CSmall) createUndeployPollingTask(vcdResponseTaskElement *vcloudclient.TaskElem) (task *taskmanager.Task) {
	task = s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusOutsourced)
	task.PrivateMetaData = s.ProcurementMeta
	task.SetPrivateMeta(VCDTaskElementHrefMetaName, vcdResponseTaskElement.Href)
	task.SetPrivateMeta(taskmanager.TaskActionMetaName, TaskActionUnDeploy)
	return
}

//PollForTasks - this is a method for polling the current long poll task queue and acting on it
func (s *Sku2CSmall) PollForTasks() {
	var (
		err  error
		task *taskmanager.Task
	)
	if task, err = s.TaskManager.FindAndStallTaskForCaller(SkuName2CSmall); err == nil {
		log.Println("we found a task: ", task)
		s.handleTaskTypes(task)

	} else {
		log.Println("Error (2c.small poller): ", err.Error())
	}
}

func (s *Sku2CSmall) handleTaskTypes(task *taskmanager.Task) {
	saveTask := true
	switch task.GetPrivateMeta(taskmanager.TaskActionMetaName) {
	case TaskActionUnDeploy:
		s.processVCDTask(task, s.deployNew2CSmall)

	case TaskActionDeploy:
		s.processVCDTask(task, s.deployComplete)

	case TaskActionSelfDestruct:
		s.processSelfDestructTask(task)

	default:
		log.Println("not a valid task action")
		saveTask = false
	}

	if saveTask {
		s.TaskManager.SaveTask(task)
	}
}

func (s *Sku2CSmall) processSelfDestructTask(task *taskmanager.Task) {
	s.expireLongRunningTask(task)
	s.ProcurementMeta = task.PrivateMetaData
	s.ReStock()
	s.ProcurementMeta = nil
}

func (s *Sku2CSmall) deployComplete(task *taskmanager.Task) {
	inventoryID := fmt.Sprintf("%s", task.GetPrivateMeta(InventoryIDFieldName))
	s.setInventoryStatusToAvailable(inventoryID)
}

func (s *Sku2CSmall) setInventoryStatusToAvailable(inventoryID string) {
	if inventoryTask, err := s.TaskManager.FindTask(inventoryID); err == nil {
		inventoryTask.Status = taskmanager.TaskStatusAvailable
		s.TaskManager.SaveTask(inventoryTask)
	}
}

func (s *Sku2CSmall) processVCDTask(task *taskmanager.Task, successCallback func(*taskmanager.Task)) {
	var (
		err            error
		vcdTaskElement *vcloudclient.TaskElem
	)
	if vcdTaskURI := fmt.Sprintf("%s", task.GetPrivateMeta(VCDTaskElementHrefMetaName)); vcdTaskURI != "" {

		if s.Client == nil {
			httpClient := vcloudclient.DefaultClient()
			s.Client = vcloudclient.NewVCDClient(httpClient, fmt.Sprintf("%s", task.GetPrivateMeta(VCDBaseURIField)))
		}

		if vcdTaskElement, err = s.Client.PollTaskURL(vcdTaskURI); err == nil {
			s.evaluateVCDTaskStatus(vcdTaskElement.Status, task, successCallback)

		} else {
			log.Println("Error (poll taskUrl VCD): ", err.Error())
		}
	}
}

func (s *Sku2CSmall) evaluateVCDTaskStatus(status string, task *taskmanager.Task, successCallback func(*taskmanager.Task)) {
	task.Status = status

	switch status {
	case vcloudclient.TaskStatusSuccess:
		s.expireLongRunningTask(task)
		successCallback(task)

	case vcloudclient.TaskStatusError, vcloudclient.TaskStatusAborted, vcloudclient.TaskStatusCanceled:
		s.expireLongRunningTask(task)
	}
}

func (s *Sku2CSmall) deployNew2CSmall(task *taskmanager.Task) {
	var (
		newTask *taskmanager.Task
	)

	if vapp, err := s.deployVappFromTemplate(task); err == nil {
		newTask = s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusOutsourced)
		newTask.SetPrivateMeta(VCDTaskElementHrefMetaName, vapp.Tasks.Task.Href)
		newTask.SetPrivateMeta(taskmanager.TaskActionMetaName, TaskActionDeploy)

	} else {
		newTask = s.TaskManager.NewTask(SkuName2CSmall, taskmanager.TaskLongPollQueue, StatusFailed)
	}
	s.TaskManager.SaveTask(newTask)
	task.SetPublicMeta(taskmanager.TaskChildID, newTask.ID)
	s.TaskManager.SaveTask(task)
}

func (s *Sku2CSmall) deployVappFromTemplate(task *taskmanager.Task) (vapp *vcloudclient.VApp, err error) {
	var (
		username     = fmt.Sprintf("%s", task.GetPrivateMeta(VCDUsernameField))
		password     = fmt.Sprintf("%s", task.GetPrivateMeta(VCDPasswordField))
		templatename = fmt.Sprintf("%s", task.GetPrivateMeta(VCDTemplateNameField))
		vappTemplate *vcloudclient.VAppTemplateRecord
	)
	s.Client.Auth(username, password)

	if vappTemplate, err = s.Client.QueryTemplate(templatename); err == nil {
		vapp, err = s.Client.DeployVApp(templatename, vappTemplate.Href, vappTemplate.Vdc)
	}
	return
}

func (s *Sku2CSmall) expireLongRunningTask(task *taskmanager.Task) {
	task.Expires = taskmanager.ExpiredTask
	s.TaskManager.SaveTask(task)
}
