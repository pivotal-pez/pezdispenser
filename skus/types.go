package skus

type (
	//Sku - interface for a sku object
	Sku interface {
		Procurement(meta map[string]interface{}) (status string, taskMeta map[string]interface{})
		ReStock(meta map[string]interface{}) (status string, taskMeta map[string]interface{})
	}
	//Sku2CSmall - a object representing a 2csmall sku
	Sku2CSmall struct {
	}
)
