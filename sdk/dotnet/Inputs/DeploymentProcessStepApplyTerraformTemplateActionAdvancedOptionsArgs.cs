// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowAdditionalPluginDownloads")]
        public Input<bool>? AllowAdditionalPluginDownloads { get; set; }

        [Input("applyParameters")]
        public Input<string>? ApplyParameters { get; set; }

        [Input("initParameters")]
        public Input<string>? InitParameters { get; set; }

        [Input("pluginCacheDirectory")]
        public Input<string>? PluginCacheDirectory { get; set; }

        [Input("workspace")]
        public Input<string>? Workspace { get; set; }

        public DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptionsArgs()
        {
        }
        public static new DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptionsArgs Empty => new DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptionsArgs();
    }
}
