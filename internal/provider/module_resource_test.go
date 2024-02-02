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
					resource.TestCheckResourceAttr("terrareg_module.example", "id", "module-basic-example/basic-example2/awsnew"),
					resource.TestCheckResourceAttr("terrareg_module.example", "namespace", "module-basic-example"),
					resource.TestCheckResourceAttr("terrareg_module.example", "name", "basic-example2"),
					resource.TestCheckResourceAttr("terrareg_module.example", "provider_name", "awsnew"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccModuleResource_full(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: buildTestProviderConfig(testAccNamespaceResourceConfig_full),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("terrareg_module.example2", "id", "module-basic-example2/basic-example3/aws"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "namespace", "module-basic-example2"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "name", "basic-example3"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "provider_name", "aws"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "git_tag_format", "v{version}3"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "repo_base_url_template", "https://somecustom-domain.com/{namespace}/{module}-{provider}"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "repo_clone_url_template", "ssh://git@some-custom-domain.com/{namespace}/{module}-{provider}.git"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "repo_browse_url_template", "https://some-custom-domain.com/{namespace}/{module}-{provider}/tree/{tag}/{path}"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "terrareg_module.example2",
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
				Config: buildTestProviderConfig(testAccNamespaceResourceConfig_full_read),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("terrareg_module.example2", "id", "module-basic-example2/basic-example3/aws"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "namespace", "module-basic-example2"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "name", "basic-example3"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "provider_name", "aws"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "git_tag_format", "v{version}4"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "repo_base_url_template", "https://somecustom-domain2.com/{namespace}/{module}-{provider}"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "repo_clone_url_template", "ssh://git@some-custom-domain2.com/{namespace}/{module}-{provider}.git"),
					resource.TestCheckResourceAttr("terrareg_module.example2", "repo_browse_url_template", "https://some-custom-domain2.com/{namespace}/{module}-{provider}/tree/{tag}/{path}"),
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
 name = "module-basic-example"
}
data "terrareg_git_provider" "this" {
  name = "Gitlab"
}

resource "terrareg_module" "example" {
  namespace      = terrareg_namespace.this.name
  name           = "basic-example2"
  provider_name  = "awsnew"
  
  git_provider_id = data.terrareg_git_provider.this.id
  git_tag_format  = "v{version}-new"
}
`

const testAccNamespaceResourceConfig_full = `
resource "terrareg_namespace" "this" {
  name = "module-basic-example2"
}
data "terrareg_git_provider" "this" {
  name = "Gitlab"
}

resource "terrareg_module" "example2" {
  namespace      = terrareg_namespace.this.name
  name           = "basic-example3"
  provider_name  = "aws"

  git_provider_id = data.terrareg_git_provider.this.id
  git_tag_format  = "v{version}3"

  repo_base_url_template = "https://somecustom-domain.com/{namespace}/{module}-{provider}"
  repo_clone_url_template = "ssh://git@some-custom-domain.com/{namespace}/{module}-{provider}.git"
  repo_browse_url_template = "https://some-custom-domain.com/{namespace}/{module}-{provider}/tree/{tag}/{path}"
}
`

const testAccNamespaceResourceConfig_full_read = `
resource "terrareg_namespace" "this" {
 name = "module-basic-example2"
}
data "terrareg_git_provider" "this" {
  name = "Gitlab"
}

resource "terrareg_module" "example2" {
  namespace      = terrareg_namespace.this.name
  name           = "basic-example3"
  provider_name  = "aws"

  git_provider_id = data.terrareg_git_provider.this.id
  git_tag_format  = "v{version}4"

  repo_base_url_template = "https://somecustom-domain2.com/{namespace}/{module}-{provider}"
  repo_clone_url_template = "ssh://git@some-custom-domain2.com/{namespace}/{module}-{provider}.git"
  repo_browse_url_template = "https://some-custom-domain2.com/{namespace}/{module}-{provider}/tree/{tag}/{path}"
}
`
