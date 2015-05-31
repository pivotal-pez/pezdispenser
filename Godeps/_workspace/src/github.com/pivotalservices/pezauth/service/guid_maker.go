package pezauth

import "code.google.com/p/go-uuid/uuid"

//Create - creates a new random guid
func (s *GUIDMake) Create() string {
	r := uuid.New()
	return r
}
