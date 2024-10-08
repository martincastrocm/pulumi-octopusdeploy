// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource manages machine policies in Octopus Deploy.
 */
export class MachinePolicy extends pulumi.CustomResource {
    /**
     * Get an existing MachinePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MachinePolicyState, opts?: pulumi.CustomResourceOptions): MachinePolicy {
        return new MachinePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/machinePolicy:MachinePolicy';

    /**
     * Returns true if the given object is an instance of MachinePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachinePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachinePolicy.__pulumiType;
    }

    /**
     * In nanoseconds. Minimum value: 10000000000 (10 seconds).
     */
    public readonly connectionConnectTimeout!: pulumi.Output<number | undefined>;
    public readonly connectionRetryCountLimit!: pulumi.Output<number | undefined>;
    /**
     * In nanoseconds.
     */
    public readonly connectionRetrySleepInterval!: pulumi.Output<number | undefined>;
    /**
     * In nanoseconds.
     */
    public readonly connectionRetryTimeLimit!: pulumi.Output<number | undefined>;
    /**
     * The description of this machine policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    public readonly machineCleanupPolicy!: pulumi.Output<outputs.MachinePolicyMachineCleanupPolicy>;
    public readonly machineConnectivityPolicy!: pulumi.Output<outputs.MachinePolicyMachineConnectivityPolicy>;
    public readonly machineHealthCheckPolicy!: pulumi.Output<outputs.MachinePolicyMachineHealthCheckPolicy>;
    public readonly machineUpdatePolicy!: pulumi.Output<outputs.MachinePolicyMachineUpdatePolicy>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * In nanoseconds.
     */
    public readonly pollingRequestQueueTimeout!: pulumi.Output<number | undefined>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;

    /**
     * Create a MachinePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MachinePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MachinePolicyArgs | MachinePolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MachinePolicyState | undefined;
            resourceInputs["connectionConnectTimeout"] = state ? state.connectionConnectTimeout : undefined;
            resourceInputs["connectionRetryCountLimit"] = state ? state.connectionRetryCountLimit : undefined;
            resourceInputs["connectionRetrySleepInterval"] = state ? state.connectionRetrySleepInterval : undefined;
            resourceInputs["connectionRetryTimeLimit"] = state ? state.connectionRetryTimeLimit : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["machineCleanupPolicy"] = state ? state.machineCleanupPolicy : undefined;
            resourceInputs["machineConnectivityPolicy"] = state ? state.machineConnectivityPolicy : undefined;
            resourceInputs["machineHealthCheckPolicy"] = state ? state.machineHealthCheckPolicy : undefined;
            resourceInputs["machineUpdatePolicy"] = state ? state.machineUpdatePolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pollingRequestQueueTimeout"] = state ? state.pollingRequestQueueTimeout : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
        } else {
            const args = argsOrState as MachinePolicyArgs | undefined;
            resourceInputs["connectionConnectTimeout"] = args ? args.connectionConnectTimeout : undefined;
            resourceInputs["connectionRetryCountLimit"] = args ? args.connectionRetryCountLimit : undefined;
            resourceInputs["connectionRetrySleepInterval"] = args ? args.connectionRetrySleepInterval : undefined;
            resourceInputs["connectionRetryTimeLimit"] = args ? args.connectionRetryTimeLimit : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["machineCleanupPolicy"] = args ? args.machineCleanupPolicy : undefined;
            resourceInputs["machineConnectivityPolicy"] = args ? args.machineConnectivityPolicy : undefined;
            resourceInputs["machineHealthCheckPolicy"] = args ? args.machineHealthCheckPolicy : undefined;
            resourceInputs["machineUpdatePolicy"] = args ? args.machineUpdatePolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pollingRequestQueueTimeout"] = args ? args.pollingRequestQueueTimeout : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["isDefault"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MachinePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MachinePolicy resources.
 */
export interface MachinePolicyState {
    /**
     * In nanoseconds. Minimum value: 10000000000 (10 seconds).
     */
    connectionConnectTimeout?: pulumi.Input<number>;
    connectionRetryCountLimit?: pulumi.Input<number>;
    /**
     * In nanoseconds.
     */
    connectionRetrySleepInterval?: pulumi.Input<number>;
    /**
     * In nanoseconds.
     */
    connectionRetryTimeLimit?: pulumi.Input<number>;
    /**
     * The description of this machine policy.
     */
    description?: pulumi.Input<string>;
    isDefault?: pulumi.Input<boolean>;
    machineCleanupPolicy?: pulumi.Input<inputs.MachinePolicyMachineCleanupPolicy>;
    machineConnectivityPolicy?: pulumi.Input<inputs.MachinePolicyMachineConnectivityPolicy>;
    machineHealthCheckPolicy?: pulumi.Input<inputs.MachinePolicyMachineHealthCheckPolicy>;
    machineUpdatePolicy?: pulumi.Input<inputs.MachinePolicyMachineUpdatePolicy>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * In nanoseconds.
     */
    pollingRequestQueueTimeout?: pulumi.Input<number>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MachinePolicy resource.
 */
export interface MachinePolicyArgs {
    /**
     * In nanoseconds. Minimum value: 10000000000 (10 seconds).
     */
    connectionConnectTimeout?: pulumi.Input<number>;
    connectionRetryCountLimit?: pulumi.Input<number>;
    /**
     * In nanoseconds.
     */
    connectionRetrySleepInterval?: pulumi.Input<number>;
    /**
     * In nanoseconds.
     */
    connectionRetryTimeLimit?: pulumi.Input<number>;
    /**
     * The description of this machine policy.
     */
    description?: pulumi.Input<string>;
    machineCleanupPolicy?: pulumi.Input<inputs.MachinePolicyMachineCleanupPolicy>;
    machineConnectivityPolicy?: pulumi.Input<inputs.MachinePolicyMachineConnectivityPolicy>;
    machineHealthCheckPolicy?: pulumi.Input<inputs.MachinePolicyMachineHealthCheckPolicy>;
    machineUpdatePolicy?: pulumi.Input<inputs.MachinePolicyMachineUpdatePolicy>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * In nanoseconds.
     */
    pollingRequestQueueTimeout?: pulumi.Input<number>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
}
