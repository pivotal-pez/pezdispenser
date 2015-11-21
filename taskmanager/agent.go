package taskmanager

func (s *Agent) Run(process func(agent *Agent)) {
	go func() {
		process(s)
	}()
}
