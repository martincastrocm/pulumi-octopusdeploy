// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing certificates.
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
//			_, err := octopusdeploy.GetCertificates(ctx, &octopusdeploy.GetCertificatesArgs{
//				Archived: pulumi.StringRef("false"),
//				Ids: []string{
//					"Certificates-123",
//					"Certificates-321",
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
func GetCertificates(ctx *pulumi.Context, args *GetCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCertificatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCertificatesResult
	err := ctx.Invoke("octopusdeploy:index/getCertificates:getCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesArgs struct {
	Archived *string `pulumi:"archived"`
	// A filter to define the first result.
	FirstResult *string `pulumi:"firstResult"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter used to order the search results.
	OrderBy *string `pulumi:"orderBy"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter of terms used the search operation.
	Search *string `pulumi:"search"`
	// A filter to specify the number of items to skip in the response.
	Skip    *int    `pulumi:"skip"`
	SpaceId *string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A filter to search by a tenant ID.
	Tenant *string `pulumi:"tenant"`
}

// A collection of values returned by getCertificates.
type GetCertificatesResult struct {
	// A filter to search for resources that have been archived.
	Archived *string `pulumi:"archived"`
	// A list of certificates that match the filter(s).
	Certificates []GetCertificatesCertificate `pulumi:"certificates"`
	// A filter to define the first result.
	FirstResult *string `pulumi:"firstResult"`
	// An auto-generated identifier that includes the timestamp when this data source was last modified.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A filter used to order the search results.
	OrderBy *string `pulumi:"orderBy"`
	// A filter to search by the partial match of a name.
	PartialName *string `pulumi:"partialName"`
	// A filter of terms used the search operation.
	Search *string `pulumi:"search"`
	// A filter to specify the number of items to skip in the response.
	Skip *int `pulumi:"skip"`
	// The space ID associated with this resource.
	SpaceId string `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take *int `pulumi:"take"`
	// A filter to search by a tenant ID.
	Tenant *string `pulumi:"tenant"`
}

func GetCertificatesOutput(ctx *pulumi.Context, args GetCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCertificatesResult, error) {
			args := v.(GetCertificatesArgs)
			r, err := GetCertificates(ctx, &args, opts...)
			var s GetCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCertificatesResultOutput)
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesOutputArgs struct {
	Archived pulumi.StringPtrInput `pulumi:"archived"`
	// A filter to define the first result.
	FirstResult pulumi.StringPtrInput `pulumi:"firstResult"`
	// A filter to search by a list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A filter used to order the search results.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// A filter to search by the partial match of a name.
	PartialName pulumi.StringPtrInput `pulumi:"partialName"`
	// A filter of terms used the search operation.
	Search pulumi.StringPtrInput `pulumi:"search"`
	// A filter to specify the number of items to skip in the response.
	Skip    pulumi.IntPtrInput    `pulumi:"skip"`
	SpaceId pulumi.StringPtrInput `pulumi:"spaceId"`
	// A filter to specify the number of items to take (or return) in the response.
	Take pulumi.IntPtrInput `pulumi:"take"`
	// A filter to search by a tenant ID.
	Tenant pulumi.StringPtrInput `pulumi:"tenant"`
}

func (GetCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getCertificates.
type GetCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesResult)(nil)).Elem()
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutput() GetCertificatesResultOutput {
	return o
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutputWithContext(ctx context.Context) GetCertificatesResultOutput {
	return o
}

// A filter to search for resources that have been archived.
func (o GetCertificatesResultOutput) Archived() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.Archived }).(pulumi.StringPtrOutput)
}

// A list of certificates that match the filter(s).
func (o GetCertificatesResultOutput) Certificates() GetCertificatesCertificateArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []GetCertificatesCertificate { return v.Certificates }).(GetCertificatesCertificateArrayOutput)
}

// A filter to define the first result.
func (o GetCertificatesResultOutput) FirstResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.FirstResult }).(pulumi.StringPtrOutput)
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetCertificatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A filter used to order the search results.
func (o GetCertificatesResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

// A filter to search by the partial match of a name.
func (o GetCertificatesResultOutput) PartialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.PartialName }).(pulumi.StringPtrOutput)
}

// A filter of terms used the search operation.
func (o GetCertificatesResultOutput) Search() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.Search }).(pulumi.StringPtrOutput)
}

// A filter to specify the number of items to skip in the response.
func (o GetCertificatesResultOutput) Skip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *int { return v.Skip }).(pulumi.IntPtrOutput)
}

// The space ID associated with this resource.
func (o GetCertificatesResultOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.SpaceId }).(pulumi.StringOutput)
}

// A filter to specify the number of items to take (or return) in the response.
func (o GetCertificatesResultOutput) Take() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *int { return v.Take }).(pulumi.IntPtrOutput)
}

// A filter to search by a tenant ID.
func (o GetCertificatesResultOutput) Tenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.Tenant }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCertificatesResultOutput{})
}
