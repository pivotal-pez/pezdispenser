package integrations

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func NewCollectionDialer(url string, dbname string, collectionname string) (collection Collection, err error) {
	var session *mgo.Session

	if session, err = mgo.Dial(url); err == nil {
		session.SetMode(mgo.Monotonic, true)
		db := session.DB(dbname)
		collection = &CollectionRepo{
			Col:     db.C(collectionname),
			session: session,
		}
	}
	return
}

func (s *CollectionRepo) FindOne(id string, result interface{}) (err error) {

	if bson.IsObjectIdHex(id) {
		hex := bson.ObjectIdHex(id)
		err = s.Col.FindId(hex).One(result)

	} else {
		err = ErrInvalidID
	}
	return
}

func (s *CollectionRepo) UpsertID(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = s.Col.UpsertId(id, update)
	return
}

func (s *CollectionRepo) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

func (s *CollectionRepo) Count() (int, error) {
	return s.Col.Count()
}
