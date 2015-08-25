package pezdispenser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"labix.org/v2/mgo"

	"encoding/json"

	"github.com/pivotal-pez/pezdispenser/service/_integrations"
	"github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
	"labix.org/v2/mgo/bson"
)

func NewLease(taskCollection integrations.Collection) *Lease {
	return &Lease{
		taskCollection: taskCollection,
		taskManager:    NewTaskManager(taskCollection),
	}
}

//Post - handle a post lease call
func (s *Lease) Post(logger *log.Logger, req *http.Request) (statusCode int, response interface{}) {
	var (
		err       error
		newTaskID = bson.NewObjectId().Hex()
		timestamp = time.Now()
		task      = &Task{
			ID:         bson.ObjectIdHex(newTaskID),
			Status:     TaskStatusStarted,
			Timestamp:  timestamp,
			MetaData:   make(map[string]interface{}),
			Profile:    TaskLeaseProcurement,
			CallerName: CallerPostLease,
		}
	)
	statusCode = http.StatusNotFound
	s.taskCollection.Wake()
	logger.Println("collection dialed successfully")

	if _, err = s.taskCollection.UpsertID(newTaskID, task); err == nil {
		s.SetTask(task)
		logger.Println("task created")

		if err = s.InitFromHTTPRequest(req); err == nil {
			logger.Println("restocking...")
			s.ReStock()
			statusCode = http.StatusCreated
			response = s
		}
	}

	if err != nil {
		response = map[string]string{"error": err.Error()}
	}
	return
}

//ReStock - this will reclaim resources for a given lease
func (s *Lease) ReStock() {
	switch s.Sku {
	case Sku2cSmall:
		s.Task.Status = TaskStatusUnavailable

		if s.InventoryAvailable() {
			httpClient := vcloudclient.DefaultClient()
			baseURI := fmt.Sprintf("%s", s.ProcurementMeta["base_uri"])
			sku := &skus.Sku2CSmall{
				Client: vcloudclient.NewVCDClient(httpClient, baseURI),
			}
			s.Task.Status, s.ConsumerMeta = sku.ReStock(s.ProcurementMeta)
		}
		s.taskManager.SaveTask(s.Task)

	default:
		s.Task.Status = TaskStatusUnavailable
		s.taskManager.SaveTask(s.Task)
	}
}

//Procurement - method to issue a procurement request for the given lease item.
func (s *Lease) Procurement() {
	switch s.Sku {
	case Sku2cSmall:
		s.Task.Status = TaskStatusUnavailable

		if s.InventoryAvailable() {
			sku := new(skus.Sku2CSmall)
			s.Task.Status, s.ConsumerMeta = sku.Procurement(s.ProcurementMeta)
		}
		s.taskManager.SaveTask(s.Task)

	default:
		s.Task.Status = TaskStatusUnavailable
		s.taskManager.SaveTask(s.Task)
	}
}

//InitFromHTTPRequest - initialize a lease from the http request object body
func (s *Lease) InitFromHTTPRequest(req *http.Request) (err error) {

	if req.Body != nil {

		if body, err := ioutil.ReadAll(req.Body); err == nil {
			err = json.Unmarshal(body, s)
		}
	} else {
		err = ErrEmptyBody
	}
	return
}

//SetTask - add a task to the lease object
func (s *Lease) SetTask(task *Task) {
	s.Task = task
	s.taskManager.SaveTask(s.Task)
}

//InventoryAvailable - lets check if a inventory management task exists for this
//inventory item. if one does not let's created it, if it does exist lets check
//its Status to see if we it available or not, return true or false on outcome
func (s *Lease) InventoryAvailable() (available bool) {
	task := new(Task)
	available = false

	if err := s.taskCollection.FindOne(s.InventoryID, task); task.Status == TaskStatusAvailable {
		available = true

	} else if err == mgo.ErrNotFound {
		task.ID = bson.ObjectIdHex(s.InventoryID)
		task.Timestamp = time.Now()
		task.Status = TaskStatusAvailable
		task.MetaData = s.ProcurementMeta
		s.taskManager.SaveTask(task)
		available = true
	}
	return
}
