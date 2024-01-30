package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccModuleResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: buildTestProviderConfig(testAccNamespaceResourceConfig_basic),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("terrareg_module.example", "id", "module-basic-example/basic-example/aws"),
					resource.TestCheckResourceAttr("terrareg_module.example", "namespace", "module-basic-example"),
					resource.TestCheckResourceAttr("terrareg_module.example", "name", "basic-example"),
					resource.TestCheckResourceAttr("terrareg_module.example", "provider_name", "aws"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "terrareg_module.example",
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
				Config: buildTestProviderConfig(testAccNamespaceResourceConfig_basic_read),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("terrareg_module.example2", "id", "module-basic-example-import/basic-example2/aws"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "namespace", "module-basic-example-import"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "name", "basic-example2"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "provider_name", "aws"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

const testAccNamespaceResourceConfig_basic = `
resource "terrareg_namespace" "this" {
  name = "module-basic-example"
}
data "terrareg_git_provider" "this" {
  name = "Gitlab"
}

resource "terrareg_module" "example" {
  namespace      = terrareg_namespace.this.name
  name           = "basic-example"
  provider_name  = "aws"

  git_provider_id = data.terrareg_git_provider.this.id
  git_tag_format  = "v{version}"
}
`

const testAccNamespaceResourceConfig_basic_read = `
resource "terrareg_namespace" "this" {
 name = "module-basic-example-import"
}
data "terrareg_git_provider" "this" {
  name = "Gitlab"
}
  
resource "terrareg_module" "example2" {
  namespace      = terrareg_namespace.this.name
  name           = "basic-example2"
  provider_name  = "aws"
  
  git_provider_id = data.terrareg_git_provider.this.id
  git_tag_format  = "v{version}"
}
`
