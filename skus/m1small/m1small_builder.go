package m1small

import "github.com/pivotal-pez/pezdispenser/skurepo"

// New -- return a new SKU provider
func (s *SkuM1SmallBuilder) New(tm skurepo.TaskManager, meta map[string]interface{}) skurepo.Sku {
	var procurementMeta = make(map[string]interface{})
	var userIdentifier string

	if meta != nil {

		if v, ok := meta[procurementMetaFieldName]; ok {
			procurementMeta = v.(map[string]interface{})
		}

		if v, ok := meta[userIdentifierMetaFieldName]; ok {
			userIdentifier = v.(string)
		}
	}

	return &SkuM1Small{
		Client:          s.Client,
		ProcurementMeta: procurementMeta,
		TaskManager:     tm,
		UserIdentifier:  userIdentifier,
	}
}
