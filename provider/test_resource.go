package provider

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func testResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "This the name of a test resource that can be provisioned by Terraform",
				ValidateFunc: validateName,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A description of the test resource",
			},
		},
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// v = value of the attribute
// k = name of attribute
func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string

	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("name must be a string"))
		return warns, errs
	}

	length := regexp.MustCompile("[a-zA-Z0-9]{3,}")
	if length.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name must be at least 3 characters in length"))
		return warns, errs
	}

	return warns, errs
}