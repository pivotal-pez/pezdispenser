package pezdispenser_test

import (
	"net/http"

	. "github.com/pivotal-pez/pezdispenser/service"
	"github.com/xchapter7x/cloudcontroller-client"
	"github.com/xchapter7x/goutil"
	"gopkg.in/mgo.v2"
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
