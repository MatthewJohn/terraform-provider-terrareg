package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/matthewjohn/terraform-provider-terrareg/internal/terrareg"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GitProvidersDataSource{}

func NewGitProvidersDataSource() datasource.DataSource {
	return &GitProvidersDataSource{}
}

// GitProvidersDataSource defines the data source implementation.
type GitProvidersDataSource struct {
	client *terrareg.TerraregClient
}

// GitProvidersDataSourceModel describes the data source data model.
type GitProvidersDataSourceModel struct {
	Id           types.String                `tfsdk:"id"`
	GitProviders []terrareg.GitProviderModel `tfsdk:"git_providers"`
}

func (d *GitProvidersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_git_providers"
}

func (d *GitProvidersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Data source for obtaining all git providers",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Internal ID",
			},
			"git_providers": schema.ListAttribute{
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":   types.Int64Type,
						"name": types.StringType,
					},
				},
				MarkdownDescription: "List of Git Providers, including id and name",
				Computed:            true,
			},
		},
	}
}

func (d *GitProvidersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*terrareg.TerraregClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *terrareg.TerraregClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *GitProvidersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GitProvidersDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	gitProviders, err := d.client.GetGitProviders()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
		return
	}

	data.GitProviders = gitProviders

	// Create fake ID, required by Terraform
	data.Id = types.StringValue("this")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
