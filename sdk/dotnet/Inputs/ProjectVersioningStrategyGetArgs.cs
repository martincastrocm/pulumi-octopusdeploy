// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class ProjectVersioningStrategyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("donorPackageStepId")]
        public Input<string>? DonorPackageStepId { get; set; }

        [Input("donorPackages")]
        private InputList<Inputs.ProjectVersioningStrategyDonorPackageGetArgs>? _donorPackages;
        public InputList<Inputs.ProjectVersioningStrategyDonorPackageGetArgs> DonorPackages
        {
            get => _donorPackages ?? (_donorPackages = new InputList<Inputs.ProjectVersioningStrategyDonorPackageGetArgs>());
            set => _donorPackages = value;
        }

        [Input("template")]
        public Input<string>? Template { get; set; }

        public ProjectVersioningStrategyGetArgs()
        {
        }
        public static new ProjectVersioningStrategyGetArgs Empty => new ProjectVersioningStrategyGetArgs();
    }
}
