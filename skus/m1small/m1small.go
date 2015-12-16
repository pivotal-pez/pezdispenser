package m1small

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"github.com/xchapter7x/lo"
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
		if clnt, err := s.GetInnkeeperClient(); err == nil {
			if phinfo, err := clnt.ProvisionHost("PAO", "4D.lowmem.R7", 1, "pez-stage", "centos67"); err == nil {
				ag.GetTask().Update(func(t *taskmanager.Task) interface{} {
					t.Status = taskmanager.AgentTaskStatusComplete
					t.SetPublicMeta(ProvisionHostInfoMetaName, phinfo)
					return t
				})
			} else {
				return err
			}
		}
		return err
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

// InitInnkeeperClient -- initialize innkeeper client based on cf configuration
func (s *SkuM1Small) InitInnkeeperClient() (clnt innkeeperclient.InnkeeperClient, err error) {
	if appEnv, err := cfenv.Current(); err == nil {

		if taskService, err := appEnv.Services.WithName("innkeeper-service"); err == nil {
			clnt = &innkeeperclient.IkClient{
				URI:      taskService.Credentials["uri"].(string),
				User:     taskService.Credentials["user"].(string),
				Password: taskService.Credentials["password"].(string),
			}
		}
	}
	return
}

// GetInnkeeperClient -- get an innkeeper client and cache it in the object
func (s *SkuM1Small) GetInnkeeperClient() (innkeeperclient.InnkeeperClient, error) {
	var err error
	if s.Client == nil {
		if clnt, err := s.InitInnkeeperClient(); err == nil {
			s.Client = clnt
		} else {
			lo.G.Error("error parsing current cfenv: ", err.Error())
		}
	}
	return s.Client, err
}
