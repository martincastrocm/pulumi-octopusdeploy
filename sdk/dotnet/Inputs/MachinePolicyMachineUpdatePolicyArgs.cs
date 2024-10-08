// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class MachinePolicyMachineUpdatePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The behaviour of how Calamari is updated. Valid values are `UpdateAlways`, `UpdateOnDeployment` and `UpdateOnNewMachine`.
        /// </summary>
        [Input("calamariUpdateBehavior")]
        public Input<string>? CalamariUpdateBehavior { get; set; }

        /// <summary>
        /// The behaviour of how Kubernetes agent machines are updated. Valid values are `NeverUpdate` and `Update`.
        /// </summary>
        [Input("kubernetesAgentUpdateBehavior")]
        public Input<string>? KubernetesAgentUpdateBehavior { get; set; }

        /// <summary>
        /// The Account ID to perform any Tentacle updates under.
        /// </summary>
        [Input("tentacleUpdateAccountId")]
        public Input<string>? TentacleUpdateAccountId { get; set; }

        /// <summary>
        /// The behaviour of how Tentacle machines are updated. Valid values are `NeverUpdate` and `Update`.
        /// </summary>
        [Input("tentacleUpdateBehavior")]
        public Input<string>? TentacleUpdateBehavior { get; set; }

        public MachinePolicyMachineUpdatePolicyArgs()
        {
        }
        public static new MachinePolicyMachineUpdatePolicyArgs Empty => new MachinePolicyMachineUpdatePolicyArgs();
    }
}
