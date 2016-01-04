package pezdispenser

import (
	"sync"

	"github.com/fatih/structs"
	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	//Register m1small
	"github.com/pivotal-pez/pezdispenser/skus/m1small"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

func init() {
	m1small.Init()
}

var onceLoadInventoryPoller sync.Once

//GetAvailableInventory - this should return available inventory and start a long task poller
func GetAvailableInventory(taskCollection integrations.Collection) (inventory map[string]skurepo.SkuBuilder) {
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
				lease := &Lease{
					ProcurementMeta: make(map[string]interface{}),
				}
				sku := v.New(taskmanager.NewTaskManager(taskCollection), structs.Map(lease))
				sku.PollForTasks()
			}
		}()
	}
}
