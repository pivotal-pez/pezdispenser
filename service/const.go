package pezdispenser

import "errors"

const (
	//TaskCollectionName - collection name for tasks
	TaskCollectionName = "dispenser_tasks"
	//SuccessStatusResponseTaskByID - success statuscode for gettaskbyidcontroller
	SuccessStatusResponseTaskByID = 200
	//FailureStatusResponseTaskByID - failure statuscode for gettaskbyidcontroller
	FailureStatusResponseTaskByID = 404
)

var (
	//ErrNoMatchInStore - error when there is no matching org in the datastore
	ErrNoMatchInStore = errors.New("Could not find a matching user org or connection failure")
	//ErrCanNotAddOrgRec - error when we can not add a new org record to the datastore
	ErrCanNotAddOrgRec = errors.New("Could not add a new org record")
)
