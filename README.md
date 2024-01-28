# Terrareg Terraform Provider

Terraform provider for configuring [Terrareg](https://github.com/matthewjohn/terrareg)

## Developing

### Build
```
go build
```

### Local testing

After building, copy the provider binary into a directory (e.g. ~/.terraform.d/plugins/).

Setup `~/.terraformrc` with the following block:
```
provider_installation {
  dev_overrides {
    "matthewjohn/terrareg" = "/home/<username>/.terraform.d/plugins/"
  }
}
```

# License

This project and all associated code is covered by GNU General Public License v3.0.
