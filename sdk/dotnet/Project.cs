// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    /// <summary>
    /// This resource manages projects in Octopus Deploy.
    /// 
    /// &gt; Credentials are stored in state as plaintext. Read more about sensitive data in state.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = Pulumi.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Octopusdeploy.Project("example", new()
    ///     {
    ///         AutoCreateRelease = false,
    ///         DefaultGuidedFailureMode = "EnvironmentDefault",
    ///         DefaultToSkipIfAlreadyInstalled = false,
    ///         Description = "The development project.",
    ///         DiscreteChannelRelease = false,
    ///         IsDisabled = false,
    ///         IsDiscreteChannelRelease = false,
    ///         IsVersionControlled = false,
    ///         LifecycleId = "Lifecycles-123",
    ///         ProjectGroupId = "ProjectGroups-123",
    ///         TenantedDeploymentParticipation = "TenantedOrUntenanted",
    ///         IncludedLibraryVariableSets = new[]
    ///         {
    ///             "LibraryVariablesSets-456",
    ///             "LibraryVariablesSets-789",
    ///         },
    ///         ConnectivityPolicies = new[]
    ///         {
    ///             new Octopusdeploy.Inputs.ProjectConnectivityPolicyArgs
    ///             {
    ///                 AllowDeploymentsToNoTargets = false,
    ///                 ExcludeUnhealthyTargets = false,
    ///                 SkipMachineBehavior = "SkipUnavailableMachines",
    ///             },
    ///         },
    ///         JiraServiceManagementExtensionSettings = new[]
    ///         {
    ///             new Octopusdeploy.Inputs.ProjectJiraServiceManagementExtensionSettingArgs
    ///             {
    ///                 ConnectionId = "133d7fe602514060a48bc42ee9870f99",
    ///                 IsEnabled = false,
    ///                 ServiceDeskProjectName = "Test Service Desk Project (OK to Delete)",
    ///             },
    ///         },
    ///         ServicenowExtensionSettings = new[]
    ///         {
    ///             new Octopusdeploy.Inputs.ProjectServicenowExtensionSettingArgs
    ///             {
    ///                 ConnectionId = "989034685e2c48c4b06a29286c9ef5cc",
    ///                 IsEnabled = false,
    ///                 IsStateAutomaticallyTransitioned = false,
    ///                 StandardChangeTemplateName = "Standard Change Template Name (OK to Delete)",
    ///             },
    ///         },
    ///         Templates = new[]
    ///         {
    ///             new Octopusdeploy.Inputs.ProjectTemplateArgs
    ///             {
    ///                 DefaultValue = "example-default-value",
    ///                 HelpText = "example-help-test",
    ///                 Label = "example-label",
    ///                 Name = "example-template-value",
    ///                 DisplaySettings = 
    ///                 {
    ///                     { "Octopus.ControlType", "SingleLineText" },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/project:Project [options] octopusdeploy_project.&lt;name&gt; &lt;project-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        [Output("allowDeploymentsToNoTargets")]
        public Output<bool?> AllowDeploymentsToNoTargets { get; private set; } = null!;

        [Output("autoCreateRelease")]
        public Output<bool> AutoCreateRelease { get; private set; } = null!;

        [Output("autoDeployReleaseOverrides")]
        public Output<ImmutableArray<Outputs.ProjectAutoDeployReleaseOverride>> AutoDeployReleaseOverrides { get; private set; } = null!;

        /// <summary>
        /// The ID of the project this project was cloned from.
        /// </summary>
        [Output("clonedFromProjectId")]
        public Output<string?> ClonedFromProjectId { get; private set; } = null!;

        [Output("connectivityPolicies")]
        public Output<ImmutableArray<Outputs.ProjectConnectivityPolicy>> ConnectivityPolicies { get; private set; } = null!;

        [Output("defaultGuidedFailureMode")]
        public Output<string> DefaultGuidedFailureMode { get; private set; } = null!;

        [Output("defaultToSkipIfAlreadyInstalled")]
        public Output<bool> DefaultToSkipIfAlreadyInstalled { get; private set; } = null!;

        [Output("deploymentChangesTemplate")]
        public Output<string> DeploymentChangesTemplate { get; private set; } = null!;

        [Output("deploymentProcessId")]
        public Output<string> DeploymentProcessId { get; private set; } = null!;

        /// <summary>
        /// The description of this project.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Output("discreteChannelRelease")]
        public Output<bool> DiscreteChannelRelease { get; private set; } = null!;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        [Output("gitAnonymousPersistenceSettings")]
        public Output<ImmutableArray<Outputs.ProjectGitAnonymousPersistenceSetting>> GitAnonymousPersistenceSettings { get; private set; } = null!;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        [Output("gitLibraryPersistenceSettings")]
        public Output<ImmutableArray<Outputs.ProjectGitLibraryPersistenceSetting>> GitLibraryPersistenceSettings { get; private set; } = null!;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        [Output("gitUsernamePasswordPersistenceSettings")]
        public Output<ImmutableArray<Outputs.ProjectGitUsernamePasswordPersistenceSetting>> GitUsernamePasswordPersistenceSettings { get; private set; } = null!;

        /// <summary>
        /// The list of included library variable set IDs.
        /// </summary>
        [Output("includedLibraryVariableSets")]
        public Output<ImmutableArray<string>> IncludedLibraryVariableSets { get; private set; } = null!;

        [Output("isDisabled")]
        public Output<bool> IsDisabled { get; private set; } = null!;

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Output("isDiscreteChannelRelease")]
        public Output<bool> IsDiscreteChannelRelease { get; private set; } = null!;

        [Output("isVersionControlled")]
        public Output<bool> IsVersionControlled { get; private set; } = null!;

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this project.
        /// </summary>
        [Output("jiraServiceManagementExtensionSettings")]
        public Output<ImmutableArray<Outputs.ProjectJiraServiceManagementExtensionSetting>> JiraServiceManagementExtensionSettings { get; private set; } = null!;

        /// <summary>
        /// The lifecycle ID associated with this project.
        /// </summary>
        [Output("lifecycleId")]
        public Output<string> LifecycleId { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project group ID associated with this project.
        /// </summary>
        [Output("projectGroupId")]
        public Output<string> ProjectGroupId { get; private set; } = null!;

        [Output("releaseCreationStrategies")]
        public Output<ImmutableArray<Outputs.ProjectReleaseCreationStrategy>> ReleaseCreationStrategies { get; private set; } = null!;

        [Output("releaseNotesTemplate")]
        public Output<string> ReleaseNotesTemplate { get; private set; } = null!;

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this project.
        /// </summary>
        [Output("servicenowExtensionSettings")]
        public Output<ImmutableArray<Outputs.ProjectServicenowExtensionSetting>> ServicenowExtensionSettings { get; private set; } = null!;

        /// <summary>
        /// A human-readable, unique identifier, used to identify a project.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this project.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        [Output("templates")]
        public Output<ImmutableArray<Outputs.ProjectTemplate>> Templates { get; private set; } = null!;

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Output("tenantedDeploymentParticipation")]
        public Output<string> TenantedDeploymentParticipation { get; private set; } = null!;

        [Output("variableSetId")]
        public Output<string> VariableSetId { get; private set; } = null!;

        [Output("versioningStrategies")]
        public Output<ImmutableArray<Outputs.ProjectVersioningStrategy>> VersioningStrategies { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/project:Project", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/OctopusDeploy/terraform-provider-octopusdeploy/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowDeploymentsToNoTargets")]
        public Input<bool>? AllowDeploymentsToNoTargets { get; set; }

        [Input("autoCreateRelease")]
        public Input<bool>? AutoCreateRelease { get; set; }

        [Input("autoDeployReleaseOverrides")]
        private InputList<Inputs.ProjectAutoDeployReleaseOverrideArgs>? _autoDeployReleaseOverrides;
        public InputList<Inputs.ProjectAutoDeployReleaseOverrideArgs> AutoDeployReleaseOverrides
        {
            get => _autoDeployReleaseOverrides ?? (_autoDeployReleaseOverrides = new InputList<Inputs.ProjectAutoDeployReleaseOverrideArgs>());
            set => _autoDeployReleaseOverrides = value;
        }

        /// <summary>
        /// The ID of the project this project was cloned from.
        /// </summary>
        [Input("clonedFromProjectId")]
        public Input<string>? ClonedFromProjectId { get; set; }

        [Input("connectivityPolicies")]
        private InputList<Inputs.ProjectConnectivityPolicyArgs>? _connectivityPolicies;
        public InputList<Inputs.ProjectConnectivityPolicyArgs> ConnectivityPolicies
        {
            get => _connectivityPolicies ?? (_connectivityPolicies = new InputList<Inputs.ProjectConnectivityPolicyArgs>());
            set => _connectivityPolicies = value;
        }

        [Input("defaultGuidedFailureMode")]
        public Input<string>? DefaultGuidedFailureMode { get; set; }

        [Input("defaultToSkipIfAlreadyInstalled")]
        public Input<bool>? DefaultToSkipIfAlreadyInstalled { get; set; }

        [Input("deploymentChangesTemplate")]
        public Input<string>? DeploymentChangesTemplate { get; set; }

        /// <summary>
        /// The description of this project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Input("discreteChannelRelease")]
        public Input<bool>? DiscreteChannelRelease { get; set; }

        [Input("gitAnonymousPersistenceSettings")]
        private InputList<Inputs.ProjectGitAnonymousPersistenceSettingArgs>? _gitAnonymousPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.ProjectGitAnonymousPersistenceSettingArgs> GitAnonymousPersistenceSettings
        {
            get => _gitAnonymousPersistenceSettings ?? (_gitAnonymousPersistenceSettings = new InputList<Inputs.ProjectGitAnonymousPersistenceSettingArgs>());
            set => _gitAnonymousPersistenceSettings = value;
        }

        [Input("gitLibraryPersistenceSettings")]
        private InputList<Inputs.ProjectGitLibraryPersistenceSettingArgs>? _gitLibraryPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.ProjectGitLibraryPersistenceSettingArgs> GitLibraryPersistenceSettings
        {
            get => _gitLibraryPersistenceSettings ?? (_gitLibraryPersistenceSettings = new InputList<Inputs.ProjectGitLibraryPersistenceSettingArgs>());
            set => _gitLibraryPersistenceSettings = value;
        }

        [Input("gitUsernamePasswordPersistenceSettings")]
        private InputList<Inputs.ProjectGitUsernamePasswordPersistenceSettingArgs>? _gitUsernamePasswordPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.ProjectGitUsernamePasswordPersistenceSettingArgs> GitUsernamePasswordPersistenceSettings
        {
            get => _gitUsernamePasswordPersistenceSettings ?? (_gitUsernamePasswordPersistenceSettings = new InputList<Inputs.ProjectGitUsernamePasswordPersistenceSettingArgs>());
            set => _gitUsernamePasswordPersistenceSettings = value;
        }

        [Input("includedLibraryVariableSets")]
        private InputList<string>? _includedLibraryVariableSets;

        /// <summary>
        /// The list of included library variable set IDs.
        /// </summary>
        public InputList<string> IncludedLibraryVariableSets
        {
            get => _includedLibraryVariableSets ?? (_includedLibraryVariableSets = new InputList<string>());
            set => _includedLibraryVariableSets = value;
        }

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Input("isDiscreteChannelRelease")]
        public Input<bool>? IsDiscreteChannelRelease { get; set; }

        [Input("isVersionControlled")]
        public Input<bool>? IsVersionControlled { get; set; }

        [Input("jiraServiceManagementExtensionSettings")]
        private InputList<Inputs.ProjectJiraServiceManagementExtensionSettingArgs>? _jiraServiceManagementExtensionSettings;

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this project.
        /// </summary>
        public InputList<Inputs.ProjectJiraServiceManagementExtensionSettingArgs> JiraServiceManagementExtensionSettings
        {
            get => _jiraServiceManagementExtensionSettings ?? (_jiraServiceManagementExtensionSettings = new InputList<Inputs.ProjectJiraServiceManagementExtensionSettingArgs>());
            set => _jiraServiceManagementExtensionSettings = value;
        }

        /// <summary>
        /// The lifecycle ID associated with this project.
        /// </summary>
        [Input("lifecycleId", required: true)]
        public Input<string> LifecycleId { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project group ID associated with this project.
        /// </summary>
        [Input("projectGroupId", required: true)]
        public Input<string> ProjectGroupId { get; set; } = null!;

        [Input("releaseCreationStrategies")]
        private InputList<Inputs.ProjectReleaseCreationStrategyArgs>? _releaseCreationStrategies;
        public InputList<Inputs.ProjectReleaseCreationStrategyArgs> ReleaseCreationStrategies
        {
            get => _releaseCreationStrategies ?? (_releaseCreationStrategies = new InputList<Inputs.ProjectReleaseCreationStrategyArgs>());
            set => _releaseCreationStrategies = value;
        }

        [Input("releaseNotesTemplate")]
        public Input<string>? ReleaseNotesTemplate { get; set; }

        [Input("servicenowExtensionSettings")]
        private InputList<Inputs.ProjectServicenowExtensionSettingArgs>? _servicenowExtensionSettings;

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this project.
        /// </summary>
        public InputList<Inputs.ProjectServicenowExtensionSettingArgs> ServicenowExtensionSettings
        {
            get => _servicenowExtensionSettings ?? (_servicenowExtensionSettings = new InputList<Inputs.ProjectServicenowExtensionSettingArgs>());
            set => _servicenowExtensionSettings = value;
        }

        /// <summary>
        /// A human-readable, unique identifier, used to identify a project.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// The space ID associated with this project.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("templates")]
        private InputList<Inputs.ProjectTemplateArgs>? _templates;
        public InputList<Inputs.ProjectTemplateArgs> Templates
        {
            get => _templates ?? (_templates = new InputList<Inputs.ProjectTemplateArgs>());
            set => _templates = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation")]
        public Input<string>? TenantedDeploymentParticipation { get; set; }

        [Input("versioningStrategies")]
        private InputList<Inputs.ProjectVersioningStrategyArgs>? _versioningStrategies;
        public InputList<Inputs.ProjectVersioningStrategyArgs> VersioningStrategies
        {
            get => _versioningStrategies ?? (_versioningStrategies = new InputList<Inputs.ProjectVersioningStrategyArgs>());
            set => _versioningStrategies = value;
        }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        [Input("allowDeploymentsToNoTargets")]
        public Input<bool>? AllowDeploymentsToNoTargets { get; set; }

        [Input("autoCreateRelease")]
        public Input<bool>? AutoCreateRelease { get; set; }

        [Input("autoDeployReleaseOverrides")]
        private InputList<Inputs.ProjectAutoDeployReleaseOverrideGetArgs>? _autoDeployReleaseOverrides;
        public InputList<Inputs.ProjectAutoDeployReleaseOverrideGetArgs> AutoDeployReleaseOverrides
        {
            get => _autoDeployReleaseOverrides ?? (_autoDeployReleaseOverrides = new InputList<Inputs.ProjectAutoDeployReleaseOverrideGetArgs>());
            set => _autoDeployReleaseOverrides = value;
        }

        /// <summary>
        /// The ID of the project this project was cloned from.
        /// </summary>
        [Input("clonedFromProjectId")]
        public Input<string>? ClonedFromProjectId { get; set; }

        [Input("connectivityPolicies")]
        private InputList<Inputs.ProjectConnectivityPolicyGetArgs>? _connectivityPolicies;
        public InputList<Inputs.ProjectConnectivityPolicyGetArgs> ConnectivityPolicies
        {
            get => _connectivityPolicies ?? (_connectivityPolicies = new InputList<Inputs.ProjectConnectivityPolicyGetArgs>());
            set => _connectivityPolicies = value;
        }

        [Input("defaultGuidedFailureMode")]
        public Input<string>? DefaultGuidedFailureMode { get; set; }

        [Input("defaultToSkipIfAlreadyInstalled")]
        public Input<bool>? DefaultToSkipIfAlreadyInstalled { get; set; }

        [Input("deploymentChangesTemplate")]
        public Input<string>? DeploymentChangesTemplate { get; set; }

        [Input("deploymentProcessId")]
        public Input<string>? DeploymentProcessId { get; set; }

        /// <summary>
        /// The description of this project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Input("discreteChannelRelease")]
        public Input<bool>? DiscreteChannelRelease { get; set; }

        [Input("gitAnonymousPersistenceSettings")]
        private InputList<Inputs.ProjectGitAnonymousPersistenceSettingGetArgs>? _gitAnonymousPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.ProjectGitAnonymousPersistenceSettingGetArgs> GitAnonymousPersistenceSettings
        {
            get => _gitAnonymousPersistenceSettings ?? (_gitAnonymousPersistenceSettings = new InputList<Inputs.ProjectGitAnonymousPersistenceSettingGetArgs>());
            set => _gitAnonymousPersistenceSettings = value;
        }

        [Input("gitLibraryPersistenceSettings")]
        private InputList<Inputs.ProjectGitLibraryPersistenceSettingGetArgs>? _gitLibraryPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.ProjectGitLibraryPersistenceSettingGetArgs> GitLibraryPersistenceSettings
        {
            get => _gitLibraryPersistenceSettings ?? (_gitLibraryPersistenceSettings = new InputList<Inputs.ProjectGitLibraryPersistenceSettingGetArgs>());
            set => _gitLibraryPersistenceSettings = value;
        }

        [Input("gitUsernamePasswordPersistenceSettings")]
        private InputList<Inputs.ProjectGitUsernamePasswordPersistenceSettingGetArgs>? _gitUsernamePasswordPersistenceSettings;

        /// <summary>
        /// Provides Git-related persistence settings for a version-controlled project.
        /// </summary>
        public InputList<Inputs.ProjectGitUsernamePasswordPersistenceSettingGetArgs> GitUsernamePasswordPersistenceSettings
        {
            get => _gitUsernamePasswordPersistenceSettings ?? (_gitUsernamePasswordPersistenceSettings = new InputList<Inputs.ProjectGitUsernamePasswordPersistenceSettingGetArgs>());
            set => _gitUsernamePasswordPersistenceSettings = value;
        }

        [Input("includedLibraryVariableSets")]
        private InputList<string>? _includedLibraryVariableSets;

        /// <summary>
        /// The list of included library variable set IDs.
        /// </summary>
        public InputList<string> IncludedLibraryVariableSets
        {
            get => _includedLibraryVariableSets ?? (_includedLibraryVariableSets = new InputList<string>());
            set => _includedLibraryVariableSets = value;
        }

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// Treats releases of different channels to the same environment as a separate deployment dimension
        /// </summary>
        [Input("isDiscreteChannelRelease")]
        public Input<bool>? IsDiscreteChannelRelease { get; set; }

        [Input("isVersionControlled")]
        public Input<bool>? IsVersionControlled { get; set; }

        [Input("jiraServiceManagementExtensionSettings")]
        private InputList<Inputs.ProjectJiraServiceManagementExtensionSettingGetArgs>? _jiraServiceManagementExtensionSettings;

        /// <summary>
        /// Provides extension settings for the Jira Service Management (JSM) integration for this project.
        /// </summary>
        public InputList<Inputs.ProjectJiraServiceManagementExtensionSettingGetArgs> JiraServiceManagementExtensionSettings
        {
            get => _jiraServiceManagementExtensionSettings ?? (_jiraServiceManagementExtensionSettings = new InputList<Inputs.ProjectJiraServiceManagementExtensionSettingGetArgs>());
            set => _jiraServiceManagementExtensionSettings = value;
        }

        /// <summary>
        /// The lifecycle ID associated with this project.
        /// </summary>
        [Input("lifecycleId")]
        public Input<string>? LifecycleId { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project group ID associated with this project.
        /// </summary>
        [Input("projectGroupId")]
        public Input<string>? ProjectGroupId { get; set; }

        [Input("releaseCreationStrategies")]
        private InputList<Inputs.ProjectReleaseCreationStrategyGetArgs>? _releaseCreationStrategies;
        public InputList<Inputs.ProjectReleaseCreationStrategyGetArgs> ReleaseCreationStrategies
        {
            get => _releaseCreationStrategies ?? (_releaseCreationStrategies = new InputList<Inputs.ProjectReleaseCreationStrategyGetArgs>());
            set => _releaseCreationStrategies = value;
        }

        [Input("releaseNotesTemplate")]
        public Input<string>? ReleaseNotesTemplate { get; set; }

        [Input("servicenowExtensionSettings")]
        private InputList<Inputs.ProjectServicenowExtensionSettingGetArgs>? _servicenowExtensionSettings;

        /// <summary>
        /// Provides extension settings for the ServiceNow integration for this project.
        /// </summary>
        public InputList<Inputs.ProjectServicenowExtensionSettingGetArgs> ServicenowExtensionSettings
        {
            get => _servicenowExtensionSettings ?? (_servicenowExtensionSettings = new InputList<Inputs.ProjectServicenowExtensionSettingGetArgs>());
            set => _servicenowExtensionSettings = value;
        }

        /// <summary>
        /// A human-readable, unique identifier, used to identify a project.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        /// <summary>
        /// The space ID associated with this project.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("templates")]
        private InputList<Inputs.ProjectTemplateGetArgs>? _templates;
        public InputList<Inputs.ProjectTemplateGetArgs> Templates
        {
            get => _templates ?? (_templates = new InputList<Inputs.ProjectTemplateGetArgs>());
            set => _templates = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation")]
        public Input<string>? TenantedDeploymentParticipation { get; set; }

        [Input("variableSetId")]
        public Input<string>? VariableSetId { get; set; }

        [Input("versioningStrategies")]
        private InputList<Inputs.ProjectVersioningStrategyGetArgs>? _versioningStrategies;
        public InputList<Inputs.ProjectVersioningStrategyGetArgs> VersioningStrategies
        {
            get => _versioningStrategies ?? (_versioningStrategies = new InputList<Inputs.ProjectVersioningStrategyGetArgs>());
            set => _versioningStrategies = value;
        }

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
