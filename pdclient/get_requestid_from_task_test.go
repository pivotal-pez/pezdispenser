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
			requestID, _ = GetRequestIDFromTaskResponse(tr)
		})
		It("then it should return properly parse out and return that requestid", func() {
			Ω(requestID).Should(Equal(controlRequestID))
		})
	})
	Context("when an object not containing a populated data array is returned", func() {
		It("then it should not panic", func() {

			tr := getNoDataTaskResponse()
			Ω(func() {
				GetRequestIDFromTaskResponse(tr)
			}).ShouldNot(Panic())
		})
		Context("when calling getrequestidfromtaskresponse without panic with an empty data oboject", func() {

			var requestID string
			var err error
			BeforeEach(func() {
				tr := getNoDataTaskResponse()
				requestID, err = GetRequestIDFromTaskResponse(tr)
			})
			It("then it should not return an error", func() {
				Ω(requestID).Should(BeEmpty())
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
})

func getNoDataTaskResponse() (taskResponse TaskResponse) {
	var taskResponseString = `{
		"ID": "56748ece7bc989001d000002",
		"Timestamp": 1450479310174826830,
		"Expires": 0,
		"Status": "complete",
		"Profile": "agent_task_long_running",
		"CallerName": "m1.small",
		"MetaData": {}
	}`
	json.Unmarshal([]byte(taskResponseString), &taskResponse)
	return

}

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
			}
		}
	}`, requestID)
	json.Unmarshal([]byte(taskResponseString), &taskResponse)
	return
}
