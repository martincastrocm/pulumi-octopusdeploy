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
    public sealed class RunbookProcessStepApplyTerraformTemplateActionAdvancedOptions
    {
        public readonly bool? AllowAdditionalPluginDownloads;
        public readonly string? ApplyParameters;
        public readonly string? InitParameters;
        public readonly string? PluginCacheDirectory;
        public readonly string? Workspace;

        [OutputConstructor]
        private RunbookProcessStepApplyTerraformTemplateActionAdvancedOptions(
            bool? allowAdditionalPluginDownloads,

            string? applyParameters,

            string? initParameters,

            string? pluginCacheDirectory,

            string? workspace)
        {
            AllowAdditionalPluginDownloads = allowAdditionalPluginDownloads;
            ApplyParameters = applyParameters;
            InitParameters = initParameters;
            PluginCacheDirectory = pluginCacheDirectory;
            Workspace = workspace;
        }
    }
}
