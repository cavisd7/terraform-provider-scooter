package main

import (
	"github.com/cavisd7/terraform-provider-scooter/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var version = "v0.0.2"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
