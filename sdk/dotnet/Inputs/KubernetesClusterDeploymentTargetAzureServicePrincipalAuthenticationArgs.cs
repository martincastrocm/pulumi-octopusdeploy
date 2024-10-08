// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        [Input("clusterResourceGroup", required: true)]
        public Input<string> ClusterResourceGroup { get; set; } = null!;

        public KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationArgs()
        {
        }
        public static new KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationArgs Empty => new KubernetesClusterDeploymentTargetAzureServicePrincipalAuthenticationArgs();
    }
}
