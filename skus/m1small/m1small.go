package m1small

import (
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
func (s *SkuM1Small) Procurement() *taskmanager.Task {
	agent := taskmanager.NewAgent(s.TaskManager, SkuName)
	task := agent.GetTask()

	agent.Run(func(ag *taskmanager.Agent) (err error) {
		if phinfo, err := s.Client.ProvisionHost("PAO", "4D.lowmem.R7", 1, "pez-stage", "centos67"); err == nil {
			ag.GetTask().Update(func(t *taskmanager.Task) interface{} {
				t.Status = taskmanager.AgentTaskStatusComplete
				t.SetPublicMeta(ProvisionHostInfoMetaName, phinfo)
				return t
			})
		}
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

	return &SkuM1Small{
		Client:          s.Client,
		ProcurementMeta: procurementMeta,
		TaskManager:     tm,
	}
}
