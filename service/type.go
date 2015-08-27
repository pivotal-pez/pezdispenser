package pezdispenser

import (
	"time"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"labix.org/v2/mgo"
)

type (
	//MongoCollectionGetter - Getting collections in mongo
	MongoCollectionGetter interface {
		Collection() Persistence
	}

	//MongoCollection - interface to a collection in mongo
	MongoCollection interface {
		Remove(selector interface{}) error
		Find(query interface{}) *mgo.Query
		Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	}

	//MongoCollectionWrapper - interface to wrap mongo collections with additional persistence functions
	MongoCollectionWrapper struct {
		Persistence
		col MongoCollection
	}

	//Persistence - interface to a persistence store of some kind
	Persistence interface {
		Remove(selector interface{}) error
		FindOne(query interface{}, result interface{}) (err error)
		Upsert(selector interface{}, update interface{}) (err error)
	}

	//Lease - this represents a lease object
	Lease struct {
		taskCollection  integrations.Collection
		taskManager     *taskmanager.TaskManager
		availableSkus   map[string]skus.Sku
		ID              string                 `json:"_id"`
		InventoryID     string                 `json:"inventory_id"`
		UserName        string                 `json:"username"`
		Sku             string                 `json:"sku"`
		LeaseDuration   float64                `json:"lease_duration"`
		LeaseEndDate    time.Time              `json:"lease_end_date"`
		LeaseStartDate  time.Time              `json:"lease_start_date"`
		ConsumerMeta    map[string]interface{} `json:"consumer_meta"`
		ProcurementMeta map[string]interface{} `json:"procurement_meta"`
		Task            *taskmanager.Task      `json:"task"`
	}
)
