package vcloud_client

func NewVCDAuth() *VCDAuth {
	return &VCDAuth{
		token: "random",
	}
}

func (s *VCDAuth) GetToken() (token string) {
	token = s.token
	return
}
