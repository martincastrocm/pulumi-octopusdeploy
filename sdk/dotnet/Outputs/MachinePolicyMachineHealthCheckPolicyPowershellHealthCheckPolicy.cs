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
    public sealed class MachinePolicyMachineHealthCheckPolicyPowershellHealthCheckPolicy
    {
        public readonly string? RunType;
        public readonly string? ScriptBody;

        [OutputConstructor]
        private MachinePolicyMachineHealthCheckPolicyPowershellHealthCheckPolicy(
            string? runType,

            string? scriptBody)
        {
            RunType = runType;
            ScriptBody = scriptBody;
        }
    }
}
