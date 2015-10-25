package fakes

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pivotal-pez/pezdispenser/service"
	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/skus/2csmall"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xchapter7x/cloudcontroller-client"
	"github.com/xchapter7x/goutil"
)

const (
	FakeCollectionHasChanges       = 1
	FakeCollectionHasNoChanges     = 0
	FakeCollectionHasNilChangeInfo = -1
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

//MakeFakeSku2CSmall ---
func MakeFakeSku2CSmall(status string) (*s2csmall.Sku2CSmall, *taskmanager.Task, *taskmanager.Task) {
	s := new(s2csmall.Sku2CSmall)
	spyTask := &taskmanager.Task{
		ID:      bson.NewObjectId(),
		Expires: time.Now().UnixNano(),
		PrivateMetaData: map[string]interface{}{
			s2csmall.VCDTaskElementHrefMetaName: "vcdTask.url.com/hithere",
			taskmanager.TaskActionMetaName:      s2csmall.TaskActionUnDeploy,
		},
	}
	s.Client = &FakeVCDClient{
		FakeVApp:               new(vcloudclient.VApp),
		FakeVAppTemplateRecord: new(vcloudclient.VAppTemplateRecord),
		ErrPollTaskURL:         nil,
		FakeTaskElem: &vcloudclient.TaskElem{
			Status: status,
		},
	}
	myFakeManager := &FakeTaskManager{
		ReturnedTask: spyTask,
		SpyTaskSaved: new(taskmanager.Task),
	}
	s.TaskManager = myFakeManager
	return s, spyTask, myFakeManager.SpyTaskSaved
}

//FakeSku -- a fake sku object
type FakeSku struct {
	ProcurementTask *taskmanager.Task
	ReStockTask     *taskmanager.Task
}

//Procurement --
func (s *FakeSku) Procurement() (task *taskmanager.Task) {
	return s.ProcurementTask
}

//ReStock --
func (s *FakeSku) ReStock() (task *taskmanager.Task) {
	return s.ReStockTask
}

//PollForTasks --
func (s *FakeSku) PollForTasks() {

}

//New --
func (s *FakeSku) New(tm skurepo.TaskManager, procurementMeta map[string]interface{}) skurepo.Sku {
	return s
}

//FakeVCDClient - this is a fake vcdclient object
type FakeVCDClient struct {
	FakeVAppTemplateRecord *vcloudclient.VAppTemplateRecord
	FakeVApp               *vcloudclient.VApp
	ErrUnDeployFake        error
	ErrDeployFake          error
	ErrQueryFake           error
	ErrAuthFake            error
	ErrPollTaskURL         error
	FakeTaskElem           *vcloudclient.TaskElem
}

//PollTaskURL -- fake a poll url call
func (s *FakeVCDClient) PollTaskURL(taskURL string) (task *vcloudclient.TaskElem, err error) {
	return s.FakeTaskElem, s.ErrPollTaskURL
}

//DeployVApp - fake out calling deploy vapp
func (s *FakeVCDClient) DeployVApp(templateName, templateHref, vcdHref string) (vapp *vcloudclient.VApp, err error) {
	return s.FakeVApp, s.ErrUnDeployFake
}

//UnDeployVApp - executes a fake undeploy on a fake client
func (s *FakeVCDClient) UnDeployVApp(vappID string) (task *vcloudclient.TaskElem, err error) {
	return &s.FakeVApp.Tasks.Task, s.ErrDeployFake
}

//Auth - fake out making an auth call
func (s *FakeVCDClient) Auth(username, password string) (err error) {
	return s.ErrAuthFake
}

//QueryTemplate - fake querying for a template
func (s *FakeVCDClient) QueryTemplate(templateName string) (vappTemplate *vcloudclient.VAppTemplateRecord, err error) {
	return s.FakeVAppTemplateRecord, s.ErrDeployFake
}

//FakeResponseBody - a fake response body object
type FakeResponseBody struct {
	io.Reader
}

//Close - close fake body
func (FakeResponseBody) Close() error { return nil }

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

//FakeNewCollectionDialer -
func FakeNewCollectionDialer(c taskmanager.Task) func(url, dbname, collectionname string) (col integrations.Collection, err error) {
	return func(url, dbname, collectionname string) (col integrations.Collection, err error) {
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

//HTTPClient -
func (s *MockHeritageClient) HTTPClient() ccclient.ClientDoer {
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

//FakeTaskManager - this is a fake representation of the task manager
type FakeTaskManager struct {
	SpyTaskSaved *taskmanager.Task
	ReturnedTask *taskmanager.Task
	ReturnedErr  error
}

//SaveTask --
func (s *FakeTaskManager) SaveTask(t *taskmanager.Task) (*taskmanager.Task, error) {

	if s.SpyTaskSaved != nil {
		*s.SpyTaskSaved = *t
	}
	fmt.Println("we have saved this", s.SpyTaskSaved)
	return t, s.ReturnedErr
}

//FindAndStallTaskForCaller --
func (s *FakeTaskManager) FindAndStallTaskForCaller(callerName string) (t *taskmanager.Task, err error) {
	return s.ReturnedTask, s.ReturnedErr
}

//FindTask --
func (s *FakeTaskManager) FindTask(id string) (t *taskmanager.Task, err error) {
	return s.ReturnedTask, s.ReturnedErr
}

//NewTask --
func (s *FakeTaskManager) NewTask(callerName string, profile taskmanager.ProfileType, status string) (t *taskmanager.Task) {
	t = new(taskmanager.Task)
	t.CallerName = callerName
	t.Profile = profile
	t.Status = status
	t.ID = bson.NewObjectId()
	t.Timestamp = time.Now().UnixNano()
	t.MetaData = make(map[string]interface{})
	t.PrivateMetaData = make(map[string]interface{})
	return
}

//FakeTask -
type FakeTask struct {
	ID              bson.ObjectId          `bson:"_id"`
	Timestamp       time.Time              `bson:"timestamp"`
	Status          string                 `bson:"status"`
	MetaData        map[string]interface{} `bson:"metadata"`
	PrivateMetaData map[string]interface{} `bson:"private_metadata"`
}

//NewFakeCollection ====
func NewFakeCollection(updated int) *FakeCollection {
	fakeCol := new(FakeCollection)

	if updated == -1 {
		fakeCol.FakeChangeInfo = nil
	} else {
		fakeCol.FakeChangeInfo = &mgo.ChangeInfo{
			Updated: updated,
		}
	}
	return fakeCol
}

//FakeCollection -
type FakeCollection struct {
	mgo.Collection
	ControlTask      taskmanager.Task
	ErrControl       error
	FakeChangeInfo   *mgo.ChangeInfo
	ErrFindAndModify error
}

//Close -
func (s *FakeCollection) Close() {

}

//Wake -
func (s *FakeCollection) Wake() {

}

//FindAndModify -
func (s *FakeCollection) FindAndModify(selector interface{}, update interface{}, result interface{}) (info *mgo.ChangeInfo, err error) {
	return s.FakeChangeInfo, s.ErrFindAndModify
}

//UpsertID -
func (s *FakeCollection) UpsertID(id interface{}, result interface{}) (changInfo *mgo.ChangeInfo, err error) {
	return
}

//FindOne -
func (s *FakeCollection) FindOne(id string, result interface{}) (err error) {
	err = s.ErrControl
	*(result.(*taskmanager.Task)) = s.ControlTask
	return
}
