package innkeeperclient

type (
	logger interface {
		Println(...interface{})
	}
	// InnkeeperClient api client
	InnkeeperClient interface {
		ProvisionHost(sku string, tenantid string) (result *ProvisionHostResponse, err error)
		GetStatus(requestID string) (resp *GetStatusResponse, err error)
		GetTenants() (result GetTenantsResponse, err error)
		DeProvisionHost(requestID string) (resp *GetStatusResponse, err error)
	}
	// IkClient api struct
	IkClient struct {
		URI      string
		User     string
		Password string
		Log      logger
	}
	// ProvisionHostResponse --
	ProvisionHostResponse struct {
		Data    []RequestData `json:"data"`
		Message string        `json:"message"`
		Status  string        `json:"status"`
	}

	//Data --
	Data struct {
		Credentials interface{}            `json:"credentials"`
		Status      string                 `json:"status"`
		Storage     map[string]interface{} `json:"storage"`
		Hosts       map[string]interface{} `json:"hosts"`
		Network     map[string]interface{} `json:"network"`
	}
	//GetStatusResponse -- a status response object
	GetStatusResponse struct {
		Data    Data   `json:"data"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	// GetTenantsResponse --
	GetTenantsResponse struct {
		Data    []TenantData `json:"data"`
		Message string       `json:"message"`
		Status  string       `json:"status"`
	}
	//RequestData - a request data object
	RequestData struct {
		RequestID string `json:"requestid"`
	}

	//TenantData - a tenant data object
	TenantData struct {
		Slotid   int    `json:"slotid"`
		Tenantid string `json:"tenantid"`
	}
)
