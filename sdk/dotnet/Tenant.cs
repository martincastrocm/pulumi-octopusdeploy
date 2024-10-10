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
    /// This resource manages tenants in Octopus Deploy.
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/tenant:Tenant")]
    public partial class Tenant : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the tenant from which this tenant was cloned.
        /// </summary>
        [Output("clonedFromTenantId")]
        public Output<string> ClonedFromTenantId { get; private set; } = null!;

        /// <summary>
        /// The description of this tenant.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this tenant.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        [Output("tenantTags")]
        public Output<ImmutableArray<string>> TenantTags { get; private set; } = null!;


        /// <summary>
        /// Create a Tenant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tenant(string name, TenantArgs? args = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tenant:Tenant", name, args ?? new TenantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Tenant(string name, Input<string> id, TenantState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/tenant:Tenant", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Tenant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Tenant Get(string name, Input<string> id, TenantState? state = null, CustomResourceOptions? options = null)
        {
            return new Tenant(name, id, state, options);
        }
    }

    public sealed class TenantArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the tenant from which this tenant was cloned.
        /// </summary>
        [Input("clonedFromTenantId")]
        public Input<string>? ClonedFromTenantId { get; set; }

        /// <summary>
        /// The description of this tenant.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The space ID associated with this tenant.
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

        public TenantArgs()
        {
        }
        public static new TenantArgs Empty => new TenantArgs();
    }

    public sealed class TenantState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the tenant from which this tenant was cloned.
        /// </summary>
        [Input("clonedFromTenantId")]
        public Input<string>? ClonedFromTenantId { get; set; }

        /// <summary>
        /// The description of this tenant.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The space ID associated with this tenant.
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

        public TenantState()
        {
        }
        public static new TenantState Empty => new TenantState();
    }
}
