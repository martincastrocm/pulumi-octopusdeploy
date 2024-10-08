// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class RunbookProcessStepDeployKubernetesSecretActionGitDependencyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the default branch of the repository.
        /// </summary>
        [Input("defaultBranch", required: true)]
        public Input<string> DefaultBranch { get; set; } = null!;

        [Input("filePathFilters")]
        private InputList<string>? _filePathFilters;

        /// <summary>
        /// List of file path filters used to narrow down the directory where files are to be sourced from. Supports glob patten syntax.
        /// </summary>
        public InputList<string> FilePathFilters
        {
            get => _filePathFilters ?? (_filePathFilters = new InputList<string>());
            set => _filePathFilters = value;
        }

        /// <summary>
        /// ID of an existing Git credential.
        /// </summary>
        [Input("gitCredentialId")]
        public Input<string>? GitCredentialId { get; set; }

        /// <summary>
        /// The Git credential authentication type.
        /// </summary>
        [Input("gitCredentialType", required: true)]
        public Input<string> GitCredentialType { get; set; } = null!;

        /// <summary>
        /// The Git URI for the repository where this resource is sourced from.
        /// </summary>
        [Input("repositoryUri", required: true)]
        public Input<string> RepositoryUri { get; set; } = null!;

        public RunbookProcessStepDeployKubernetesSecretActionGitDependencyArgs()
        {
        }
        public static new RunbookProcessStepDeployKubernetesSecretActionGitDependencyArgs Empty => new RunbookProcessStepDeployKubernetesSecretActionGitDependencyArgs();
    }
}
