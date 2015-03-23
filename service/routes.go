package pezdispenser

import (
	"fmt"

	"github.com/go-martini/martini"
)

//Constants to construct routes with
const (
	APIVersion1 = "v1"
	indexRoute  = "/"
	leasePath   = "lease"
	lockPath    = "lock"
	typePath    = "type"
	itemPath    = "item"
	ItemGUID    = "inventoryItemGUID"
	TypeGUID    = "inventoryTypeGUID"
)

//formatted strings based on constants, to be used in URLs
var (
	URLLeaseBaseV1 = fmt.Sprintf("/%s/%s", APIVersion1, leasePath)
	URLLockBaseV1  = fmt.Sprintf("/%s/%s", APIVersion1, lockPath)
	URLTypeGUID    = fmt.Sprintf("/%s/:%s", typePath, TypeGUID)
	URLItemGUID    = fmt.Sprintf("/%s/:%s", itemPath, ItemGUID)
	URLLeases      = "/list"
)

//InitRoutes - initialize the mappings for controllers against valid routes
func InitRoutes(m *martini.ClassicMartini) {
	itemLeaseController := NewLeaseController(APIVersion1, Item)
	typeLeaseController := NewLeaseController(APIVersion1, Type)
	listLeaseController := NewLeaseController(APIVersion1, List)
	lockController := NewLockController(APIVersion1)

	m.Group("/", func(r martini.Router) {
		r.Get("info", func() string {
			return "the dispenser service will give you candy"
		})
	})

	m.Group(URLLeaseBaseV1, func(r martini.Router) {
		r.Post(URLTypeGUID, typeLeaseController.Post())
		r.Get(URLTypeGUID, typeLeaseController.Get())

		r.Post(URLItemGUID, itemLeaseController.Post())
		r.Get(URLItemGUID, itemLeaseController.Get())
		r.Delete(URLItemGUID, itemLeaseController.Delete())

		r.Get(URLLeases, listLeaseController.Get())
	})

	m.Group(URLLockBaseV1, func(r martini.Router) {
		r.Post(URLItemGUID, lockController.Post())
		r.Get(URLItemGUID, lockController.Get())
	})
}
