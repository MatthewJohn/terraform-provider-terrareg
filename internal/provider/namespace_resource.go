package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/matthewjohn/terraform-provider-terrareg/internal/terrareg"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NamespaceResource{}
var _ resource.ResourceWithImportState = &NamespaceResource{}

func NewNamespaceResource() resource.Resource {
	return &NamespaceResource{}
}

// NamespaceResource defines the resource implementation.
type NamespaceResource struct {
	client *terrareg.TerraregClient
}

// NamespaceResourceModel describes the resource data model.
type NamespaceResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	DisplayName types.String `tfsdk:"display_name"`
}

func (r *NamespaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_namespace"
}

func (r *NamespaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Namespace resource",

		Attributes: map[string]schema.Attribute{
			// ID attribute required for unit testing
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Namespace name",
				Required:            true,
			},
			"display_name": schema.StringAttribute{
				MarkdownDescription: "User-friendly Namespace display name",
				Optional:            true,
			},
		},
	}
}

func (r *NamespaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*terrareg.TerraregClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NamespaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NamespaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.CreateNamespace(terrareg.NamespaceConfigModel{
		Name:        data.Name.ValueString(),
		DisplayName: data.DisplayName.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create namespace, got error: %s", err))
		return
	}

	// Set ID attribute
	data.ID = data.Name

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NamespaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NamespaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	namespace, err := r.client.GetNamespace(data.ID.ValueString())
	// If namespace was not found, set Name to empty value
	if err == terrareg.ErrNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read namespace, got error: %s", err))
		return
	}

	// Update attributes
	if data.DisplayName.ValueString() != namespace.DisplayName {
		data.DisplayName = types.StringValue(namespace.DisplayName)
	}

	// Update ID, if it does not match
	if data.ID.IsUnknown() || data.ID.ValueString() != data.Name.ValueString() {
		data.ID = data.Name
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NamespaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NamespaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get old namespace name
	var name types.String
	diags := req.State.GetAttribute(ctx, path.Root("name"), &name)
	resp.Diagnostics.Append(diags...)

	err := r.client.UpdateNamespace(
		name.ValueString(),
		terrareg.NamespaceConfigModel{
			Name:        data.Name.ValueString(),
			DisplayName: data.DisplayName.ValueString(),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update namespace, got error: %s", err))
		return
	}

	if data.ID.IsUnknown() || data.ID.ValueString() != data.Name.ValueString() {
		data.ID = data.Name
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NamespaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NamespaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteNamespace(data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete namespace, got error: %s", err))
		return
	}
}

func (r *NamespaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
