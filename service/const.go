package pezdispenser

import (
	"errors"
	"net/http"
)

const (
	//TaskStatusAvailable - this means the task is in an avaiable state
	TaskStatusAvailable = "available"
	//TaskStatusRestocking - reclaiming inventory and restocking
	TaskStatusRestocking = "restocking"
	//TaskStatusUnavailable - unavailable procurement request
	TaskStatusUnavailable = "unavailable"
	//TaskStatusStarted - started this task
	TaskStatusStarted = "started"
	//TaskStatusProcurement - task is now in procurement
	TaskStatusProcurement = "in_procurement"
	//TaskCollectionName - collection name for tasks
	TaskCollectionName = "dispenser_tasks"
	//SuccessStatusResponseTaskByID - success statuscode for gettaskbyidcontroller
	SuccessStatusResponseTaskByID = http.StatusOK
	//FailureStatusResponseTaskByID - failure statuscode for gettaskbyidcontroller
	FailureStatusResponseTaskByID = http.StatusNotFound
	//CallerPostLease --
	CallerPostLease = "post_lease"
	//LeaseExpiresFieldName ----
	LeaseExpiresFieldName = "lease_expires"
	//InventoryIDFieldName ---
	InventoryIDFieldName = "inventory_id"
)

var (
	//ErrNoMatchInStore - error when there is no matching org in the datastore
	ErrNoMatchInStore = errors.New("Could not find a matching user org or connection failure")
	//ErrCanNotAddOrgRec - error when we can not add a new org record to the datastore
	ErrCanNotAddOrgRec = errors.New("Could not add a new org record")
	//ErrEmptyBody - no data in request body
	ErrEmptyBody = errors.New("request body is empty or invalid")
)
