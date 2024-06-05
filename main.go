package main

import (
	"context"
	"flag"
	"log"

	"github.com/borgeslima/ssd-provider/sdd"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "1.0.0"
)

func main() {

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "github.com/borgeslima/ssd-provider",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), sdd.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
