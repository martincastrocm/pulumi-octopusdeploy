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
    /// This resource manages channels in Octopus Deploy.
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
    ///     var example = new Octopusdeploy.Channel("example", new()
    ///     {
    ///         ProjectId = "Projects-123",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/channel:Channel [options] octopusdeploy_channel.&lt;name&gt; &lt;channel-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/channel:Channel")]
    public partial class Channel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this channel.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates if this is the default channel for the associated project.
        /// </summary>
        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The lifecycle ID associated with this channel.
        /// </summary>
        [Output("lifecycleId")]
        public Output<string?> LifecycleId { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project ID associated with this channel.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// A list of rules associated with this channel.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.ChannelRule>> Rules { get; private set; } = null!;

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
        /// Create a Channel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Channel(string name, ChannelArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/channel:Channel", name, args ?? new ChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Channel(string name, Input<string> id, ChannelState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/channel:Channel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Channel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Channel Get(string name, Input<string> id, ChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new Channel(name, id, state, options);
        }
    }

    public sealed class ChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this channel.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates if this is the default channel for the associated project.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The lifecycle ID associated with this channel.
        /// </summary>
        [Input("lifecycleId")]
        public Input<string>? LifecycleId { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project ID associated with this channel.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.ChannelRuleArgs>? _rules;

        /// <summary>
        /// A list of rules associated with this channel.
        /// </summary>
        public InputList<Inputs.ChannelRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ChannelRuleArgs>());
            set => _rules = value;
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

        public ChannelArgs()
        {
        }
        public static new ChannelArgs Empty => new ChannelArgs();
    }

    public sealed class ChannelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this channel.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates if this is the default channel for the associated project.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The lifecycle ID associated with this channel.
        /// </summary>
        [Input("lifecycleId")]
        public Input<string>? LifecycleId { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project ID associated with this channel.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("rules")]
        private InputList<Inputs.ChannelRuleGetArgs>? _rules;

        /// <summary>
        /// A list of rules associated with this channel.
        /// </summary>
        public InputList<Inputs.ChannelRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ChannelRuleGetArgs>());
            set => _rules = value;
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

        public ChannelState()
        {
        }
        public static new ChannelState Empty => new ChannelState();
    }
}
