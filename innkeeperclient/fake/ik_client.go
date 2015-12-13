package fakeinnkeeperclient

import (
	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
)

// IKClient -- fake!
type IKClient struct {
	innkeeperclient.InnkeeperClient
	FakeStatus []string
	FakeMessage []string
	cnt int
}

// ProvisionHost -- 
func (s *IKClient) ProvisionHost(geoLoc string, sku string, count int, tenantid string, osarg string) (result *innkeeperclient.ProvisionHostResponse, err error) {
	result = new(innkeeperclient.ProvisionHostResponse)
	result.Status = s.FakeStatus[s.cnt]
	result.Message = s.FakeMessage[s.cnt]
	s.cnt++
	return
}