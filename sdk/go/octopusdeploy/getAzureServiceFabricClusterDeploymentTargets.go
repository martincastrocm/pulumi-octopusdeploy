// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing Azure service fabric cluster deployment targets.
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
//			_, err := octopusdeploy.GetAzureServiceFabricClusterDeploymentTargets(ctx, &octopusdeploy.GetAzureServiceFabricClusterDeploymentTargetsArgs{
//				HealthStatuses: []string{
//					"Healthy",
//					"Unavailable",
//				},
//				Ids: []string{
//					"Machines-123",
//					"Machines-321",
//				},
//				PartialName: pulumi.StringRef("Defau"),
//				Skip:        pulumi.IntRef(5),
//				Take:        pulumi.IntRef(100),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetAzureServiceFabricClusterDeploymentTargets(ctx *pulumi.Context, args *GetAzureServiceFabricClusterDeploymentTargetsArgs, opts ...pulumi.InvokeOption) (*GetAzureServiceFabricClusterDeploymentTargetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAzureServiceFabricClusterDeploymentTargetsResult
	err := ctx.Invoke("octopusdeploy:index/getAzureServiceFabricClusterDeploymentTargets:getAzureServiceFabricClusterDeploymentTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAzureServiceFabricClusterDeploymentTargets.
type GetAzureServiceFabricClusterDeploymentTargetsArgs struct {
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

// A collection of values returned by getAzureServiceFabricClusterDeploymentTargets.
type GetAzureServiceFabricClusterDeploymentTargetsResult struct {
	// A list of Azure service fabric cluster deployment targets that match the filter(s).
	AzureServiceFabricClusterDeploymentTargets []GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTarget `pulumi:"azureServiceFabricClusterDeploymentTargets"`
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

func GetAzureServiceFabricClusterDeploymentTargetsOutput(ctx *pulumi.Context, args GetAzureServiceFabricClusterDeploymentTargetsOutputArgs, opts ...pulumi.InvokeOption) GetAzureServiceFabricClusterDeploymentTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAzureServiceFabricClusterDeploymentTargetsResult, error) {
			args := v.(GetAzureServiceFabricClusterDeploymentTargetsArgs)
			r, err := GetAzureServiceFabricClusterDeploymentTargets(ctx, &args, opts...)
			var s GetAzureServiceFabricClusterDeploymentTargetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAzureServiceFabricClusterDeploymentTargetsResultOutput)
}

// A collection of arguments for invoking getAzureServiceFabricClusterDeploymentTargets.
type GetAzureServiceFabricClusterDeploymentTargetsOutputArgs struct {
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

func (GetAzureServiceFabricClusterDeploymentTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAzureServiceFabricClusterDeploymentTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getAzureServiceFabricClusterDeploymentTargets.
type GetAzureServiceFabricClusterDeploymentTargetsResultOutput struct{ *pulumi.OutputState }

func (GetAzureServiceFabricClusterDeploymentTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAzureServiceFabricClusterDeploymentTargetsResult)(nil)).Elem()
}

func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) ToGetAzureServiceFabricClusterDeploymentTargetsResultOutput() GetAzureServiceFabricClusterDeploymentTargetsResultOutput {
	return o
}

func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) ToGetAzureServiceFabricClusterDeploymentTargetsResultOutputWithContext(ctx context.Context) GetAzureServiceFabricClusterDeploymentTargetsResultOutput {
	return o
}

// A list of Azure service fabric cluster deployment targets that match the filter(s).
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) AzureServiceFabricClusterDeploymentTargets() GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTarget {
		return v.AzureServiceFabricClusterDeploymentTargets
	}).(GetAzureServiceFabricClusterDeploymentTargetsAzureServiceFabricClusterDeploymentTargetArrayOutput)
}

// A filter to search by deployment ID.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of environment IDs.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []string { return v.Environments }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) HealthStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []string { return v.HealthStatuses }).(pulumi.StringArrayOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to search by the disabled status of a resource.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// A filter to search by name.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A filter to search by the partial match of a name.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to search by a list of role IDs.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// A list of shell names to match in the query and/or search
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) ShellNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []string { return v.ShellNames }).(pulumi.StringArrayOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// The space ID associated with this resource.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A filter to search by a list of tenant tags.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []string { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// A filter to search by a list of tenant IDs.
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The thumbprint of the deployment target to match in the query and/or search
func (o GetAzureServiceFabricClusterDeploymentTargetsResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAzureServiceFabricClusterDeploymentTargetsResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAzureServiceFabricClusterDeploymentTargetsResultOutput{})
}
