package pezauth

import "gopkg.in/mgo.v2"

type (
	mongoCollection interface {
		Find(query interface{}) *mgo.Query
		Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	}
	mongoCollectionWrapper struct {
		Persistence
		col mongoCollection
	}
)

func newMongoCollectionWrapper(c mongoCollection) Persistence {
	return &mongoCollectionWrapper{
		col: c,
	}
}

//FindOne - combining the Find and One calls of a mongo collection object
func (s *mongoCollectionWrapper) FindOne(query interface{}, result interface{}) (err error) {

	if err = s.col.Find(query).One(result); err != nil {
		err = ErrNoMatchInStore
	}
	return
}

//Upsert - allow us to call upsert on mongo collection object
func (s *mongoCollectionWrapper) Upsert(selector interface{}, update interface{}) (err error) {

	if _, err = s.col.Upsert(selector, update); err != nil {
		err = ErrCanNotAddOrgRec
	}
	return
}
