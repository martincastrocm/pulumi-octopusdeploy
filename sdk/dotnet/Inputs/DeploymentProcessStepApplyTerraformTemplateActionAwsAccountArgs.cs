// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepApplyTerraformTemplateActionAwsAccountArgs : global::Pulumi.ResourceArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// (see below for nested schema)
        /// </summary>
        [Input("role")]
        public Input<Inputs.DeploymentProcessStepApplyTerraformTemplateActionAwsAccountRoleArgs>? Role { get; set; }

        [Input("useInstanceRole")]
        public Input<bool>? UseInstanceRole { get; set; }

        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public DeploymentProcessStepApplyTerraformTemplateActionAwsAccountArgs()
        {
        }
        public static new DeploymentProcessStepApplyTerraformTemplateActionAwsAccountArgs Empty => new DeploymentProcessStepApplyTerraformTemplateActionAwsAccountArgs();
    }
}
