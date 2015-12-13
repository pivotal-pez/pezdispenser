package taskmanager

import(
	
)
//SetPrivateMeta - set a private meta data record
func (s *Task) SetPrivateMeta(name string, value interface{}) {
	if s.PrivateMetaData == nil {
		s.PrivateMetaData = make(map[string]interface{})
	}
	
	s.PrivateMetaData[name] = value
}

//SetPublicMeta - set a public metadata record
func (s *Task) SetPublicMeta(name string, value interface{}) {
	if s.MetaData == nil {
		s.MetaData = make(map[string]interface{})
	}
	s.MetaData[name] = value
}

//GetPublicMeta - get a public metadata record
func (s *Task) GetPublicMeta(name string) interface{} {
	return s.MetaData[name]
}

//GetPrivateMeta - get a private meta record
func (s *Task) GetPrivateMeta(name string) interface{} {
	return s.PrivateMetaData[name]
}

//GetRedactedVersion - returns a redacted version of this task, removing private info
func (s *Task) GetRedactedVersion() RedactedTask {
	return RedactedTask{
		ID:         s.ID,
		Timestamp:  s.Timestamp,
		Expires:    s.Expires,
		Status:     s.Status,
		Profile:    s.Profile,
		CallerName: s.CallerName,
		MetaData:   s.MetaData,
	}
}
// Update -- Safe way to update a task
func (s *Task) Update(update func(*Task) (interface{})) (interface{}){
	s.Mutex.Lock()
	var ret = update(s)
	s.TaskManager.SaveTask(s)
	s.Mutex.Unlock()
	return ret
}

// Read -- Safe way to read from a task
func (s *Task) Read(read func(*Task) (interface{})) (interface{}){
	s.Mutex.RLock()
	var ret = read(s)
	s.Mutex.RUnlock()
	return ret
}
