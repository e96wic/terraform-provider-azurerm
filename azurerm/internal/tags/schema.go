package tags

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// DataSourceSchema returns the Schema which should be used for Tags on a Data Source
func DataSourceSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
	}
}

// ForceNewSchema returns the Schema which should be used for Tags when changes
// require recreation of the resource
func ForceNewSchema() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeMap,
		Optional:     true,
		Computed:     true,
		ForceNew:     true,
		ValidateFunc: Validate,
	}
}

// Schema returns the Schema used for Tags
func Schema() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeMap,
		Optional:     true,
		Computed:     true,
		ValidateFunc: Validate,
	}
}
