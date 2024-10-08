// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing listening tentacle deployment targets.
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
//			_, err := octopusdeploy.GetListeningTentacleDeploymentTargets(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetListeningTentacleDeploymentTargets(ctx *pulumi.Context, args *GetListeningTentacleDeploymentTargetsArgs, opts ...pulumi.InvokeOption) (*GetListeningTentacleDeploymentTargetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetListeningTentacleDeploymentTargetsResult
	err := ctx.Invoke("octopusdeploy:index/getListeningTentacleDeploymentTargets:getListeningTentacleDeploymentTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListeningTentacleDeploymentTargets.
type GetListeningTentacleDeploymentTargetsArgs struct {
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

// A collection of values returned by getListeningTentacleDeploymentTargets.
type GetListeningTentacleDeploymentTargetsResult struct {
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
	// A list of listening tentacle deployment targets that match the filter(s).
	ListeningTentacleDeploymentTargets []GetListeningTentacleDeploymentTargetsListeningTentacleDeploymentTarget `pulumi:"listeningTentacleDeploymentTargets"`
	// A filter to search by name.
	Name *string `pulumi:"name"`
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

func GetListeningTentacleDeploymentTargetsOutput(ctx *pulumi.Context, args GetListeningTentacleDeploymentTargetsOutputArgs, opts ...pulumi.InvokeOption) GetListeningTentacleDeploymentTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetListeningTentacleDeploymentTargetsResult, error) {
			args := v.(GetListeningTentacleDeploymentTargetsArgs)
			r, err := GetListeningTentacleDeploymentTargets(ctx, &args, opts...)
			var s GetListeningTentacleDeploymentTargetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetListeningTentacleDeploymentTargetsResultOutput)
}

// A collection of arguments for invoking getListeningTentacleDeploymentTargets.
type GetListeningTentacleDeploymentTargetsOutputArgs struct {
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

func (GetListeningTentacleDeploymentTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListeningTentacleDeploymentTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getListeningTentacleDeploymentTargets.
type GetListeningTentacleDeploymentTargetsResultOutput struct{ *pulumi.OutputState }

func (GetListeningTentacleDeploymentTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListeningTentacleDeploymentTargetsResult)(nil)).Elem()
}

func (o GetListeningTentacleDeploymentTargetsResultOutput) ToGetListeningTentacleDeploymentTargetsResultOutput() GetListeningTentacleDeploymentTargetsResultOutput {
	return o
}

func (o GetListeningTentacleDeploymentTargetsResultOutput) ToGetListeningTentacleDeploymentTargetsResultOutputWithContext(ctx context.Context) GetListeningTentacleDeploymentTargetsResultOutput {
	return o
}

// A filter to search by deployment ID.
func (o GetListeningTentacleDeploymentTargetsResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of environment IDs.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []string { return v.Environments }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o GetListeningTentacleDeploymentTargetsResultOutput) HealthStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []string { return v.HealthStatuses }).(pulumi.StringArrayOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by the disabled status of a resource.
func (o GetListeningTentacleDeploymentTargetsResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// A list of listening tentacle deployment targets that match the filter(s).
func (o GetListeningTentacleDeploymentTargetsResultOutput) ListeningTentacleDeploymentTargets() GetListeningTentacleDeploymentTargetsListeningTentacleDeploymentTargetArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []GetListeningTentacleDeploymentTargetsListeningTentacleDeploymentTarget {
		return v.ListeningTentacleDeploymentTargets
	}).(GetListeningTentacleDeploymentTargetsListeningTentacleDeploymentTargetArrayOutput)
}

// A filter to search by name.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A filter to search by the partial match of a name.
func (o GetListeningTentacleDeploymentTargetsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of role IDs.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// A list of shell names to match in the query and/or search
func (o GetListeningTentacleDeploymentTargetsResultOutput) ShellNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []string { return v.ShellNames }).(pulumi.StringArrayOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// The space ID associated with this resource.
func (o GetListeningTentacleDeploymentTargetsResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A filter to search by a list of tenant tags.
func (o GetListeningTentacleDeploymentTargetsResultOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []string { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of tenant IDs.
func (o GetListeningTentacleDeploymentTargetsResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The thumbprint of the deployment target to match in the query and/or search
func (o GetListeningTentacleDeploymentTargetsResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListeningTentacleDeploymentTargetsResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetListeningTentacleDeploymentTargetsResultOutput{})
}
