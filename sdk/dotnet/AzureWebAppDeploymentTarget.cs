// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    /// <summary>
    /// This resource manages Azure web app deployment targets in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = Pulumi.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.AzureWebAppDeploymentTarget("example", new()
    ///     {
    ///         AccountId = "Accounts-123",
    ///         Environments = new[]
    ///         {
    ///             "Environments-123",
    ///         },
    ///         ResourceGroupName = "resource-group-name",
    ///         Roles = new[]
    ///         {
    ///             "Development Team",
    ///             "System Administrators",
    ///         },
    ///         TenantedDeploymentParticipation = "Untenanted",
    ///         WebAppName = "web-app-name",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget [options] octopusdeploy_azure_web_app_deployment_target.&lt;name&gt; &lt;machine-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget")]
    public partial class AzureWebAppDeploymentTarget : global::Pulumi.CustomResource
    {
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.AzureWebAppDeploymentTargetEndpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        [Output("hasLatestCalamari")]
        public Output<bool> HasLatestCalamari { get; private set; } = null!;

        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        [Output("healthStatus")]
        public Output<string> HealthStatus { get; private set; } = null!;

        [Output("isDisabled")]
        public Output<bool> IsDisabled { get; private set; } = null!;

        [Output("isInProcess")]
        public Output<bool> IsInProcess { get; private set; } = null!;

        [Output("machinePolicyId")]
        public Output<string> MachinePolicyId { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("operatingSystem")]
        public Output<string> OperatingSystem { get; private set; } = null!;

        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        [Output("shellName")]
        public Output<string> ShellName { get; private set; } = null!;

        [Output("shellVersion")]
        public Output<string> ShellVersion { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        [Output("statusSummary")]
        public Output<string> StatusSummary { get; private set; } = null!;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        [Output("tenantTags")]
        public Output<ImmutableArray<string>> TenantTags { get; private set; } = null!;

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Output("tenantedDeploymentParticipation")]
        public Output<string> TenantedDeploymentParticipation { get; private set; } = null!;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        [Output("tenants")]
        public Output<ImmutableArray<string>> Tenants { get; private set; } = null!;

        [Output("thumbprint")]
        public Output<string> Thumbprint { get; private set; } = null!;

        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;

        [Output("webAppName")]
        public Output<string> WebAppName { get; private set; } = null!;

        [Output("webAppSlotName")]
        public Output<string?> WebAppSlotName { get; private set; } = null!;


        /// <summary>
        /// Create a AzureWebAppDeploymentTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AzureWebAppDeploymentTarget(string name, AzureWebAppDeploymentTargetArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget", name, args ?? new AzureWebAppDeploymentTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AzureWebAppDeploymentTarget(string name, Input<string> id, AzureWebAppDeploymentTargetState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/azureWebAppDeploymentTarget:AzureWebAppDeploymentTarget", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/OctopusDeploy/terraform-provider-octopusdeploy/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AzureWebAppDeploymentTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AzureWebAppDeploymentTarget Get(string name, Input<string> id, AzureWebAppDeploymentTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new AzureWebAppDeploymentTarget(name, id, state, options);
        }
    }

    public sealed class AzureWebAppDeploymentTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("endpoints")]
        private InputList<Inputs.AzureWebAppDeploymentTargetEndpointArgs>? _endpoints;
        public InputList<Inputs.AzureWebAppDeploymentTargetEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.AzureWebAppDeploymentTargetEndpointArgs>());
            set => _endpoints = value;
        }

        [Input("environments", required: true)]
        private InputList<string>? _environments;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        [Input("healthStatus")]
        public Input<string>? HealthStatus { get; set; }

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        [Input("machinePolicyId")]
        public Input<string>? MachinePolicyId { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("roles", required: true)]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("shellName")]
        public Input<string>? ShellName { get; set; }

        [Input("shellVersion")]
        public Input<string>? ShellVersion { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        [Input("statusSummary")]
        public Input<string>? StatusSummary { get; set; }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation")]
        public Input<string>? TenantedDeploymentParticipation { get; set; }

        [Input("tenants")]
        private InputList<string>? _tenants;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public InputList<string> Tenants
        {
            get => _tenants ?? (_tenants = new InputList<string>());
            set => _tenants = value;
        }

        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        [Input("uri")]
        public Input<string>? Uri { get; set; }

        [Input("webAppName", required: true)]
        public Input<string> WebAppName { get; set; } = null!;

        [Input("webAppSlotName")]
        public Input<string>? WebAppSlotName { get; set; }

        public AzureWebAppDeploymentTargetArgs()
        {
        }
        public static new AzureWebAppDeploymentTargetArgs Empty => new AzureWebAppDeploymentTargetArgs();
    }

    public sealed class AzureWebAppDeploymentTargetState : global::Pulumi.ResourceArgs
    {
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.AzureWebAppDeploymentTargetEndpointGetArgs>? _endpoints;
        public InputList<Inputs.AzureWebAppDeploymentTargetEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.AzureWebAppDeploymentTargetEndpointGetArgs>());
            set => _endpoints = value;
        }

        [Input("environments")]
        private InputList<string>? _environments;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        [Input("hasLatestCalamari")]
        public Input<bool>? HasLatestCalamari { get; set; }

        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        [Input("healthStatus")]
        public Input<string>? HealthStatus { get; set; }

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        [Input("isInProcess")]
        public Input<bool>? IsInProcess { get; set; }

        [Input("machinePolicyId")]
        public Input<string>? MachinePolicyId { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("shellName")]
        public Input<string>? ShellName { get; set; }

        [Input("shellVersion")]
        public Input<string>? ShellVersion { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        [Input("statusSummary")]
        public Input<string>? StatusSummary { get; set; }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation")]
        public Input<string>? TenantedDeploymentParticipation { get; set; }

        [Input("tenants")]
        private InputList<string>? _tenants;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public InputList<string> Tenants
        {
            get => _tenants ?? (_tenants = new InputList<string>());
            set => _tenants = value;
        }

        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        [Input("uri")]
        public Input<string>? Uri { get; set; }

        [Input("webAppName")]
        public Input<string>? WebAppName { get; set; }

        [Input("webAppSlotName")]
        public Input<string>? WebAppSlotName { get; set; }

        public AzureWebAppDeploymentTargetState()
        {
        }
        public static new AzureWebAppDeploymentTargetState Empty => new AzureWebAppDeploymentTargetState();
    }
}
