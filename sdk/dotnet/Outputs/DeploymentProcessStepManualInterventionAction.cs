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
    public sealed class DeploymentProcessStepManualInterventionAction
    {
        /// <summary>
        /// Represents the template that is associated with this action.
        /// </summary>
        public readonly Outputs.DeploymentProcessStepManualInterventionActionActionTemplate? ActionTemplate;
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
        public readonly ImmutableArray<Outputs.DeploymentProcessStepManualInterventionActionContainer> Containers;
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
        public readonly Outputs.DeploymentProcessStepManualInterventionActionGitDependency? GitDependency;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The instructions for the user to follow
        /// </summary>
        public readonly string Instructions;
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
        public readonly ImmutableArray<Outputs.DeploymentProcessStepManualInterventionActionPackage> Packages;
        /// <summary>
        /// A list of properties associated with this package.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;
        /// <summary>
        /// The teams responsible to resolve this step. If no teams are specified, all users who have permission to deploy the project can resolve it.
        /// </summary>
        public readonly string? ResponsibleTeams;
        /// <summary>
        /// The human-readable unique identifier for this resource.
        /// </summary>
        public readonly string? Slug;
        public readonly int? SortOrder;
        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> TenantTags;

        [OutputConstructor]
        private DeploymentProcessStepManualInterventionAction(
            Outputs.DeploymentProcessStepManualInterventionActionActionTemplate? actionTemplate,

            bool? canBeUsedForProjectVersioning,

            ImmutableArray<string> channels,

            string? condition,

            ImmutableArray<Outputs.DeploymentProcessStepManualInterventionActionContainer> containers,

            ImmutableArray<string> environments,

            ImmutableArray<string> excludedEnvironments,

            ImmutableArray<string> features,

            Outputs.DeploymentProcessStepManualInterventionActionGitDependency? gitDependency,

            string? id,

            string instructions,

            bool? isDisabled,

            bool? isRequired,

            string name,

            string? notes,

            ImmutableArray<Outputs.DeploymentProcessStepManualInterventionActionPackage> packages,

            ImmutableDictionary<string, string>? properties,

            string? responsibleTeams,

            string? slug,

            int? sortOrder,

            ImmutableArray<string> tenantTags)
        {
            ActionTemplate = actionTemplate;
            CanBeUsedForProjectVersioning = canBeUsedForProjectVersioning;
            Channels = channels;
            Condition = condition;
            Containers = containers;
            Environments = environments;
            ExcludedEnvironments = excludedEnvironments;
            Features = features;
            GitDependency = gitDependency;
            Id = id;
            Instructions = instructions;
            IsDisabled = isDisabled;
            IsRequired = isRequired;
            Name = name;
            Notes = notes;
            Packages = packages;
            Properties = properties;
            ResponsibleTeams = responsibleTeams;
            Slug = slug;
            SortOrder = sortOrder;
            TenantTags = tenantTags;
        }
    }
}
