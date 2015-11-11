package pezdispenser

import (
	"fmt"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
)

//SetupDB - setup your db connection and return a collection interface
func SetupDB(dialer integrations.CollectionDialer, URI string, DBName string, collectionName string) (collection integrations.Collection) {
	var (
		err error
	)

	if collection, err = dialer(URI, DBName, collectionName); err != nil {
		panic(fmt.Sprintf("can not dial connection due to error: %s URI:%s col:%s db:%s", err.Error(), URI, collectionName, DBName))
	}
	return
}
