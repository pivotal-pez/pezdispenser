package vcloudclient

import (
	"encoding/xml"
	"net/http"
	"time"
)

type (
	//VCDAuth - vcd authentication object
	VCDClient struct {
		Token  string
		client httpClientDoer
	}
	//QueryResultRecords - root level query result xml object
	QueryResultRecords struct {
		XMLName            xml.Name             `xml:"QueryResultRecords"`
		VAppTemplateRecord []VAppTemplateRecord `xml:"VAppTemplateRecord"`
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
