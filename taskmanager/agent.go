package taskmanager

func NewAgent(t *Task) *Agent {
	return &Agent{
		task: t,
	}
}

func (s *Agent) Run(process func(agent *Agent)) (task *Task) {
	go func() {
		process(s)
	}()
	return s.task
}
