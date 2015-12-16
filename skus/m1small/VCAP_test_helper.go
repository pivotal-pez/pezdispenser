package m1small

import "fmt"

var vcapServices = `{
"user-provided": [
        {
          "name": "pezvalidator-service",
          "label": "user-provided",
          "tags": [],
          "credentials": {
            "target-url": "https://hcfdev.pezapp.io/valid-key"
          },
          "syslog_drain_url": ""
        },
        {
          "name": "innkeeper-service",
          "label": "user-provided",
          "tags": [],
          "credentials": {
            "enable": "%d",
            "uri": "%s",
            "password": "%s",
            "user": "%s"
          },
          "syslog_drain_url": ""
        }
      ]
}
`
var vcapApplication = `{
      "limits": {
        "mem": 1024,
        "disk": 1024,
        "fds": 16384
      },
      "application_version": "0",
      "application_name": "r",
      "application_uris": [
        "pivotal.io"
      ],
      "version": "0",
      "name": "r",
      "space_name": "z",
      "space_id": "4",
      "uris": [
        "pivotal.io"
      ],
      "users": null
		}`

// GetDefaultVCAPApplicationString -- vcap application env string
func GetDefaultVCAPApplicationString() string {
	return vcapApplication
}

// GetDefaultVCAPServicesString -- vcap services env string, with default innkeeper
func GetDefaultVCAPServicesString() string {
	return GetVCAPServicesString(1, "http://innkeeper.pivotal.io", "admin", "somepass")
}

// GetVCAPServicesString - vcap services env string with specific innkeeper params
func GetVCAPServicesString(enable int, uri string, user string, password string) string {
	// default uri = http://innkeeper.pivotal.io
	return fmt.Sprintf(vcapServices, enable, uri, password, user)
}
