package skurepo

var (
	//Repo -- repo holds the registered sku interfaces
	Repo = make(map[string]SkuBuilder)
)
