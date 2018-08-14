package heroku

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/heroku/heroku-go/v3"
)

func dataSourceHerokuAppConfigVars() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHerokuAppConfigVarsRead,
		Schema: map[string]*schema.Schema{

			"app": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"all_config_vars": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceHerokuAppConfigVarsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*heroku.Service)

	appName := d.Get("app").(string)
	d.Set("name", appName)
	d.SetId(appName)
	configVarInfo, err := client.ConfigVarInfoForApp(context.TODO(), appName)
	if err != nil {
		return err
	}

	if configVarInfo == nil {
		return fmt.Errorf(`Error getting config vars for %s`, appName)
	}
	configMap := make(map[string]string)

	for configKey, configValue := range configVarInfo {
		configMap[configKey] = *configValue
	}

	d.Set("all_config_vars", configMap)

	return nil
}
