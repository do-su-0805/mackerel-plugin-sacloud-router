package main

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func main() {

	client := sacloud.NewClient("test", "test")
	fmt.Printf("%+v", client)
}
