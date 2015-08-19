package vcloudclient

import (
	"encoding/xml"
	"net/http"
	"time"
)

type (
	//VCDAuth - vcd authentication object
	VCDClient struct {
		//BaseURI
		BaseURI string
		//Token
		Token  string
		client httpClientDoer
	}
	//VApp - xml object used in response from deploy app api call
	VApp struct {
		//XMLName
		XMLName xml.Name `xml:"VApp"`
		//Link
		Link []interface{} `xml:"Link"`
		//Description
		Description string `xml:"Description"`
		//Tasks
		Tasks TasksElem `xml:"Tasks"`
		//Files
		Files []interface{} `xml:"Files"`
		//VAppParent
		VAppParent interface{} `xml:"VAppParent"`
		//Owner
		Owner interface{} `xml:"Owner"`
		//InMaintenanceMode
		InMaintenanceMode bool `xml:"InMaintenanceMode"`
		//Children
		Children []interface{} `xml:"Children"`
	}
	//TasksElem
	TasksElem struct {
		Task TaskElem `xml:"Task"`
	}
	//TaskElem - object representing a XML task response object from the api
	TaskElem struct {
		//Href
		Href string `xml:"href,attr"`
		//Status
		Status string `xml:"status,attr"`
		//StartTime
		StartTime   time.Time `xml:"startTime,attr"`
		Description string    `xml:"Description"`
	}
	//QueryResultRecords - root level query result xml object
	QueryResultRecords struct {
		XMLName            xml.Name             `xml:"QueryResultRecords"`
		VAppTemplateRecord []VAppTemplateRecord `xml:"VAppTemplateRecord"`
		Link               []interface{}        `xml:"Link"`
	}
	//VAppTemplateRecord - vapp result from query
	VAppTemplateRecord struct {
		//VdcName
		VdcName string `xml:"vdcName,attr"`
		//Vdc
		Vdc string `xml:"vdc,attr"`
		//StorageProfileName
		StorageProfileName string `xml:"storageProfileName,attr"`
		//Status
		Status string `xml:"status,attr"`
		//OwnerName
		OwnerName string `xml:"ownerName,attr"`
		//Org
		Org string `xml:"org,attr"`
		//Name
		Name string `xml:"name,attr"`
		//IsPublished
		IsPublished bool `xml:"isPublished,attr"`
		//IsGoldMaster
		IsGoldMaster bool `xml:"isGoldMaster,attr"`
		//IsExpired
		IsExpired bool `xml:"isExpired,attr"`
		//IsEnabled
		IsEnabled bool `xml:"isEnabled,attr"`
		//IsDeployed
		IsDeployed bool `xml:"isDeployed,attr"`
		//IsBusy
		IsBusy bool `xml:"isBusy,attr"`
		//CreationDate
		CreationDate time.Time `xml:"creationDate,attr"`
		//CatalogName
		CatalogName string `xml:"catalogName,attr"`
		//Href
		Href string `xml:"href,attr"`
		//IsInCatalog
		IsInCatalog bool `xml:"isInCatalog,attr"`
		//CpuAllocationMhz
		CpuAllocationMhz float64 `xml:"cpuAllocationMhz,attr"`
		//CpuAllocationInMhz
		CpuAllocationInMhz float64 `xml:"cpuAllocationInMhz,attr"`
		//Task
		Task string `xml:"task,attr"`
		//NumberOfShadowVMs
		NumberOfShadowVMs float64 `xml:"numberOfShadowVMs,attr"`
		//AutoDeleteDate
		AutoDeleteDate time.Time `xml:"autoDeleteDate,attr"`
		//IsAutoDeleteNotified
		IsAutoDeleteNotified bool `xml:"isAutoDeleteNotified,attr"`
		//NumberOfVMs
		NumberOfVMs float64 `xml:"numberOfVMs,attr"`
		//IsAutoUndeployNotified
		IsAutoUndeployNotified bool `xml:"isAutoUndeployNotified,attr"`
		//TaskStatusName
		TaskStatusName string `xml:"taskStatusName,attr"`
		//IsVdcEnabled
		IsVdcEnabled bool `xml:"isVdcEnabled,attr"`
		//HonorBootOrder
		HonorBootOrder bool `xml:"honorBootOrder,attr"`
		//TaskStatus
		TaskStatus string `xml:"taskStatus,attr"`
		//StorageKB
		StorageKB float64 `xml:"storageKB,attr"`
		//TaskDetails
		TaskDetails string `xml:"taskDetails,attr"`
		//NumberOfCpus
		NumberOfCpus float64 `xml:"numberOfCpus,attr"`
		//MemoryAllocationMB
		MemoryAllocationMB float64 `xml:"memoryAllocationMB,attr"`
	}
	httpClientDoer interface {
		Do(req *http.Request) (resp *http.Response, err error)
	}
)
