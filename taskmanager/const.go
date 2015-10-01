package taskmanager

import "errors"

const (
	//TaskLeaseProcurement --
	TaskLeaseProcurement ProfileType = "lease_procurement"
	//TaskLeaseReStock --
	TaskLeaseReStock ProfileType = "lease_restock"
	//TaskInventoryLedger --
	TaskInventoryLedger ProfileType = "inventory_ledger"
	//TaskLongPollQueue --
	TaskLongPollQueue ProfileType = "longpoll_queue"
	//TaskChildID -- child task spawned from current task
	TaskChildID = "child_task_id"
	//TaskActionMetaName --
	TaskActionMetaName = "task_action"

	//ExpiredTask -
	ExpiredTask int64 = 0
	//TaskStatusAvailable --- task status is set to available
	TaskStatusAvailable = "available"
)

var (
	//ErrNoResults - no results found in query
	ErrNoResults = errors.New("no results found")
)
