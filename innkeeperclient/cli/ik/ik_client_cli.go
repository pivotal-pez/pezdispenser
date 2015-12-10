package main

import (
	"fmt"

	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
)

func main() {
	clnt := innkeeperclient.New()
	gtinfo, _ := clnt.GetTenants()
	fmt.Println(gtinfo)
	phinfo, _ := clnt.ProvisionHost("PAO", "4D.lowmem.R7", 1, "pez-stage", "centos67")
	fmt.Println(phinfo)

}
