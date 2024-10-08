// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing machine policies.
func GetMachinePolicies(ctx *pulumi.Context, args *GetMachinePoliciesArgs, opts ...pulumi.InvokeOption) (*GetMachinePoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMachinePoliciesResult
	err := ctx.Invoke("octopusdeploy:index/getMachinePolicies:getMachinePolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMachinePolicies.
type GetMachinePoliciesArgs struct {
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip    *int    `pulumi:"skip"`
	SpaceId *string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
}

// A collection of values returned by getMachinePolicies.
type GetMachinePoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A list of machine policies that match the filter(s).
	MachinePolicies []GetMachinePoliciesMachinePolicy `pulumi:"machinePolicies"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// The space ID associated with this resource.
	SpaceId string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
}

func GetMachinePoliciesOutput(ctx *pulumi.Context, args GetMachinePoliciesOutputArgs, opts ...pulumi.InvokeOption) GetMachinePoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMachinePoliciesResult, error) {
			args := v.(GetMachinePoliciesArgs)
			r, err := GetMachinePolicies(ctx, &args, opts...)
			var s GetMachinePoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMachinePoliciesResultOutput)
}

// A collection of arguments for invoking getMachinePolicies.
type GetMachinePoliciesOutputArgs struct {
	// A filter to search by a list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A filter to search by the partial match of a name.
	PartialName pulumi.StringPtrInput `pulumi:"partialName"`
	// A filter to specify the number of items to skip in the response.
	Skip    pulumi.IntPtrInput    `pulumi:"skip"`
	SpaceId pulumi.StringPtrInput `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take pulumi.IntPtrInput `pulumi:"take"`
}

func (GetMachinePoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMachinePoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getMachinePolicies.
type GetMachinePoliciesResultOutput struct{ *pulumi.OutputState }

func (GetMachinePoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMachinePoliciesResult)(nil)).Elem()
}

func (o GetMachinePoliciesResultOutput) ToGetMachinePoliciesResultOutput() GetMachinePoliciesResultOutput {
	return o
}

func (o GetMachinePoliciesResultOutput) ToGetMachinePoliciesResultOutputWithContext(ctx context.Context) GetMachinePoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetMachinePoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMachinePoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetMachinePoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMachinePoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of machine policies that match the filter(s).
func (o GetMachinePoliciesResultOutput) MachinePolicies() GetMachinePoliciesMachinePolicyArrayOutput {
	return o.ApplyT(func(v GetMachinePoliciesResult) []GetMachinePoliciesMachinePolicy { return v.MachinePolicies }).(GetMachinePoliciesMachinePolicyArrayOutput)
}

// A filter to search by the partial match of a name.
func (o GetMachinePoliciesResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMachinePoliciesResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetMachinePoliciesResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetMachinePoliciesResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// The space ID associated with this resource.
func (o GetMachinePoliciesResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMachinePoliciesResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetMachinePoliciesResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetMachinePoliciesResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMachinePoliciesResultOutput{})
}
