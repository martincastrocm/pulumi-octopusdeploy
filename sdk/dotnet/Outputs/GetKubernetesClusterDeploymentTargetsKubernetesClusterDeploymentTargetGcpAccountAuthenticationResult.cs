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
    public sealed class GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetGcpAccountAuthenticationResult
    {
        public readonly string AccountId;
        public readonly string ClusterName;
        public readonly bool? ImpersonateServiceAccount;
        public readonly string Project;
        public readonly string? Region;
        public readonly string? ServiceAccountEmails;
        public readonly bool? UseVmServiceAccount;
        public readonly string? Zone;

        [OutputConstructor]
        private GetKubernetesClusterDeploymentTargetsKubernetesClusterDeploymentTargetGcpAccountAuthenticationResult(
            string accountId,

            string clusterName,

            bool? impersonateServiceAccount,

            string project,

            string? region,

            string? serviceAccountEmails,

            bool? useVmServiceAccount,

            string? zone)
        {
            AccountId = accountId;
            ClusterName = clusterName;
            ImpersonateServiceAccount = impersonateServiceAccount;
            Project = project;
            Region = region;
            ServiceAccountEmails = serviceAccountEmails;
            UseVmServiceAccount = useVmServiceAccount;
            Zone = zone;
        }
    }
}
