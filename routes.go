package pezdispenser

import (
	"fmt"

	"github.com/go-martini/martini"
)

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

var (
	UrlLeaseBaseV1 = fmt.Sprintf("/%s/%s", ApiVersion1, leasePath)
	UrlLockBaseV1  = fmt.Sprintf("/%s/%s", ApiVersion1, lockPath)
	UrlTypeGuid    = fmt.Sprintf("/%s/:%s", typePath, TypeGuid)
	UrlItemGuid    = fmt.Sprintf("/%s/:%s", itemPath, InvGuid)
)

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
