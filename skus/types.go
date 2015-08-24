package skus

import "github.com/pivotal-pez/pezdispenser/vcloudclient"

type (
	//Sku - interface for a sku object
	Sku interface {
		Procurement(meta map[string]interface{}) (status string, taskMeta map[string]interface{})
		ReStock(meta map[string]interface{}) (status string, taskMeta map[string]interface{})
	}
	//Sku2CSmall - a object representing a 2csmall sku
	Sku2CSmall struct {
		Client vcdClient
	}

	vcdClient interface {
		DeployVApp(templateName, templateHref, vcdHref string) (vapp *vcloudclient.VApp, err error)
		Auth(username, password string) (err error)
		QueryTemplate(templateName string) (vappTemplate *vcloudclient.VAppTemplateRecord, err error)
	}
)
