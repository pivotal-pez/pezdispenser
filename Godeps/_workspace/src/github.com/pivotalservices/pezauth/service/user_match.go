package pezauth

//NewUserMatch - creates a new usermatch struct pointer
func NewUserMatch() *UserMatch {
	return new(UserMatch)
}

//UserInfo - accepts a userinfo object grabbed from google auth
func (s *UserMatch) UserInfo(userInfo map[string]interface{}) *UserMatch {
	s.userInfo = userInfo
	return s
}

//UserName - takes a username which is passed as part of the rest call
func (s *UserMatch) UserName(username string) *UserMatch {
	s.username = username
	return s
}

//OnSuccess - function to run if they are allowed to make the calls
func (s *UserMatch) OnSuccess(successFunc func()) *UserMatch {
	s.successFunc = successFunc
	return s
}

//OnFailure - function to call if they are not allowed to make the call
func (s *UserMatch) OnFailure(failFunc func()) *UserMatch {
	s.failFunc = failFunc
	return s
}

//Run - executes the check and run of success or failure function
func (s *UserMatch) Run() (err error) {
	var hasValidEmail = false

	for _, email := range s.userInfo["emails"].([]interface{}) {

		if email.(map[string]interface{})["value"].(string) == s.username {
			hasValidEmail = true
			s.successFunc()
		}
	}

	if !hasValidEmail {
		s.failFunc()
		err = ErrNotValidActionForUser
	}
	return
}
