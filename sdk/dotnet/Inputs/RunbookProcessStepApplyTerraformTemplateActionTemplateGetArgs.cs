// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class RunbookProcessStepApplyTerraformTemplateActionTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalVariableFiles")]
        public Input<string>? AdditionalVariableFiles { get; set; }

        [Input("directory")]
        public Input<string>? Directory { get; set; }

        [Input("runAutomaticFileSubstitution")]
        public Input<bool>? RunAutomaticFileSubstitution { get; set; }

        [Input("targetFiles")]
        public Input<string>? TargetFiles { get; set; }

        public RunbookProcessStepApplyTerraformTemplateActionTemplateGetArgs()
        {
        }
        public static new RunbookProcessStepApplyTerraformTemplateActionTemplateGetArgs Empty => new RunbookProcessStepApplyTerraformTemplateActionTemplateGetArgs();
    }
}
