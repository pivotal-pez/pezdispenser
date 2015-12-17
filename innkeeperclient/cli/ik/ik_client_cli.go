package main

import (
	"fmt"

	"github.com/pivotal-pez/pezdispenser/skus/m1small"
)

func main() {
	var sku = &m1small.SkuM1Small{}
	clnt, err := sku.GetInnkeeperClient()
	if err == nil {
		gtinfo, _ := clnt.GetTenants()
		fmt.Println(gtinfo)
	}
}
