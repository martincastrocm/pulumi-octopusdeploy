// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing offline package drop deployment targets.
func GetOfflinePackageDropDeploymentTargets(ctx *pulumi.Context, args *GetOfflinePackageDropDeploymentTargetsArgs, opts ...pulumi.InvokeOption) (*GetOfflinePackageDropDeploymentTargetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOfflinePackageDropDeploymentTargetsResult
	err := ctx.Invoke("octopusdeploy:index/getOfflinePackageDropDeploymentTargets:getOfflinePackageDropDeploymentTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOfflinePackageDropDeploymentTargets.
type GetOfflinePackageDropDeploymentTargetsArgs struct {
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

// A collection of values returned by getOfflinePackageDropDeploymentTargets.
type GetOfflinePackageDropDeploymentTargetsResult struct {
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
	// A list of offline package drop deployment targets that match the filter(s).
	OfflinePackageDropDeploymentTargets []GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTarget `pulumi:"offlinePackageDropDeploymentTargets"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
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

func GetOfflinePackageDropDeploymentTargetsOutput(ctx *pulumi.Context, args GetOfflinePackageDropDeploymentTargetsOutputArgs, opts ...pulumi.InvokeOption) GetOfflinePackageDropDeploymentTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOfflinePackageDropDeploymentTargetsResult, error) {
			args := v.(GetOfflinePackageDropDeploymentTargetsArgs)
			r, err := GetOfflinePackageDropDeploymentTargets(ctx, &args, opts...)
			var s GetOfflinePackageDropDeploymentTargetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOfflinePackageDropDeploymentTargetsResultOutput)
}

// A collection of arguments for invoking getOfflinePackageDropDeploymentTargets.
type GetOfflinePackageDropDeploymentTargetsOutputArgs struct {
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

func (GetOfflinePackageDropDeploymentTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOfflinePackageDropDeploymentTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getOfflinePackageDropDeploymentTargets.
type GetOfflinePackageDropDeploymentTargetsResultOutput struct{ *pulumi.OutputState }

func (GetOfflinePackageDropDeploymentTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOfflinePackageDropDeploymentTargetsResult)(nil)).Elem()
}

func (o GetOfflinePackageDropDeploymentTargetsResultOutput) ToGetOfflinePackageDropDeploymentTargetsResultOutput() GetOfflinePackageDropDeploymentTargetsResultOutput {
	return o
}

func (o GetOfflinePackageDropDeploymentTargetsResultOutput) ToGetOfflinePackageDropDeploymentTargetsResultOutputWithContext(ctx context.Context) GetOfflinePackageDropDeploymentTargetsResultOutput {
	return o
}

// A filter to search by deployment ID.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of environment IDs.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []string { return v.Environments }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) HealthStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []string { return v.HealthStatuses }).(pulumi.StringArrayOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by the disabled status of a resource.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// A filter to search by name.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of offline package drop deployment targets that match the filter(s).
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) OfflinePackageDropDeploymentTargets() GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTargetArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTarget {
		return v.OfflinePackageDropDeploymentTargets
	}).(GetOfflinePackageDropDeploymentTargetsOfflinePackageDropDeploymentTargetArrayOutput)
}

// A filter to search by the partial match of a name.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of role IDs.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// A list of shell names to match in the query and/or search
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) ShellNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []string { return v.ShellNames }).(pulumi.StringArrayOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// The space ID associated with this resource.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A filter to search by a list of tenant tags.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []string { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of tenant IDs.
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The thumbprint of the deployment target to match in the query and/or search
func (o GetOfflinePackageDropDeploymentTargetsResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOfflinePackageDropDeploymentTargetsResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOfflinePackageDropDeploymentTargetsResultOutput{})
}
