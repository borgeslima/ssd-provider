package main

import (
	"context"
	"flag"
	"log"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/borgeslima/ssd-provider/sdd/provider"
)


var (

	version string = "1.0.0"

)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "Definido como true para executar o provedor com suporte para depuradores como delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "github.com/borgeslima/ssd-provider",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.Provider(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}