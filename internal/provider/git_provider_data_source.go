package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/dockstudios/terraform-provider-terrareg/internal/terrareg"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GitProviderDataSource{}

func NewGitProviderDataSource() datasource.DataSource {
	return &GitProviderDataSource{}
}

// GitProviderDataSource defines the data source implementation.
type GitProviderDataSource struct {
	client *terrareg.TerraregClient
}

// GitProviderDataSourceModel describes the data source data model.
type GitProviderDataSourceModel struct {
	Id   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *GitProviderDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_git_provider"
}

func (d *GitProviderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Data source for obtaining git provider",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				// Mark as computed, but also Optional, to allow the user to define the value
				Computed:            true,
				MarkdownDescription: "Internal ID",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				// Mark as computed, but also Optional, to allow the user to define the value
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Internal ID",
			},
		},
	}
}

func (d GitProviderDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.Conflicting(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *GitProviderDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *GitProviderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GitProviderDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() && data.Name.IsNull() {
		resp.Diagnostics.AddError("Client Error", "Either 'id' or 'name' must be provided to terrareg_git_provider data source.")
		return
	}

	gitProviders, err := d.client.GetGitProviders()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
		return
	}

	foundMatch := false
	for _, v := range gitProviders {
		if !data.Name.IsNull() {
			if data.Name.Equal(types.StringValue(v.Name)) {
				data.Id = types.Int64Value(v.ID)
				foundMatch = true
				break
			}
		}
		if !data.Id.IsNull() {
			if data.Id.Equal(types.Int64Value(v.ID)) {
				data.Name = types.StringValue(v.Name)
				foundMatch = true
				break
			}
		}
	}
	if !foundMatch {
		resp.Diagnostics.AddError("Client Error", "Unable to find git provider with matching details")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
