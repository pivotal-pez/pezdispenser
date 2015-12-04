package m1small

import (
	"bytes"
	"log"

	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

/*
 @implements skurepo.Sku
*/

// PollForTasks -- no longer needed, agent.Run in already asynchronous
func (s *SkuM1Small) PollForTasks() {
	return
}

// Procurement -- use agent to run async task
func (s *SkuM1Small) Procurement() (*taskmanager.Task) {
	agent := taskmanager.NewAgent(s.TaskManager, SkuName)
	task := agent.GetTask()
	
	agent.Run( func (ag *taskmanager.Agent) (err error){
			phinfo, err := s.Client.ProvisionHost("PAO", "4D.lowmem.R7", 1, "pez-stage", "centos67")
			tsk := ag.GetTask()
			tsk.Status = "ok"
			tsk.SetPublicMeta("phinfo", phinfo)
			return
		})
	return task
}

// ReStock -- WARNING not implemented
func (s *SkuM1Small) ReStock() (tm *taskmanager.Task) {
	return
}

// New -- return a new SKU provider
func (s *SkuM1Small) New(tm skurepo.TaskManager, procurementMeta map[string]interface{}) skurepo.Sku {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	return &SkuM1Small{
		Client:          innkeeperclient.New(logger),
		ProcurementMeta: procurementMeta,
		TaskManager:     tm,
	}
}
