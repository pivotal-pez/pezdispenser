package pezdispenser

//NewMongoCollectionWrapper - wraps a Mongo collection in as a Peristence interface implementation
func NewMongoCollectionWrapper(c MongoCollection) Persistence {
	return &MongoCollectionWrapper{
		col: c,
	}
}

//FindOne - combining the Find and One calls of a Mongo collection object
func (s *MongoCollectionWrapper) FindOne(query interface{}, result interface{}) (err error) {
	if err = s.col.Find(query).One(result); err != nil {
		err = ErrNoMatchInStore
	}
	return
}

//Upsert - allow us to call upsert on Mongo collection object
func (s *MongoCollectionWrapper) Upsert(selector interface{}, update interface{}) (err error) {
	if _, err = s.col.Upsert(selector, update); err != nil {
		err = ErrCanNotAddOrgRec
	}
	return
}

//Remove - removes the matching selector from collection
func (s *MongoCollectionWrapper) Remove(selector interface{}) error {
	return s.col.Remove(selector)
}
