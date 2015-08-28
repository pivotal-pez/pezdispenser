package skus

import (
	"fmt"

	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

//New - create a new instance of the given object type, initialized with some vars
func (s *Sku2CSmall) New(tm TaskManager, procurementMeta map[string]interface{}) Sku {
	httpClient := vcloudclient.DefaultClient()
	baseURI := fmt.Sprintf("%s", procurementMeta["base_uri"])

	return &Sku2CSmall{
		Client:          vcloudclient.NewVCDClient(httpClient, baseURI),
		ProcurementMeta: procurementMeta,
		TaskManager:     tm,
		Name:            SkuName2CSmall,
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
	user := fmt.Sprintf("%s", s.ProcurementMeta["vcd_username"])
	pass := fmt.Sprintf("%s", s.ProcurementMeta["vcd_password"])
	vAppID := fmt.Sprintf("%s", s.ProcurementMeta["vapp_id"])
	s.Client.Auth(user, pass)

	if vcdResponseTaskElement, err := s.Client.UnDeployVApp(vAppID); err == nil {
		status = StatusOutsourced
		task := s.TaskManager.NewTask(s.Name, taskmanager.TaskLongPollQueue, StatusProcessing)
		task.MetaData = s.ProcurementMeta
		task.MetaData[VCDTaskElementHrefMetaName] = vcdResponseTaskElement.Href
		task.MetaData[taskmanager.TaskActionMetaName] = TaskActionUnDeploy
		taskMeta[VCDTaskElementHrefMetaName] = vcdResponseTaskElement.Href
		taskMeta[taskmanager.TaskActionMetaName] = TaskActionUnDeploy
		taskMeta["subtask_id"] = task.ID.Hex()
		s.TaskManager.SaveTask(task)

	} else {
		status = StatusFailed
	}
	return
}

//PollForTasks - this is a method for polling the current long poll task queue and acting on it
func (s *Sku2CSmall) PollForTasks() {
	task, err := s.TaskManager.FindAndStallTaskForCaller(s.Name)
	fmt.Println("task:", task)
	fmt.Println("error:", err)
}
