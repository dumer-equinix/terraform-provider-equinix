package equinix

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccFabricReadPort(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricReadPortConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.equinix_fabric_port.test", "name", fmt.Sprint("ops-user100-CX-SV5-NL-Qinq-BO-10G-PRI-JP-180")),
				),
			},
		},
	})
}

func testAccFabricReadPortConfig() string {
	return fmt.Sprint(`data "equinix_fabric_port" "test" {
	uuid = "c4d9350e-77c5-7c5d-1ce0-306a5c00a600"
	}`)
}
