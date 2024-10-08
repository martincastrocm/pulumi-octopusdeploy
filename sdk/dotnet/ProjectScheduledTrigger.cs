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
    /// This resource manages a scheduled trigger for a project or runbook in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = Pulumi.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var onceDailyExample = new Octopusdeploy.ProjectScheduledTrigger("onceDailyExample", new()
    ///     {
    ///         DeployNewReleaseAction = new Octopusdeploy.Inputs.ProjectScheduledTriggerDeployNewReleaseActionArgs
    ///         {
    ///             DestinationEnvironmentId = "environments-123",
    ///         },
    ///         Description = "This is a once daily schedule",
    ///         OnceDailySchedule = new Octopusdeploy.Inputs.ProjectScheduledTriggerOnceDailyScheduleArgs
    ///         {
    ///             DaysOfWeeks = new[]
    ///             {
    ///                 "Tuesday",
    ///                 "Wednesday",
    ///                 "Monday",
    ///             },
    ///             StartTime = "2024-03-22T09:00:00",
    ///         },
    ///         ProjectId = "projects-123",
    ///         SpaceId = "spaces-123",
    ///     });
    /// 
    ///     var continuousExample = new Octopusdeploy.ProjectScheduledTrigger("continuousExample", new()
    ///     {
    ///         ContinuousDailySchedule = new Octopusdeploy.Inputs.ProjectScheduledTriggerContinuousDailyScheduleArgs
    ///         {
    ///             DaysOfWeeks = new[]
    ///             {
    ///                 "Monday",
    ///                 "Tuesday",
    ///                 "Friday",
    ///             },
    ///             HourInterval = 3,
    ///             Interval = "OnceHourly",
    ///             RunAfter = "2024-03-22T09:00:00",
    ///             RunUntil = "2024-03-29T13:00:00",
    ///         },
    ///         DeployNewReleaseAction = new Octopusdeploy.Inputs.ProjectScheduledTriggerDeployNewReleaseActionArgs
    ///         {
    ///             DestinationEnvironmentId = "environments-123",
    ///         },
    ///         Description = "This is a continuous daily schedule",
    ///         ProjectId = "projects-123",
    ///         SpaceId = "spaces-123",
    ///     });
    /// 
    ///     var deployLatestExample = new Octopusdeploy.ProjectScheduledTrigger("deployLatestExample", new()
    ///     {
    ///         CronExpressionSchedule = new Octopusdeploy.Inputs.ProjectScheduledTriggerCronExpressionScheduleArgs
    ///         {
    ///             CronExpression = "0 0 06 * * Mon-Fri",
    ///         },
    ///         DeployLatestReleaseAction = new Octopusdeploy.Inputs.ProjectScheduledTriggerDeployLatestReleaseActionArgs
    ///         {
    ///             DestinationEnvironmentId = "environments-123",
    ///             ShouldRedeploy = true,
    ///             SourceEnvironmentId = "environments-321",
    ///         },
    ///         ProjectId = "projects-123",
    ///         SpaceId = "spaces-123",
    ///     });
    /// 
    ///     var deployNewExample = new Octopusdeploy.ProjectScheduledTrigger("deployNewExample", new()
    ///     {
    ///         CronExpressionSchedule = new Octopusdeploy.Inputs.ProjectScheduledTriggerCronExpressionScheduleArgs
    ///         {
    ///             CronExpression = "0 0 06 * * Mon-Fri",
    ///         },
    ///         DeployNewReleaseAction = new Octopusdeploy.Inputs.ProjectScheduledTriggerDeployNewReleaseActionArgs
    ///         {
    ///             DestinationEnvironmentId = "environments-123",
    ///         },
    ///         ProjectId = "projects-123",
    ///         SpaceId = "spaces-123",
    ///     });
    /// 
    ///     var runbookExample = new Octopusdeploy.ProjectScheduledTrigger("runbookExample", new()
    ///     {
    ///         CronExpressionSchedule = new Octopusdeploy.Inputs.ProjectScheduledTriggerCronExpressionScheduleArgs
    ///         {
    ///             CronExpression = "0 0 06 * * Mon-Fri",
    ///         },
    ///         Description = "This is a Cron schedule",
    ///         ProjectId = "projects-123",
    ///         RunRunbookAction = new Octopusdeploy.Inputs.ProjectScheduledTriggerRunRunbookActionArgs
    ///         {
    ///             RunbookId = "runbooks-123",
    ///             TargetEnvironmentIds = new[]
    ///             {
    ///                 "environments-123",
    ///                 "environments-321",
    ///             },
    ///         },
    ///         SpaceId = "spaces-123",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/projectScheduledTrigger:ProjectScheduledTrigger [options] octopusdeploy_project_scheduled_trigger.&lt;name&gt; &lt;trigger-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/projectScheduledTrigger:ProjectScheduledTrigger")]
    public partial class ProjectScheduledTrigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The channel ID to use when creating the release. Will use the default channel if left blank.
        /// </summary>
        [Output("channelId")]
        public Output<string?> ChannelId { get; private set; } = null!;

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Output("continuousDailySchedule")]
        public Output<Outputs.ProjectScheduledTriggerContinuousDailySchedule?> ContinuousDailySchedule { get; private set; } = null!;

        /// <summary>
        /// The cron expression schedule for the trigger.
        /// </summary>
        [Output("cronExpressionSchedule")]
        public Output<Outputs.ProjectScheduledTriggerCronExpressionSchedule?> CronExpressionSchedule { get; private set; } = null!;

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Output("daysPerMonthSchedule")]
        public Output<Outputs.ProjectScheduledTriggerDaysPerMonthSchedule?> DaysPerMonthSchedule { get; private set; } = null!;

        /// <summary>
        /// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
        /// </summary>
        [Output("deployLatestReleaseAction")]
        public Output<Outputs.ProjectScheduledTriggerDeployLatestReleaseAction?> DeployLatestReleaseAction { get; private set; } = null!;

        /// <summary>
        /// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
        /// </summary>
        [Output("deployNewReleaseAction")]
        public Output<Outputs.ProjectScheduledTriggerDeployNewReleaseAction?> DeployNewReleaseAction { get; private set; } = null!;

        /// <summary>
        /// A description of the trigger.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the trigger is disabled.
        /// </summary>
        [Output("isDisabled")]
        public Output<bool?> IsDisabled { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Output("onceDailySchedule")]
        public Output<Outputs.ProjectScheduledTriggerOnceDailySchedule?> OnceDailySchedule { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to attach the trigger.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
        /// </summary>
        [Output("runRunbookAction")]
        public Output<Outputs.ProjectScheduledTriggerRunRunbookAction?> RunRunbookAction { get; private set; } = null!;

        /// <summary>
        /// The space ID where this trigger's project exists.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// The IDs of the tenants to deploy to.
        /// </summary>
        [Output("tenantIds")]
        public Output<ImmutableArray<string>> TenantIds { get; private set; } = null!;

        /// <summary>
        /// The timezone for the trigger.
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectScheduledTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectScheduledTrigger(string name, ProjectScheduledTriggerArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/projectScheduledTrigger:ProjectScheduledTrigger", name, args ?? new ProjectScheduledTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectScheduledTrigger(string name, Input<string> id, ProjectScheduledTriggerState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/projectScheduledTrigger:ProjectScheduledTrigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectScheduledTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectScheduledTrigger Get(string name, Input<string> id, ProjectScheduledTriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectScheduledTrigger(name, id, state, options);
        }
    }

    public sealed class ProjectScheduledTriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The channel ID to use when creating the release. Will use the default channel if left blank.
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Input("continuousDailySchedule")]
        public Input<Inputs.ProjectScheduledTriggerContinuousDailyScheduleArgs>? ContinuousDailySchedule { get; set; }

        /// <summary>
        /// The cron expression schedule for the trigger.
        /// </summary>
        [Input("cronExpressionSchedule")]
        public Input<Inputs.ProjectScheduledTriggerCronExpressionScheduleArgs>? CronExpressionSchedule { get; set; }

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Input("daysPerMonthSchedule")]
        public Input<Inputs.ProjectScheduledTriggerDaysPerMonthScheduleArgs>? DaysPerMonthSchedule { get; set; }

        /// <summary>
        /// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
        /// </summary>
        [Input("deployLatestReleaseAction")]
        public Input<Inputs.ProjectScheduledTriggerDeployLatestReleaseActionArgs>? DeployLatestReleaseAction { get; set; }

        /// <summary>
        /// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
        /// </summary>
        [Input("deployNewReleaseAction")]
        public Input<Inputs.ProjectScheduledTriggerDeployNewReleaseActionArgs>? DeployNewReleaseAction { get; set; }

        /// <summary>
        /// A description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether the trigger is disabled.
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Input("onceDailySchedule")]
        public Input<Inputs.ProjectScheduledTriggerOnceDailyScheduleArgs>? OnceDailySchedule { get; set; }

        /// <summary>
        /// The ID of the project to attach the trigger.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
        /// </summary>
        [Input("runRunbookAction")]
        public Input<Inputs.ProjectScheduledTriggerRunRunbookActionArgs>? RunRunbookAction { get; set; }

        /// <summary>
        /// The space ID where this trigger's project exists.
        /// </summary>
        [Input("spaceId", required: true)]
        public Input<string> SpaceId { get; set; } = null!;

        [Input("tenantIds")]
        private InputList<string>? _tenantIds;

        /// <summary>
        /// The IDs of the tenants to deploy to.
        /// </summary>
        public InputList<string> TenantIds
        {
            get => _tenantIds ?? (_tenantIds = new InputList<string>());
            set => _tenantIds = value;
        }

        /// <summary>
        /// The timezone for the trigger.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public ProjectScheduledTriggerArgs()
        {
        }
        public static new ProjectScheduledTriggerArgs Empty => new ProjectScheduledTriggerArgs();
    }

    public sealed class ProjectScheduledTriggerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The channel ID to use when creating the release. Will use the default channel if left blank.
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Input("continuousDailySchedule")]
        public Input<Inputs.ProjectScheduledTriggerContinuousDailyScheduleGetArgs>? ContinuousDailySchedule { get; set; }

        /// <summary>
        /// The cron expression schedule for the trigger.
        /// </summary>
        [Input("cronExpressionSchedule")]
        public Input<Inputs.ProjectScheduledTriggerCronExpressionScheduleGetArgs>? CronExpressionSchedule { get; set; }

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Input("daysPerMonthSchedule")]
        public Input<Inputs.ProjectScheduledTriggerDaysPerMonthScheduleGetArgs>? DaysPerMonthSchedule { get; set; }

        /// <summary>
        /// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
        /// </summary>
        [Input("deployLatestReleaseAction")]
        public Input<Inputs.ProjectScheduledTriggerDeployLatestReleaseActionGetArgs>? DeployLatestReleaseAction { get; set; }

        /// <summary>
        /// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
        /// </summary>
        [Input("deployNewReleaseAction")]
        public Input<Inputs.ProjectScheduledTriggerDeployNewReleaseActionGetArgs>? DeployNewReleaseAction { get; set; }

        /// <summary>
        /// A description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether the trigger is disabled.
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The daily schedule for the trigger.
        /// </summary>
        [Input("onceDailySchedule")]
        public Input<Inputs.ProjectScheduledTriggerOnceDailyScheduleGetArgs>? OnceDailySchedule { get; set; }

        /// <summary>
        /// The ID of the project to attach the trigger.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
        /// </summary>
        [Input("runRunbookAction")]
        public Input<Inputs.ProjectScheduledTriggerRunRunbookActionGetArgs>? RunRunbookAction { get; set; }

        /// <summary>
        /// The space ID where this trigger's project exists.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        [Input("tenantIds")]
        private InputList<string>? _tenantIds;

        /// <summary>
        /// The IDs of the tenants to deploy to.
        /// </summary>
        public InputList<string> TenantIds
        {
            get => _tenantIds ?? (_tenantIds = new InputList<string>());
            set => _tenantIds = value;
        }

        /// <summary>
        /// The timezone for the trigger.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public ProjectScheduledTriggerState()
        {
        }
        public static new ProjectScheduledTriggerState Empty => new ProjectScheduledTriggerState();
    }
}
