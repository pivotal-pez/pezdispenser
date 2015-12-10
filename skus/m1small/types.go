package m1small

import (
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"github.com/xchapter7x/lo"
)

func isEnabled() bool {

	if appEnv, err := cfenv.Current(); err == nil {
		if taskService, err := appEnv.Services.WithName("innkeeper-service"); err == nil {
			if taskService.Credentials["enable"].(string) == "1" {
				return true
			}
		}
	}
	lo.G.Error("m1small not enabled")
	return false
}
func init() {
	if isEnabled() {
		s := new(SkuM1Small)
		s.Client = innkeeperclient.New()
		skurepo.Register(SkuName, s)
	}
}

type (
	//SkuM1Small - a object representing a m1small sku implements skurepo.Sku
	SkuM1Small struct {
		Client          innkeeperclient.InnkeeperClient
		TaskManager     taskmanager.TaskManagerInterface
		ProcurementMeta map[string]interface{}
	}
)