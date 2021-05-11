package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cavisd7/terraform-provider-scooter/api/client"
	"github.com/cavisd7/terraform-provider-scooter/api/server"
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
	if !length.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name must be at least 3 characters in length"))
		return warns, errs
	}

	return warns, errs
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	item := server.Item{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	err := apiClient.NewItem(&item)
	if err != nil {
		return err
	}

	d.SetId(item.Name)

	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error occurred while looking for item: %s", itemId)
		}
	}

	d.SetId(item.Name)
	d.Set("name", item.Name)
	d.Set("description", item.Description)
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	item := server.Item{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}

	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteItem(itemId)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
