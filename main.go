package main

import (
	"fmt"

	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

type SacloudRouterPlugin struct {
	Prefix string
}

func main() {

	client := sacloud.NewClient("test", "test")
	s := SacloudRouterPlugin{Prefix: "TEST"}
	plugin := mp.NewMackerelPlugin(s)
	fmt.Printf("%+v", client)
}
