package skus

//Procurement - this method will walk the procurement flow for the 2csmall
//object
func (s *Sku2CSmall) Procurement(meta map[string]interface{}) (status string, taskMeta map[string]interface{}) {
	status = StatusComplete
	return
}

//ReStock - this method will walk the restock flow for the 2csmall object
func (s *Sku2CSmall) ReStock(meta map[string]interface{}) (status string, taskMeta map[string]interface{}) {
	status = StatusComplete
	return
}
