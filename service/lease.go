package pezdispenser

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"labix.org/v2/mgo"

	"encoding/json"

	"github.com/pivotal-pez/pezdispenser/service/_integrations"
	"github.com/pivotal-pez/pezdispenser/skus"
	"labix.org/v2/mgo/bson"
)

func NewLease(taskCollection integrations.Collection) *Lease {
	return &Lease{
		taskCollection: taskCollection,
	}
}

//Post - handle a post lease call
func (s *Lease) Post(logger *log.Logger, req *http.Request) (statusCode int, response interface{}) {
	var (
		err       error
		newTaskID = bson.NewObjectId().Hex()
		timestamp = time.Now()
		task      = &Task{
			ID:        bson.ObjectIdHex(newTaskID),
			Status:    TaskStatusStarted,
			Timestamp: timestamp,
			MetaData:  make(map[string]interface{}),
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
	s.Task.Status = TaskStatusUnavailable
	s.saveTask()
}

//Procurement - method to issue a procurement request for the given lease item.
func (s *Lease) Procurement() {
	switch s.Sku {
	case Sku2cSmall:
		s.Task.Status = TaskStatusUnavailable

		if s.InventoryAvailable() {
			sku := new(skus.Sku2CSmall)
			s.Task.Status, s.Task.MetaData = sku.Procurement(s.ProcurementMeta)
		}
		s.saveTask()

	default:
		s.Task.Status = TaskStatusUnavailable
		s.saveTask()
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
	s.saveTask()
}

func (s *Lease) saveTask() {
	s.taskCollection.UpsertID(s.Task.ID, s.Task)
}

//InventoryAvailable - lets check if a inventory management task exists for this
//inventory item. if one does not let's created it, if it does exist lets check
//its Status to see if we it available or not, return true or false on outcome
func (s *Lease) InventoryAvailable() (available bool) {
	task := Task{}
	available = false

	if err := s.taskCollection.FindOne(s.InventoryID, &task); task.Status == TaskStatusAvailable {
		available = true

	} else if err == mgo.ErrNotFound {
		task.ID = bson.ObjectIdHex(s.InventoryID)
		task.Timestamp = time.Now()
		task.Status = TaskStatusAvailable
		task.MetaData = s.ProcurementMeta
		s.taskCollection.UpsertID(task.ID, task)
		available = true
	}
	return
}
