// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class ExternalFeedCreateReleaseTriggerPackageArgs : global::Pulumi.ResourceArgs
    {
        [Input("deploymentActionSlug", required: true)]
        public Input<string> DeploymentActionSlug { get; set; } = null!;

        [Input("packageReference", required: true)]
        public Input<string> PackageReference { get; set; } = null!;

        public ExternalFeedCreateReleaseTriggerPackageArgs()
        {
        }
        public static new ExternalFeedCreateReleaseTriggerPackageArgs Empty => new ExternalFeedCreateReleaseTriggerPackageArgs();
    }
}
