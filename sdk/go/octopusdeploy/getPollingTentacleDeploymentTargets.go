// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing polling tentacle deployment targets.
func GetPollingTentacleDeploymentTargets(ctx *pulumi.Context, args *GetPollingTentacleDeploymentTargetsArgs, opts ...pulumi.InvokeOption) (*GetPollingTentacleDeploymentTargetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPollingTentacleDeploymentTargetsResult
	err := ctx.Invoke("octopusdeploy:index/getPollingTentacleDeploymentTargets:getPollingTentacleDeploymentTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPollingTentacleDeploymentTargets.
type GetPollingTentacleDeploymentTargetsArgs struct {
	// A filter to search by deployment ID.
	DeploymentId *string  `pulumi:"deploymentId"`
	Environments []string `pulumi:"environments"`
	// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatuses []string `pulumi:"healthStatuses"`
	// A filter to search by a list of IDs.
	Ids        []string `pulumi:"ids"`
	IsDisabled *bool    `pulumi:"isDisabled"`
	Name       *string  `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName *string  `pulumi:"partialName"`
	Roles       []string `pulumi:"roles"`
	// A list of shell names to match in the query and/or search
	ShellNames []string `pulumi:"shellNames"`
	// A filter to specify the number of items to skip in the response.
	Skip    *int    `pulumi:"skip"`
	SpaceId *string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take       *int     `pulumi:"take"`
	TenantTags []string `pulumi:"tenantTags"`
	Tenants    []string `pulumi:"tenants"`
	Thumbprint *string  `pulumi:"thumbprint"`
}

// A collection of values returned by getPollingTentacleDeploymentTargets.
type GetPollingTentacleDeploymentTargetsResult struct {
	// A filter to search by deployment ID.
	DeploymentId *string `pulumi:"deploymentId"`
	// A filter to search by a list of environment IDs.
	Environments []string `pulumi:"environments"`
	// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatuses []string `pulumi:"healthStatuses"`
	// An auto-generated identifier that includes the timestamp when this data source was last modified.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by the disabled status of a resource.
	IsDisabled *bool `pulumi:"isDisabled"`
	// A filter to search by name.
	Name *string `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A list of polling tentacle deployment targets that match the filter(s).
	PollingTentacleDeploymentTargets []GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTarget `pulumi:"pollingTentacleDeploymentTargets"`
	// A filter to search by a list of role IDs.
	Roles []string `pulumi:"roles"`
	// A list of shell names to match in the query and/or search
	ShellNames []string `pulumi:"shellNames"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// The space ID associated with this resource.
	SpaceId string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A filter to search by a list of tenant tags.
	TenantTags []string `pulumi:"tenantTags"`
	// A filter to search by a list of tenant IDs.
	Tenants []string `pulumi:"tenants"`
	// The thumbprint of the deployment target to match in the query and/or search
	Thumbprint *string `pulumi:"thumbprint"`
}

func GetPollingTentacleDeploymentTargetsOutput(ctx *pulumi.Context, args GetPollingTentacleDeploymentTargetsOutputArgs, opts ...pulumi.InvokeOption) GetPollingTentacleDeploymentTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPollingTentacleDeploymentTargetsResult, error) {
			args := v.(GetPollingTentacleDeploymentTargetsArgs)
			r, err := GetPollingTentacleDeploymentTargets(ctx, &args, opts...)
			var s GetPollingTentacleDeploymentTargetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPollingTentacleDeploymentTargetsResultOutput)
}

// A collection of arguments for invoking getPollingTentacleDeploymentTargets.
type GetPollingTentacleDeploymentTargetsOutputArgs struct {
	// A filter to search by deployment ID.
	DeploymentId pulumi.StringPtrInput   `pulumi:"deploymentId"`
	Environments pulumi.StringArrayInput `pulumi:"environments"`
	// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
	HealthStatuses pulumi.StringArrayInput `pulumi:"healthStatuses"`
	// A filter to search by a list of IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	IsDisabled pulumi.BoolPtrInput     `pulumi:"isDisabled"`
	Name       pulumi.StringPtrInput   `pulumi:"name"`
	// A filter to search by the partial match of a name.
	PartialName pulumi.StringPtrInput   `pulumi:"partialName"`
	Roles       pulumi.StringArrayInput `pulumi:"roles"`
	// A list of shell names to match in the query and/or search
	ShellNames pulumi.StringArrayInput `pulumi:"shellNames"`
	// A filter to specify the number of items to skip in the response.
	Skip    pulumi.IntPtrInput    `pulumi:"skip"`
	SpaceId pulumi.StringPtrInput `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take       pulumi.IntPtrInput      `pulumi:"take"`
	TenantTags pulumi.StringArrayInput `pulumi:"tenantTags"`
	Tenants    pulumi.StringArrayInput `pulumi:"tenants"`
	Thumbprint pulumi.StringPtrInput   `pulumi:"thumbprint"`
}

func (GetPollingTentacleDeploymentTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPollingTentacleDeploymentTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getPollingTentacleDeploymentTargets.
type GetPollingTentacleDeploymentTargetsResultOutput struct{ *pulumi.OutputState }

func (GetPollingTentacleDeploymentTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPollingTentacleDeploymentTargetsResult)(nil)).Elem()
}

func (o GetPollingTentacleDeploymentTargetsResultOutput) ToGetPollingTentacleDeploymentTargetsResultOutput() GetPollingTentacleDeploymentTargetsResultOutput {
	return o
}

func (o GetPollingTentacleDeploymentTargetsResultOutput) ToGetPollingTentacleDeploymentTargetsResultOutputWithContext(ctx context.Context) GetPollingTentacleDeploymentTargetsResultOutput {
	return o
}

// A filter to search by deployment ID.
func (o GetPollingTentacleDeploymentTargetsResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of environment IDs.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []string { return v.Environments }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o GetPollingTentacleDeploymentTargetsResultOutput) HealthStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []string { return v.HealthStatuses }).(pulumi.StringArrayOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by the disabled status of a resource.
func (o GetPollingTentacleDeploymentTargetsResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// A filter to search by name.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A filter to search by the partial match of a name.
func (o GetPollingTentacleDeploymentTargetsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A list of polling tentacle deployment targets that match the filter(s).
func (o GetPollingTentacleDeploymentTargetsResultOutput) PollingTentacleDeploymentTargets() GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTarget {
		return v.PollingTentacleDeploymentTargets
	}).(GetPollingTentacleDeploymentTargetsPollingTentacleDeploymentTargetArrayOutput)
}

// A filter to search by a list of role IDs.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// A list of shell names to match in the query and/or search
func (o GetPollingTentacleDeploymentTargetsResultOutput) ShellNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []string { return v.ShellNames }).(pulumi.StringArrayOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// The space ID associated with this resource.
func (o GetPollingTentacleDeploymentTargetsResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A filter to search by a list of tenant tags.
func (o GetPollingTentacleDeploymentTargetsResultOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []string { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of tenant IDs.
func (o GetPollingTentacleDeploymentTargetsResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The thumbprint of the deployment target to match in the query and/or search
func (o GetPollingTentacleDeploymentTargetsResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPollingTentacleDeploymentTargetsResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPollingTentacleDeploymentTargetsResultOutput{})
}
