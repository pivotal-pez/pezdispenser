package pezdispenser

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"labix.org/v2/mgo"

	"encoding/json"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"labix.org/v2/mgo/bson"
)

//NewLease - create and return a new lease object
func NewLease(taskCollection integrations.Collection, availableSkus map[string]skus.Sku) *Lease {

	return &Lease{
		taskCollection: taskCollection,
		taskManager:    taskmanager.NewTaskManager(taskCollection),
		availableSkus:  availableSkus,
		Task:           new(taskmanager.Task).GetRedactedVersion(),
	}
}

//Delete - handle a delete lease call
func (s *Lease) Delete(logger *log.Logger, req *http.Request) (statusCode int, response interface{}) {
	var (
		err error
	)
	statusCode = http.StatusNotFound
	s.taskCollection.Wake()
	logger.Println("collection dialed successfully")

	if err = s.InitFromHTTPRequest(req); err == nil {
		logger.Println("restocking inventory...")
		s.ReStock()
		statusCode = http.StatusAccepted
		response = s

	} else {
		response = map[string]string{"error": err.Error()}
	}
	return
}

//Post - handle a post lease call
func (s *Lease) Post(logger *log.Logger, req *http.Request) (statusCode int, response interface{}) {
	var (
		err error
	)
	statusCode = http.StatusNotFound
	s.taskCollection.Wake()
	logger.Println("collection dialed successfully")

	if err = s.InitFromHTTPRequest(req); err == nil {
		logger.Println("obtaining lease...")
		s.Procurement()
		statusCode = http.StatusCreated
		response = s

	} else {
		response = map[string]string{"error": err.Error()}
	}
	return
}

//ReStock - this will reclaim resources for a given lease
func (s *Lease) ReStock() {

	if skuConstructor, ok := s.availableSkus[s.Sku]; ok {
		s.ProcurementMeta[InventoryIDFieldName] = s.InventoryID
		sku := skuConstructor.New(s.taskManager, s.ProcurementMeta)
		s.Task = sku.ReStock().GetRedactedVersion()

	} else {
		s.Task.Status = TaskStatusUnavailable
	}
}

//Procurement - method to issue a procurement request for the given lease item.
func (s *Lease) Procurement() {

	if skuConstructor, ok := s.availableSkus[s.Sku]; ok {
		s.Task.Status = TaskStatusUnavailable

		if s.InventoryAvailable() == true {
			s.ProcurementMeta[LeaseExpiresFieldName] = s.LeaseEndDate
			s.ProcurementMeta[InventoryIDFieldName] = s.InventoryID
			sku := skuConstructor.New(s.taskManager, s.ProcurementMeta)
			s.Task = sku.Procurement().GetRedactedVersion()
		}

	} else {
		s.Task.Status = TaskStatusUnavailable
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

	if s.ProcurementMeta == nil {
		s.ProcurementMeta = make(map[string]interface{})
	}
	return
}

//InventoryAvailable - lets check if a inventory management task exists for this
//inventory item. if one does not let's created it, if it does exist lets check
//its Status to see if we it available or not, return true or false on outcome
func (s *Lease) InventoryAvailable() (available bool) {
	task := new(taskmanager.Task)
	available = false

	ci, err := s.taskCollection.FindAndModify(
		bson.M{
			"_id":    s.InventoryID,
			"status": TaskStatusAvailable,
		},
		bson.M{
			"status": TaskStatusUnavailable,
		},
		task,
	)
	if ci.Updated == 1 {
		available = true

	} else if err == nil && ci.Updated <= 0 {

		if err := s.taskCollection.FindOne(s.InventoryID, task); err == mgo.ErrNotFound {
			task.ID = bson.ObjectIdHex(s.InventoryID)
			task.Timestamp = time.Now().UnixNano()
			task.Status = TaskStatusAvailable
			task.MetaData = s.ProcurementMeta
			s.taskManager.SaveTask(task)
			available = true
		}
	}
	return
}
