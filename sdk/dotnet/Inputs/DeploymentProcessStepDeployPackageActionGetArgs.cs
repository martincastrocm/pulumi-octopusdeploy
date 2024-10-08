// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class DeploymentProcessStepDeployPackageActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents the template that is associated with this action.
        /// </summary>
        [Input("actionTemplate")]
        public Input<Inputs.DeploymentProcessStepDeployPackageActionActionTemplateGetArgs>? ActionTemplate { get; set; }

        [Input("canBeUsedForProjectVersioning")]
        public Input<bool>? CanBeUsedForProjectVersioning { get; set; }

        [Input("channels")]
        private InputList<string>? _channels;

        /// <summary>
        /// The channels associated with this deployment action.
        /// </summary>
        public InputList<string> Channels
        {
            get => _channels ?? (_channels = new InputList<string>());
            set => _channels = value;
        }

        /// <summary>
        /// The condition associated with this deployment action.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("containers")]
        private InputList<Inputs.DeploymentProcessStepDeployPackageActionContainerGetArgs>? _containers;

        /// <summary>
        /// The deployment action container associated with this deployment action.
        /// </summary>
        public InputList<Inputs.DeploymentProcessStepDeployPackageActionContainerGetArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.DeploymentProcessStepDeployPackageActionContainerGetArgs>());
            set => _containers = value;
        }

        [Input("environments")]
        private InputList<string>? _environments;

        /// <summary>
        /// The environments within which this deployment action will run.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        [Input("excludedEnvironments")]
        private InputList<string>? _excludedEnvironments;

        /// <summary>
        /// The environments that this step will be skipped in
        /// </summary>
        public InputList<string> ExcludedEnvironments
        {
            get => _excludedEnvironments ?? (_excludedEnvironments = new InputList<string>());
            set => _excludedEnvironments = value;
        }

        [Input("features")]
        private InputList<string>? _features;

        /// <summary>
        /// A list of enabled features for this action.
        /// </summary>
        public InputList<string> Features
        {
            get => _features ?? (_features = new InputList<string>());
            set => _features = value;
        }

        /// <summary>
        /// Configuration for resource sourcing from a git repository.
        /// </summary>
        [Input("gitDependency")]
        public Input<Inputs.DeploymentProcessStepDeployPackageActionGitDependencyGetArgs>? GitDependency { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Indicates the disabled status of this deployment action.
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// Indicates the required status of this deployment action.
        /// </summary>
        [Input("isRequired")]
        public Input<bool>? IsRequired { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The notes associated with this deployment action.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("packages")]
        private InputList<Inputs.DeploymentProcessStepDeployPackageActionPackageGetArgs>? _packages;

        /// <summary>
        /// The package assocated with this action.
        /// </summary>
        public InputList<Inputs.DeploymentProcessStepDeployPackageActionPackageGetArgs> Packages
        {
            get => _packages ?? (_packages = new InputList<Inputs.DeploymentProcessStepDeployPackageActionPackageGetArgs>());
            set => _packages = value;
        }

        /// <summary>
        /// The package assocated with this action.
        /// </summary>
        [Input("primaryPackage", required: true)]
        public Input<Inputs.DeploymentProcessStepDeployPackageActionPrimaryPackageGetArgs> PrimaryPackage { get; set; } = null!;

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// A list of properties associated with this package.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The human-readable unique identifier for this resource.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("sortOrder")]
        public Input<int>? SortOrder { get; set; }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        /// <summary>
        /// Deploy a windows service feature
        /// </summary>
        [Input("windowsService")]
        public Input<Inputs.DeploymentProcessStepDeployPackageActionWindowsServiceGetArgs>? WindowsService { get; set; }

        public DeploymentProcessStepDeployPackageActionGetArgs()
        {
        }
        public static new DeploymentProcessStepDeployPackageActionGetArgs Empty => new DeploymentProcessStepDeployPackageActionGetArgs();
    }
}
