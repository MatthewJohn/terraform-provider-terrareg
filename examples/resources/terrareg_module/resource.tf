resource "terrareg_namespace" "this" {
  name = "example-namespace"
}
data "terrareg_git_provider" "this" {
  name = "Gitlab"
}

resource "terrareg_module" "example" {
  namespace      = terrareg_namespace.this.name
  name           = "example"
  provider_name  = "aws"

  git_provider_id = data.terrareg_git_provider.this.id
  git_tag_format  = "v{version}"
}


