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
    public sealed class RunbookProcessStepApplyTerraformTemplateActionTemplate
    {
        public readonly string? AdditionalVariableFiles;
        public readonly string? Directory;
        public readonly bool? RunAutomaticFileSubstitution;
        public readonly string? TargetFiles;

        [OutputConstructor]
        private RunbookProcessStepApplyTerraformTemplateActionTemplate(
            string? additionalVariableFiles,

            string? directory,

            bool? runAutomaticFileSubstitution,

            string? targetFiles)
        {
            AdditionalVariableFiles = additionalVariableFiles;
            Directory = directory;
            RunAutomaticFileSubstitution = runAutomaticFileSubstitution;
            TargetFiles = targetFiles;
        }
    }
}
