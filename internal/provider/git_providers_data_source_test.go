package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGitProvidersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: buildTestProviderConfig(testAccGitProvidersDataSourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.terrareg_git_providers.this", "git_providers.0.id", "1"),
					resource.TestCheckResourceAttr("data.terrareg_git_providers.this", "git_providers.0.name", "Github"),
					resource.TestCheckResourceAttr("data.terrareg_git_providers.this", "git_providers.1.id", "2"),
					resource.TestCheckResourceAttr("data.terrareg_git_providers.this", "git_providers.1.name", "Bitbucket"),
					resource.TestCheckResourceAttr("data.terrareg_git_providers.this", "git_providers.2.id", "3"),
					resource.TestCheckResourceAttr("data.terrareg_git_providers.this", "git_providers.2.name", "Gitlab"),
				),
			},
		},
	})
}

const testAccGitProvidersDataSourceConfig = `
data "terrareg_git_providers" "this" { }
`
