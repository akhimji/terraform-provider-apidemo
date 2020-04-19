package main

import (
	"github.com/alyarctiq/terraform-provider-apidemo/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_ADDRESS", ""),
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_PORT", ""),
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_KEY", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"apidemo_entry": resourceServer(),
		},
		ConfigureFunc: providerConfigure,
	}

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname := d.Get("address").(string)
	port := d.Get("port").(string)
	key := d.Get("key").(string)
	return client.CreateClient(hostname, port, key), nil
}
