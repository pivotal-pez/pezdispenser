package fakeinnkeeperclient

import (
	"sync/atomic"

	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
)

// IKClient -- fake!
type IKClient struct {
	innkeeperclient.InnkeeperClient
	FakeStatus         []string
	FakeMessage        []string
	FakeData           []innkeeperclient.RequestData
	cnt                int
	SpyStatusCallCount *int64
}

// ProvisionHost --
func (s *IKClient) ProvisionHost(sku string, tenantid string) (result *innkeeperclient.ProvisionHostResponse, err error) {
	result = new(innkeeperclient.ProvisionHostResponse)
	result.Data = make([]innkeeperclient.RequestData, 1)
	if len(s.FakeStatus) > s.cnt {
		result.Status = s.FakeStatus[s.cnt]
	}
	if len(s.FakeData) > s.cnt {
		result.Data[0] = s.FakeData[s.cnt]
	}
	if len(s.FakeMessage) > s.cnt {
		result.Message = s.FakeMessage[s.cnt]
	}
	s.cnt++
	return
}

//GetStatus --
func (s *IKClient) GetStatus(requestID string) (resp *innkeeperclient.GetStatusResponse, err error) {
	atomic.AddInt64(s.SpyStatusCallCount, 1)
	return
}
