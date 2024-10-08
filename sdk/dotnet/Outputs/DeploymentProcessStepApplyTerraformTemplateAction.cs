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
    public sealed class DeploymentProcessStepApplyTerraformTemplateAction
    {
        /// <summary>
        /// Represents the template that is associated with this action.
        /// </summary>
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionActionTemplate? ActionTemplate;
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptions AdvancedOptions;
        /// <summary>
        /// (see below for nested schema)
        /// </summary>
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionAwsAccount? AwsAccount;
        /// <summary>
        /// (see below for nested schema)
        /// </summary>
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionAzureAccount? AzureAccount;
        public readonly bool? CanBeUsedForProjectVersioning;
        /// <summary>
        /// The channels associated with this deployment action.
        /// </summary>
        public readonly ImmutableArray<string> Channels;
        /// <summary>
        /// The condition associated with this deployment action.
        /// </summary>
        public readonly string? Condition;
        /// <summary>
        /// The deployment action container associated with this deployment action.
        /// </summary>
        public readonly ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionContainer> Containers;
        /// <summary>
        /// The environments within which this deployment action will run.
        /// </summary>
        public readonly ImmutableArray<string> Environments;
        /// <summary>
        /// The environments that this step will be skipped in
        /// </summary>
        public readonly ImmutableArray<string> ExcludedEnvironments;
        /// <summary>
        /// A list of enabled features for this action.
        /// </summary>
        public readonly ImmutableArray<string> Features;
        /// <summary>
        /// Configuration for resource sourcing from a git repository.
        /// </summary>
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionGitDependency? GitDependency;
        /// <summary>
        /// (see below for nested schema)
        /// </summary>
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccount? GoogleCloudAccount;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string? Id;
        public readonly string? InlineTemplate;
        /// <summary>
        /// Indicates the disabled status of this deployment action.
        /// </summary>
        public readonly bool? IsDisabled;
        /// <summary>
        /// Indicates the required status of this deployment action.
        /// </summary>
        public readonly bool? IsRequired;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The notes associated with this deployment action.
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// The package assocated with this action.
        /// </summary>
        public readonly ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionPackage> Packages;
        /// <summary>
        /// The package assocated with this action.
        /// </summary>
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionPrimaryPackage? PrimaryPackage;
        /// <summary>
        /// A list of properties associated with this package.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;
        /// <summary>
        /// Whether this step runs on a worker or on the target
        /// </summary>
        public readonly bool? RunOnServer;
        /// <summary>
        /// The human-readable unique identifier for this resource.
        /// </summary>
        public readonly string? Slug;
        public readonly int? SortOrder;
        /// <summary>
        /// (see below for nested schema)
        /// </summary>
        public readonly Outputs.DeploymentProcessStepApplyTerraformTemplateActionTemplate? Template;
        public readonly string? TemplateParameters;
        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> TenantTags;
        /// <summary>
        /// The worker pool associated with this deployment action.
        /// </summary>
        public readonly string? WorkerPoolId;
        /// <summary>
        /// The worker pool variable associated with this deployment action.
        /// </summary>
        public readonly string? WorkerPoolVariable;

        [OutputConstructor]
        private DeploymentProcessStepApplyTerraformTemplateAction(
            Outputs.DeploymentProcessStepApplyTerraformTemplateActionActionTemplate? actionTemplate,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionAdvancedOptions advancedOptions,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionAwsAccount? awsAccount,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionAzureAccount? azureAccount,

            bool? canBeUsedForProjectVersioning,

            ImmutableArray<string> channels,

            string? condition,

            ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionContainer> containers,

            ImmutableArray<string> environments,

            ImmutableArray<string> excludedEnvironments,

            ImmutableArray<string> features,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionGitDependency? gitDependency,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionGoogleCloudAccount? googleCloudAccount,

            string? id,

            string? inlineTemplate,

            bool? isDisabled,

            bool? isRequired,

            string name,

            string? notes,

            ImmutableArray<Outputs.DeploymentProcessStepApplyTerraformTemplateActionPackage> packages,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionPrimaryPackage? primaryPackage,

            ImmutableDictionary<string, string>? properties,

            bool? runOnServer,

            string? slug,

            int? sortOrder,

            Outputs.DeploymentProcessStepApplyTerraformTemplateActionTemplate? template,

            string? templateParameters,

            ImmutableArray<string> tenantTags,

            string? workerPoolId,

            string? workerPoolVariable)
        {
            ActionTemplate = actionTemplate;
            AdvancedOptions = advancedOptions;
            AwsAccount = awsAccount;
            AzureAccount = azureAccount;
            CanBeUsedForProjectVersioning = canBeUsedForProjectVersioning;
            Channels = channels;
            Condition = condition;
            Containers = containers;
            Environments = environments;
            ExcludedEnvironments = excludedEnvironments;
            Features = features;
            GitDependency = gitDependency;
            GoogleCloudAccount = googleCloudAccount;
            Id = id;
            InlineTemplate = inlineTemplate;
            IsDisabled = isDisabled;
            IsRequired = isRequired;
            Name = name;
            Notes = notes;
            Packages = packages;
            PrimaryPackage = primaryPackage;
            Properties = properties;
            RunOnServer = runOnServer;
            Slug = slug;
            SortOrder = sortOrder;
            Template = template;
            TemplateParameters = templateParameters;
            TenantTags = tenantTags;
            WorkerPoolId = workerPoolId;
            WorkerPoolVariable = workerPoolVariable;
        }
    }
}
