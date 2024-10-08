// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class MachinePolicyMachineHealthCheckPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bashHealthCheckPolicy", required: true)]
        public Input<Inputs.MachinePolicyMachineHealthCheckPolicyBashHealthCheckPolicyGetArgs> BashHealthCheckPolicy { get; set; } = null!;

        [Input("healthCheckCron")]
        public Input<string>? HealthCheckCron { get; set; }

        [Input("healthCheckCronTimezone")]
        public Input<string>? HealthCheckCronTimezone { get; set; }

        /// <summary>
        /// In nanoseconds.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        [Input("powershellHealthCheckPolicy", required: true)]
        public Input<Inputs.MachinePolicyMachineHealthCheckPolicyPowershellHealthCheckPolicyGetArgs> PowershellHealthCheckPolicy { get; set; } = null!;

        public MachinePolicyMachineHealthCheckPolicyGetArgs()
        {
        }
        public static new MachinePolicyMachineHealthCheckPolicyGetArgs Empty => new MachinePolicyMachineHealthCheckPolicyGetArgs();
    }
}
