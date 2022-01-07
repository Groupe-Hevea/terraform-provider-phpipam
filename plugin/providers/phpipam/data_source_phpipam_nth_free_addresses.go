package phpipam

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourcePHPIPAMNthFreeAddresses() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePHPIPAMNthFreeAddressesRead,
		Schema: map[string]*schema.Schema{
			"subnet_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourcePHPIPAMNthFreeAddressesRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*ProviderPHPIPAMClient).subnetsController

	s, err := c.GetSubnetByID(d.Get("subnet_id").(int))
	if err != nil {
		return err
	}
	used_ips, err := c.GetAddressesInSubnet(d.Get("subnet_id").(int))
	if err != nil {

		// return err
		return errors.New("[ERROR] RED ALERT")
	}

	n := d.Get("number").(int)

	nthFreeAddresses, err := FilterUsedIPAddresses(fmt.Sprintf("%s/%d", s.SubnetAddress, s.Mask), &used_ips, n)
	if err != nil {
		return errors.New("[ERROR] failed to retrieve nth first ip addresses")
	}

	log.Printf("[INFO] Found %d free addresses", len(*nthFreeAddresses))

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	d.Set("ip_addresses", nthFreeAddresses)

	return nil
}
