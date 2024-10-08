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
    /// This resource manages Azure OpenID Connect accounts in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.AzureOpenidConnect("example", new()
    ///     {
    ///         AccountTestSubjectKeys = new[]
    ///         {
    ///             "space",
    ///             "type",
    ///         },
    ///         ApplicationId = "00000000-0000-0000-0000-000000000000",
    ///         Audience = "api://AzureADTokenExchange",
    ///         ExecutionSubjectKeys = new[]
    ///         {
    ///             "space",
    ///             "project",
    ///         },
    ///         HealthSubjectKeys = new[]
    ///         {
    ///             "space",
    ///             "target",
    ///             "type",
    ///         },
    ///         SubscriptionId = "00000000-0000-0000-0000-000000000000",
    ///         TenantId = "00000000-0000-0000-0000-000000000000",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/azureOpenidConnect:AzureOpenidConnect [options] octopusdeploy_azure_openid_connect.&lt;name&gt; &lt;account-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/azureOpenidConnect:AzureOpenidConnect")]
    public partial class AzureOpenidConnect : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Keys to include in an account test. Valid options are: `space`, `account`, `type`
        /// </summary>
        [Output("accountTestSubjectKeys")]
        public Output<ImmutableArray<string>> AccountTestSubjectKeys { get; private set; } = null!;

        /// <summary>
        /// The application ID of this resource.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Federated credentials audience, this value is used to establish a connection between external workload identities and Microsoft Entra ID.
        /// </summary>
        [Output("audience")]
        public Output<string?> Audience { get; private set; } = null!;

        /// <summary>
        /// The authentication endpoint URI for this resource.
        /// </summary>
        [Output("authenticationEndpoint")]
        public Output<string?> AuthenticationEndpoint { get; private set; } = null!;

        /// <summary>
        /// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
        /// </summary>
        [Output("azureEnvironment")]
        public Output<string> AzureEnvironment { get; private set; } = null!;

        /// <summary>
        /// The description of this Azure OpenID Connect account.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        /// <summary>
        /// Keys to include in a deployment or runbook. Valid options are `space`, `environment`, `project`, `tenant`, `runbook`, `account`, `type`
        /// </summary>
        [Output("executionSubjectKeys")]
        public Output<ImmutableArray<string>> ExecutionSubjectKeys { get; private set; } = null!;

        /// <summary>
        /// Keys to include in a health check. Valid options are `space`, `account`, `target`, `type`
        /// </summary>
        [Output("healthSubjectKeys")]
        public Output<ImmutableArray<string>> HealthSubjectKeys { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource manager endpoint URI for this resource.
        /// </summary>
        [Output("resourceManagerEndpoint")]
        public Output<string?> ResourceManagerEndpoint { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// The subscription ID of this resource.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The tenant ID of this resource.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

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


        /// <summary>
        /// Create a AzureOpenidConnect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AzureOpenidConnect(string name, AzureOpenidConnectArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/azureOpenidConnect:AzureOpenidConnect", name, args ?? new AzureOpenidConnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AzureOpenidConnect(string name, Input<string> id, AzureOpenidConnectState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/azureOpenidConnect:AzureOpenidConnect", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AzureOpenidConnect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AzureOpenidConnect Get(string name, Input<string> id, AzureOpenidConnectState? state = null, CustomResourceOptions? options = null)
        {
            return new AzureOpenidConnect(name, id, state, options);
        }
    }

    public sealed class AzureOpenidConnectArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountTestSubjectKeys")]
        private InputList<string>? _accountTestSubjectKeys;

        /// <summary>
        /// Keys to include in an account test. Valid options are: `space`, `account`, `type`
        /// </summary>
        public InputList<string> AccountTestSubjectKeys
        {
            get => _accountTestSubjectKeys ?? (_accountTestSubjectKeys = new InputList<string>());
            set => _accountTestSubjectKeys = value;
        }

        /// <summary>
        /// The application ID of this resource.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Federated credentials audience, this value is used to establish a connection between external workload identities and Microsoft Entra ID.
        /// </summary>
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        /// <summary>
        /// The authentication endpoint URI for this resource.
        /// </summary>
        [Input("authenticationEndpoint")]
        public Input<string>? AuthenticationEndpoint { get; set; }

        /// <summary>
        /// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
        /// </summary>
        [Input("azureEnvironment")]
        public Input<string>? AzureEnvironment { get; set; }

        /// <summary>
        /// The description of this Azure OpenID Connect account.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        [Input("executionSubjectKeys")]
        private InputList<string>? _executionSubjectKeys;

        /// <summary>
        /// Keys to include in a deployment or runbook. Valid options are `space`, `environment`, `project`, `tenant`, `runbook`, `account`, `type`
        /// </summary>
        public InputList<string> ExecutionSubjectKeys
        {
            get => _executionSubjectKeys ?? (_executionSubjectKeys = new InputList<string>());
            set => _executionSubjectKeys = value;
        }

        [Input("healthSubjectKeys")]
        private InputList<string>? _healthSubjectKeys;

        /// <summary>
        /// Keys to include in a health check. Valid options are `space`, `account`, `target`, `type`
        /// </summary>
        public InputList<string> HealthSubjectKeys
        {
            get => _healthSubjectKeys ?? (_healthSubjectKeys = new InputList<string>());
            set => _healthSubjectKeys = value;
        }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource manager endpoint URI for this resource.
        /// </summary>
        [Input("resourceManagerEndpoint")]
        public Input<string>? ResourceManagerEndpoint { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The subscription ID of this resource.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        /// <summary>
        /// The tenant ID of this resource.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

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

        public AzureOpenidConnectArgs()
        {
        }
        public static new AzureOpenidConnectArgs Empty => new AzureOpenidConnectArgs();
    }

    public sealed class AzureOpenidConnectState : global::Pulumi.ResourceArgs
    {
        [Input("accountTestSubjectKeys")]
        private InputList<string>? _accountTestSubjectKeys;

        /// <summary>
        /// Keys to include in an account test. Valid options are: `space`, `account`, `type`
        /// </summary>
        public InputList<string> AccountTestSubjectKeys
        {
            get => _accountTestSubjectKeys ?? (_accountTestSubjectKeys = new InputList<string>());
            set => _accountTestSubjectKeys = value;
        }

        /// <summary>
        /// The application ID of this resource.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Federated credentials audience, this value is used to establish a connection between external workload identities and Microsoft Entra ID.
        /// </summary>
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        /// <summary>
        /// The authentication endpoint URI for this resource.
        /// </summary>
        [Input("authenticationEndpoint")]
        public Input<string>? AuthenticationEndpoint { get; set; }

        /// <summary>
        /// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
        /// </summary>
        [Input("azureEnvironment")]
        public Input<string>? AzureEnvironment { get; set; }

        /// <summary>
        /// The description of this Azure OpenID Connect account.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        [Input("executionSubjectKeys")]
        private InputList<string>? _executionSubjectKeys;

        /// <summary>
        /// Keys to include in a deployment or runbook. Valid options are `space`, `environment`, `project`, `tenant`, `runbook`, `account`, `type`
        /// </summary>
        public InputList<string> ExecutionSubjectKeys
        {
            get => _executionSubjectKeys ?? (_executionSubjectKeys = new InputList<string>());
            set => _executionSubjectKeys = value;
        }

        [Input("healthSubjectKeys")]
        private InputList<string>? _healthSubjectKeys;

        /// <summary>
        /// Keys to include in a health check. Valid options are `space`, `account`, `target`, `type`
        /// </summary>
        public InputList<string> HealthSubjectKeys
        {
            get => _healthSubjectKeys ?? (_healthSubjectKeys = new InputList<string>());
            set => _healthSubjectKeys = value;
        }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource manager endpoint URI for this resource.
        /// </summary>
        [Input("resourceManagerEndpoint")]
        public Input<string>? ResourceManagerEndpoint { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The subscription ID of this resource.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// The tenant ID of this resource.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public AzureOpenidConnectState()
        {
        }
        public static new AzureOpenidConnectState Empty => new AzureOpenidConnectState();
    }
}
