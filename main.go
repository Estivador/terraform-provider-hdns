package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/Estivador/terraform-provider-hdns/hdns"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return hdns.Provider()
		},
	})
}
