// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides information about existing Azure cloud service deployment targets.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getAzureCloudServiceDeploymentTargets({
 *     healthStatuses: [
 *         "Healthy",
 *         "Unavailable",
 *     ],
 *     ids: [
 *         "Machines-123",
 *         "Machines-321",
 *     ],
 *     partialName: "Defau",
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAzureCloudServiceDeploymentTargets(args?: GetAzureCloudServiceDeploymentTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetAzureCloudServiceDeploymentTargetsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("octopusdeploy:index/getAzureCloudServiceDeploymentTargets:getAzureCloudServiceDeploymentTargets", {
        "deploymentId": args.deploymentId,
        "environments": args.environments,
        "healthStatuses": args.healthStatuses,
        "ids": args.ids,
        "isDisabled": args.isDisabled,
        "name": args.name,
        "partialName": args.partialName,
        "roles": args.roles,
        "shellNames": args.shellNames,
        "skip": args.skip,
        "spaceId": args.spaceId,
        "take": args.take,
        "tenantTags": args.tenantTags,
        "tenants": args.tenants,
        "thumbprint": args.thumbprint,
    }, opts);
}

/**
 * A collection of arguments for invoking getAzureCloudServiceDeploymentTargets.
 */
export interface GetAzureCloudServiceDeploymentTargetsArgs {
    /**
     * A filter to search by deployment ID.
     */
    deploymentId?: string;
    environments?: string[];
    /**
     * A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatuses?: string[];
    /**
     * A filter to search by a list of IDs.
     */
    ids?: string[];
    isDisabled?: boolean;
    name?: string;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: string;
    roles?: string[];
    /**
     * A list of shell names to match in the query and/or search
     */
    shellNames?: string[];
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: number;
    spaceId?: string;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: number;
    tenantTags?: string[];
    tenants?: string[];
    thumbprint?: string;
}

/**
 * A collection of values returned by getAzureCloudServiceDeploymentTargets.
 */
export interface GetAzureCloudServiceDeploymentTargetsResult {
    /**
     * A list of Azure cloud service deployment targets that match the filter(s).
     */
    readonly azureCloudServiceDeploymentTargets: outputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTarget[];
    /**
     * A filter to search by deployment ID.
     */
    readonly deploymentId?: string;
    /**
     * A filter to search by a list of environment IDs.
     */
    readonly environments?: string[];
    /**
     * A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    readonly healthStatuses?: string[];
    /**
     * An auto-generated identifier that includes the timestamp when this data source was last modified.
     */
    readonly id: string;
    /**
     * A filter to search by a list of IDs.
     */
    readonly ids?: string[];
    /**
     * A filter to search by the disabled status of a resource.
     */
    readonly isDisabled?: boolean;
    /**
     * A filter to search by name.
     */
    readonly name?: string;
    /**
     * A filter to search by the partial match of a name.
     */
    readonly partialName?: string;
    /**
     * A filter to search by a list of role IDs.
     */
    readonly roles?: string[];
    /**
     * A list of shell names to match in the query and/or search
     */
    readonly shellNames?: string[];
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
    /**
     * A filter to search by a list of tenant tags.
     */
    readonly tenantTags?: string[];
    /**
     * A filter to search by a list of tenant IDs.
     */
    readonly tenants?: string[];
    /**
     * The thumbprint of the deployment target to match in the query and/or search
     */
    readonly thumbprint?: string;
}
/**
 * Provides information about existing Azure cloud service deployment targets.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = octopusdeploy.getAzureCloudServiceDeploymentTargets({
 *     healthStatuses: [
 *         "Healthy",
 *         "Unavailable",
 *     ],
 *     ids: [
 *         "Machines-123",
 *         "Machines-321",
 *     ],
 *     partialName: "Defau",
 *     skip: 5,
 *     take: 100,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAzureCloudServiceDeploymentTargetsOutput(args?: GetAzureCloudServiceDeploymentTargetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAzureCloudServiceDeploymentTargetsResult> {
    return pulumi.output(args).apply((a: any) => getAzureCloudServiceDeploymentTargets(a, opts))
}

/**
 * A collection of arguments for invoking getAzureCloudServiceDeploymentTargets.
 */
export interface GetAzureCloudServiceDeploymentTargetsOutputArgs {
    /**
     * A filter to search by deployment ID.
     */
    deploymentId?: pulumi.Input<string>;
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
     */
    healthStatuses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to search by a list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    isDisabled?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * A filter to search by the partial match of a name.
     */
    partialName?: pulumi.Input<string>;
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of shell names to match in the query and/or search
     */
    shellNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A filter to specify the number of items to skip in the response.
     */
    skip?: pulumi.Input<number>;
    spaceId?: pulumi.Input<string>;
    /**
     * A filter to specify the number of items to take (or return) in the response.
     */
    take?: pulumi.Input<number>;
    tenantTags?: pulumi.Input<pulumi.Input<string>[]>;
    tenants?: pulumi.Input<pulumi.Input<string>[]>;
    thumbprint?: pulumi.Input<string>;
}
