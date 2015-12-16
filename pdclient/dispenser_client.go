package pdclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xchapter7x/lo"
)

//NewClient - constructor for a new dispenser client
func NewClient(apiKey string, url string, client clientDoer) *PDClient {
	return &PDClient{
		APIKey: apiKey,
		client: client,
		URL:    url,
	}
}

func (s *PDClient) PostLease(leaseId, inventoryId, skuId string, leaseDaysDuration int64) (leaseCreateResponse LeaseCreateResponseBody, res *http.Response, err error) {
	var body io.Reader
	if body, err = s.getRequestBody(leaseId, inventoryId, skuId, leaseDaysDuration); err == nil {
		req, _ := s.createRequest("POST", fmt.Sprintf("%s/v1/lease", s.URL), body)
		res, err = s.client.Do(req)
		resBodyBytes, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(resBodyBytes, &leaseCreateResponse)

	} else {
		lo.G.Error("request body error: ", err.Error())
	}
	return
}

func (s *PDClient) getRequestBody(leaseId, inventoryId, skuId string, durationDays int64) (body io.Reader, err error) {
	var (
		now       = time.Now()
		bodyBytes []byte
	)
	expire := now.Add(time.Duration(durationDays) * 24 * time.Hour)
	leaseBody := LeaseRequestBody{
		LeaseID:        leaseId,
		InventoryID:    inventoryId,
		Sku:            skuId,
		LeaseDuration:  durationDays,
		LeaseEndDate:   expire.UnixNano(),
		LeaseStartDate: now.UnixNano(),
	}
	if bodyBytes, err = json.Marshal(leaseBody); err == nil {
		body = bytes.NewBuffer(bodyBytes)
	}
	return
}

func (s *PDClient) createRequest(method string, urlStr string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, urlStr, body)
	req.Header.Add("X-API-KEY", s.APIKey)
	return
}
