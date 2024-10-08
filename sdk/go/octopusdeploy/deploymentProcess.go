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

// This resource manages deployment processes in Octopus Deploy.
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
//			// basic deployment process with 2 run a script steps
//			_, err := octopusdeploy.NewDeploymentProcess(ctx, "example", &octopusdeploy.DeploymentProcessArgs{
//				ProjectId: pulumi.String("Projects-123"),
//				Steps: octopusdeploy.DeploymentProcessStepArray{
//					&octopusdeploy.DeploymentProcessStepArgs{
//						Condition:          pulumi.String("Success"),
//						Name:               pulumi.String("Hello world (using PowerShell)"),
//						PackageRequirement: pulumi.String("LetOctopusDecide"),
//						RunScriptActions: octopusdeploy.DeploymentProcessStepRunScriptActionArray{
//							&octopusdeploy.DeploymentProcessStepRunScriptActionArgs{
//								CanBeUsedForProjectVersioning: pulumi.Bool(false),
//								Condition:                     pulumi.String("Success"),
//								IsDisabled:                    pulumi.Bool(false),
//								IsRequired:                    pulumi.Bool(true),
//								Name:                          pulumi.String("Hello world (using PowerShell)"),
//								RunOnServer:                   pulumi.Bool(true),
//								ScriptBody:                    pulumi.String("  Write-Host 'Hello world, using PowerShell'\n  #TODO: Experiment with steps of your own :)\n  Write-Host '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'\n\n"),
//							},
//						},
//						StartTrigger: pulumi.String("StartAfterPrevious"),
//					},
//					&octopusdeploy.DeploymentProcessStepArgs{
//						Condition:          pulumi.String("Success"),
//						Name:               pulumi.String("Hello world (using Bash)"),
//						PackageRequirement: pulumi.String("LetOctopusDecide"),
//						RunScriptActions: octopusdeploy.DeploymentProcessStepRunScriptActionArray{
//							&octopusdeploy.DeploymentProcessStepRunScriptActionArgs{
//								CanBeUsedForProjectVersioning: pulumi.Bool(false),
//								Condition:                     pulumi.String("Success"),
//								IsDisabled:                    pulumi.Bool(false),
//								IsRequired:                    pulumi.Bool(true),
//								Name:                          pulumi.String("Hello world (using Bash)"),
//								RunOnServer:                   pulumi.Bool(true),
//								ScriptBody:                    pulumi.String("  echo 'Hello world, using Bash'\n  #TODO: Experiment with steps of your own :)\n  echo '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'\n\n"),
//							},
//						},
//						StartTrigger: pulumi.String("StartWithPrevious"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// basic deployment process with 2 run a script steps as child steps
//			_, err = octopusdeploy.NewDeploymentProcess(ctx, "childStepExample", &octopusdeploy.DeploymentProcessArgs{
//				ProjectId: pulumi.String("Projects-123"),
//				Steps: octopusdeploy.DeploymentProcessStepArray{
//					&octopusdeploy.DeploymentProcessStepArgs{
//						Condition:          pulumi.String("Success"),
//						Name:               pulumi.String("Hello world (using PowerShell)"),
//						PackageRequirement: pulumi.String("LetOctopusDecide"),
//						RunScriptActions: octopusdeploy.DeploymentProcessStepRunScriptActionArray{
//							&octopusdeploy.DeploymentProcessStepRunScriptActionArgs{
//								CanBeUsedForProjectVersioning: pulumi.Bool(false),
//								Condition:                     pulumi.String("Success"),
//								IsDisabled:                    pulumi.Bool(false),
//								IsRequired:                    pulumi.Bool(true),
//								Name:                          pulumi.String("Hello world (using PowerShell)"),
//								ScriptBody:                    pulumi.String("  Write-Host 'Hello world, using PowerShell'\n  #TODO: Experiment with steps of your own :)\n  Write-Host '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'\n\n"),
//								SortOrder:                     pulumi.Int(1),
//							},
//							&octopusdeploy.DeploymentProcessStepRunScriptActionArgs{
//								CanBeUsedForProjectVersioning: pulumi.Bool(false),
//								Condition:                     pulumi.String("Success"),
//								IsDisabled:                    pulumi.Bool(false),
//								IsRequired:                    pulumi.Bool(true),
//								Name:                          pulumi.String("Hello world (using Bash)"),
//								ScriptBody:                    pulumi.String("  echo 'Hello world, using Bash'\n  #TODO: Experiment with steps of your own :)\n  echo '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'\n\n"),
//								SortOrder:                     pulumi.Int(2),
//							},
//						},
//						StartTrigger: pulumi.String("StartAfterPrevious"),
//						TargetRoles: pulumi.StringArray{
//							pulumi.String("hello-world"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// rolling deployment process with a step with 2 run a script steps as child steps deploying to 2 targets in parallel
//			_, err = octopusdeploy.NewDeploymentProcess(ctx, "childStepRollingDeploymentExample", &octopusdeploy.DeploymentProcessArgs{
//				ProjectId: pulumi.String("Projects-123"),
//				Steps: octopusdeploy.DeploymentProcessStepArray{
//					&octopusdeploy.DeploymentProcessStepArgs{
//						Condition:          pulumi.String("Success"),
//						Name:               pulumi.String("Hello world (using PowerShell)"),
//						PackageRequirement: pulumi.String("LetOctopusDecide"),
//						RunScriptActions: octopusdeploy.DeploymentProcessStepRunScriptActionArray{
//							&octopusdeploy.DeploymentProcessStepRunScriptActionArgs{
//								CanBeUsedForProjectVersioning: pulumi.Bool(false),
//								Condition:                     pulumi.String("Success"),
//								IsDisabled:                    pulumi.Bool(false),
//								IsRequired:                    pulumi.Bool(true),
//								Name:                          pulumi.String("Hello world (using PowerShell)"),
//								ScriptBody:                    pulumi.String("  Write-Host 'Hello world, using PowerShell'\n  #TODO: Experiment with steps of your own :)\n  Write-Host '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'\n\n"),
//								SortOrder:                     pulumi.Int(1),
//							},
//							&octopusdeploy.DeploymentProcessStepRunScriptActionArgs{
//								CanBeUsedForProjectVersioning: pulumi.Bool(false),
//								Condition:                     pulumi.String("Success"),
//								IsDisabled:                    pulumi.Bool(false),
//								IsRequired:                    pulumi.Bool(true),
//								Name:                          pulumi.String("Hello world (using Bash)"),
//								ScriptBody:                    pulumi.String("  echo 'Hello world, using Bash'\n  #TODO: Experiment with steps of your own :)\n  echo '[Learn more about the types of steps available in Octopus](https://g.octopushq.com/OnboardingAddStepsLearnMore)'\n\n"),
//								SortOrder:                     pulumi.Int(2),
//							},
//						},
//						StartTrigger: pulumi.String("StartAfterPrevious"),
//						TargetRoles: pulumi.StringArray{
//							pulumi.String("hello-world"),
//						},
//						WindowSize: pulumi.String("2"),
//					},
//				},
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
// $ pulumi import octopusdeploy:index/deploymentProcess:DeploymentProcess [options] octopusdeploy_deployment_process.<name> <deployment-process-id>
// ```
type DeploymentProcess struct {
	pulumi.CustomResourceState

	// The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
	Branch         pulumi.StringOutput    `pulumi:"branch"`
	LastSnapshotId pulumi.StringPtrOutput `pulumi:"lastSnapshotId"`
	// The project ID associated with this deployment process.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The space ID associated with this resource.
	SpaceId pulumi.StringOutput              `pulumi:"spaceId"`
	Steps   DeploymentProcessStepArrayOutput `pulumi:"steps"`
	// The version number of this deployment process.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewDeploymentProcess registers a new resource with the given unique name, arguments, and options.
func NewDeploymentProcess(ctx *pulumi.Context,
	name string, args *DeploymentProcessArgs, opts ...pulumi.ResourceOption) (*DeploymentProcess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeploymentProcess
	err := ctx.RegisterResource("octopusdeploy:index/deploymentProcess:DeploymentProcess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentProcess gets an existing DeploymentProcess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentProcess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentProcessState, opts ...pulumi.ResourceOption) (*DeploymentProcess, error) {
	var resource DeploymentProcess
	err := ctx.ReadResource("octopusdeploy:index/deploymentProcess:DeploymentProcess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentProcess resources.
type deploymentProcessState struct {
	// The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
	Branch         *string `pulumi:"branch"`
	LastSnapshotId *string `pulumi:"lastSnapshotId"`
	// The project ID associated with this deployment process.
	ProjectId *string `pulumi:"projectId"`
	// The space ID associated with this resource.
	SpaceId *string                 `pulumi:"spaceId"`
	Steps   []DeploymentProcessStep `pulumi:"steps"`
	// The version number of this deployment process.
	Version *int `pulumi:"version"`
}

type DeploymentProcessState struct {
	// The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
	Branch         pulumi.StringPtrInput
	LastSnapshotId pulumi.StringPtrInput
	// The project ID associated with this deployment process.
	ProjectId pulumi.StringPtrInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	Steps   DeploymentProcessStepArrayInput
	// The version number of this deployment process.
	Version pulumi.IntPtrInput
}

func (DeploymentProcessState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentProcessState)(nil)).Elem()
}

type deploymentProcessArgs struct {
	// The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
	Branch         *string `pulumi:"branch"`
	LastSnapshotId *string `pulumi:"lastSnapshotId"`
	// The project ID associated with this deployment process.
	ProjectId string `pulumi:"projectId"`
	// The space ID associated with this resource.
	SpaceId *string                 `pulumi:"spaceId"`
	Steps   []DeploymentProcessStep `pulumi:"steps"`
	// The version number of this deployment process.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a DeploymentProcess resource.
type DeploymentProcessArgs struct {
	// The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
	Branch         pulumi.StringPtrInput
	LastSnapshotId pulumi.StringPtrInput
	// The project ID associated with this deployment process.
	ProjectId pulumi.StringInput
	// The space ID associated with this resource.
	SpaceId pulumi.StringPtrInput
	Steps   DeploymentProcessStepArrayInput
	// The version number of this deployment process.
	Version pulumi.IntPtrInput
}

func (DeploymentProcessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentProcessArgs)(nil)).Elem()
}

type DeploymentProcessInput interface {
	pulumi.Input

	ToDeploymentProcessOutput() DeploymentProcessOutput
	ToDeploymentProcessOutputWithContext(ctx context.Context) DeploymentProcessOutput
}

func (*DeploymentProcess) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProcess)(nil)).Elem()
}

func (i *DeploymentProcess) ToDeploymentProcessOutput() DeploymentProcessOutput {
	return i.ToDeploymentProcessOutputWithContext(context.Background())
}

func (i *DeploymentProcess) ToDeploymentProcessOutputWithContext(ctx context.Context) DeploymentProcessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentProcessOutput)
}

// DeploymentProcessArrayInput is an input type that accepts DeploymentProcessArray and DeploymentProcessArrayOutput values.
// You can construct a concrete instance of `DeploymentProcessArrayInput` via:
//
//	DeploymentProcessArray{ DeploymentProcessArgs{...} }
type DeploymentProcessArrayInput interface {
	pulumi.Input

	ToDeploymentProcessArrayOutput() DeploymentProcessArrayOutput
	ToDeploymentProcessArrayOutputWithContext(context.Context) DeploymentProcessArrayOutput
}

type DeploymentProcessArray []DeploymentProcessInput

func (DeploymentProcessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeploymentProcess)(nil)).Elem()
}

func (i DeploymentProcessArray) ToDeploymentProcessArrayOutput() DeploymentProcessArrayOutput {
	return i.ToDeploymentProcessArrayOutputWithContext(context.Background())
}

func (i DeploymentProcessArray) ToDeploymentProcessArrayOutputWithContext(ctx context.Context) DeploymentProcessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentProcessArrayOutput)
}

// DeploymentProcessMapInput is an input type that accepts DeploymentProcessMap and DeploymentProcessMapOutput values.
// You can construct a concrete instance of `DeploymentProcessMapInput` via:
//
//	DeploymentProcessMap{ "key": DeploymentProcessArgs{...} }
type DeploymentProcessMapInput interface {
	pulumi.Input

	ToDeploymentProcessMapOutput() DeploymentProcessMapOutput
	ToDeploymentProcessMapOutputWithContext(context.Context) DeploymentProcessMapOutput
}

type DeploymentProcessMap map[string]DeploymentProcessInput

func (DeploymentProcessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeploymentProcess)(nil)).Elem()
}

func (i DeploymentProcessMap) ToDeploymentProcessMapOutput() DeploymentProcessMapOutput {
	return i.ToDeploymentProcessMapOutputWithContext(context.Background())
}

func (i DeploymentProcessMap) ToDeploymentProcessMapOutputWithContext(ctx context.Context) DeploymentProcessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentProcessMapOutput)
}

type DeploymentProcessOutput struct{ *pulumi.OutputState }

func (DeploymentProcessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProcess)(nil)).Elem()
}

func (o DeploymentProcessOutput) ToDeploymentProcessOutput() DeploymentProcessOutput {
	return o
}

func (o DeploymentProcessOutput) ToDeploymentProcessOutputWithContext(ctx context.Context) DeploymentProcessOutput {
	return o
}

// The branch name associated with this deployment process (i.e. `main`). This value is optional and only applies to associated projects that are stored in version control.
func (o DeploymentProcessOutput) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentProcess) pulumi.StringOutput { return v.Branch }).(pulumi.StringOutput)
}

func (o DeploymentProcessOutput) LastSnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentProcess) pulumi.StringPtrOutput { return v.LastSnapshotId }).(pulumi.StringPtrOutput)
}

// The project ID associated with this deployment process.
func (o DeploymentProcessOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentProcess) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The space ID associated with this resource.
func (o DeploymentProcessOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentProcess) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

func (o DeploymentProcessOutput) Steps() DeploymentProcessStepArrayOutput {
	return o.ApplyT(func(v *DeploymentProcess) DeploymentProcessStepArrayOutput { return v.Steps }).(DeploymentProcessStepArrayOutput)
}

// The version number of this deployment process.
func (o DeploymentProcessOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *DeploymentProcess) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type DeploymentProcessArrayOutput struct{ *pulumi.OutputState }

func (DeploymentProcessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeploymentProcess)(nil)).Elem()
}

func (o DeploymentProcessArrayOutput) ToDeploymentProcessArrayOutput() DeploymentProcessArrayOutput {
	return o
}

func (o DeploymentProcessArrayOutput) ToDeploymentProcessArrayOutputWithContext(ctx context.Context) DeploymentProcessArrayOutput {
	return o
}

func (o DeploymentProcessArrayOutput) Index(i pulumi.IntInput) DeploymentProcessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeploymentProcess {
		return vs[0].([]*DeploymentProcess)[vs[1].(int)]
	}).(DeploymentProcessOutput)
}

type DeploymentProcessMapOutput struct{ *pulumi.OutputState }

func (DeploymentProcessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeploymentProcess)(nil)).Elem()
}

func (o DeploymentProcessMapOutput) ToDeploymentProcessMapOutput() DeploymentProcessMapOutput {
	return o
}

func (o DeploymentProcessMapOutput) ToDeploymentProcessMapOutputWithContext(ctx context.Context) DeploymentProcessMapOutput {
	return o
}

func (o DeploymentProcessMapOutput) MapIndex(k pulumi.StringInput) DeploymentProcessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeploymentProcess {
		return vs[0].(map[string]*DeploymentProcess)[vs[1].(string)]
	}).(DeploymentProcessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentProcessInput)(nil)).Elem(), &DeploymentProcess{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentProcessArrayInput)(nil)).Elem(), DeploymentProcessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentProcessMapInput)(nil)).Elem(), DeploymentProcessMap{})
	pulumi.RegisterOutputType(DeploymentProcessOutput{})
	pulumi.RegisterOutputType(DeploymentProcessArrayOutput{})
	pulumi.RegisterOutputType(DeploymentProcessMapOutput{})
}
