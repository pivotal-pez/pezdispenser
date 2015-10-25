package s2csmall

import (
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

func init() {
	skurepo.Register(SkuName2CSmall, new(Sku2CSmall))
}

type (
	//Sku2CSmall - a object representing a 2csmall sku
	Sku2CSmall struct {
		Client          vcdClient
		TaskManager     skurepo.TaskManager
		ProcurementMeta map[string]interface{}
	}

	vcdClient interface {
		UnDeployVApp(vappID string) (task *vcloudclient.TaskElem, err error)
		DeployVApp(templateName, templateHref, vcdHref string) (vapp *vcloudclient.VApp, err error)
		Auth(username, password string) (err error)
		QueryTemplate(templateName string) (vappTemplate *vcloudclient.VAppTemplateRecord, err error)
		PollTaskURL(string) (*vcloudclient.TaskElem, error)
	}
)
