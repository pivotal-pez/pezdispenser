package pezdispenser

import (
	"gopkg.in/mgo.v2"
)

type (
	mongoCollectionGetter interface {
		Collection() Persistence
	}

	mongoCollection interface {
		Remove(selector interface{}) error
		Find(query interface{}) *mgo.Query
		Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	}

	mongoCollectionWrapper struct {
		Persistence
		col mongoCollection
	}

	//Persistence - interface to a persistence store of some kind
	Persistence interface {
		Remove(selector interface{}) error
		FindOne(query interface{}, result interface{}) (err error)
		Upsert(selector interface{}, update interface{}) (err error)
	}
)
