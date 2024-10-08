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

// This resource manages cloud region deployment targets in Octopus Deploy.
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
//			_, err := octopusdeploy.NewCloudRegionDeploymentTarget(ctx, "example", &octopusdeploy.CloudRegionDeploymentTargetArgs{
//				DefaultWorkerPoolId: pulumi.String("WorkerPools-123"),
//				Environments: pulumi.StringArray{
//					pulumi.String("Environments-123"),
//					pulumi.String("Environment-321"),
//				},
//				Roles: pulumi.StringArray{
//					pulumi.String("Development Team"),
//					pulumi.String("System Administrators"),
//				},
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
// $ pulumi import octopusdeploy:index/cloudRegionDeploymentTarget:CloudRegionDeploymentTarget [options] octopusdeploy_cloud_region_deployment_target.<name> <machine-id>
// ```
type CloudRegionDeploymentTarget struct {
	pulumi.CustomResourceState

	DefaultWorkerPoolId pulumi.StringPtrOutput                         `pulumi:"defaultWorkerPoolId"`
	Endpoints           CloudRegionDeploymentTargetEndpointArrayOutput `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments      pulumi.StringArrayOutput `pulumi:"environments"`
	HasLatestCalamari pulumi.BoolOutput        `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringOutput `pulumi:"healthStatus"`
	IsDisabled      pulumi.BoolOutput   `pulumi:"isDisabled"`
	IsInProcess     pulumi.BoolOutput   `pulumi:"isInProcess"`
	MachinePolicyId pulumi.StringOutput `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name            pulumi.StringOutput      `pulumi:"name"`
	OperatingSystem pulumi.StringOutput      `pulumi:"operatingSystem"`
	Roles           pulumi.StringArrayOutput `pulumi:"roles"`
	ShellName       pulumi.StringOutput      `pulumi:"shellName"`
	ShellVersion    pulumi.StringOutput      `pulumi:"shellVersion"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringOutput `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary pulumi.StringOutput `pulumi:"statusSummary"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringOutput `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants    pulumi.StringArrayOutput `pulumi:"tenants"`
	Thumbprint pulumi.StringOutput      `pulumi:"thumbprint"`
	Uri        pulumi.StringOutput      `pulumi:"uri"`
}

// NewCloudRegionDeploymentTarget registers a new resource with the given unique name, arguments, and options.
func NewCloudRegionDeploymentTarget(ctx *pulumi.Context,
	name string, args *CloudRegionDeploymentTargetArgs, opts ...pulumi.ResourceOption) (*CloudRegionDeploymentTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudRegionDeploymentTarget
	err := ctx.RegisterResource("octopusdeploy:index/cloudRegionDeploymentTarget:CloudRegionDeploymentTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudRegionDeploymentTarget gets an existing CloudRegionDeploymentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudRegionDeploymentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudRegionDeploymentTargetState, opts ...pulumi.ResourceOption) (*CloudRegionDeploymentTarget, error) {
	var resource CloudRegionDeploymentTarget
	err := ctx.ReadResource("octopusdeploy:index/cloudRegionDeploymentTarget:CloudRegionDeploymentTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudRegionDeploymentTarget resources.
type cloudRegionDeploymentTargetState struct {
	DefaultWorkerPoolId *string                               `pulumi:"defaultWorkerPoolId"`
	Endpoints           []CloudRegionDeploymentTargetEndpoint `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments      []string `pulumi:"environments"`
	HasLatestCalamari *bool    `pulumi:"hasLatestCalamari"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	IsInProcess     *bool   `pulumi:"isInProcess"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name            *string  `pulumi:"name"`
	OperatingSystem *string  `pulumi:"operatingSystem"`
	Roles           []string `pulumi:"roles"`
	ShellName       *string  `pulumi:"shellName"`
	ShellVersion    *string  `pulumi:"shellVersion"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status *string `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary *string `pulumi:"statusSummary"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants    []string `pulumi:"tenants"`
	Thumbprint *string  `pulumi:"thumbprint"`
	Uri        *string  `pulumi:"uri"`
}

type CloudRegionDeploymentTargetState struct {
	DefaultWorkerPoolId pulumi.StringPtrInput
	Endpoints           CloudRegionDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments      pulumi.StringArrayInput
	HasLatestCalamari pulumi.BoolPtrInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	IsInProcess     pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name            pulumi.StringPtrInput
	OperatingSystem pulumi.StringPtrInput
	Roles           pulumi.StringArrayInput
	ShellName       pulumi.StringPtrInput
	ShellVersion    pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringPtrInput
	// A summary elaborating on the status of this resource.
	StatusSummary pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants    pulumi.StringArrayInput
	Thumbprint pulumi.StringPtrInput
	Uri        pulumi.StringPtrInput
}

func (CloudRegionDeploymentTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRegionDeploymentTargetState)(nil)).Elem()
}

type cloudRegionDeploymentTargetArgs struct {
	DefaultWorkerPoolId *string                               `pulumi:"defaultWorkerPoolId"`
	Endpoints           []CloudRegionDeploymentTargetEndpoint `pulumi:"endpoints"`
	// A list of environment IDs associated with this resource.
	Environments []string `pulumi:"environments"`
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    *string `pulumi:"healthStatus"`
	IsDisabled      *bool   `pulumi:"isDisabled"`
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name            *string  `pulumi:"name"`
	OperatingSystem *string  `pulumi:"operatingSystem"`
	Roles           []string `pulumi:"roles"`
	ShellName       *string  `pulumi:"shellName"`
	ShellVersion    *string  `pulumi:"shellVersion"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status *string `pulumi:"status"`
	// A summary elaborating on the status of this resource.
	StatusSummary *string `pulumi:"statusSummary"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants    []string `pulumi:"tenants"`
	Thumbprint *string  `pulumi:"thumbprint"`
	Uri        *string  `pulumi:"uri"`
}

// The set of arguments for constructing a CloudRegionDeploymentTarget resource.
type CloudRegionDeploymentTargetArgs struct {
	DefaultWorkerPoolId pulumi.StringPtrInput
	Endpoints           CloudRegionDeploymentTargetEndpointArrayInput
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayInput
	// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatus    pulumi.StringPtrInput
	IsDisabled      pulumi.BoolPtrInput
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name            pulumi.StringPtrInput
	OperatingSystem pulumi.StringPtrInput
	Roles           pulumi.StringArrayInput
	ShellName       pulumi.StringPtrInput
	ShellVersion    pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
	Status pulumi.StringPtrInput
	// A summary elaborating on the status of this resource.
	StatusSummary pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants    pulumi.StringArrayInput
	Thumbprint pulumi.StringPtrInput
	Uri        pulumi.StringPtrInput
}

func (CloudRegionDeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRegionDeploymentTargetArgs)(nil)).Elem()
}

type CloudRegionDeploymentTargetInput interface {
	pulumi.Input

	ToCloudRegionDeploymentTargetOutput() CloudRegionDeploymentTargetOutput
	ToCloudRegionDeploymentTargetOutputWithContext(ctx context.Context) CloudRegionDeploymentTargetOutput
}

func (*CloudRegionDeploymentTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRegionDeploymentTarget)(nil)).Elem()
}

func (i *CloudRegionDeploymentTarget) ToCloudRegionDeploymentTargetOutput() CloudRegionDeploymentTargetOutput {
	return i.ToCloudRegionDeploymentTargetOutputWithContext(context.Background())
}

func (i *CloudRegionDeploymentTarget) ToCloudRegionDeploymentTargetOutputWithContext(ctx context.Context) CloudRegionDeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudRegionDeploymentTargetOutput)
}

// CloudRegionDeploymentTargetArrayInput is an input type that accepts CloudRegionDeploymentTargetArray and CloudRegionDeploymentTargetArrayOutput values.
// You can construct a concrete instance of `CloudRegionDeploymentTargetArrayInput` via:
//
//	CloudRegionDeploymentTargetArray{ CloudRegionDeploymentTargetArgs{...} }
type CloudRegionDeploymentTargetArrayInput interface {
	pulumi.Input

	ToCloudRegionDeploymentTargetArrayOutput() CloudRegionDeploymentTargetArrayOutput
	ToCloudRegionDeploymentTargetArrayOutputWithContext(context.Context) CloudRegionDeploymentTargetArrayOutput
}

type CloudRegionDeploymentTargetArray []CloudRegionDeploymentTargetInput

func (CloudRegionDeploymentTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudRegionDeploymentTarget)(nil)).Elem()
}

func (i CloudRegionDeploymentTargetArray) ToCloudRegionDeploymentTargetArrayOutput() CloudRegionDeploymentTargetArrayOutput {
	return i.ToCloudRegionDeploymentTargetArrayOutputWithContext(context.Background())
}

func (i CloudRegionDeploymentTargetArray) ToCloudRegionDeploymentTargetArrayOutputWithContext(ctx context.Context) CloudRegionDeploymentTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudRegionDeploymentTargetArrayOutput)
}

// CloudRegionDeploymentTargetMapInput is an input type that accepts CloudRegionDeploymentTargetMap and CloudRegionDeploymentTargetMapOutput values.
// You can construct a concrete instance of `CloudRegionDeploymentTargetMapInput` via:
//
//	CloudRegionDeploymentTargetMap{ "key": CloudRegionDeploymentTargetArgs{...} }
type CloudRegionDeploymentTargetMapInput interface {
	pulumi.Input

	ToCloudRegionDeploymentTargetMapOutput() CloudRegionDeploymentTargetMapOutput
	ToCloudRegionDeploymentTargetMapOutputWithContext(context.Context) CloudRegionDeploymentTargetMapOutput
}

type CloudRegionDeploymentTargetMap map[string]CloudRegionDeploymentTargetInput

func (CloudRegionDeploymentTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudRegionDeploymentTarget)(nil)).Elem()
}

func (i CloudRegionDeploymentTargetMap) ToCloudRegionDeploymentTargetMapOutput() CloudRegionDeploymentTargetMapOutput {
	return i.ToCloudRegionDeploymentTargetMapOutputWithContext(context.Background())
}

func (i CloudRegionDeploymentTargetMap) ToCloudRegionDeploymentTargetMapOutputWithContext(ctx context.Context) CloudRegionDeploymentTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudRegionDeploymentTargetMapOutput)
}

type CloudRegionDeploymentTargetOutput struct{ *pulumi.OutputState }

func (CloudRegionDeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRegionDeploymentTarget)(nil)).Elem()
}

func (o CloudRegionDeploymentTargetOutput) ToCloudRegionDeploymentTargetOutput() CloudRegionDeploymentTargetOutput {
	return o
}

func (o CloudRegionDeploymentTargetOutput) ToCloudRegionDeploymentTargetOutputWithContext(ctx context.Context) CloudRegionDeploymentTargetOutput {
	return o
}

func (o CloudRegionDeploymentTargetOutput) DefaultWorkerPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringPtrOutput { return v.DefaultWorkerPoolId }).(pulumi.StringPtrOutput)
}

func (o CloudRegionDeploymentTargetOutput) Endpoints() CloudRegionDeploymentTargetEndpointArrayOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) CloudRegionDeploymentTargetEndpointArrayOutput {
		return v.Endpoints
	}).(CloudRegionDeploymentTargetEndpointArrayOutput)
}

// A list of environment IDs associated with this resource.
func (o CloudRegionDeploymentTargetOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o CloudRegionDeploymentTargetOutput) HasLatestCalamari() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.BoolOutput { return v.HasLatestCalamari }).(pulumi.BoolOutput)
}

// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o CloudRegionDeploymentTargetOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.HealthStatus }).(pulumi.StringOutput)
}

func (o CloudRegionDeploymentTargetOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o CloudRegionDeploymentTargetOutput) IsInProcess() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.BoolOutput { return v.IsInProcess }).(pulumi.BoolOutput)
}

func (o CloudRegionDeploymentTargetOutput) MachinePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.MachinePolicyId }).(pulumi.StringOutput)
}

// The name of this resource.
func (o CloudRegionDeploymentTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudRegionDeploymentTargetOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.OperatingSystem }).(pulumi.StringOutput)
}

func (o CloudRegionDeploymentTargetOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o CloudRegionDeploymentTargetOutput) ShellName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.ShellName }).(pulumi.StringOutput)
}

func (o CloudRegionDeploymentTargetOutput) ShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.ShellVersion }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o CloudRegionDeploymentTargetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
func (o CloudRegionDeploymentTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A summary elaborating on the status of this resource.
func (o CloudRegionDeploymentTargetOutput) StatusSummary() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.StatusSummary }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o CloudRegionDeploymentTargetOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o CloudRegionDeploymentTargetOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.TenantedDeploymentParticipation }).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o CloudRegionDeploymentTargetOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

func (o CloudRegionDeploymentTargetOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CloudRegionDeploymentTargetOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRegionDeploymentTarget) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

type CloudRegionDeploymentTargetArrayOutput struct{ *pulumi.OutputState }

func (CloudRegionDeploymentTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudRegionDeploymentTarget)(nil)).Elem()
}

func (o CloudRegionDeploymentTargetArrayOutput) ToCloudRegionDeploymentTargetArrayOutput() CloudRegionDeploymentTargetArrayOutput {
	return o
}

func (o CloudRegionDeploymentTargetArrayOutput) ToCloudRegionDeploymentTargetArrayOutputWithContext(ctx context.Context) CloudRegionDeploymentTargetArrayOutput {
	return o
}

func (o CloudRegionDeploymentTargetArrayOutput) Index(i pulumi.IntInput) CloudRegionDeploymentTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudRegionDeploymentTarget {
		return vs[0].([]*CloudRegionDeploymentTarget)[vs[1].(int)]
	}).(CloudRegionDeploymentTargetOutput)
}

type CloudRegionDeploymentTargetMapOutput struct{ *pulumi.OutputState }

func (CloudRegionDeploymentTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudRegionDeploymentTarget)(nil)).Elem()
}

func (o CloudRegionDeploymentTargetMapOutput) ToCloudRegionDeploymentTargetMapOutput() CloudRegionDeploymentTargetMapOutput {
	return o
}

func (o CloudRegionDeploymentTargetMapOutput) ToCloudRegionDeploymentTargetMapOutputWithContext(ctx context.Context) CloudRegionDeploymentTargetMapOutput {
	return o
}

func (o CloudRegionDeploymentTargetMapOutput) MapIndex(k pulumi.StringInput) CloudRegionDeploymentTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudRegionDeploymentTarget {
		return vs[0].(map[string]*CloudRegionDeploymentTarget)[vs[1].(string)]
	}).(CloudRegionDeploymentTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudRegionDeploymentTargetInput)(nil)).Elem(), &CloudRegionDeploymentTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudRegionDeploymentTargetArrayInput)(nil)).Elem(), CloudRegionDeploymentTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudRegionDeploymentTargetMapInput)(nil)).Elem(), CloudRegionDeploymentTargetMap{})
	pulumi.RegisterOutputType(CloudRegionDeploymentTargetOutput{})
	pulumi.RegisterOutputType(CloudRegionDeploymentTargetArrayOutput{})
	pulumi.RegisterOutputType(CloudRegionDeploymentTargetMapOutput{})
}
