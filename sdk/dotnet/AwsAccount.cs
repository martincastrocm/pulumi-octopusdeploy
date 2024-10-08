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
    /// This resource manages AWS accounts in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.AwsAccount("example", new()
    ///     {
    ///         AccessKey = "access-key",
    ///         SecretKey = "###########",
    ///     });
    /// 
    ///     // required; get from secure environment/store
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/awsAccount:AwsAccount [options] octopusdeploy_aws_account.&lt;name&gt; &lt;account-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/awsAccount:AwsAccount")]
    public partial class AwsAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access key associated with this AWS account.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// A user-friendly description of this AWS account.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        /// <summary>
        /// The name of this AWS account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The secret key associated with this resource.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

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
        /// Create a AwsAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AwsAccount(string name, AwsAccountArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/awsAccount:AwsAccount", name, args ?? new AwsAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AwsAccount(string name, Input<string> id, AwsAccountState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/awsAccount:AwsAccount", name, state, MakeResourceOptions(options, id))
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
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AwsAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AwsAccount Get(string name, Input<string> id, AwsAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new AwsAccount(name, id, state, options);
        }
    }

    public sealed class AwsAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key associated with this AWS account.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// A user-friendly description of this AWS account.
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

        /// <summary>
        /// The name of this AWS account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secretKey", required: true)]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret key associated with this resource.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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

        public AwsAccountArgs()
        {
        }
        public static new AwsAccountArgs Empty => new AwsAccountArgs();
    }

    public sealed class AwsAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key associated with this AWS account.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// A user-friendly description of this AWS account.
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

        /// <summary>
        /// The name of this AWS account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret key associated with this resource.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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

        public AwsAccountState()
        {
        }
        public static new AwsAccountState Empty => new AwsAccountState();
    }
}
