package taskmanager

const (
	//TaskLeaseProcurement --
	TaskLeaseProcurement ProfileType = "lease_procurement"
	//TaskLeaseReStock --
	TaskLeaseReStock ProfileType = "lease_restock"
	//TaskInventoryLedger --
	TaskInventoryLedger ProfileType = "inventory_ledger"
	//TaskLongPollQueue --
	TaskLongPollQueue ProfileType = "longpoll_queue"

	//TaskActionMetaName --
	TaskActionMetaName = "task_action"

	//ExpiredTask -
	ExpiredTask int64 = 0
)
