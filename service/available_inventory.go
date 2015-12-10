package pezdispenser

import (
	"sync"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	//Register m1small
	_ "github.com/pivotal-pez/pezdispenser/skus/m1small"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var onceLoadInventoryPoller sync.Once

//GetAvailableInventory - this should return available inventory and start a long task poller
func GetAvailableInventory(taskCollection integrations.Collection) (inventory map[string]skurepo.Sku) {
	inventory = skurepo.GetRegistry()

	onceLoadInventoryPoller.Do(func() {
		startTaskPollingForRegisteredSkus(taskCollection)
	})
	return
}

func startTaskPollingForRegisteredSkus(taskCollection integrations.Collection) {
	for _, v := range skurepo.GetRegistry() {
		go func() {
			for {
				sku := v.New(taskmanager.NewTaskManager(taskCollection), nil)
				sku.PollForTasks()
			}
		}()
	}
}
