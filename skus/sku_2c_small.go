package skus

import (
	"fmt"

	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

func New2CSmallSku(client vcdClient, tm taskManager) *Sku2CSmall {
	return &Sku2CSmall{
		Client:      client,
		TaskManager: tm,
		name:        "2c.small",
	}
}

//Procurement - this method will walk the procurement flow for the 2csmall
//object
func (s *Sku2CSmall) Procurement(meta map[string]interface{}) (status string, taskMeta map[string]interface{}) {
	status = StatusComplete
	return
}

//ReStock - this method will walk the restock flow for the 2csmall object
func (s *Sku2CSmall) ReStock(meta map[string]interface{}) (status string, taskMeta map[string]interface{}) {
	taskMeta = make(map[string]interface{})
	user := fmt.Sprintf("%s", meta["vcd_username"])
	pass := fmt.Sprintf("%s", meta["vcd_password"])
	vAppID := fmt.Sprintf("%s", meta["vapp_id"])
	s.Client.Auth(user, pass)

	if vcdResponseTaskElement, err := s.Client.UnDeployVApp(vAppID); err == nil {
		status = StatusProcessing
		task := s.TaskManager.NewTask(s.name, taskmanager.TaskLongPollQueue, status)
		task.MetaData = meta
		task.MetaData[VCDTaskElementHrefMetaName] = vcdResponseTaskElement.Href
		task.MetaData[taskmanager.TaskActionMetaName] = TaskActionUnDeploy
		s.TaskManager.SaveTask(task)

	} else {
		status = StatusFailed
	}
	return
}
