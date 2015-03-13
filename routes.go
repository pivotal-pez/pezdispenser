package pezdispenser

import (
	"fmt"

	"github.com/go-martini/martini"
)

//Constants to construct routes with
const (
	ApiVersion1 = "v1"
	indexRoute  = "/"
	leasePath   = "lease"
	lockPath    = "lock"
	typePath    = "type"
	itemPath    = "item"
	InvGuid     = "inventoryTypeGuid"
	TypeGuid    = "inventoryItemGuid"
)

//formatted strings based on constants, to be used in URLs
var (
	UrlLeaseBaseV1 = fmt.Sprintf("/%s/%s", ApiVersion1, leasePath)
	UrlLockBaseV1  = fmt.Sprintf("/%s/%s", ApiVersion1, lockPath)
	UrlTypeGuid    = fmt.Sprintf("/%s/:%s", typePath, TypeGuid)
	UrlItemGuid    = fmt.Sprintf("/%s/:%s", itemPath, InvGuid)
)

//InitRoutes - initialize the mappings for controllers against valid routes
func InitRoutes(m *martini.ClassicMartini) {

	m.Group(UrlLeaseBaseV1, func(r martini.Router) {
		itemLeaseController := NewLeaseController(ApiVersion1, Item)
		typeLeaseController := NewLeaseController(ApiVersion1, Type)
		r.Post(UrlTypeGuid, typeLeaseController.Post())
		r.Post(UrlItemGuid, itemLeaseController.Post())
		r.Delete(UrlItemGuid, itemLeaseController.Delete())
	})

	m.Group(UrlLockBaseV1, func(r martini.Router) {
		lockController := NewLockController(ApiVersion1)
		r.Post(UrlItemGuid, lockController.Post())
		r.Get(UrlItemGuid, lockController.Get())
	})
}
