package pezdispenser

import (
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

//NewLease - create and return a new lease object
func NewLease(taskCollection integrations.Collection, availableSkus map[string]skurepo.Sku) *Lease {

	return &Lease{
		taskCollection: taskCollection,
		taskManager:    taskmanager.NewTaskManager(taskCollection),
		availableSkus:  availableSkus,
		Task:           taskmanager.RedactedTask{},
	}
}

//Delete - handle a delete lease call
func (s *Lease) Delete(logger *log.Logger, req *http.Request) (statusCode int, response interface{}) {
	var (
		err error
	)
	statusCode = http.StatusNotFound
	s.taskCollection.Wake()

	if err = s.InitFromHTTPRequest(req); err == nil {
		logger.Println("restocking inventory...")
		s.ReStock()
		statusCode = http.StatusAccepted
		response = s.Task

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
		logger.Println("obtaining lease...", s)
		s.Procurement()
		statusCode = http.StatusCreated
		response = s.Task

	} else {
		response = map[string]string{"error": err.Error()}
	}
	return
}

//ReStock - this will reclaim resources for a given lease
func (s *Lease) ReStock() (skuTask *taskmanager.Task) {

	if skuConstructor, ok := s.availableSkus[s.Sku]; ok {
		s.ProcurementMeta[InventoryIDFieldName] = s.InventoryID
		sku := skuConstructor.New(s.taskManager, s.ProcurementMeta)
		skuTask = sku.ReStock()
		s.Task = skuTask.GetRedactedVersion()

	} else {
		s.Task.Status = TaskStatusUnavailable
	}
	return
}

//Procurement - method to issue a procurement request for the given lease item.
func (s *Lease) Procurement() (skuTask *taskmanager.Task) {

	if skuConstructor, ok := s.availableSkus[s.Sku]; ok {
		s.ProcurementMeta[LeaseExpiresFieldName] = s.LeaseEndDate
		s.ProcurementMeta[InventoryIDFieldName] = s.InventoryID
		sku := skuConstructor.New(s.taskManager, s.ProcurementMeta)
		GLogger.Println("here is my sku: ", sku)
		skuTask = sku.Procurement()
		tt := skuTask.Read(func(t *taskmanager.Task) interface{} {
			tt := *t
			return tt
		})
		GLogger.Println("here is my task after procurement: ", tt)
		s.Task = skuTask.GetRedactedVersion()

	} else {
		GLogger.Println("No Sku Match: ", s.Sku, s.availableSkus)
		s.Task.Status = TaskStatusUnavailable
	}
	return
}

//InitFromHTTPRequest - initialize a lease from the http request object body
func (s *Lease) InitFromHTTPRequest(req *http.Request) (err error) {

	if req.Body != nil {

		if body, err := ioutil.ReadAll(req.Body); err == nil {

			if err = json.Unmarshal(body, s); err != nil {
				GLogger.Println(err)
			}
		}
	} else {
		err = ErrEmptyBody
	}

	if s.ProcurementMeta == nil {
		s.ProcurementMeta = make(map[string]interface{})
	}
	return
}
