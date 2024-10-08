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

// This resource manages Kubernetes agent deployment targets in Octopus Deploy.
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
//			_, err := octopusdeploy.NewKubernetesAgentDeploymentTarget(ctx, "minimal", &octopusdeploy.KubernetesAgentDeploymentTargetArgs{
//				Environments: pulumi.StringArray{
//					pulumi.String("environments-1"),
//				},
//				Roles: pulumi.StringArray{
//					pulumi.String("role-1"),
//					pulumi.String("role-2"),
//				},
//				Thumbprint: pulumi.String("96203ED84246201C26A2F4360D7CBC36AC1D232D"),
//				Uri:        pulumi.String("poll://kcxzcv2fpsxkn6tk9u6d/"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewKubernetesAgentDeploymentTarget(ctx, "optionals", &octopusdeploy.KubernetesAgentDeploymentTargetArgs{
//				DefaultNamespace: pulumi.String("kubernetes-namespace"),
//				Environments: pulumi.StringArray{
//					pulumi.String("environments-1"),
//				},
//				IsDisabled:      pulumi.Bool(true),
//				MachinePolicyId: pulumi.String("machinepolicies-1"),
//				Roles: pulumi.StringArray{
//					pulumi.String("role-1"),
//					pulumi.String("role-2"),
//				},
//				Thumbprint:    pulumi.String("96203ED84246201C26A2F4360D7CBC36AC1D232D"),
//				UpgradeLocked: pulumi.Bool(true),
//				Uri:           pulumi.String("poll://kcxzcv2fpsxkn6tk9u6d/"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = octopusdeploy.NewKubernetesAgentDeploymentTarget(ctx, "tenantedAgent", &octopusdeploy.KubernetesAgentDeploymentTargetArgs{
//				Environments: pulumi.StringArray{
//					pulumi.String("environments-1"),
//				},
//				Roles: pulumi.StringArray{
//					pulumi.String("role-1"),
//					pulumi.String("role-2"),
//				},
//				TenantTags: pulumi.StringArray{
//					pulumi.String("TagSets-1/Tags-1"),
//				},
//				TenantedDeploymentParticipation: pulumi.String("Tenanted"),
//				Tenants: pulumi.StringArray{
//					pulumi.String("tenants-1"),
//				},
//				Thumbprint: pulumi.String("96203ED84246201C26A2F4360D7CBC36AC1D232D"),
//				Uri:        pulumi.String("poll://kcxzcv2fpsxkn6tk9u6d/"),
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
// $ pulumi import octopusdeploy:index/kubernetesAgentDeploymentTarget:KubernetesAgentDeploymentTarget [options] octopusdeploy_kubernetes_agent_deployment_target.<name> <machine-id>
// ```
type KubernetesAgentDeploymentTarget struct {
	pulumi.CustomResourceState

	// Name of the Helm release that the agent belongs to.
	AgentHelmReleaseName pulumi.StringOutput `pulumi:"agentHelmReleaseName"`
	// Name of the Kubernetes namespace where the agent is installed.
	AgentKubernetesNamespace pulumi.StringOutput `pulumi:"agentKubernetesNamespace"`
	// Current Tentacle version of the agent
	AgentTentacleVersion pulumi.StringOutput `pulumi:"agentTentacleVersion"`
	// Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
	AgentUpgradeStatus pulumi.StringOutput `pulumi:"agentUpgradeStatus"`
	// Current Helm chart version of the agent.
	AgentVersion pulumi.StringOutput `pulumi:"agentVersion"`
	// The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
	CommunicationMode pulumi.StringPtrOutput `pulumi:"communicationMode"`
	// Optional default namespace that will be used when using Kubernetes deployment steps, can be overrides within step configurations.
	DefaultNamespace pulumi.StringOutput `pulumi:"defaultNamespace"`
	// A list of environment IDs this Kubernetes agent can deploy to.
	Environments pulumi.StringArrayOutput `pulumi:"environments"`
	// Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
	IsDisabled pulumi.BoolPtrOutput `pulumi:"isDisabled"`
	// Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
	MachinePolicyId pulumi.StringOutput `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of target roles that are associated to this Kubernetes agent.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayOutput `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringOutput `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayOutput `pulumi:"tenants"`
	// The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
	UpgradeLocked pulumi.BoolPtrOutput `pulumi:"upgradeLocked"`
	// The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewKubernetesAgentDeploymentTarget registers a new resource with the given unique name, arguments, and options.
func NewKubernetesAgentDeploymentTarget(ctx *pulumi.Context,
	name string, args *KubernetesAgentDeploymentTargetArgs, opts ...pulumi.ResourceOption) (*KubernetesAgentDeploymentTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Thumbprint == nil {
		return nil, errors.New("invalid value for required argument 'Thumbprint'")
	}
	if args.Uri == nil {
		return nil, errors.New("invalid value for required argument 'Uri'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubernetesAgentDeploymentTarget
	err := ctx.RegisterResource("octopusdeploy:index/kubernetesAgentDeploymentTarget:KubernetesAgentDeploymentTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesAgentDeploymentTarget gets an existing KubernetesAgentDeploymentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesAgentDeploymentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesAgentDeploymentTargetState, opts ...pulumi.ResourceOption) (*KubernetesAgentDeploymentTarget, error) {
	var resource KubernetesAgentDeploymentTarget
	err := ctx.ReadResource("octopusdeploy:index/kubernetesAgentDeploymentTarget:KubernetesAgentDeploymentTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesAgentDeploymentTarget resources.
type kubernetesAgentDeploymentTargetState struct {
	// Name of the Helm release that the agent belongs to.
	AgentHelmReleaseName *string `pulumi:"agentHelmReleaseName"`
	// Name of the Kubernetes namespace where the agent is installed.
	AgentKubernetesNamespace *string `pulumi:"agentKubernetesNamespace"`
	// Current Tentacle version of the agent
	AgentTentacleVersion *string `pulumi:"agentTentacleVersion"`
	// Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
	AgentUpgradeStatus *string `pulumi:"agentUpgradeStatus"`
	// Current Helm chart version of the agent.
	AgentVersion *string `pulumi:"agentVersion"`
	// The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
	CommunicationMode *string `pulumi:"communicationMode"`
	// Optional default namespace that will be used when using Kubernetes deployment steps, can be overrides within step configurations.
	DefaultNamespace *string `pulumi:"defaultNamespace"`
	// A list of environment IDs this Kubernetes agent can deploy to.
	Environments []string `pulumi:"environments"`
	// Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
	IsDisabled *bool `pulumi:"isDisabled"`
	// Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// A list of target roles that are associated to this Kubernetes agent.
	Roles []string `pulumi:"roles"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
	// The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
	Thumbprint *string `pulumi:"thumbprint"`
	// If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
	UpgradeLocked *bool `pulumi:"upgradeLocked"`
	// The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
	Uri *string `pulumi:"uri"`
}

type KubernetesAgentDeploymentTargetState struct {
	// Name of the Helm release that the agent belongs to.
	AgentHelmReleaseName pulumi.StringPtrInput
	// Name of the Kubernetes namespace where the agent is installed.
	AgentKubernetesNamespace pulumi.StringPtrInput
	// Current Tentacle version of the agent
	AgentTentacleVersion pulumi.StringPtrInput
	// Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
	AgentUpgradeStatus pulumi.StringPtrInput
	// Current Helm chart version of the agent.
	AgentVersion pulumi.StringPtrInput
	// The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
	CommunicationMode pulumi.StringPtrInput
	// Optional default namespace that will be used when using Kubernetes deployment steps, can be overrides within step configurations.
	DefaultNamespace pulumi.StringPtrInput
	// A list of environment IDs this Kubernetes agent can deploy to.
	Environments pulumi.StringArrayInput
	// Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
	IsDisabled pulumi.BoolPtrInput
	// Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// A list of target roles that are associated to this Kubernetes agent.
	Roles pulumi.StringArrayInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
	// The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
	Thumbprint pulumi.StringPtrInput
	// If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
	UpgradeLocked pulumi.BoolPtrInput
	// The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
	Uri pulumi.StringPtrInput
}

func (KubernetesAgentDeploymentTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesAgentDeploymentTargetState)(nil)).Elem()
}

type kubernetesAgentDeploymentTargetArgs struct {
	// The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
	CommunicationMode *string `pulumi:"communicationMode"`
	// Optional default namespace that will be used when using Kubernetes deployment steps, can be overrides within step configurations.
	DefaultNamespace *string `pulumi:"defaultNamespace"`
	// A list of environment IDs this Kubernetes agent can deploy to.
	Environments []string `pulumi:"environments"`
	// Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
	IsDisabled *bool `pulumi:"isDisabled"`
	// Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
	MachinePolicyId *string `pulumi:"machinePolicyId"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// A list of target roles that are associated to this Kubernetes agent.
	Roles []string `pulumi:"roles"`
	// The space ID associated with this resource.
	SpaceId *string `pulumi:"spaceId"`
	// A list of tenant tags associated with this resource.
	TenantTags []string `pulumi:"tenantTags"`
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation *string `pulumi:"tenantedDeploymentParticipation"`
	// A list of tenant IDs associated with this resource.
	Tenants []string `pulumi:"tenants"`
	// The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
	Thumbprint string `pulumi:"thumbprint"`
	// If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
	UpgradeLocked *bool `pulumi:"upgradeLocked"`
	// The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
	Uri string `pulumi:"uri"`
}

// The set of arguments for constructing a KubernetesAgentDeploymentTarget resource.
type KubernetesAgentDeploymentTargetArgs struct {
	// The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
	CommunicationMode pulumi.StringPtrInput
	// Optional default namespace that will be used when using Kubernetes deployment steps, can be overrides within step configurations.
	DefaultNamespace pulumi.StringPtrInput
	// A list of environment IDs this Kubernetes agent can deploy to.
	Environments pulumi.StringArrayInput
	// Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
	IsDisabled pulumi.BoolPtrInput
	// Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
	MachinePolicyId pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// A list of target roles that are associated to this Kubernetes agent.
	Roles pulumi.StringArrayInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	// A list of tenant tags associated with this resource.
	TenantTags pulumi.StringArrayInput
	// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
	TenantedDeploymentParticipation pulumi.StringPtrInput
	// A list of tenant IDs associated with this resource.
	Tenants pulumi.StringArrayInput
	// The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
	Thumbprint pulumi.StringInput
	// If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
	UpgradeLocked pulumi.BoolPtrInput
	// The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
	Uri pulumi.StringInput
}

func (KubernetesAgentDeploymentTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesAgentDeploymentTargetArgs)(nil)).Elem()
}

type KubernetesAgentDeploymentTargetInput interface {
	pulumi.Input

	ToKubernetesAgentDeploymentTargetOutput() KubernetesAgentDeploymentTargetOutput
	ToKubernetesAgentDeploymentTargetOutputWithContext(ctx context.Context) KubernetesAgentDeploymentTargetOutput
}

func (*KubernetesAgentDeploymentTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesAgentDeploymentTarget)(nil)).Elem()
}

func (i *KubernetesAgentDeploymentTarget) ToKubernetesAgentDeploymentTargetOutput() KubernetesAgentDeploymentTargetOutput {
	return i.ToKubernetesAgentDeploymentTargetOutputWithContext(context.Background())
}

func (i *KubernetesAgentDeploymentTarget) ToKubernetesAgentDeploymentTargetOutputWithContext(ctx context.Context) KubernetesAgentDeploymentTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAgentDeploymentTargetOutput)
}

// KubernetesAgentDeploymentTargetArrayInput is an input type that accepts KubernetesAgentDeploymentTargetArray and KubernetesAgentDeploymentTargetArrayOutput values.
// You can construct a concrete instance of `KubernetesAgentDeploymentTargetArrayInput` via:
//
//	KubernetesAgentDeploymentTargetArray{ KubernetesAgentDeploymentTargetArgs{...} }
type KubernetesAgentDeploymentTargetArrayInput interface {
	pulumi.Input

	ToKubernetesAgentDeploymentTargetArrayOutput() KubernetesAgentDeploymentTargetArrayOutput
	ToKubernetesAgentDeploymentTargetArrayOutputWithContext(context.Context) KubernetesAgentDeploymentTargetArrayOutput
}

type KubernetesAgentDeploymentTargetArray []KubernetesAgentDeploymentTargetInput

func (KubernetesAgentDeploymentTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesAgentDeploymentTarget)(nil)).Elem()
}

func (i KubernetesAgentDeploymentTargetArray) ToKubernetesAgentDeploymentTargetArrayOutput() KubernetesAgentDeploymentTargetArrayOutput {
	return i.ToKubernetesAgentDeploymentTargetArrayOutputWithContext(context.Background())
}

func (i KubernetesAgentDeploymentTargetArray) ToKubernetesAgentDeploymentTargetArrayOutputWithContext(ctx context.Context) KubernetesAgentDeploymentTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAgentDeploymentTargetArrayOutput)
}

// KubernetesAgentDeploymentTargetMapInput is an input type that accepts KubernetesAgentDeploymentTargetMap and KubernetesAgentDeploymentTargetMapOutput values.
// You can construct a concrete instance of `KubernetesAgentDeploymentTargetMapInput` via:
//
//	KubernetesAgentDeploymentTargetMap{ "key": KubernetesAgentDeploymentTargetArgs{...} }
type KubernetesAgentDeploymentTargetMapInput interface {
	pulumi.Input

	ToKubernetesAgentDeploymentTargetMapOutput() KubernetesAgentDeploymentTargetMapOutput
	ToKubernetesAgentDeploymentTargetMapOutputWithContext(context.Context) KubernetesAgentDeploymentTargetMapOutput
}

type KubernetesAgentDeploymentTargetMap map[string]KubernetesAgentDeploymentTargetInput

func (KubernetesAgentDeploymentTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesAgentDeploymentTarget)(nil)).Elem()
}

func (i KubernetesAgentDeploymentTargetMap) ToKubernetesAgentDeploymentTargetMapOutput() KubernetesAgentDeploymentTargetMapOutput {
	return i.ToKubernetesAgentDeploymentTargetMapOutputWithContext(context.Background())
}

func (i KubernetesAgentDeploymentTargetMap) ToKubernetesAgentDeploymentTargetMapOutputWithContext(ctx context.Context) KubernetesAgentDeploymentTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAgentDeploymentTargetMapOutput)
}

type KubernetesAgentDeploymentTargetOutput struct{ *pulumi.OutputState }

func (KubernetesAgentDeploymentTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesAgentDeploymentTarget)(nil)).Elem()
}

func (o KubernetesAgentDeploymentTargetOutput) ToKubernetesAgentDeploymentTargetOutput() KubernetesAgentDeploymentTargetOutput {
	return o
}

func (o KubernetesAgentDeploymentTargetOutput) ToKubernetesAgentDeploymentTargetOutputWithContext(ctx context.Context) KubernetesAgentDeploymentTargetOutput {
	return o
}

// Name of the Helm release that the agent belongs to.
func (o KubernetesAgentDeploymentTargetOutput) AgentHelmReleaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.AgentHelmReleaseName }).(pulumi.StringOutput)
}

// Name of the Kubernetes namespace where the agent is installed.
func (o KubernetesAgentDeploymentTargetOutput) AgentKubernetesNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.AgentKubernetesNamespace }).(pulumi.StringOutput)
}

// Current Tentacle version of the agent
func (o KubernetesAgentDeploymentTargetOutput) AgentTentacleVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.AgentTentacleVersion }).(pulumi.StringOutput)
}

// Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
func (o KubernetesAgentDeploymentTargetOutput) AgentUpgradeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.AgentUpgradeStatus }).(pulumi.StringOutput)
}

// Current Helm chart version of the agent.
func (o KubernetesAgentDeploymentTargetOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

// The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
func (o KubernetesAgentDeploymentTargetOutput) CommunicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringPtrOutput { return v.CommunicationMode }).(pulumi.StringPtrOutput)
}

// Optional default namespace that will be used when using Kubernetes deployment steps, can be overrides within step configurations.
func (o KubernetesAgentDeploymentTargetOutput) DefaultNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.DefaultNamespace }).(pulumi.StringOutput)
}

// A list of environment IDs this Kubernetes agent can deploy to.
func (o KubernetesAgentDeploymentTargetOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

// Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
func (o KubernetesAgentDeploymentTargetOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
func (o KubernetesAgentDeploymentTargetOutput) MachinePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.MachinePolicyId }).(pulumi.StringOutput)
}

// The name of this resource.
func (o KubernetesAgentDeploymentTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of target roles that are associated to this Kubernetes agent.
func (o KubernetesAgentDeploymentTargetOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The space ID associated with this resource.
func (o KubernetesAgentDeploymentTargetOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// A list of tenant tags associated with this resource.
func (o KubernetesAgentDeploymentTargetOutput) TenantTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringArrayOutput { return v.TenantTags }).(pulumi.StringArrayOutput)
}

// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
func (o KubernetesAgentDeploymentTargetOutput) TenantedDeploymentParticipation() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.TenantedDeploymentParticipation }).(pulumi.StringOutput)
}

// A list of tenant IDs associated with this resource.
func (o KubernetesAgentDeploymentTargetOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

// The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
func (o KubernetesAgentDeploymentTargetOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
func (o KubernetesAgentDeploymentTargetOutput) UpgradeLocked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.BoolPtrOutput { return v.UpgradeLocked }).(pulumi.BoolPtrOutput)
}

// The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
func (o KubernetesAgentDeploymentTargetOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAgentDeploymentTarget) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

type KubernetesAgentDeploymentTargetArrayOutput struct{ *pulumi.OutputState }

func (KubernetesAgentDeploymentTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesAgentDeploymentTarget)(nil)).Elem()
}

func (o KubernetesAgentDeploymentTargetArrayOutput) ToKubernetesAgentDeploymentTargetArrayOutput() KubernetesAgentDeploymentTargetArrayOutput {
	return o
}

func (o KubernetesAgentDeploymentTargetArrayOutput) ToKubernetesAgentDeploymentTargetArrayOutputWithContext(ctx context.Context) KubernetesAgentDeploymentTargetArrayOutput {
	return o
}

func (o KubernetesAgentDeploymentTargetArrayOutput) Index(i pulumi.IntInput) KubernetesAgentDeploymentTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubernetesAgentDeploymentTarget {
		return vs[0].([]*KubernetesAgentDeploymentTarget)[vs[1].(int)]
	}).(KubernetesAgentDeploymentTargetOutput)
}

type KubernetesAgentDeploymentTargetMapOutput struct{ *pulumi.OutputState }

func (KubernetesAgentDeploymentTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesAgentDeploymentTarget)(nil)).Elem()
}

func (o KubernetesAgentDeploymentTargetMapOutput) ToKubernetesAgentDeploymentTargetMapOutput() KubernetesAgentDeploymentTargetMapOutput {
	return o
}

func (o KubernetesAgentDeploymentTargetMapOutput) ToKubernetesAgentDeploymentTargetMapOutputWithContext(ctx context.Context) KubernetesAgentDeploymentTargetMapOutput {
	return o
}

func (o KubernetesAgentDeploymentTargetMapOutput) MapIndex(k pulumi.StringInput) KubernetesAgentDeploymentTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubernetesAgentDeploymentTarget {
		return vs[0].(map[string]*KubernetesAgentDeploymentTarget)[vs[1].(string)]
	}).(KubernetesAgentDeploymentTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAgentDeploymentTargetInput)(nil)).Elem(), &KubernetesAgentDeploymentTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAgentDeploymentTargetArrayInput)(nil)).Elem(), KubernetesAgentDeploymentTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAgentDeploymentTargetMapInput)(nil)).Elem(), KubernetesAgentDeploymentTargetMap{})
	pulumi.RegisterOutputType(KubernetesAgentDeploymentTargetOutput{})
	pulumi.RegisterOutputType(KubernetesAgentDeploymentTargetArrayOutput{})
	pulumi.RegisterOutputType(KubernetesAgentDeploymentTargetMapOutput{})
}
