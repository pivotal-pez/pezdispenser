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
	s.task.Status = AgentTaskStatusRunning
	s.statusEmitter <- s.task.Status
	s.taskManager.SaveTask(s.task)
	go s.startTaskPoller()
	go s.listenForPoll()

	go func(agent Agent) {
		s := &agent
		fmt.Println("Running Agent process")
		err := process(s)
		fmt.Println("Done Agent process", err)
		
		s.taskPollEmitter <- false

		select {
		case <-s.processComplete:
			if err == nil {
				s.task.Status = AgentTaskStatusComplete

			} else {
				s.task.Status = fmt.Sprintf("status: %s, error: %s", AgentTaskStatusFailed, err.Error())
			}
			s.task.Expires = 0
			s.taskManager.SaveTask(s.task)
			s.statusEmitter <- s.task.Status
		}
	}(*s)
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
		s.task.Expires = time.Now().Add(AgentTaskPollerTimeout).UnixNano()
		s.taskManager.SaveTask(s.task)
	}
	s.killTaskPoller <- true
}
