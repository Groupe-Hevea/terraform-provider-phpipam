package phpipam

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const testAccDataSourcePHPIPAMFirstFreeSubnetConfig = `
data "phpipam_subnet" "subnet_by_cidr" {
	subnet_address = "10.10.1.0"
	subnet_mask = 24
}

data "phpipam_first_free_subnet" "next" {
	subnet_id = "${data.phpipam_subnet.subnet_by_cidr.subnet_id}"
	subnet_mask = 26
}
`

const testAccDataSourcePHPIPAMFirstFreeSubnetNoFreeConfig = `
resource "phpipam_subnet" "subnet" {
  section_id     = 1
  subnet_address = "10.10.3.0"
  subnet_mask    = 24
}

resource "phpipam_subnet" "subnet1"  {
  section_id     = 1
	subnet_address = "10.10.3.0"
  subnet_mask    = 25
}

resource "phpipam_subnet" "subnet2"  {
  section_id     = 1
	subnet_address = "10.10.3.128"
  subnet_mask    = 25
}

data "phpipam_first_free_subnet" "next" {
	subnet_id = "${data.phpipam_subnet.subnet_by_cidr.subnet_id}"
	subnet_mask = 25
}

  depends_on = [
    "phpipam_subnet.subnet1",
    "phpipam_subnet.subnet2",
  ]
}
`

func TestAccDataSourcePHPIPAMFirstFreeSubnet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourcePHPIPAMFirstFreeAddressConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.phpipam_first_free_subnet.next", "subnet", "10.10.1.0/26"),
				),
			},
		},
	})
}

func TestAccDataSourcePHPIPAMFirstFreeSubnetNoFree(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccDataSourcePHPIPAMFirstFreeSubnetNoFreeConfig,
				ExpectError: regexp.MustCompile("Subnet has no free IP addresses"),
			},
		},
	})
}
