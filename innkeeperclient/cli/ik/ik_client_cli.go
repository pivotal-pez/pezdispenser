package main

import (
	"fmt"

	"github.com/pivotal-pez/pezdispenser/skus/m1small"
)

func main() {
	var sku = &m1small.SkuM1Small{}
	clnt := sku.GetInnkeeperClient()
	gtinfo, _ := clnt.GetTenants()
	fmt.Println(gtinfo)
	phinfo, _ := clnt.ProvisionHost("PAO", "4D.lowmem.R7", 1, "pez-stage", "centos67")
	fmt.Println(phinfo)

}
