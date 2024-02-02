package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/dockstudios/terraform-provider-terrareg/internal/terrareg"
)

var _ provider.Provider = &TerraregProvider{}

type TerraregProvider struct {
	TerraregProviderModel
	version string
}

// TerraregProviderModel describes the provider data model.
type TerraregProviderModel struct {
	Url    types.String `tfsdk:"url"`
	ApiKey types.String `tfsdk:"api_key"`
}

func (p *TerraregProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "terrareg"
	resp.Version = p.version
}

func (p *TerraregProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				MarkdownDescription: "Terrareg url (e.g. https://terrareg.example.com)",
				Required:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API Key for authenticating to Terrareg (currently supports admin auth token)",
				Optional:            true,
			},
		},
	}
}

func (p *TerraregProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data TerraregProviderModel

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Url.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Unknown Terrareg url",
			"The provider cannot create the Terrareg API client as there is an unknown configuration value for the Terrareg URL. "+
				"Either target apply the source of the value first or set the value statically in the configuration.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	url := data.Url.ValueString()
	if url == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Missing Url.",
			"The provider must be configured with a URL. "+
				"Set the host value in the configuration. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Terrareg client")

	api, err := terrareg.NewClient(url, data.ApiKey.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Terrareg API Client",
			"An unexpected error occurred when creating the Terrareg API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Terrareg Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = api
	resp.ResourceData = api
}

func (p *TerraregProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNamespaceResource,
	}
}

func (p *TerraregProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// NewExampleDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TerraregProvider{
			version: version,
		}
	}
}
