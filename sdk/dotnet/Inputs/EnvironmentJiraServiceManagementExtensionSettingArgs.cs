// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class EnvironmentJiraServiceManagementExtensionSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether or not this extension is enabled for this project.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        public EnvironmentJiraServiceManagementExtensionSettingArgs()
        {
        }
        public static new EnvironmentJiraServiceManagementExtensionSettingArgs Empty => new EnvironmentJiraServiceManagementExtensionSettingArgs();
    }
}
