GO_FILES = $(shell find internal -type f -name '*.go' ! -name '*_test.go')

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m


terraform-provider-terrareg: $(GO_FILES)
	go build

~/.terraform.d/plugins/terraform-provider-terrareg: terraform-provider-terrareg
	cp terraform-provider-terrareg ~/.terraform.d/plugins/

dev: ~/.terraform.d/plugins/terraform-provider-terrareg

