package consul

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceConsulAgentCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceConsulAgentCheckCreate,
		Update: resourceConsulAgentCheckCreate,
		Read:   resourceConsulAgentCheckRead,
		Delete: resourceConsulAgentCheckDelete,

		Schema: map[string]*schema.Schema{
			"Node": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"CheckID": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"Name": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"Status": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"Notes": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"Output": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"ServiceID": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"ServiceName": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
		},
	}
}
