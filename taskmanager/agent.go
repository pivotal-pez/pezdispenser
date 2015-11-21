package taskmanager

import (
	"fmt"
	"time"
)

//NewAgent -- creates a new initialized agent object
func NewAgent(t TaskManagerInterface, callerName string) *Agent {
	return &Agent{
		statusEmitter: make(chan string, 1),
		taskManager:   t,
		task:          t.NewTask(callerName, TaskAgentLongRunning, AgentTaskStatusInitializing),
	}
}

//Run - this begins the running of an agent's async process
func (s *Agent) Run(process func(agent *Agent) error) {
	s.task.Status = AgentTaskStatusRunning
	s.statusEmitter <- s.task.Status
	s.taskManager.SaveTask(s.task)

	go func() {

		if err := process(s); err == nil {
			s.task.Status = AgentTaskStatusComplete

		} else {
			s.task.Status = fmt.Sprintf("status: %s, error: %s", AgentTaskStatusFailed, err.Error())
		}
		s.task.Expires = 0
		s.statusEmitter <- s.task.Status
		s.taskManager.SaveTask(s.task)
	}()
}

//GetTask - get the agents task object
func (s *Agent) GetTask() *Task {
	return s.task
}

//GetStatus - returns a status emitting channel
func (s *Agent) GetStatus() chan string {
	return s.statusEmitter
}

func (s *Agent) executeTaskPoller() {
	for {
		time.Sleep(time.Duration(AgentTaskPollerInterval) * time.Second)
	}
}
