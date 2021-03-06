package taskmanager

import (
	"fmt"
	"time"
)

//NewAgent -- creates a new initialized agent object
func NewAgent(t TaskManagerInterface, callerName string) *Agent {
	return &Agent{
		killTaskPoller:  make(chan bool, 1),
		processComplete: make(chan bool, 1),
		taskPollEmitter: make(chan bool, 1),
		statusEmitter:   make(chan string, 1),
		taskManager:     t,
		task:            t.NewTask(callerName, TaskAgentLongRunning, AgentTaskStatusInitializing),
	}
}

//Run - this begins the running of an agent's async process
func (s *Agent) Run(process func(*Agent) error) {
	s.task.Update(func(t *Task) interface{} {
		t.taskManager = s.taskManager
		t.Status = AgentTaskStatusRunning
		return t
	})
	s.statusEmitter <- AgentTaskStatusRunning
	go s.startTaskPoller()
	go s.listenForPoll()

	go func(agent Agent) {
		s := &agent
		s.processExitHanderlDecorate(process)
		<-s.processComplete
	}(*s)
}

func (s *Agent) processExitHanderlDecorate(process func(*Agent) error) {
	err := process(s)
	fmt.Println("Done Agent process", process, err)

	s.taskPollEmitter <- false

	status := AgentTaskStatusComplete
	if err != nil {
		status = fmt.Sprintf("status: %s, error: %s", AgentTaskStatusFailed, err.Error())
	}
	s.task.Update(func(t *Task) interface{} {
		t.Status = status
		t.Expires = 0
		return t
	})
	s.statusEmitter <- status
}

//GetTask - get the agents task object
func (s *Agent) GetTask() *Task {
	return s.task
}

//GetStatus - returns a status emitting channel
func (s *Agent) GetStatus() chan string {
	return s.statusEmitter
}

func (s *Agent) startTaskPoller() {
ForLoop:
	for {
		select {
		case <-s.killTaskPoller:
			s.processComplete <- true
			break ForLoop
		default:
			s.taskPollEmitter <- true
		}
		time.Sleep(AgentTaskPollerInterval)
	}
}

func (s *Agent) listenForPoll() {
	for <-s.taskPollEmitter {
		s.task.Update(func(t *Task) interface{} {
			t.Expires = time.Now().Add(AgentTaskPollerTimeout).UnixNano()
			return t
		})
	}
	s.killTaskPoller <- true
}
