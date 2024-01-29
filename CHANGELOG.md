# Changelog

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
