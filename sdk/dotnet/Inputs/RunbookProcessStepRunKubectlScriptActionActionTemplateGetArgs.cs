// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class RunbookProcessStepRunKubectlScriptActionActionTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("communityActionTemplateId")]
        public Input<string>? CommunityActionTemplateId { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("version")]
        public Input<int>? Version { get; set; }

        public RunbookProcessStepRunKubectlScriptActionActionTemplateGetArgs()
        {
        }
        public static new RunbookProcessStepRunKubectlScriptActionActionTemplateGetArgs Empty => new RunbookProcessStepRunKubectlScriptActionActionTemplateGetArgs();
    }
}
