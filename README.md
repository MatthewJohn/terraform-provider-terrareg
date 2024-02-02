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
    "dockstudios/terrareg" = "/home/<username>/.terraform.d/plugins/"
  }
}
```

### Running tests

Run tests using:
```
go test $(go list ./...) -count=1 -v
```

**Currently, no tests support this method**

To run acceptance tests, run an instance of terrareg (https://github.com/matthewjohn/terrareg) and run acceptance tests:
```
docker run -d -p 5000:5000 -e MIGRATE_DATABASE=true -e ADMIN_AUTHENTICATION_TOKEN=password

TF_ACC=1 TERRAREG_URL=http://localhost:5000 TERRAREG_API_KEY=password go test $(go list ./...) -count=1 -v
```

# License

This project and all associated code is covered by GNU General Public License v3.0.
