// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Impersonate service accounts
        /// </summary>
        [Input("impersonateServiceAccount")]
        public Input<bool>? ImpersonateServiceAccount { get; set; }

        /// <summary>
        /// This sets GOOGLE_PROJECT environment variable
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// This sets GOOGLE_REGION environment variable
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// This sets GOOGLE*IMPERSONATE*SERVICE_ACCOUNT environment variable
        /// </summary>
        [Input("serviceAccountEmails")]
        public Input<string>? ServiceAccountEmails { get; set; }

        /// <summary>
        /// When running in a Compute Engine virtual machine, use the associated VM service account
        /// </summary>
        [Input("useVmServiceAccount")]
        public Input<bool>? UseVmServiceAccount { get; set; }

        [Input("variable")]
        public Input<string>? Variable { get; set; }

        /// <summary>
        /// This sets GOOGLE_ZONE environment variable
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccountArgs()
        {
        }
        public static new DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccountArgs Empty => new DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccountArgs();
    }
}
