package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNamespaceResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccNamespaceResourceConfig("one", "Display Name One"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("terrareg_namespace.test", "name", "one"),
					resource.TestCheckResourceAttr("terrareg_namespace.test", "display_name", "Display Name One"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "terrareg_namespace.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				// ImportStateVerifyIgnore: []string{"configurable_attribute", "defaulted"},
			},
			// Update and Read testing
			{
				Config: testAccNamespaceResourceConfig("two", "Name Two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("terrareg_namespace.test", "name", "two"),
					resource.TestCheckResourceAttr("terrareg_namespace.test", "display_name", "Name Two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNamespaceResourceConfig(name string, displayName string) string {
	return buildTestProviderConfig(fmt.Sprintf(`
resource "terrareg_namespace" "test" {
  name         = %[1]q
  display_name = %[2]q
}
`, name, displayName))
}
