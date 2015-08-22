package fakes

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pivotal-pez/pezdispenser/service"
	"github.com/pivotal-pez/pezdispenser/service/_integrations"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xchapter7x/cloudcontroller-client"
	"github.com/xchapter7x/goutil"
)

const (
	//VcapServicesFormatter -
	VcapServicesFormatter = `{
				"p-mongodb": [
				{"name": "%s","label": "p-mongodb",
				"tags": ["pivotal","mongodb"],"plan": "development",
          "credentials": {
            "uri": "%s",
            "scheme": "mongodb",
            "username": "c39642c7-9bf4-4cbe-9ac9-db67a3bbc98f",
            "password": "f6ac4b827ea044393dc9ee553533154e",
            "host": "192.168.8.147",
            "port": 27017,
            "database": "70ef645b-7e3a-461c-a0d6-94d5b0e5107f"
			}}]}`
	//VcapApplicationFormatter -
	VcapApplicationFormatter = `{
				"limits":{"mem":1024,"disk":1024,"fds":16384},
				"application_version":"56637561-e847-4023-87fa-1e476cb0b7e3",
				"application_name":"dispenserdev",
				"application_uris":["dispenserdev.cfapps.pez.pivotal.io","dispenserdev.pezapp.io"],
				"version":"56637561-e847-4023-87fa-1e476cb0b7e3",
				"name":"dispenserdev",
				"space_name":"pez-dev",
				"space_id":"ea88ed9e-91f1-4763-8eef-54fe38acf603",
				"uris":["dispenserdev.cfapps.pez.pivotal.io","dispenserdev.pezapp.io"],
				"users":null
			}`
)

//FakeRenderer -
type FakeRenderer struct {
	render.Render
	SpyStatus int
	SpyValue  interface{}
}

//JSON -
func (s *FakeRenderer) JSON(status int, v interface{}) {
	s.SpyStatus = status
	s.SpyValue = v
}

var (
	//FakeKeyCheck -
	FakeKeyCheck martini.Handler = func(log *log.Logger, res http.ResponseWriter, req *http.Request) {}

	buf bytes.Buffer
	//MockLogger -
	MockLogger = log.New(&buf, "logger: ", log.Lshortfile)
)

//FakeNewCollectionDialer
func FakeNewCollectionDialer(c pezdispenser.Task) func(url, dbname, collectionname string) (col integrations.Collection, err error) {
	return func(url, dbname, collectionname string) (col integrations.Collection, err error) {
		fmt.Println("this is the one that was called")
		col = &FakeCollection{
			ControlTask: c,
		}
		return
	}
}

//MockMongo -
type MockMongo struct {
	Err    error
	Result interface{}
}

//Collection -
func (s *MockMongo) Collection() pezdispenser.Persistence {
	return &MockPersistence{
		Err:    s.Err,
		Result: s.Result,
	}
}

//Find -
func (s *MockMongo) Find(query interface{}) *mgo.Query {
	return new(mgo.Query)
}

//Remove -
func (s *MockMongo) Remove(selector interface{}) (err error) {
	err = s.Err
	return
}

//Upsert -
func (s *MockMongo) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	err = s.Err
	return
}

//MockPersistence -
type MockPersistence struct {
	Result interface{}
	Err    error
}

//Remove -
func (s *MockPersistence) Remove(selector interface{}) (err error) {
	return
}

//FindOne -
func (s *MockPersistence) FindOne(query interface{}, result interface{}) (err error) {
	goutil.Unpack([]interface{}{s.Result}, result)
	err = s.Err
	return
}

//Upsert -
func (s *MockPersistence) Upsert(selector interface{}, update interface{}) (err error) {
	return
}

//MockHeritageClient -
type MockHeritageClient struct {
	*ccclient.Client
	Res *http.Response
}

//CreateAuthRequest -
func (s *MockHeritageClient) CreateAuthRequest(verb, requestURL, path string, args interface{}) (*http.Request, error) {
	return &http.Request{}, nil
}

//CCTarget -
func (s *MockHeritageClient) CCTarget() string {
	return ccclient.URLPWSLogin
}

//HttpClient -
func (s *MockHeritageClient) HttpClient() ccclient.ClientDoer {
	return &MockClientDoer{
		Res: s.Res,
	}
}

//Login -
func (s *MockHeritageClient) Login() (c *ccclient.Client, err error) {
	return
}

//MockClientDoer -
type MockClientDoer struct {
	Req *http.Request
	Res *http.Response
	Err error
}

//Do -
func (s *MockClientDoer) Do(rq *http.Request) (rs *http.Response, e error) {
	s.Req = rq
	rs = s.Res
	e = s.Err
	return
}

//FakeTask -
type FakeTask struct {
	ID        bson.ObjectId          `bson:"_id"`
	Timestamp time.Time              `bson:"timestamp"`
	Status    string                 `bson:"status"`
	MetaData  map[string]interface{} `bson:"metadata"`
}

//FakeCollection -
type FakeCollection struct {
	mgo.Collection
	ControlTask pezdispenser.Task
}

//Close -
func (s *FakeCollection) Close() {

}

//UpsertID -
func (s *FakeCollection) UpsertID(id interface{}, result interface{}) (changInfo *mgo.ChangeInfo, err error) {
	return
}

//FindOne -
func (s *FakeCollection) FindOne(id string, result interface{}) (err error) {
	*(result.(*pezdispenser.Task)) = s.ControlTask
	return
}
