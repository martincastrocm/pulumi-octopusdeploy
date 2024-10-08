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
//			_, err := octopusdeploy.NewTokenAccount(ctx, "example", &octopusdeploy.TokenAccountArgs{
//				Token: pulumi.String("[token]"),
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
// $ pulumi import octopusdeploy:index/tokenAccount:TokenAccount [options] octopusdeploy_token_account.<name> <account-id>
// ```
type TokenAccount struct {
	pulumi.CustomResourceState

	// The description of this token account.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayOutput `pulumi:"environments"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringOutput `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayOutput `pulumi:"tenants"`
	// The token of this resource.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewTokenAccount registers a new resource with the given unique name, arguments, and options.
func NewTokenAccount(ctx *pulumi.Context,
	name string, args *TokenAccountArgs, opts ...pulumi.ResourceOption) (*TokenAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TokenAccount
	err := ctx.RegisterResource("octopusdeploy:index/tokenAccount:TokenAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTokenAccount gets an existing TokenAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTokenAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenAccountState, opts ...pulumi.ResourceOption) (*TokenAccount, error) {
	var resource TokenAccount
	err := ctx.ReadResource("octopusdeploy:index/tokenAccount:TokenAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TokenAccount resources.
type tokenAccountState struct {
	// The description of this token account.
	Description *string `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments []string `pulumi:"environments"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
	// The token of this resource.
	Token *string `pulumi:"token"`
}

type TokenAccountState struct {
	// The description of this token account.
	Description pulumi.StringPtrInput
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
	// The token of this resource.
	Token pulumi.StringPtrInput
}

func (TokenAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenAccountState)(nil)).Elem()
}

type tokenAccountArgs struct {
	// The description of this token account.
	Description *string `pulumi:"description"`
	// A list of environment IDs associated with this resource.
	Environments []string `pulumi:"environments"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
	// The token of this resource.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a TokenAccount resource.
type TokenAccountArgs struct {
	// The description of this token account.
	Description pulumi.StringPtrInput
	// A list of environment IDs associated with this resource.
	Environments pulumi.StringArrayInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
	// The token of this resource.
	Token pulumi.StringInput
}

func (TokenAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenAccountArgs)(nil)).Elem()
}

type TokenAccountInput interface {
	pulumi.Input

	ToTokenAccountOutput() TokenAccountOutput
	ToTokenAccountOutputWithContext(ctx context.Context) TokenAccountOutput
}

func (*TokenAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenAccount)(nil)).Elem()
}

func (i *TokenAccount) ToTokenAccountOutput() TokenAccountOutput {
	return i.ToTokenAccountOutputWithContext(context.Background())
}

func (i *TokenAccount) ToTokenAccountOutputWithContext(ctx context.Context) TokenAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenAccountOutput)
}

// TokenAccountArrayInput is an input type that accepts TokenAccountArray and TokenAccountArrayOutput values.
// You can construct a concrete instance of `TokenAccountArrayInput` via:
//
//	TokenAccountArray{ TokenAccountArgs{...} }
type TokenAccountArrayInput interface {
	pulumi.Input

	ToTokenAccountArrayOutput() TokenAccountArrayOutput
	ToTokenAccountArrayOutputWithContext(context.Context) TokenAccountArrayOutput
}

type TokenAccountArray []TokenAccountInput

func (TokenAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TokenAccount)(nil)).Elem()
}

func (i TokenAccountArray) ToTokenAccountArrayOutput() TokenAccountArrayOutput {
	return i.ToTokenAccountArrayOutputWithContext(context.Background())
}

func (i TokenAccountArray) ToTokenAccountArrayOutputWithContext(ctx context.Context) TokenAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenAccountArrayOutput)
}

// TokenAccountMapInput is an input type that accepts TokenAccountMap and TokenAccountMapOutput values.
// You can construct a concrete instance of `TokenAccountMapInput` via:
//
//	TokenAccountMap{ "key": TokenAccountArgs{...} }
type TokenAccountMapInput interface {
	pulumi.Input

	ToTokenAccountMapOutput() TokenAccountMapOutput
	ToTokenAccountMapOutputWithContext(context.Context) TokenAccountMapOutput
}

type TokenAccountMap map[string]TokenAccountInput

func (TokenAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TokenAccount)(nil)).Elem()
}

func (i TokenAccountMap) ToTokenAccountMapOutput() TokenAccountMapOutput {
	return i.ToTokenAccountMapOutputWithContext(context.Background())
}

func (i TokenAccountMap) ToTokenAccountMapOutputWithContext(ctx context.Context) TokenAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenAccountMapOutput)
}

type TokenAccountOutput struct{ *pulumi.OutputState }

func (TokenAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenAccount)(nil)).Elem()
}

func (o TokenAccountOutput) ToTokenAccountOutput() TokenAccountOutput {
	return o
}

func (o TokenAccountOutput) ToTokenAccountOutputWithContext(ctx context.Context) TokenAccountOutput {
	return o
}

// The description of this token account.
func (o TokenAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of environment IDs associated with this resource.
func (o TokenAccountOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

// The name of this resource.
func (o TokenAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o TokenAccountOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o TokenAccountOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o TokenAccountOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringOutput { return v.TenantedDeploymentParticipation }).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o TokenAccountOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The token of this resource.
func (o TokenAccountOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *TokenAccount) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type TokenAccountArrayOutput struct{ *pulumi.OutputState }

func (TokenAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TokenAccount)(nil)).Elem()
}

func (o TokenAccountArrayOutput) ToTokenAccountArrayOutput() TokenAccountArrayOutput {
	return o
}

func (o TokenAccountArrayOutput) ToTokenAccountArrayOutputWithContext(ctx context.Context) TokenAccountArrayOutput {
	return o
}

func (o TokenAccountArrayOutput) Index(i pulumi.IntInput) TokenAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TokenAccount {
		return vs[0].([]*TokenAccount)[vs[1].(int)]
	}).(TokenAccountOutput)
}

type TokenAccountMapOutput struct{ *pulumi.OutputState }

func (TokenAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TokenAccount)(nil)).Elem()
}

func (o TokenAccountMapOutput) ToTokenAccountMapOutput() TokenAccountMapOutput {
	return o
}

func (o TokenAccountMapOutput) ToTokenAccountMapOutputWithContext(ctx context.Context) TokenAccountMapOutput {
	return o
}

func (o TokenAccountMapOutput) MapIndex(k pulumi.StringInput) TokenAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TokenAccount {
		return vs[0].(map[string]*TokenAccount)[vs[1].(string)]
	}).(TokenAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TokenAccountInput)(nil)).Elem(), &TokenAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenAccountArrayInput)(nil)).Elem(), TokenAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenAccountMapInput)(nil)).Elem(), TokenAccountMap{})
	pulumi.RegisterOutputType(TokenAccountOutput{})
	pulumi.RegisterOutputType(TokenAccountArrayOutput{})
	pulumi.RegisterOutputType(TokenAccountMapOutput{})
}
