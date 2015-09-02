package pezdispenser

import (
	"sync"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var onceLoadInventoryPoller sync.Once

//GetAvailableInventory - this should return available inventory and start a long task poller
func GetAvailableInventory(taskCollection integrations.Collection) (inventory map[string]skus.Sku) {

	inventory = map[string]skus.Sku{
		skus.SkuName2CSmall: &skus.Sku2CSmall{
			TaskManager: taskmanager.NewTaskManager(taskCollection),
		},
	}

	onceLoadInventoryPoller.Do(func() {
		for _, v := range inventory {
			go func() {
				for {
					v.PollForTasks()
				}
			}()
		}
	})
	return
}
