package hdns

import (
	"errors"
	"time"

	"github.com/hashicorp/logutils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns the hdns terraform provider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HDNS_TOKEN", nil),
				Description: "The API token to access the Hetzner DNS.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					token := val.(string)
					if len(token) != 32 {
						errs = append(errs, errors.New("entered token is invalid (must be exactly 32 characters long)"))
					}
					return
				},
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HDNS_ENDPOINT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hdns_zone":   resourceZone(),
			"hdns_record": resourceRecord(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hdns_zone":    dataSourceHdnsZone(),
			"hdns_zones":   dataSourceHdnsZones(),
			"hdns_record":  dataSourceHdnsRecord(),
			"hdns_records": dataSourceHdnsRecords(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	opts := []hdns.ClientOption{
		hdns.WithToken(d.Get("token").(string)),
	}
	if endpoint, ok := d.GetOk("endpoint"); ok {
		opts = append(opts, hdns.WithEndpoint(endpoint.(string)))
	}
	if pollInterval, ok := d.GetOk("poll_interval"); ok {
		pollInterval, err := time.ParseDuration(pollInterval.(string))
		if err != nil {
			return nil, err
		}
		opts = append(opts, hdns.WithPollInterval(pollInterval))
	}
	if logging.LogLevel() != "" {
		writer, err := logging.LogOutput()
		if err != nil {
			return nil, err
		}
		opts = append(opts, hdns.WithDebugWriter(writer.(*logutils.LevelFilter).Writer))
	}
	return hdns.NewClient(opts...), nil
}
