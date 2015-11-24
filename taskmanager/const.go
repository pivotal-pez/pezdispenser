package taskmanager

import (
	"errors"
	"time"
)

const (
	//TaskAgentLongRunning --
	TaskAgentLongRunning ProfileType = "agent_task_long_running"
	//TaskAgentScheduledTask --
	TaskAgentScheduledTask ProfileType = "agent_scheduled_task"
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
	//AgentTaskStatusScheduled ---
	AgentTaskStatusScheduled = "scheduled"
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
	AgentTaskPollerInterval = time.Duration(30) * time.Second
	//AgentTaskPollerTimeout - time until a agent will expire its task if not polled
	AgentTaskPollerTimeout = time.Duration(5) * time.Minute
)
