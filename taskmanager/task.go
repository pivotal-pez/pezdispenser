package taskmanager

import(
	"sync"
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
	s.mutex.Lock()
	var ret = update(s)
	s.taskManager.SaveTask(s)
	s.mutex.Unlock()
	return ret
}

// Read -- Safe way to read from a task
func (s *Task) Read(read func(*Task) (interface{})) (interface{}){
	s.mutex.RLock()
	var ret = read(s)
	s.mutex.RUnlock()
	return ret
}

// Protect -- add mutex and taskmanager protection to task
func (s *Task) Protect(taskmanager TaskManagerInterface, mutex sync.RWMutex){
	s.taskManager = taskmanager
	s.mutex = mutex	
}

// Equal - define task equality
func (s Task) Equal (b Task) bool {
	return (s.ID == b.ID &&
		s.Timestamp == b.Timestamp &&
		s.Expires == b.Expires &&
		s.Status == b.Status &&
		s.Profile == b.Profile &&
		s.CallerName == b.CallerName)
}