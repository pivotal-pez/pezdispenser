package vcloudclient_test

import (
	"io"
	"net/http"
)

const (
	TaskResponseFormatter = `
<?xml version="1.0" encoding="UTF-8"?>
<Task xmlns="http://www.vmware.com/vcloud/v1.5" href="xs:anyURI" type="xs:string" name="xs:string" id="xs:string"
        status="%s" operation="xs:string" operationName="xs:string" startTime="xs:dateTime"
        endTime="xs:dateTime" expiryTime="xs:dateTime">
    <Link href="xs:anyURI" id="xs:string" type="xs:string" name="xs:string"
            rel="xs:string"/>
    <Description> xs:string </Description>
    <Tasks>
        <Task> TaskType </Task>
    </Tasks>
    <Owner href="xs:anyURI" id="xs:string" type="xs:string" name="xs:string"/>
    <Error message="xs:string" majorErrorCode="xs:int" minorErrorCode="xs:string" vendorSpecificErrorCode="xs:string"
            stackTrace="xs:string"/>
    <User href="xs:anyURI" id="xs:string" type="xs:string" name="xs:string"/>
    <Organization href="xs:anyURI" id="xs:string" type="xs:string" name="xs:string"/>
    <Progress> xs:int </Progress>
    <Params> ... </Params>
</Task>
	`
)

type fakeHttpClient struct {
	Err      error
	Response *http.Response
}

func (s *fakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	return s.Response, s.Err
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
