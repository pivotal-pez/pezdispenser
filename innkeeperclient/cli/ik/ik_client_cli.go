package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
)

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	clnt := innkeeperclient.New(logger)
	gtinfo, _ := clnt.GetTenants()
	fmt.Println(gtinfo)
	phinfo, _ := clnt.ProvisionHost("PAO", "4D.lowmem.R7", 1, "pez-stage", "centos67")
	fmt.Println(phinfo)

}
