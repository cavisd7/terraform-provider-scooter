package provider

import (
	"github.com/cavisd7/terraform-provider-scooter/api/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			// Attributes that can be provided to the provider block in Terraform
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_PORT", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"test_item": testResource(),
		},
		ConfigureFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	port := d.Get("port").(int)
	return client.NewClient(address, port), nil
}
