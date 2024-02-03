# Changelog

## [1.1.1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/compare/v1.1.0...v1.1.1) (2024-02-03)


### Bug Fixes

* Replace github user from matthewjohn to dockstudios ([6c3436b](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/6c3436bce30e549fd323641c8cb59b6f2d0ff0f2)), closes [#6](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/6)
* Update remaining references to matthewjohn after merge to main ([0e34c6e](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/0e34c6e2f0dc5767c9b97b53d25fe14e430ae5c7))

# [1.1.0](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/compare/v1.0.0...v1.1.0) (2024-02-03)


### Bug Fixes

* Add validation to ensure that either name or id is passed to git_provider data source ([028ef7a](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/028ef7a74fa7b186f49688873f5a6bed69670a9b)), closes [#2](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/2)
* Avoid unknown ID values during plan ([8eff463](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/8eff46327e2cad16d3e47fdb2aecc44a02403f8c)), closes [#2](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/2)
* Correctly mark object as deleted in READ when API returns 404 ([fe5b91e](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/fe5b91eab8f8cbdc5de4bd05bfcc511c29228a65)), closes [#2](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/2)
* Only update attributes during Read if the value has actually changed, to avoid showing plan changes ([fcb74ec](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/fcb74ecdb2109aef2d2cafa2c2a61981a298cb69)), closes [#2](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/2)
* Update namespace ID only in Read operation, using ID for obtaining information from API and updating, if it differs from name ([901a360](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/901a360482123faa243ff5aed87d524c19464808))


### Features

* Add data source for git_provider. Update ID type of git provider to int64 ([b579953](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/b5799535d6cc47f2a92531c95dfe561d5453bb79)), closes [#2](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/2)
* Add module resource ([a54c3c0](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/a54c3c0298d9fd0866dba8e01b649262ddadc2dc)), closes [#2](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/2)
* Add terrareg_git_providers data source for obtaining all git providers ([a34887d](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/a34887d42cbe6dceddd9a5484ed525ceef513cf6)), closes [#2](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/2)

# 1.0.0 (2024-01-29)


### Bug Fixes

* Handle empty body in requests, fixing 'Read' method of namespace ([0b696e1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/0b696e15137f8d5086d1a0faa17c881f3e337697))
* Handle error returned by json.Encode in makeRequest ([1340f22](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/1340f22aaa287909f00f3370ec742fdb4ff69a4a)), closes [#1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/1)
* Remove invalid declaration of err ([2ef7240](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/2ef72403d0c95e138f3cc1742b495e3a3c2aa7e1)), closes [#1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/1)
* Send empty JSON map to DELETE endpoint to fix request ([a5d5b36](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/a5d5b3687a42685d38d27062bc21062b4271daf7)), closes [#1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/1)
* Set ID property during import, fixing error about ID not being set ([78db9c9](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/78db9c95834f018ae48c0d56f61f16e1a7796f59)), closes [#1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/1)
* Update display name attribute in read operation for namespace ([069799e](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/069799e9bd0bf45afb4c03b101ba3666b7b33a9b)), closes [#1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/1)
* Update implementation of JSON encoding to handle empty data ([02e3092](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/02e3092d86571b7eb8949f2fd82648f954896c49)), closes [#1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/1)


### Features

* Initial base layout of provider and functionality to create namespace ([2a499c9](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/commit/2a499c9f1484c2cff41b0a546074c2b9e2c08eec)), closes [#1](https://gitlab.dockstudios.co.uk/pub/terra/terraform-provider-terrareg/issues/1)
