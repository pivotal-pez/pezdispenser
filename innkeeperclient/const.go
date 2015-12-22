package innkeeperclient

const (
	//StatusReady - ready status
	StatusReady = "ready"
	//StatusRunning - status running
	StatusRunning = "running"
	//StatusSuccess - success status
	StatusSuccess = "success"

	//RouteProvisionHost - route to provision host endpoint
	RouteProvisionHost = "api/v1/Provision"
	//RouteGetStatus - route to getstatus endpoint
	RouteGetStatus = "api/v1/StatusDetails"
	//RouteDeProvisionHost - deprovision endpoint
	RouteDeProvisionHost = "api/v1/Deprovision"

	requestIDGetParam = "requestid"
)
