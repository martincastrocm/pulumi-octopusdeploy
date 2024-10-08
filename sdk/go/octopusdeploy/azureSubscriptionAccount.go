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

// This resource manages Azure subscription accounts in Octopus Deploy.
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
//			_, err := octopusdeploy.NewAzureSubscriptionAccount(ctx, "example", &octopusdeploy.AzureSubscriptionAccountArgs{
//				SubscriptionId: pulumi.String("00000000-0000-0000-0000-000000000000"),
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
// $ pulumi import octopusdeploy:index/azureSubscriptionAccount:AzureSubscriptionAccount [options] octopusdeploy_azure_subscription_account.<name> <account-id>
// ```
type AzureSubscriptionAccount struct {
	pulumi.CustomResourceState

	// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
	AzureEnvironment      pulumi.StringOutput `pulumi:"azureEnvironment"`
	Certificate           pulumi.StringOutput `pulumi:"certificate"`
	CertificateThumbprint pulumi.StringOutput `pulumi:"certificateThumbprint"`
	// The description of this Azure subscription account.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments       pulumi.StringArrayOutput `pulumi:"environments"`
	ManagementEndpoint pulumi.StringPtrOutput   `pulumi:"managementEndpoint"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// The storage endpoint suffix associated with this Azure subscription account.
	StorageEndpointSuffix pulumi.StringPtrOutput `pulumi:"storageEndpointSuffix"`
	// The subscription ID of this resource.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringOutput `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayOutput `pulumi:"tenants"`
}

// NewAzureSubscriptionAccount registers a new resource with the given unique name, arguments, and options.
func NewAzureSubscriptionAccount(ctx *pulumi.Context,
	name string, args *AzureSubscriptionAccountArgs, opts ...pulumi.ResourceOption) (*AzureSubscriptionAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	if args.Certificate != nil {
		args.Certificate = pulumi.ToSecret(args.Certificate).(pulumi.StringPtrInput)
	}
	if args.CertificateThumbprint != nil {
		args.CertificateThumbprint = pulumi.ToSecret(args.CertificateThumbprint).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"certificate",
		"certificateThumbprint",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureSubscriptionAccount
	err := ctx.RegisterResource("octopusdeploy:index/azureSubscriptionAccount:AzureSubscriptionAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureSubscriptionAccount gets an existing AzureSubscriptionAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureSubscriptionAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureSubscriptionAccountState, opts ...pulumi.ResourceOption) (*AzureSubscriptionAccount, error) {
	var resource AzureSubscriptionAccount
	err := ctx.ReadResource("octopusdeploy:index/azureSubscriptionAccount:AzureSubscriptionAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureSubscriptionAccount resources.
type azureSubscriptionAccountState struct {
	// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
	AzureEnvironment      *string `pulumi:"azureEnvironment"`
	Certificate           *string `pulumi:"certificate"`
	CertificateThumbprint *string `pulumi:"certificateThumbprint"`
	// The description of this Azure subscription account.
	Description *string `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments       []string `pulumi:"environments"`
	ManagementEndpoint *string  `pulumi:"managementEndpoint"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The storage endpoint suffix associated with this Azure subscription account.
	StorageEndpointSuffix *string `pulumi:"storageEndpointSuffix"`
	// The subscription ID of this resource.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
}

type AzureSubscriptionAccountState struct {
	// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
	AzureEnvironment      pulumi.StringPtrInput
	Certificate           pulumi.StringPtrInput
	CertificateThumbprint pulumi.StringPtrInput
	// The description of this Azure subscription account.
	Description pulumi.StringPtrInput
	// A list of environment IDs associated with this resource.
	Environments       pulumi.StringArrayInput
	ManagementEndpoint pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The storage endpoint suffix associated with this Azure subscription account.
	StorageEndpointSuffix pulumi.StringPtrInput
	// The subscription ID of this resource.
	SubscriptionId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
}

func (AzureSubscriptionAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureSubscriptionAccountState)(nil)).Elem()
}

type azureSubscriptionAccountArgs struct {
	// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
	AzureEnvironment      *string `pulumi:"azureEnvironment"`
	Certificate           *string `pulumi:"certificate"`
	CertificateThumbprint *string `pulumi:"certificateThumbprint"`
	// The description of this Azure subscription account.
	Description *string `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments       []string `pulumi:"environments"`
	ManagementEndpoint *string  `pulumi:"managementEndpoint"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// The storage endpoint suffix associated with this Azure subscription account.
	StorageEndpointSuffix *string `pulumi:"storageEndpointSuffix"`
	// The subscription ID of this resource.
	SubscriptionId string `pulumi:"subscriptionId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
}

// The set of arguments for constructing a AzureSubscriptionAccount resource.
type AzureSubscriptionAccountArgs struct {
	// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
	AzureEnvironment      pulumi.StringPtrInput
	Certificate           pulumi.StringPtrInput
	CertificateThumbprint pulumi.StringPtrInput
	// The description of this Azure subscription account.
	Description pulumi.StringPtrInput
	// A list of environment IDs associated with this resource.
	Environments       pulumi.StringArrayInput
	ManagementEndpoint pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// The storage endpoint suffix associated with this Azure subscription account.
	StorageEndpointSuffix pulumi.StringPtrInput
	// The subscription ID of this resource.
	SubscriptionId pulumi.StringInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
}

func (AzureSubscriptionAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureSubscriptionAccountArgs)(nil)).Elem()
}

type AzureSubscriptionAccountInput interface {
	pulumi.Input

	ToAzureSubscriptionAccountOutput() AzureSubscriptionAccountOutput
	ToAzureSubscriptionAccountOutputWithContext(ctx context.Context) AzureSubscriptionAccountOutput
}

func (*AzureSubscriptionAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSubscriptionAccount)(nil)).Elem()
}

func (i *AzureSubscriptionAccount) ToAzureSubscriptionAccountOutput() AzureSubscriptionAccountOutput {
	return i.ToAzureSubscriptionAccountOutputWithContext(context.Background())
}

func (i *AzureSubscriptionAccount) ToAzureSubscriptionAccountOutputWithContext(ctx context.Context) AzureSubscriptionAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSubscriptionAccountOutput)
}

// AzureSubscriptionAccountArrayInput is an input type that accepts AzureSubscriptionAccountArray and AzureSubscriptionAccountArrayOutput values.
// You can construct a concrete instance of `AzureSubscriptionAccountArrayInput` via:
//
//	AzureSubscriptionAccountArray{ AzureSubscriptionAccountArgs{...} }
type AzureSubscriptionAccountArrayInput interface {
	pulumi.Input

	ToAzureSubscriptionAccountArrayOutput() AzureSubscriptionAccountArrayOutput
	ToAzureSubscriptionAccountArrayOutputWithContext(context.Context) AzureSubscriptionAccountArrayOutput
}

type AzureSubscriptionAccountArray []AzureSubscriptionAccountInput

func (AzureSubscriptionAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureSubscriptionAccount)(nil)).Elem()
}

func (i AzureSubscriptionAccountArray) ToAzureSubscriptionAccountArrayOutput() AzureSubscriptionAccountArrayOutput {
	return i.ToAzureSubscriptionAccountArrayOutputWithContext(context.Background())
}

func (i AzureSubscriptionAccountArray) ToAzureSubscriptionAccountArrayOutputWithContext(ctx context.Context) AzureSubscriptionAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSubscriptionAccountArrayOutput)
}

// AzureSubscriptionAccountMapInput is an input type that accepts AzureSubscriptionAccountMap and AzureSubscriptionAccountMapOutput values.
// You can construct a concrete instance of `AzureSubscriptionAccountMapInput` via:
//
//	AzureSubscriptionAccountMap{ "key": AzureSubscriptionAccountArgs{...} }
type AzureSubscriptionAccountMapInput interface {
	pulumi.Input

	ToAzureSubscriptionAccountMapOutput() AzureSubscriptionAccountMapOutput
	ToAzureSubscriptionAccountMapOutputWithContext(context.Context) AzureSubscriptionAccountMapOutput
}

type AzureSubscriptionAccountMap map[string]AzureSubscriptionAccountInput

func (AzureSubscriptionAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureSubscriptionAccount)(nil)).Elem()
}

func (i AzureSubscriptionAccountMap) ToAzureSubscriptionAccountMapOutput() AzureSubscriptionAccountMapOutput {
	return i.ToAzureSubscriptionAccountMapOutputWithContext(context.Background())
}

func (i AzureSubscriptionAccountMap) ToAzureSubscriptionAccountMapOutputWithContext(ctx context.Context) AzureSubscriptionAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSubscriptionAccountMapOutput)
}

type AzureSubscriptionAccountOutput struct{ *pulumi.OutputState }

func (AzureSubscriptionAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSubscriptionAccount)(nil)).Elem()
}

func (o AzureSubscriptionAccountOutput) ToAzureSubscriptionAccountOutput() AzureSubscriptionAccountOutput {
	return o
}

func (o AzureSubscriptionAccountOutput) ToAzureSubscriptionAccountOutputWithContext(ctx context.Context) AzureSubscriptionAccountOutput {
	return o
}

// The Azure environment associated with this resource. Valid Azure environments are `AzureCloud`, `AzureChinaCloud`, `AzureGermanCloud`, or `AzureUSGovernment`.
func (o AzureSubscriptionAccountOutput) AzureEnvironment() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringOutput { return v.AzureEnvironment }).(pulumi.StringOutput)
}

func (o AzureSubscriptionAccountOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

func (o AzureSubscriptionAccountOutput) CertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringOutput { return v.CertificateThumbprint }).(pulumi.StringOutput)
}

// The description of this Azure subscription account.
func (o AzureSubscriptionAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of environment IDs associated with this resource.
func (o AzureSubscriptionAccountOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o AzureSubscriptionAccountOutput) ManagementEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringPtrOutput { return v.ManagementEndpoint }).(pulumi.StringPtrOutput)
}

// The name of this resource.
func (o AzureSubscriptionAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o AzureSubscriptionAccountOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// The storage endpoint suffix associated with this Azure subscription account.
func (o AzureSubscriptionAccountOutput) StorageEndpointSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringPtrOutput { return v.StorageEndpointSuffix }).(pulumi.StringPtrOutput)
}

// The subscription ID of this resource.
func (o AzureSubscriptionAccountOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o AzureSubscriptionAccountOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o AzureSubscriptionAccountOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringOutput { return v.TenantedDeploymentParticipation }).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o AzureSubscriptionAccountOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureSubscriptionAccount) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

type AzureSubscriptionAccountArrayOutput struct{ *pulumi.OutputState }

func (AzureSubscriptionAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureSubscriptionAccount)(nil)).Elem()
}

func (o AzureSubscriptionAccountArrayOutput) ToAzureSubscriptionAccountArrayOutput() AzureSubscriptionAccountArrayOutput {
	return o
}

func (o AzureSubscriptionAccountArrayOutput) ToAzureSubscriptionAccountArrayOutputWithContext(ctx context.Context) AzureSubscriptionAccountArrayOutput {
	return o
}

func (o AzureSubscriptionAccountArrayOutput) Index(i pulumi.IntInput) AzureSubscriptionAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureSubscriptionAccount {
		return vs[0].([]*AzureSubscriptionAccount)[vs[1].(int)]
	}).(AzureSubscriptionAccountOutput)
}

type AzureSubscriptionAccountMapOutput struct{ *pulumi.OutputState }

func (AzureSubscriptionAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureSubscriptionAccount)(nil)).Elem()
}

func (o AzureSubscriptionAccountMapOutput) ToAzureSubscriptionAccountMapOutput() AzureSubscriptionAccountMapOutput {
	return o
}

func (o AzureSubscriptionAccountMapOutput) ToAzureSubscriptionAccountMapOutputWithContext(ctx context.Context) AzureSubscriptionAccountMapOutput {
	return o
}

func (o AzureSubscriptionAccountMapOutput) MapIndex(k pulumi.StringInput) AzureSubscriptionAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureSubscriptionAccount {
		return vs[0].(map[string]*AzureSubscriptionAccount)[vs[1].(string)]
	}).(AzureSubscriptionAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSubscriptionAccountInput)(nil)).Elem(), &AzureSubscriptionAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSubscriptionAccountArrayInput)(nil)).Elem(), AzureSubscriptionAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSubscriptionAccountMapInput)(nil)).Elem(), AzureSubscriptionAccountMap{})
	pulumi.RegisterOutputType(AzureSubscriptionAccountOutput{})
	pulumi.RegisterOutputType(AzureSubscriptionAccountArrayOutput{})
	pulumi.RegisterOutputType(AzureSubscriptionAccountMapOutput{})
}
