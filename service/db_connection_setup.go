package pezdispenser

import (
	"fmt"

	"labix.org/v2/mgo"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
)

func SetupDB(dialer integrations.CollectionDialer, URI string, collectionName string) (collection integrations.Collection) {
	var (
		err      error
		dialInfo *mgo.DialInfo
	)

	if dialInfo, err = ParseURL(URI); err != nil || dialInfo.Database == "" {
		panic(fmt.Sprintf("can not parse given URI %s due to error: %s", URI, err.Error()))
	}

	if collection, err = dialer(URI, dialInfo.Database, collectionName); err != nil {
		panic(fmt.Sprintf("can not dial connection due to error: %s URI:%s col:%s db:%s", err.Error(), URI, collectionName, dialInfo.Database))
	}
	return
}
