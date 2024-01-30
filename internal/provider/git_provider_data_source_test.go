package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGitProviderDataSource_by_id(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: buildTestProviderConfig(testAccGitProviderDataSourceConfig_by_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.terrareg_git_provider.this", "id", "3"),
					resource.TestCheckResourceAttr("data.terrareg_git_provider.this", "name", "Gitlab"),
				),
			},
		},
	})
}

func TestAccGitProviderDataSource_by_name(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: buildTestProviderConfig(testAccGitProviderDataSourceConfig_by_name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.terrareg_git_provider.this", "id", "2"),
					resource.TestCheckResourceAttr("data.terrareg_git_provider.this", "name", "Bitbucket"),
				),
			},
		},
	})
}

func TestAccGitProviderDataSource_id_and_name(t *testing.T) {
	errorRegex, err := regexp.Compile(".*These attributes cannot be configured together: \\[id,name\\].*")
	if err != nil {
		t.Error(err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config:      buildTestProviderConfig(testAccGitProviderDataSourceConfig_id_and_name),
				ExpectError: errorRegex,
			},
		},
	})
}

func TestAccGitProviderDataSource_no_arguments(t *testing.T) {
	errorRegex, err := regexp.Compile(".*Either 'id' or 'name' must be provided to terrareg_git_provider data source.*")
	if err != nil {
		t.Error(err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config:      buildTestProviderConfig(testAccGitProviderDataSourceConfig_no_arguments),
				ExpectError: errorRegex,
			},
		},
	})
}

const testAccGitProviderDataSourceConfig_by_id = `
data "terrareg_git_provider" "this" {
  id = 3
}
`

const testAccGitProviderDataSourceConfig_by_name = `
data "terrareg_git_provider" "this" {
  name = "Bitbucket"
}
`

const testAccGitProviderDataSourceConfig_id_and_name = `
data "terrareg_git_provider" "this" {
  id   = 3
  name = "Bitbucket"
}
`

const testAccGitProviderDataSourceConfig_no_arguments = `
data "terrareg_git_provider" "this" {
}
`
