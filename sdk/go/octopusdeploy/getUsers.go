// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing users.
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
//			_, err := octopusdeploy.GetUsers(ctx, &octopusdeploy.GetUsersArgs{
//				Ids: []string{
//					"Users-123",
//					"Users-321",
//				},
//				Skip: pulumi.IntRef(5),
//				Take: pulumi.IntRef(100),
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
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("octopusdeploy:index/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// A filter with which to search.
	Filter *string `pulumi:"filter"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A Space ID to filter by. Will revert what is specified on the provider if not set.
	SpaceId *string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	// A filter with which to search.
	Filter *string `pulumi:"filter"`
	// An auto-generated identifier that includes the timestamp when this data source was last modified.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// A Space ID to filter by. Will revert what is specified on the provider if not set.
	SpaceId *string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A list of users that match the filter(s).
	Users []GetUsersUser `pulumi:"users"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsersResult, error) {
			args := v.(GetUsersArgs)
			r, err := GetUsers(ctx, &args, opts...)
			var s GetUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	// A filter with which to search.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// A filter to search by a list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A filter to specify the number of items to skip in the response.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// A Space ID to filter by. Will revert what is specified on the provider if not set.
	SpaceId pulumi.StringPtrInput `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take pulumi.IntPtrInput `pulumi:"take"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

// A filter with which to search.
func (o GetUsersResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetUsersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetUsersResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// A Space ID to filter by. Will revert what is specified on the provider if not set.
func (o GetUsersResultOutput) SpaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.SpaceId }).(pulumi.StringPtrOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetUsersResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A list of users that match the filter(s).
func (o GetUsersResultOutput) Users() GetUsersUserArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUser { return v.Users }).(GetUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
