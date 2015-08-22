package pezdispenser_test

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "github.com/pivotal-pez/pezdispenser/service"
	"github.com/xchapter7x/cloudcontroller-client"
	"github.com/xchapter7x/goutil"
	"gopkg.in/mgo.v2"
)

const (
	vcapServicesFormatter = `{
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
	vcapApplicationFormatter = `{
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

type fakeRenderer struct {
	render.Render
	SpyStatus int
	SpyValue  interface{}
}

func (s *fakeRenderer) JSON(status int, v interface{}) {
	s.SpyStatus = status
	s.SpyValue = v
}

var (
	fakeKeyCheck martini.Handler = func(log *log.Logger, res http.ResponseWriter, req *http.Request) {}
)

type mockMongo struct {
	err    error
	result interface{}
}

func (s *mockMongo) Collection() Persistence {
	return &mockPersistence{
		err:    s.err,
		result: s.result,
	}
}

func (s *mockMongo) Find(query interface{}) *mgo.Query {
	return new(mgo.Query)
}

func (s *mockMongo) Remove(selector interface{}) (err error) {
	err = s.err
	return
}

func (s *mockMongo) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	err = s.err
	return
}

type mockPersistence struct {
	result interface{}
	err    error
}

func (s *mockPersistence) Remove(selector interface{}) (err error) {
	return
}

func (s *mockPersistence) FindOne(query interface{}, result interface{}) (err error) {
	goutil.Unpack([]interface{}{s.result}, result)
	err = s.err
	return
}

func (s *mockPersistence) Upsert(selector interface{}, update interface{}) (err error) {
	return
}

type mockHeritageClient struct {
	*ccclient.Client
	res *http.Response
}

func (s *mockHeritageClient) CreateAuthRequest(verb, requestURL, path string, args interface{}) (*http.Request, error) {
	return &http.Request{}, nil
}

func (s *mockHeritageClient) CCTarget() string {
	return ccclient.URLPWSLogin
}

func (s *mockHeritageClient) HttpClient() ccclient.ClientDoer {
	return &mockClientDoer{
		res: s.res,
	}
}

func (s *mockHeritageClient) Login() (c *ccclient.Client, err error) {
	return
}

type mockClientDoer struct {
	req *http.Request
	res *http.Response
	err error
}

func (s *mockClientDoer) Do(rq *http.Request) (rs *http.Response, e error) {
	s.req = rq
	rs = s.res
	e = s.err
	return
}
