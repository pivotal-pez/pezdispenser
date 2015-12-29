package m1small

import (
	"fmt"
	"time"

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
	if task, err := s.TaskManager.FindAndStallTaskForCaller(SkuName); err == nil {
		fmt.Println(task)
	}
	return
}

//StartPoller --
func (s *SkuM1Small) StartPoller(requestID string, task *taskmanager.Task) (err error) {
	var (
		clnt innkeeperclient.InnkeeperClient
		resp innkeeperclient.GetStatusResponse
	)

	if clnt, err = s.GetInnkeeperClient(); err == nil {

		if resp, err = s.waitForStatusComplete(requestID, clnt); err == nil {
			s.updateTaskForStatusComplete(task, resp)
		}
	}
	return
}

func (s *SkuM1Small) updateTaskForStatusComplete(task *taskmanager.Task, resp innkeeperclient.GetStatusResponse) {
	task.Update(func(t *taskmanager.Task) interface{} {
		t.Status = taskmanager.AgentTaskStatusComplete
		t.Expires = taskmanager.ExpiredTask
		t.SetPublicMeta(GetStatusInfoMetaName, resp)
		return t
	})
}

func (s *SkuM1Small) waitForStatusComplete(requestID string, clnt innkeeperclient.InnkeeperClient) (resp innkeeperclient.GetStatusResponse, err error) {
	respLocal := &innkeeperclient.GetStatusResponse{}

	for {

		if respLocal, err = clnt.GetStatus(requestID); err != nil {
			lo.G.Error("get status yielded error: ", err)

		} else {
			resp = *respLocal
		}

		if resp.Data.Status != taskmanager.AgentTaskStatusComplete {
			time.Sleep(taskmanager.AgentTaskPollerInterval)

		} else {
			break
		}
	}
	return
}

// Procurement -- use agent to run async task
func (s *SkuM1Small) Procurement() *taskmanager.Task {
	agent := taskmanager.NewAgent(s.TaskManager, SkuName)
	task := agent.GetTask()

	agent.Run(func(ag *taskmanager.Agent) (err error) {
		if clnt, err := s.GetInnkeeperClient(); err == nil {
			if phinfo, err := clnt.ProvisionHost(ClientSkuName, ClientLeaseOwner); err == nil {
				lo.G.Debug("provisionhost response: ", phinfo)

				ag.GetTask().Update(func(t *taskmanager.Task) interface{} {
					t.Status = taskmanager.AgentTaskStatusComplete
					t.SetPublicMeta(ProvisionHostInfoMetaName, phinfo)
					return t
				})
				go s.StartPoller(phinfo.Data[0].RequestID, task)
			} else {
				return err
			}
		}
		return err
	})

	return task
}

// ReStock -- this will grab a requestid from procurementMeta and call the innkeeper client to deprovision.
func (s *SkuM1Small) ReStock() (tm *taskmanager.Task) {
	lo.G.Debug("ProcurementMeta: ", s.ProcurementMeta)
	requestID := s.ProcurementMeta[ProcurementMetaFieldRequestID].(string)
	lo.G.Debug("requestID: ", requestID)

	if clnt, err := s.GetInnkeeperClient(); err == nil {

		var err error
		var res *innkeeperclient.GetStatusResponse
		if res, err = clnt.DeProvisionHost(requestID); err != nil && res.Status != "success" {
			lo.G.Error("de-provision requestid call error: ", requestID, err, res)
		}
		lo.G.Debug("deprovision call results: ", requestID, err, res)
	}
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
