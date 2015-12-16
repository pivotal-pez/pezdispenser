package pdclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

//NewClient - constructor for a new dispenser client
func NewClient(apiKey string, url string, client clientDoer) *PDClient {
	return &PDClient{
		APIKey: apiKey,
		client: client,
		URL:    url,
	}
}

func (s *PDClient) PostLease(leaseId, inventoryId, skuId string) {
	req, _ := s.createRequest("POST", s.URL, s.getRequestBody(leaseId, inventoryId, skuId))
	s.client.Do(req)
}

func (s *PDClient) getRequestBody(leaseId, inventoryId, skuId string) (body io.Reader) {
	var durationDays int64 = 14
	now := time.Now()
	expire := now.Add(time.Duration(durationDays) * 24 * time.Hour)
	body = bytes.NewBufferString(
		fmt.Sprintf(`{
			"lease_id":"%s",
			"inventory_id":"%s",
			"username":"joe@user.net",
			"sku":"%s",
			"lease_duration":%d,
			"lease_end_date":%d,
			"lease_start_date":%d,
			"procurement_meta":{}`,
			leaseId,
			inventoryId,
			skuId,
			durationDays,
			expire.UnixNano(),
			now.UnixNano()))
	return
}

func (s *PDClient) createRequest(method string, urlStr string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, urlStr, body)
	req.Header.Add("X-API-KEY", s.APIKey)
	return
}
