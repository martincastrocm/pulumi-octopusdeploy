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
    public sealed class PollingTentacleDeploymentTargetEndpointDestination
    {
        public readonly string? DestinationType;
        public readonly string? DropFolderPath;

        [OutputConstructor]
        private PollingTentacleDeploymentTargetEndpointDestination(
            string? destinationType,

            string? dropFolderPath)
        {
            DestinationType = destinationType;
            DropFolderPath = dropFolderPath;
        }
    }
}
