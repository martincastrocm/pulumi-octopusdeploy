// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about existing machine policies.
 */
export function getMachinePolicies(args?: GetMachinePoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetMachinePoliciesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("octopusdeploy:index/getMachinePolicies:getMachinePolicies", {
        "ids": args.ids,
        "partialName": args.partialName,
        "skip": args.skip,
        "spaceId": args.spaceId,
        "take": args.take,
    }, opts);
}

/**
 * A collection of arguments for invoking getMachinePolicies.
 */
export interface GetMachinePoliciesArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: string;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: number;
    spaceId?: string;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: number;
}

/**
 * A collection of values returned by getMachinePolicies.
 */
export interface GetMachinePoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A list of machine policies that match the filter(s).
     */
    readonly machinePolicies: outputs.GetMachinePoliciesMachinePolicy[];
    /**
     * A filter to search by the partial match of a name.
     */
    readonly partialName?: string;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    readonly skip?: number;
    /**
     * The space ID associated with this resource.
     */
    readonly spaceId: string;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    readonly take?: number;
}
/**
 * Provides information about existing machine policies.
 */
export function getMachinePoliciesOutput(args?: GetMachinePoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMachinePoliciesResult> {
    return pulumi.output(args).apply((a: any) => getMachinePolicies(a, opts))
}

/**
 * A collection of arguments for invoking getMachinePolicies.
 */
export interface GetMachinePoliciesOutputArgs {
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: pulumi.Input<string>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    spaceId?: pulumi.Input<string>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
}
