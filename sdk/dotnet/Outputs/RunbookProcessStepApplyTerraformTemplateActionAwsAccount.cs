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
    public sealed class RunbookProcessStepApplyTerraformTemplateActionAwsAccount
    {
        public readonly string? Region;
        /// <summary>
        /// (see below for nested schema)
        /// </summary>
        public readonly Outputs.RunbookProcessStepApplyTerraformTemplateActionAwsAccountRole? Role;
        public readonly bool? UseInstanceRole;
        public readonly string? Variable;

        [OutputConstructor]
        private RunbookProcessStepApplyTerraformTemplateActionAwsAccount(
            string? region,

            Outputs.RunbookProcessStepApplyTerraformTemplateActionAwsAccountRole? role,

            bool? useInstanceRole,

            string? variable)
        {
            Region = region;
            Role = role;
            UseInstanceRole = useInstanceRole;
            Variable = variable;
        }
    }
}
