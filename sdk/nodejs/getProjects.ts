// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about existing Octopus Deploy projects.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getProjects({
 *     clonedFromProjectId: "Projects-456",
 *     ids: [
 *         "Projects-123",
 *         "Projects-321",
 *     ],
 *     isClone: true,
 *     name: "Default",
 *     partialName: "Defau",
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("octopusdeploy:index/getProjects:getProjects", {
        "clonedFromProjectId": args.clonedFromProjectId,
        "ids": args.ids,
        "isClone": args.isClone,
        "name": args.name,
        "partialName": args.partialName,
        "skip": args.skip,
        "spaceId": args.spaceId,
        "take": args.take,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsArgs {
    /**
     * A filter to search for cloned resources by a project ID.
     */
    clonedFromProjectId?: string;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    /**
     * A filter to search for cloned resources.
     */
    isClone?: boolean;
    /**
     * A filter to search by name
     */
    name?: string;
    /**
     * A filter to search by a partial name.
     */
    partialName?: string;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: number;
    /**
     * A Space ID to filter by. Will revert what is specified on the provider if not set
     */
    spaceId?: string;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: number;
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    /**
     * A filter to search for cloned resources by a project ID.
     */
    readonly clonedFromProjectId?: string;
    /**
     * An auto-generated identifier that includes the timestamp when this data source was last modified.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A filter to search for cloned resources.
     */
    readonly isClone?: boolean;
    /**
     * A filter to search by name
     */
    readonly name?: string;
    /**
     * A filter to search by a partial name.
     */
    readonly partialName?: string;
    /**
     * A list of projects that match the filter(s).
     */
    readonly projects: outputs.GetProjectsProject[];
    /**
     * A filter to specify the number of items to skip in the response.
     */
    readonly skip?: number;
    /**
     * A Space ID to filter by. Will revert what is specified on the provider if not set
     */
    readonly spaceId?: string;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    readonly take?: number;
}
/**
 * Provides information about existing Octopus Deploy projects.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getProjects({
 *     clonedFromProjectId: "Projects-456",
 *     ids: [
 *         "Projects-123",
 *         "Projects-321",
 *     ],
 *     isClone: true,
 *     name: "Default",
 *     partialName: "Defau",
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 */
export function getProjectsOutput(args?: GetProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("octopusdeploy:index/getProjects:getProjects", {
        "clonedFromProjectId": args.clonedFromProjectId,
        "ids": args.ids,
        "isClone": args.isClone,
        "name": args.name,
        "partialName": args.partialName,
        "skip": args.skip,
        "spaceId": args.spaceId,
        "take": args.take,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsOutputArgs {
    /**
     * A filter to search for cloned resources by a project ID.
     */
    clonedFromProjectId?: pulumi.Input<string>;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search for cloned resources.
     */
    isClone?: pulumi.Input<boolean>;
    /**
     * A filter to search by name
     */
    name?: pulumi.Input<string>;
    /**
     * A filter to search by a partial name.
     */
    partialName?: pulumi.Input<string>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    /**
     * A Space ID to filter by. Will revert what is specified on the provider if not set
     */
    spaceId?: pulumi.Input<string>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
}
