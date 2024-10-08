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
    public sealed class DeploymentProcessStepDeployPackageActionActionTemplate
    {
        public readonly string? CommunityActionTemplateId;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        public readonly int? Version;

        [OutputConstructor]
        private DeploymentProcessStepDeployPackageActionActionTemplate(
            string? communityActionTemplateId,

            string id,

            int? version)
        {
            CommunityActionTemplateId = communityActionTemplateId;
            Id = id;
            Version = version;
        }
    }
}
