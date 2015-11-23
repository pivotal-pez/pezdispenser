package taskmanager

import "errors"

const (
	//TaskAgentLongRunning --
	TaskAgentLongRunning ProfileType = "agent_task_long_running"
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

	//AgentTaskStatusInitializing ---
	AgentTaskStatusInitializing = "initializing"
	//AgentTaskStatusRunning ---
	AgentTaskStatusRunning = "running"
	//AgentTaskStatusComplete ---
	AgentTaskStatusComplete = "complete"
	//AgentTaskStatusFailed ---
	AgentTaskStatusFailed = "failed"
)

var (
	//ErrNoResults - no results found in query
	ErrNoResults = errors.New("no results found")
	//AgentTaskPollerInterval - time offset to poll a task from an agent
	AgentTaskPollerInterval = 30
	//AgentTaskPollerTimeout - time until a agent will expire its task if not polled
	AgentTaskPollerTimeout = 5 * 60
)
