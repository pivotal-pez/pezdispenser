package pdclient_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/pdclient"
)

var _ = Describe("given a GetRequestIDFromTaskResponse func", func() {
	Context("when called with a TaskResponse object containing request id", func() {
		var (
			controlRequestID = "63e24c92-a5da-11e5-bd74-0050569b9b57"
			requestID        string
		)
		BeforeEach(func() {
			tr := getTaskResponse(controlRequestID)
			requestID = GetRequestIDFromTaskResponse(tr)
		})
		It("then it should return properly parse out and return that requestid", func() {
			Ω(requestID).Should(Equal(controlRequestID))
		})
	})
})

func getTaskResponse(requestID string) (taskResponse TaskResponse) {
	var taskResponseString = fmt.Sprintf(`{
		"ID": "56748ece7bc989001d000002",
		"Timestamp": 1450479310174826830,
		"Expires": 0,
		"Status": "complete",
		"Profile": "agent_task_long_running",
		"CallerName": "m1.small",
		"MetaData": {
			"phinfo": {
				"data": [
					{
						"requestid": "%s"
					}
				],
				"message": "ok",
				"status": "success"
			},
			"status": {
				"data": {
					"credentials": {
						"name": "host-07-16",
						"oob_ip": "10.65.70.116",
						"oob_pw": "d3v0ps!",
						"oob_user": "pezuser"
					},
					"status": "complete"
				},
				"message": "ok",
				"status": "success"
			}
		}
	}`, requestID)
	json.Unmarshal([]byte(taskResponseString), &taskResponse)
	return
}
