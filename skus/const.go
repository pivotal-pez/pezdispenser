package skus

const (
	//StatusComplete - a status flag for complete tasks
	StatusComplete = "complete"
	//StatusFailed - a status for failed tasks
	StatusFailed = "failed"
	//StatusProcessing - a status for in process items
	StatusProcessing = "processing"
	//StatusOutsourced - this is to indicate the the task tracking has been outsourced
	StatusOutsourced = "outsourced"
	//VCDTaskElementHrefMetaName - the name of the meta data field containing the href for the vcd task
	VCDTaskElementHrefMetaName = "vcd_task_element_href"
	//TaskActionUnDeploy --
	TaskActionUnDeploy = "undeploy_vapp"
	//SkuName2CSmall --
	SkuName2CSmall = "2c.small"
)
