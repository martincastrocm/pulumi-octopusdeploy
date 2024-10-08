// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages a scheduled trigger for a project or runbook in Octopus Deploy.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := octopusdeploy.NewProjectScheduledTrigger(ctx, "onceDailyExample", &octopusdeploy.ProjectScheduledTriggerArgs{
//				DeployNewReleaseAction: &octopusdeploy.ProjectScheduledTriggerDeployNewReleaseActionArgs{
//					DestinationEnvironmentId: pulumi.String("environments-123"),
//				},
//				Description: pulumi.String("This is a once daily schedule"),
//				OnceDailySchedule: &octopusdeploy.ProjectScheduledTriggerOnceDailyScheduleArgs{
//					DaysOfWeeks: pulumi.StringArray{
//						pulumi.String("Tuesday"),
//						pulumi.String("Wednesday"),
//						pulumi.String("Monday"),
//					},
//					StartTime: pulumi.String("2024-03-22T09:00:00"),
//				},
//				ProjectId: pulumi.String("projects-123"),
//				SpaceId:   pulumi.String("spaces-123"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewProjectScheduledTrigger(ctx, "continuousExample", &octopusdeploy.ProjectScheduledTriggerArgs{
//				ContinuousDailySchedule: &octopusdeploy.ProjectScheduledTriggerContinuousDailyScheduleArgs{
//					DaysOfWeeks: pulumi.StringArray{
//						pulumi.String("Monday"),
//						pulumi.String("Tuesday"),
//						pulumi.String("Friday"),
//					},
//					HourInterval: pulumi.Int(3),
//					Interval:     pulumi.String("OnceHourly"),
//					RunAfter:     pulumi.String("2024-03-22T09:00:00"),
//					RunUntil:     pulumi.String("2024-03-29T13:00:00"),
//				},
//				DeployNewReleaseAction: &octopusdeploy.ProjectScheduledTriggerDeployNewReleaseActionArgs{
//					DestinationEnvironmentId: pulumi.String("environments-123"),
//				},
//				Description: pulumi.String("This is a continuous daily schedule"),
//				ProjectId:   pulumi.String("projects-123"),
//				SpaceId:     pulumi.String("spaces-123"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewProjectScheduledTrigger(ctx, "deployLatestExample", &octopusdeploy.ProjectScheduledTriggerArgs{
//				CronExpressionSchedule: &octopusdeploy.ProjectScheduledTriggerCronExpressionScheduleArgs{
//					CronExpression: pulumi.String("0 0 06 * * Mon-Fri"),
//				},
//				DeployLatestReleaseAction: &octopusdeploy.ProjectScheduledTriggerDeployLatestReleaseActionArgs{
//					DestinationEnvironmentId: pulumi.String("environments-123"),
//					ShouldRedeploy:           pulumi.Bool(true),
//					SourceEnvironmentId:      pulumi.String("environments-321"),
//				},
//				ProjectId: pulumi.String("projects-123"),
//				SpaceId:   pulumi.String("spaces-123"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewProjectScheduledTrigger(ctx, "deployNewExample", &octopusdeploy.ProjectScheduledTriggerArgs{
//				CronExpressionSchedule: &octopusdeploy.ProjectScheduledTriggerCronExpressionScheduleArgs{
//					CronExpression: pulumi.String("0 0 06 * * Mon-Fri"),
//				},
//				DeployNewReleaseAction: &octopusdeploy.ProjectScheduledTriggerDeployNewReleaseActionArgs{
//					DestinationEnvironmentId: pulumi.String("environments-123"),
//				},
//				ProjectId: pulumi.String("projects-123"),
//				SpaceId:   pulumi.String("spaces-123"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewProjectScheduledTrigger(ctx, "runbookExample", &octopusdeploy.ProjectScheduledTriggerArgs{
//				CronExpressionSchedule: &octopusdeploy.ProjectScheduledTriggerCronExpressionScheduleArgs{
//					CronExpression: pulumi.String("0 0 06 * * Mon-Fri"),
//				},
//				Description: pulumi.String("This is a Cron schedule"),
//				ProjectId:   pulumi.String("projects-123"),
//				RunRunbookAction: &octopusdeploy.ProjectScheduledTriggerRunRunbookActionArgs{
//					RunbookId: pulumi.String("runbooks-123"),
//					TargetEnvironmentIds: pulumi.StringArray{
//						pulumi.String("environments-123"),
//						pulumi.String("environments-321"),
//					},
//				},
//				SpaceId: pulumi.String("spaces-123"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// ```sh
// $ pulumi import octopusdeploy:index/projectScheduledTrigger:ProjectScheduledTrigger [options] octopusdeploy_project_scheduled_trigger.<name> <trigger-id>
// ```
type ProjectScheduledTrigger struct {
	pulumi.CustomResourceState

	// The channel ID to use when creating the release. Will use the default channel if left blank.
	ChannelId pulumi.StringPtrOutput `pulumi:"channelId"`
	// The daily schedule for the trigger.
	ContinuousDailySchedule ProjectScheduledTriggerContinuousDailySchedulePtrOutput `pulumi:"continuousDailySchedule"`
	// The cron expression schedule for the trigger.
	CronExpressionSchedule ProjectScheduledTriggerCronExpressionSchedulePtrOutput `pulumi:"cronExpressionSchedule"`
	// The daily schedule for the trigger.
	DaysPerMonthSchedule ProjectScheduledTriggerDaysPerMonthSchedulePtrOutput `pulumi:"daysPerMonthSchedule"`
	// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
	DeployLatestReleaseAction ProjectScheduledTriggerDeployLatestReleaseActionPtrOutput `pulumi:"deployLatestReleaseAction"`
	// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
	DeployNewReleaseAction ProjectScheduledTriggerDeployNewReleaseActionPtrOutput `pulumi:"deployNewReleaseAction"`
	// A description of the trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether the trigger is disabled.
	IsDisabled pulumi.BoolPtrOutput `pulumi:"isDisabled"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The daily schedule for the trigger.
	OnceDailySchedule ProjectScheduledTriggerOnceDailySchedulePtrOutput `pulumi:"onceDailySchedule"`
	// The ID of the project to attach the trigger.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
	RunRunbookAction ProjectScheduledTriggerRunRunbookActionPtrOutput `pulumi:"runRunbookAction"`
	// The space ID where this trigger's project exists.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// The IDs of the tenants to deploy to.
	TenantIds pulumi.StringArrayOutput `pulumi:"tenantIds"`
	// The timezone for the trigger.
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewProjectScheduledTrigger registers a new resource with the given unique name, arguments, and options.
func NewProjectScheduledTrigger(ctx *pulumi.Context,
	name string, args *ProjectScheduledTriggerArgs, opts ...pulumi.ResourceOption) (*ProjectScheduledTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.SpaceId == nil {
		return nil, errors.New("invalid value for required argument 'SpaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectScheduledTrigger
	err := ctx.RegisterResource("octopusdeploy:index/projectScheduledTrigger:ProjectScheduledTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectScheduledTrigger gets an existing ProjectScheduledTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectScheduledTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectScheduledTriggerState, opts ...pulumi.ResourceOption) (*ProjectScheduledTrigger, error) {
	var resource ProjectScheduledTrigger
	err := ctx.ReadResource("octopusdeploy:index/projectScheduledTrigger:ProjectScheduledTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectScheduledTrigger resources.
type projectScheduledTriggerState struct {
	// The channel ID to use when creating the release. Will use the default channel if left blank.
	ChannelId *string `pulumi:"channelId"`
	// The daily schedule for the trigger.
	ContinuousDailySchedule *ProjectScheduledTriggerContinuousDailySchedule `pulumi:"continuousDailySchedule"`
	// The cron expression schedule for the trigger.
	CronExpressionSchedule *ProjectScheduledTriggerCronExpressionSchedule `pulumi:"cronExpressionSchedule"`
	// The daily schedule for the trigger.
	DaysPerMonthSchedule *ProjectScheduledTriggerDaysPerMonthSchedule `pulumi:"daysPerMonthSchedule"`
	// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
	DeployLatestReleaseAction *ProjectScheduledTriggerDeployLatestReleaseAction `pulumi:"deployLatestReleaseAction"`
	// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
	DeployNewReleaseAction *ProjectScheduledTriggerDeployNewReleaseAction `pulumi:"deployNewReleaseAction"`
	// A description of the trigger.
	Description *string `pulumi:"description"`
	// Indicates whether the trigger is disabled.
	IsDisabled *bool `pulumi:"isDisabled"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The daily schedule for the trigger.
	OnceDailySchedule *ProjectScheduledTriggerOnceDailySchedule `pulumi:"onceDailySchedule"`
	// The ID of the project to attach the trigger.
	ProjectId *string `pulumi:"projectId"`
	// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
	RunRunbookAction *ProjectScheduledTriggerRunRunbookAction `pulumi:"runRunbookAction"`
	// The space ID where this trigger's project exists.
	SpaceId *string `pulumi:"spaceId"`
	// The IDs of the tenants to deploy to.
	TenantIds []string `pulumi:"tenantIds"`
	// The timezone for the trigger.
	Timezone *string `pulumi:"timezone"`
}

type ProjectScheduledTriggerState struct {
	// The channel ID to use when creating the release. Will use the default channel if left blank.
	ChannelId pulumi.StringPtrInput
	// The daily schedule for the trigger.
	ContinuousDailySchedule ProjectScheduledTriggerContinuousDailySchedulePtrInput
	// The cron expression schedule for the trigger.
	CronExpressionSchedule ProjectScheduledTriggerCronExpressionSchedulePtrInput
	// The daily schedule for the trigger.
	DaysPerMonthSchedule ProjectScheduledTriggerDaysPerMonthSchedulePtrInput
	// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
	DeployLatestReleaseAction ProjectScheduledTriggerDeployLatestReleaseActionPtrInput
	// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
	DeployNewReleaseAction ProjectScheduledTriggerDeployNewReleaseActionPtrInput
	// A description of the trigger.
	Description pulumi.StringPtrInput
	// Indicates whether the trigger is disabled.
	IsDisabled pulumi.BoolPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The daily schedule for the trigger.
	OnceDailySchedule ProjectScheduledTriggerOnceDailySchedulePtrInput
	// The ID of the project to attach the trigger.
	ProjectId pulumi.StringPtrInput
	// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
	RunRunbookAction ProjectScheduledTriggerRunRunbookActionPtrInput
	// The space ID where this trigger's project exists.
	SpaceId pulumi.StringPtrInput
	// The IDs of the tenants to deploy to.
	TenantIds pulumi.StringArrayInput
	// The timezone for the trigger.
	Timezone pulumi.StringPtrInput
}

func (ProjectScheduledTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectScheduledTriggerState)(nil)).Elem()
}

type projectScheduledTriggerArgs struct {
	// The channel ID to use when creating the release. Will use the default channel if left blank.
	ChannelId *string `pulumi:"channelId"`
	// The daily schedule for the trigger.
	ContinuousDailySchedule *ProjectScheduledTriggerContinuousDailySchedule `pulumi:"continuousDailySchedule"`
	// The cron expression schedule for the trigger.
	CronExpressionSchedule *ProjectScheduledTriggerCronExpressionSchedule `pulumi:"cronExpressionSchedule"`
	// The daily schedule for the trigger.
	DaysPerMonthSchedule *ProjectScheduledTriggerDaysPerMonthSchedule `pulumi:"daysPerMonthSchedule"`
	// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
	DeployLatestReleaseAction *ProjectScheduledTriggerDeployLatestReleaseAction `pulumi:"deployLatestReleaseAction"`
	// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
	DeployNewReleaseAction *ProjectScheduledTriggerDeployNewReleaseAction `pulumi:"deployNewReleaseAction"`
	// A description of the trigger.
	Description *string `pulumi:"description"`
	// Indicates whether the trigger is disabled.
	IsDisabled *bool `pulumi:"isDisabled"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The daily schedule for the trigger.
	OnceDailySchedule *ProjectScheduledTriggerOnceDailySchedule `pulumi:"onceDailySchedule"`
	// The ID of the project to attach the trigger.
	ProjectId string `pulumi:"projectId"`
	// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
	RunRunbookAction *ProjectScheduledTriggerRunRunbookAction `pulumi:"runRunbookAction"`
	// The space ID where this trigger's project exists.
	SpaceId string `pulumi:"spaceId"`
	// The IDs of the tenants to deploy to.
	TenantIds []string `pulumi:"tenantIds"`
	// The timezone for the trigger.
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a ProjectScheduledTrigger resource.
type ProjectScheduledTriggerArgs struct {
	// The channel ID to use when creating the release. Will use the default channel if left blank.
	ChannelId pulumi.StringPtrInput
	// The daily schedule for the trigger.
	ContinuousDailySchedule ProjectScheduledTriggerContinuousDailySchedulePtrInput
	// The cron expression schedule for the trigger.
	CronExpressionSchedule ProjectScheduledTriggerCronExpressionSchedulePtrInput
	// The daily schedule for the trigger.
	DaysPerMonthSchedule ProjectScheduledTriggerDaysPerMonthSchedulePtrInput
	// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
	DeployLatestReleaseAction ProjectScheduledTriggerDeployLatestReleaseActionPtrInput
	// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
	DeployNewReleaseAction ProjectScheduledTriggerDeployNewReleaseActionPtrInput
	// A description of the trigger.
	Description pulumi.StringPtrInput
	// Indicates whether the trigger is disabled.
	IsDisabled pulumi.BoolPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The daily schedule for the trigger.
	OnceDailySchedule ProjectScheduledTriggerOnceDailySchedulePtrInput
	// The ID of the project to attach the trigger.
	ProjectId pulumi.StringInput
	// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
	RunRunbookAction ProjectScheduledTriggerRunRunbookActionPtrInput
	// The space ID where this trigger's project exists.
	SpaceId pulumi.StringInput
	// The IDs of the tenants to deploy to.
	TenantIds pulumi.StringArrayInput
	// The timezone for the trigger.
	Timezone pulumi.StringPtrInput
}

func (ProjectScheduledTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectScheduledTriggerArgs)(nil)).Elem()
}

type ProjectScheduledTriggerInput interface {
	pulumi.Input

	ToProjectScheduledTriggerOutput() ProjectScheduledTriggerOutput
	ToProjectScheduledTriggerOutputWithContext(ctx context.Context) ProjectScheduledTriggerOutput
}

func (*ProjectScheduledTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectScheduledTrigger)(nil)).Elem()
}

func (i *ProjectScheduledTrigger) ToProjectScheduledTriggerOutput() ProjectScheduledTriggerOutput {
	return i.ToProjectScheduledTriggerOutputWithContext(context.Background())
}

func (i *ProjectScheduledTrigger) ToProjectScheduledTriggerOutputWithContext(ctx context.Context) ProjectScheduledTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectScheduledTriggerOutput)
}

// ProjectScheduledTriggerArrayInput is an input type that accepts ProjectScheduledTriggerArray and ProjectScheduledTriggerArrayOutput values.
// You can construct a concrete instance of `ProjectScheduledTriggerArrayInput` via:
//
//	ProjectScheduledTriggerArray{ ProjectScheduledTriggerArgs{...} }
type ProjectScheduledTriggerArrayInput interface {
	pulumi.Input

	ToProjectScheduledTriggerArrayOutput() ProjectScheduledTriggerArrayOutput
	ToProjectScheduledTriggerArrayOutputWithContext(context.Context) ProjectScheduledTriggerArrayOutput
}

type ProjectScheduledTriggerArray []ProjectScheduledTriggerInput

func (ProjectScheduledTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectScheduledTrigger)(nil)).Elem()
}

func (i ProjectScheduledTriggerArray) ToProjectScheduledTriggerArrayOutput() ProjectScheduledTriggerArrayOutput {
	return i.ToProjectScheduledTriggerArrayOutputWithContext(context.Background())
}

func (i ProjectScheduledTriggerArray) ToProjectScheduledTriggerArrayOutputWithContext(ctx context.Context) ProjectScheduledTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectScheduledTriggerArrayOutput)
}

// ProjectScheduledTriggerMapInput is an input type that accepts ProjectScheduledTriggerMap and ProjectScheduledTriggerMapOutput values.
// You can construct a concrete instance of `ProjectScheduledTriggerMapInput` via:
//
//	ProjectScheduledTriggerMap{ "key": ProjectScheduledTriggerArgs{...} }
type ProjectScheduledTriggerMapInput interface {
	pulumi.Input

	ToProjectScheduledTriggerMapOutput() ProjectScheduledTriggerMapOutput
	ToProjectScheduledTriggerMapOutputWithContext(context.Context) ProjectScheduledTriggerMapOutput
}

type ProjectScheduledTriggerMap map[string]ProjectScheduledTriggerInput

func (ProjectScheduledTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectScheduledTrigger)(nil)).Elem()
}

func (i ProjectScheduledTriggerMap) ToProjectScheduledTriggerMapOutput() ProjectScheduledTriggerMapOutput {
	return i.ToProjectScheduledTriggerMapOutputWithContext(context.Background())
}

func (i ProjectScheduledTriggerMap) ToProjectScheduledTriggerMapOutputWithContext(ctx context.Context) ProjectScheduledTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectScheduledTriggerMapOutput)
}

type ProjectScheduledTriggerOutput struct{ *pulumi.OutputState }

func (ProjectScheduledTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectScheduledTrigger)(nil)).Elem()
}

func (o ProjectScheduledTriggerOutput) ToProjectScheduledTriggerOutput() ProjectScheduledTriggerOutput {
	return o
}

func (o ProjectScheduledTriggerOutput) ToProjectScheduledTriggerOutputWithContext(ctx context.Context) ProjectScheduledTriggerOutput {
	return o
}

// The channel ID to use when creating the release. Will use the default channel if left blank.
func (o ProjectScheduledTriggerOutput) ChannelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.StringPtrOutput { return v.ChannelId }).(pulumi.StringPtrOutput)
}

// The daily schedule for the trigger.
func (o ProjectScheduledTriggerOutput) ContinuousDailySchedule() ProjectScheduledTriggerContinuousDailySchedulePtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) ProjectScheduledTriggerContinuousDailySchedulePtrOutput {
		return v.ContinuousDailySchedule
	}).(ProjectScheduledTriggerContinuousDailySchedulePtrOutput)
}

// The cron expression schedule for the trigger.
func (o ProjectScheduledTriggerOutput) CronExpressionSchedule() ProjectScheduledTriggerCronExpressionSchedulePtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) ProjectScheduledTriggerCronExpressionSchedulePtrOutput {
		return v.CronExpressionSchedule
	}).(ProjectScheduledTriggerCronExpressionSchedulePtrOutput)
}

// The daily schedule for the trigger.
func (o ProjectScheduledTriggerOutput) DaysPerMonthSchedule() ProjectScheduledTriggerDaysPerMonthSchedulePtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) ProjectScheduledTriggerDaysPerMonthSchedulePtrOutput {
		return v.DaysPerMonthSchedule
	}).(ProjectScheduledTriggerDaysPerMonthSchedulePtrOutput)
}

// Configuration for deploying the latest release. Can not be used with 'deploy*new*release*action' or 'run*runbook*action'.
func (o ProjectScheduledTriggerOutput) DeployLatestReleaseAction() ProjectScheduledTriggerDeployLatestReleaseActionPtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) ProjectScheduledTriggerDeployLatestReleaseActionPtrOutput {
		return v.DeployLatestReleaseAction
	}).(ProjectScheduledTriggerDeployLatestReleaseActionPtrOutput)
}

// Configuration for deploying a new release. Can not be used with 'deploy*latest*release*action' or 'run*runbook*action'.
func (o ProjectScheduledTriggerOutput) DeployNewReleaseAction() ProjectScheduledTriggerDeployNewReleaseActionPtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) ProjectScheduledTriggerDeployNewReleaseActionPtrOutput {
		return v.DeployNewReleaseAction
	}).(ProjectScheduledTriggerDeployNewReleaseActionPtrOutput)
}

// A description of the trigger.
func (o ProjectScheduledTriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether the trigger is disabled.
func (o ProjectScheduledTriggerOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// The name of this resource.
func (o ProjectScheduledTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The daily schedule for the trigger.
func (o ProjectScheduledTriggerOutput) OnceDailySchedule() ProjectScheduledTriggerOnceDailySchedulePtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) ProjectScheduledTriggerOnceDailySchedulePtrOutput {
		return v.OnceDailySchedule
	}).(ProjectScheduledTriggerOnceDailySchedulePtrOutput)
}

// The ID of the project to attach the trigger.
func (o ProjectScheduledTriggerOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Configuration for running a runbook. Can not be used with 'deploy*latest*release*action' or 'deploy*new*release*action'.
func (o ProjectScheduledTriggerOutput) RunRunbookAction() ProjectScheduledTriggerRunRunbookActionPtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) ProjectScheduledTriggerRunRunbookActionPtrOutput {
		return v.RunRunbookAction
	}).(ProjectScheduledTriggerRunRunbookActionPtrOutput)
}

// The space ID where this trigger's project exists.
func (o ProjectScheduledTriggerOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The IDs of the tenants to deploy to.
func (o ProjectScheduledTriggerOutput) TenantIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.StringArrayOutput { return v.TenantIds }).(pulumi.StringArrayOutput)
}

// The timezone for the trigger.
func (o ProjectScheduledTriggerOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectScheduledTrigger) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

type ProjectScheduledTriggerArrayOutput struct{ *pulumi.OutputState }

func (ProjectScheduledTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectScheduledTrigger)(nil)).Elem()
}

func (o ProjectScheduledTriggerArrayOutput) ToProjectScheduledTriggerArrayOutput() ProjectScheduledTriggerArrayOutput {
	return o
}

func (o ProjectScheduledTriggerArrayOutput) ToProjectScheduledTriggerArrayOutputWithContext(ctx context.Context) ProjectScheduledTriggerArrayOutput {
	return o
}

func (o ProjectScheduledTriggerArrayOutput) Index(i pulumi.IntInput) ProjectScheduledTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectScheduledTrigger {
		return vs[0].([]*ProjectScheduledTrigger)[vs[1].(int)]
	}).(ProjectScheduledTriggerOutput)
}

type ProjectScheduledTriggerMapOutput struct{ *pulumi.OutputState }

func (ProjectScheduledTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectScheduledTrigger)(nil)).Elem()
}

func (o ProjectScheduledTriggerMapOutput) ToProjectScheduledTriggerMapOutput() ProjectScheduledTriggerMapOutput {
	return o
}

func (o ProjectScheduledTriggerMapOutput) ToProjectScheduledTriggerMapOutputWithContext(ctx context.Context) ProjectScheduledTriggerMapOutput {
	return o
}

func (o ProjectScheduledTriggerMapOutput) MapIndex(k pulumi.StringInput) ProjectScheduledTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectScheduledTrigger {
		return vs[0].(map[string]*ProjectScheduledTrigger)[vs[1].(string)]
	}).(ProjectScheduledTriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectScheduledTriggerInput)(nil)).Elem(), &ProjectScheduledTrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectScheduledTriggerArrayInput)(nil)).Elem(), ProjectScheduledTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectScheduledTriggerMapInput)(nil)).Elem(), ProjectScheduledTriggerMap{})
	pulumi.RegisterOutputType(ProjectScheduledTriggerOutput{})
	pulumi.RegisterOutputType(ProjectScheduledTriggerArrayOutput{})
	pulumi.RegisterOutputType(ProjectScheduledTriggerMapOutput{})
}
