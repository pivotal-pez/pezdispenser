package vcloud_client

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
	QueryResultRecords struct {
		XMLName            xml.Name             `xml:"QueryResultRecords"`
		VAppTemplateRecord []VAppTemplateRecord `xml:"VAppTemplateRecord"`
	}
	VAppTemplateRecord struct {
		VdcName                string    `xml:"vdcName,attr"`
		Vdc                    string    `xml:"vdc,attr"`
		StorageProfileName     string    `xml:"storageProfileName,attr"`
		Status                 string    `xml:"status,attr"`
		OwnerName              string    `xml:"ownerName,attr"`
		Org                    string    `xml:"org,attr"`
		Name                   string    `xml:"name,attr"`
		IsPublished            bool      `xml:"isPublished,attr"`
		IsGoldMaster           bool      `xml:"isGoldMaster,attr"`
		IsExpired              bool      `xml:"isExpired,attr"`
		IsEnabled              bool      `xml:"isEnabled,attr"`
		IsDeployed             bool      `xml:"isDeployed,attr"`
		IsBusy                 bool      `xml:"isBusy,attr"`
		CreationDate           time.Time `xml:"creationDate,attr"`
		CatalogName            string    `xml:"catalogName,attr"`
		Href                   string    `xml:"href,attr"`
		IsInCatalog            bool      `xml:"isInCatalog,attr"`
		CpuAllocationMhz       float64   `xml:"cpuAllocationMhz,attr"`
		CpuAllocationInMhz     float64   `xml:"cpuAllocationInMhz,attr"`
		Task                   string    `xml:"task,attr"`
		NumberOfShadowVMs      float64   `xml:"numberOfShadowVMs,attr"`
		AutoDeleteDate         time.Time `xml:"autoDeleteDate,attr"`
		IsAutoDeleteNotified   bool      `xml:"isAutoDeleteNotified,attr"`
		NumberOfVMs            float64   `xml:"numberOfVMs,attr"`
		IsAutoUndeployNotified bool      `xml:"isAutoUndeployNotified,attr"`
		TaskStatusName         string    `xml:"taskStatusName,attr"`
		IsVdcEnabled           bool      `xml:"isVdcEnabled,attr"`
		HonorBootOrder         bool      `xml:"honorBootOrder,attr"`
		TaskStatus             string    `xml:"taskStatus,attr"`
		StorageKB              float64   `xml:"storageKB,attr"`
		TaskDetails            string    `xml:"taskDetails,attr"`
		NumberOfCpus           float64   `xml:"numberOfCpus,attr"`
		MemoryAllocationMB     float64   `xml:"memoryAllocationMB,attr"`
	}
	httpClientDoer interface {
		Do(req *http.Request) (resp *http.Response, err error)
	}
)
