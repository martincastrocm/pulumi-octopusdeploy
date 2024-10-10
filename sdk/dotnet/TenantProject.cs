// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    [OctopusdeployResourceType("octopusdeploy:index/tenantProject:TenantProject")]
    public partial class TenantProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The environment IDs associated with this tenant.
        /// </summary>
        [Output("environmentIds")]
        public Output<ImmutableArray<string>> EnvironmentIds { get; private set; } = null!;

        /// <summary>
        /// The project ID associated with this tenant.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this project tenant.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// The tenant ID associated with this tenant.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a TenantProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TenantProject(string name, TenantProjectArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tenantProject:TenantProject", name, args ?? new TenantProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TenantProject(string name, Input<string> id, TenantProjectState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tenantProject:TenantProject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TenantProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TenantProject Get(string name, Input<string> id, TenantProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new TenantProject(name, id, state, options);
        }
    }

    public sealed class TenantProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("environmentIds")]
        private InputList<string>? _environmentIds;

        /// <summary>
        /// The environment IDs associated with this tenant.
        /// </summary>
        public InputList<string> EnvironmentIds
        {
            get => _environmentIds ?? (_environmentIds = new InputList<string>());
            set => _environmentIds = value;
        }

        /// <summary>
        /// The project ID associated with this tenant.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The space ID associated with this project tenant.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The tenant ID associated with this tenant.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public TenantProjectArgs()
        {
        }
        public static new TenantProjectArgs Empty => new TenantProjectArgs();
    }

    public sealed class TenantProjectState : global::Pulumi.ResourceArgs
    {
        [Input("environmentIds")]
        private InputList<string>? _environmentIds;

        /// <summary>
        /// The environment IDs associated with this tenant.
        /// </summary>
        public InputList<string> EnvironmentIds
        {
            get => _environmentIds ?? (_environmentIds = new InputList<string>());
            set => _environmentIds = value;
        }

        /// <summary>
        /// The project ID associated with this tenant.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The space ID associated with this project tenant.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The tenant ID associated with this tenant.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public TenantProjectState()
        {
        }
        public static new TenantProjectState Empty => new TenantProjectState();
    }
}
