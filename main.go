// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"log"
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"ssd-provider/internal/provider/sdd"
)


var (
	version string = "0.0.1"
)

func main() {

   var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")

	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		ProviderAddr: "github.com/borgeslima/ssd-provider",
		ProviderFunc: func() *schema.Provider {
			return sdd.Provider()
		},
		Debug: debug,
	})
}
