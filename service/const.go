package pezdispenser

import "errors"

const (
	//TaskStatusAvailable - this means the task is in an avaiable state
	TaskStatusAvailable = "available"
	//TaskStatusUnavailable - unavailable procurement request
	TaskStatusUnavailable = "unavailable"
	//TaskStatusStarted - started this task
	TaskStatusStarted = "started"
	//TaskStatusProcurement - task is now in procurement
	TaskStatusProcurement = "in_procurement"
	//TaskCollectionName - collection name for tasks
	TaskCollectionName = "dispenser_tasks"
	//SuccessStatusResponseTaskByID - success statuscode for gettaskbyidcontroller
	SuccessStatusResponseTaskByID = 200
	//FailureStatusResponseTaskByID - failure statuscode for gettaskbyidcontroller
	FailureStatusResponseTaskByID = 404
	//Sku2CSmall - lease sku type indicator. to be replaced with a cleaner injected pattern
	Sku2cSmall = "2c.small"

	TaskLeaseProcurement ProfileType = "lease_procurement"
	TaskLeaseReStock     ProfileType = "lease_restock"
	TaskUnDeploy         ProfileType = "undeploy"
	TaskDeploy           ProfileType = "deploy"
	TaskInventoryLedger  ProfileType = "inventory_ledger"
	TaskGeneric          ProfileType = "generic"

	CallerPostLease = "post_lease"
)

var (
	//ErrNoMatchInStore - error when there is no matching org in the datastore
	ErrNoMatchInStore = errors.New("Could not find a matching user org or connection failure")
	//ErrCanNotAddOrgRec - error when we can not add a new org record to the datastore
	ErrCanNotAddOrgRec = errors.New("Could not add a new org record")
	//ErrEmptyBody - no data in request body
	ErrEmptyBody = errors.New("request body is empty or invalid")
)
