// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class RunbookProcessStepDeployWindowsServiceActionPackage
    {
        /// <summary>
        /// Whether to acquire this package on the server ('Server'), target ('ExecutionTarget') or not at all ('NotAcquired'). Can be an expression
        /// </summary>
        public readonly string? AcquisitionLocation;
        /// <summary>
        /// The feed ID associated with this package reference.
        /// </summary>
        public readonly string? FeedId;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID of the package.
        /// </summary>
        public readonly string PackageId;
        /// <summary>
        /// A list of properties associated with this package.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;

        [OutputConstructor]
        private RunbookProcessStepDeployWindowsServiceActionPackage(
            string? acquisitionLocation,

            string? feedId,

            string? id,

            string? name,

            string packageId,

            ImmutableDictionary<string, string>? properties)
        {
            AcquisitionLocation = acquisitionLocation;
            FeedId = feedId;
            Id = id;
            Name = name;
            PackageId = packageId;
            Properties = properties;
        }
    }
}
