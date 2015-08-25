package skus

import "fmt"

//Procurement - this method will walk the procurement flow for the 2csmall
//object
func (s *Sku2CSmall) Procurement(meta map[string]interface{}) (status string, taskMeta map[string]interface{}) {
	status = StatusComplete
	return
}

//ReStock - this method will walk the restock flow for the 2csmall object
func (s *Sku2CSmall) ReStock(meta map[string]interface{}) (status string, taskMeta map[string]interface{}) {
	taskMeta = make(map[string]interface{})
	template := fmt.Sprintf("%s", meta["template_name"])
	//baseURI := fmt.Sprintf("%s", meta["base_uri"])
	user := fmt.Sprintf("%s", meta["vcd_username"])
	pass := fmt.Sprintf("%s", meta["vcd_password"])
	vAppID := fmt.Sprintf("%s", meta["vapp_id"])
	s.Client.Auth(user, pass)
	s.Client.UnDeployVApp(vAppID)

	vappTemplate, _ := s.Client.QueryTemplate(template)
	vapp, _ := s.Client.DeployVApp(template, vappTemplate.Href, vappTemplate.Vdc)
	taskMeta["vcd_task_href"] = vapp.Tasks.Task.Href
	status = StatusOutsourced
	return
}
