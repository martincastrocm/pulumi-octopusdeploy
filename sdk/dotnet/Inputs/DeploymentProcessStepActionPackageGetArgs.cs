// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepActionPackageGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to acquire this package on the server ('Server'), target ('ExecutionTarget') or not at all ('NotAcquired'). Can be an expression
        /// </summary>
        [Input("acquisitionLocation")]
        public Input<string>? AcquisitionLocation { get; set; }

        /// <summary>
        /// Whether to extract the package during deployment
        /// </summary>
        [Input("extractDuringDeployment")]
        public Input<bool>? ExtractDuringDeployment { get; set; }

        /// <summary>
        /// The feed ID associated with this package reference.
        /// </summary>
        [Input("feedId")]
        public Input<string>? FeedId { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the package.
        /// </summary>
        [Input("packageId", required: true)]
        public Input<string> PackageId { get; set; } = null!;

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// A list of properties associated with this package.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        public DeploymentProcessStepActionPackageGetArgs()
        {
        }
        public static new DeploymentProcessStepActionPackageGetArgs Empty => new DeploymentProcessStepActionPackageGetArgs();
    }
}
