package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/dockstudios/terraform-provider-terrareg/internal/terrareg"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ModuleResource{}
var _ resource.ResourceWithImportState = &ModuleResource{}
var _ resource.ResourceWithModifyPlan = &ModuleResource{}

func NewModuleResource() resource.Resource {
	return &ModuleResource{}
}

// ModuleResource defines the resource implementation.
type ModuleResource struct {
	client *terrareg.TerraregClient
}

// ModuleResourceModel describes the resource data model.
type ModuleResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	Namespace             types.String `tfsdk:"namespace"`
	Name                  types.String `tfsdk:"name"`
	Provider              types.String `tfsdk:"provider_name"`
	GitProviderID         types.Int64  `tfsdk:"git_provider_id"`
	RepoBaseUrlTemplate   types.String `tfsdk:"repo_base_url_template"`
	RepoCloneUrlTemplate  types.String `tfsdk:"repo_clone_url_template"`
	RepoBrowseUrlTemplate types.String `tfsdk:"repo_browse_url_template"`
	ArchiveGitPath        types.Bool   `tfsdk:"archive_git_path"`
	GitTagFormat          types.String `tfsdk:"git_tag_format"`
	GitPath               types.String `tfsdk:"git_path"`
}

func (r *ModuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_module"
}

func (r *ModuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Module resource",

		Attributes: map[string]schema.Attribute{
			// ID attribute required for unit testing
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Full ID of the module",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"namespace": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Namespace of the module",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Module name",
			},
			"provider_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Module provider",
			},
			"git_provider_id": schema.Int64Attribute{
				Optional: true,
				MarkdownDescription: `Id of the Git Repository Provider to use for the module.
Set to ` + "`null`" + `for Custom.
(See https://github.com/MatthewJohn/terrareg/blob/main/docs/USER_GUIDE.md#git-providers)`,
			},
			"repo_base_url_template": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: `This URL must be valid for browsing the base of the repository.
It may include templated values, such as: {namespace}, {module}, {provider}.
E.g. https://github.com/{namespace}/{module}-{provider}
NOTE: Setting this field will override the repository provider configuration.`,
			},
			"repo_clone_url_template": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: `This URL must be valid for cloning the repository.
It may include templated values, such as: {namespace}, {module}, {provider}.
E.g. ssh://git@github.com/{namespace}/{module}-{provider}.git
NOTE: Setting this field will override the repository provider configuration.`,
			},
			"repo_browse_url_template": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: `This URL must be valid for browsing the source code of the repository at a particular tag/path.
									  It may include templated values, such as: {namespace}, {module}, {provider}.
									  It must include the following template values: {tag} and {path}
									  E.g. https://github.com/{namespace}/{module}-{provider}/tree/{tag}/{path}
									  NOTE: Setting this field will override the repository provider configuration.`,
			},
			"archive_git_path": schema.BoolAttribute{
				Optional: true,
				MarkdownDescription: `This determines whether the generated archives only contain the contents of the "Module path".
									  This is only used for providing modules from archives rather than using Git repository redirects.
									  This can be used if the source directory contains other content that you do no wish to distribute to users.
									  Ensure that there are no depdencies on Terraform outside of the "Module path", as this will not be available to users.`,
			},
			"git_tag_format": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `This value will be converted to the expected git tag for a module version.

The {version} placeholder will be used to generated the git tag when translating the module version to a git tag.
For example, using v{version} will translate to a git tag 'v1.1.1' for module version '1.1.1'
If the git tagging format in use does not contain a full semantic version, use placeholders {major}, {minor} and {patch}
to indicate which values are present in the tag - any missing values will be assumed to be '0'.
For example a git tag format of v{major}.{minor} would interpret a tag v1.2 as a module version 1.2.0,
where as a git tag format v{major}.{patch} would generate a version v1.0.2.

Note that if the {version} placeholder is not used, the module version import API must be provided with the git_tag argument and indexing with version argument is disabled.`,
			},
			"git_path": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Set the path within the repository that the module exists. Defaults to the root of the repository.",
			},
		},
	}
}

func (r *ModuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*terrareg.TerraregClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *terrareg.TerraregClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ModuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ModuleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	model := terrareg.ModuleModel{
		GitProviderID:         data.GitProviderID.ValueInt64(),
		RepoBaseUrlTemplate:   data.RepoBaseUrlTemplate.ValueString(),
		RepoCloneUrlTemplate:  data.RepoCloneUrlTemplate.ValueString(),
		RepoBrowseUrlTemplate: data.RepoBrowseUrlTemplate.ValueString(),
		GitTagFormat:          data.GitTagFormat.ValueString(),
		GitPath:               data.GitPath.ValueString(),
		ArchiveGitPath:        data.ArchiveGitPath.ValueBool(),
	}
	if !data.ArchiveGitPath.IsNull() {
		model.ArchiveGitPath = data.ArchiveGitPath.ValueBool()
	}

	id, err := r.client.CreateModule(
		data.Namespace.ValueString(),
		data.Name.ValueString(),
		data.Provider.ValueString(),
		model,
	)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create module, got error: %s", err))
		return
	}

	// Set ID attribute
	data.ID = types.StringValue(id)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ModuleResource) generateId(namespace string, name string, provider string) string {
	return fmt.Sprintf("%s/%s/%s", namespace, name, provider)
}

func (r *ModuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ModuleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use existing ID, if state is not available for namespace, name or provider
	var namespace, name, provider string
	if data.Namespace.IsUnknown() ||
		data.Namespace.IsNull() ||
		data.Name.IsUnknown() ||
		data.Name.IsNull() ||
		data.Provider.IsUnknown() ||
		data.Provider.IsNull() {

		splitId := strings.Split(data.ID.ValueString(), "/")
		if len(splitId) != 3 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("ID is an invalid format: %s", data.ID.ValueString()))
			return
		}
		namespace, name, provider = splitId[0], splitId[1], splitId[2]
	} else {
		namespace = data.Namespace.ValueString()
		name = data.Name.ValueString()
		provider = data.Provider.ValueString()
	}

	module, err := r.client.GetModule(namespace, name, provider)
	// If module was not found, set ID to empty value
	if err == terrareg.ErrNotFound {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read module, got error: %s", err))
		return
	}

	// Update attributes, if they've modified
	if data.Namespace.ValueString() != namespace {
		data.Namespace = types.StringValue(namespace)
	}
	if data.Name.ValueString() != name {
		data.Name = types.StringValue(name)
	}
	if data.Provider.ValueString() != provider {
		data.Provider = types.StringValue(provider)
	}
	if data.GitProviderID.ValueInt64() != module.GitProviderID {
		data.GitProviderID = types.Int64Value(module.GitProviderID)
	}
	if data.RepoBaseUrlTemplate.ValueString() != module.RepoBaseUrlTemplate {
		data.RepoBaseUrlTemplate = types.StringValue(module.RepoBaseUrlTemplate)
	}
	if data.RepoCloneUrlTemplate.ValueString() != module.RepoCloneUrlTemplate {
		data.RepoCloneUrlTemplate = types.StringValue(module.RepoCloneUrlTemplate)
	}
	if data.RepoBrowseUrlTemplate.ValueString() != module.RepoBrowseUrlTemplate {
		data.RepoBrowseUrlTemplate = types.StringValue(module.RepoBrowseUrlTemplate)
	}
	if data.GitTagFormat.ValueString() != module.GitTagFormat {
		data.GitTagFormat = types.StringValue(module.GitTagFormat)
	}
	if data.GitPath.ValueString() != module.GitPath {
		data.GitPath = types.StringValue(module.GitPath)
	}
	if data.ArchiveGitPath.ValueBool() != module.ArchiveGitPath {
		data.ArchiveGitPath = types.BoolValue(module.ArchiveGitPath)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ModuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ModuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ModuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Only provide namespace, name and provider, if one of the attributes
	// has been changed
	var newNamespace string
	var newName string
	var newProvider string
	if !state.Namespace.Equal(plan.Namespace) ||
		!state.Name.Equal(plan.Name) ||
		!state.Provider.Equal(plan.Provider) {

		newNamespace = plan.Namespace.ValueString()
		newName = plan.Name.ValueString()
		newProvider = plan.Provider.ValueString()
	}

	model := &terrareg.ModuleModel{
		GitProviderID:         plan.GitProviderID.ValueInt64(),
		RepoBaseUrlTemplate:   plan.RepoBaseUrlTemplate.ValueString(),
		RepoCloneUrlTemplate:  plan.RepoCloneUrlTemplate.ValueString(),
		RepoBrowseUrlTemplate: plan.RepoBrowseUrlTemplate.ValueString(),
		GitTagFormat:          plan.GitTagFormat.ValueString(),
		GitPath:               plan.GitPath.ValueString(),
	}
	if !plan.ArchiveGitPath.IsNull() {
		model.ArchiveGitPath = plan.ArchiveGitPath.ValueBool()
	}

	_, err := r.client.UpdateModule(
		state.Namespace.ValueString(),
		state.Name.ValueString(),
		state.Provider.ValueString(),
		terrareg.ModuleUpdateModel{
			Namespace:   newNamespace,
			Name:        newName,
			Provider:    newProvider,
			ModuleModel: model,
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update module, got error: %s", err))
		return
	}

	newId := types.StringValue(r.generateId(plan.Namespace.ValueString(), plan.Name.ValueString(), plan.Provider.ValueString()))
	if !plan.ID.Equal(newId) {
		plan.ID = newId
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ModuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ModuleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteModule(state.Namespace.ValueString(), state.Name.ValueString(), state.Provider.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete module, got error: %s", err))
		return
	}
}

func (r *ModuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r ModuleResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan ModuleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	// If unable to obtain plan (generally during a destroy),
	// exit early without erroring
	if diags.HasError() {
		return
	}

	if !plan.Namespace.IsNull() && !plan.Namespace.IsUnknown() && !plan.Name.IsNull() && !plan.Name.IsUnknown() && !plan.Provider.IsNull() && !plan.Provider.IsUnknown() {
		newId := r.generateId(plan.Namespace.ValueString(), plan.Name.ValueString(), plan.Provider.ValueString())

		// If plan value of ID is not unknown and needs to be modified,
		// update it.
		if !plan.ID.IsUnknown() && plan.ID.ValueString() != newId {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("id"), types.StringValue(newId))...)
		}
	}
}
