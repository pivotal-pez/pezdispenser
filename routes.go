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
)

//InitRoutes - initialize the mappings for controllers against valid routes
func InitRoutes(m *martini.ClassicMartini) {

	m.Group(URLLeaseBaseV1, func(r martini.Router) {
		itemLeaseController := NewLeaseController(APIVersion1, Item)
		typeLeaseController := NewLeaseController(APIVersion1, Type)
		r.Post(URLTypeGUID, typeLeaseController.Post())
		r.Post(URLItemGUID, itemLeaseController.Post())
		r.Delete(URLItemGUID, itemLeaseController.Delete())
	})

	m.Group(URLLockBaseV1, func(r martini.Router) {
		lockController := NewLockController(APIVersion1)
		r.Post(URLItemGUID, lockController.Post())
		r.Get(URLItemGUID, lockController.Get())
	})
}
