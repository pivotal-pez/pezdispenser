package pezdispenser

import (
	"fmt"

	"github.com/go-martini/martini"
)

const (
	ApiVersion = "v1"
	indexRoute = "/"
	leasePath  = "lease"
	lockPath   = "lock"
	typePath   = "type"
	itemPath   = "item"
	invGuid    = ":inventoryTypeGuid"
	typeGuid   = ":inventoryItemGuid"
)

var (
	UrlLeaseBase = fmt.Sprintf("/%s/%s", ApiVersion, leasePath)
	UrlLockBase  = fmt.Sprintf("/%s/%s", ApiVersion, lockPath)
	UrlTypeGuid  = fmt.Sprintf("/%s/%s", typePath, typeGuid)
	UrlItemGuid  = fmt.Sprintf("/%s/%s", itemPath, invGuid)
)

func InitRoutes(m *martini.ClassicMartini) {

	m.Group(UrlLeaseBase, func(r martini.Router) {
		r.Post(UrlTypeGuid, RandomController)
		r.Post(UrlItemGuid, RandomController)
		r.Delete(UrlItemGuid, RandomController)
	})

	m.Group(UrlLockBase, func(r martini.Router) {
		r.Post(UrlItemGuid, RandomController)
		r.Get(UrlItemGuid, RandomController)
	})
}
