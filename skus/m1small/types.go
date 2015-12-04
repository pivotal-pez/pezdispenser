package m1small

import (
	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"os"
	"fmt"
)

func isEnabled() (bool){
	for _, propSuffix := range []string{"ENABLE", "USER", "PASSWORD", "HOST"} {
		propName := "INKEEPER_" + propSuffix
		if _, found := os.LookupEnv(propName); !found {
			fmt.Println(propName + " Is not defined")
			return false
		}
	}
	return true
}
func init() {
	if isEnabled(){
		skurepo.Register(SkuName, new(SkuM1Small))
	}
}

type (
	//SkuM1Small - a object representing a m1small sku
	SkuM1Small struct {
		Client          innkeeperclient.InnkeeperClient
		TaskManager     taskmanager.TaskManagerInterface
		ProcurementMeta map[string]interface{}
	}
)
