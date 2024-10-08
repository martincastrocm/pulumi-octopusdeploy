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
    /// This resource manages GCP accounts in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.GcpAccount("example", new()
    ///     {
    ///         JsonKey = "json-key",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/gcpAccount:GcpAccount [options] octopusdeploy_gcp_account.&lt;name&gt; &lt;account-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/gcpAccount:GcpAccount")]
    public partial class GcpAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A user-friendly description of this GCP account.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        /// <summary>
        /// The JSON key associated with this GCP account.
        /// </summary>
        [Output("jsonKey")]
        public Output<string> JsonKey { get; private set; } = null!;

        /// <summary>
        /// The name of this GCP account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

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
        /// Create a GcpAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GcpAccount(string name, GcpAccountArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/gcpAccount:GcpAccount", name, args ?? new GcpAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GcpAccount(string name, Input<string> id, GcpAccountState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/gcpAccount:GcpAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/OctopusDeploy/terraform-provider-octopusdeploy/releases",
                AdditionalSecretOutputs =
                {
                    "jsonKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GcpAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GcpAccount Get(string name, Input<string> id, GcpAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new GcpAccount(name, id, state, options);
        }
    }

    public sealed class GcpAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-friendly description of this GCP account.
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

        [Input("jsonKey", required: true)]
        private Input<string>? _jsonKey;

        /// <summary>
        /// The JSON key associated with this GCP account.
        /// </summary>
        public Input<string>? JsonKey
        {
            get => _jsonKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _jsonKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of this GCP account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

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

        public GcpAccountArgs()
        {
        }
        public static new GcpAccountArgs Empty => new GcpAccountArgs();
    }

    public sealed class GcpAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-friendly description of this GCP account.
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

        [Input("jsonKey")]
        private Input<string>? _jsonKey;

        /// <summary>
        /// The JSON key associated with this GCP account.
        /// </summary>
        public Input<string>? JsonKey
        {
            get => _jsonKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _jsonKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of this GCP account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

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

        public GcpAccountState()
        {
        }
        public static new GcpAccountState Empty => new GcpAccountState();
    }
}
