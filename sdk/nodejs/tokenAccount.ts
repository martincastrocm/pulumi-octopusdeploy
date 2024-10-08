// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as octopusdeploy from "@pulumi/octopusdeploy";
 *
 * const example = new octopusdeploy.TokenAccount("example", {token: "[token]"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import octopusdeploy:index/tokenAccount:TokenAccount [options] octopusdeploy_token_account.<name> <account-id>
 * ```
 */
export class TokenAccount extends pulumi.CustomResource {
    /**
     * Get an existing TokenAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TokenAccountState, opts?: pulumi.CustomResourceOptions): TokenAccount {
        return new TokenAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/tokenAccount:TokenAccount';

    /**
     * Returns true if the given object is an instance of TokenAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TokenAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TokenAccount.__pulumiType;
    }

    /**
     * The description of this token account.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of environment IDs associated with this resource.
     */
    public readonly environments!: pulumi.Output<string[]>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The space ID associated with this resource.
     */
    public readonly spaceId!: pulumi.Output<string>;
    /**
     * A list of tenant tags associated with this resource.
     */
    public readonly tenantTags!: pulumi.Output<string[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    public readonly tenantedDeploymentParticipation!: pulumi.Output<string>;
    /**
     * A list of tenant IDs associated with this resource.
     */
    public readonly tenants!: pulumi.Output<string[]>;
    /**
     * The token of this resource.
     */
    public readonly token!: pulumi.Output<string>;

    /**
     * Create a TokenAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TokenAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TokenAccountArgs | TokenAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TokenAccountState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environments"] = state ? state.environments : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["tenantTags"] = state ? state.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = state ? state.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = state ? state.tenants : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
        } else {
            const args = argsOrState as TokenAccountArgs | undefined;
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environments"] = args ? args.environments : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["tenantTags"] = args ? args.tenantTags : undefined;
            resourceInputs["tenantedDeploymentParticipation"] = args ? args.tenantedDeploymentParticipation : undefined;
            resourceInputs["tenants"] = args ? args.tenants : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(TokenAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TokenAccount resources.
 */
export interface TokenAccountState {
    /**
     * The description of this token account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * A list of tenant tags associated with this resource.
     */
    tenantTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    tenantedDeploymentParticipation?: pulumi.Input<string>;
    /**
     * A list of tenant IDs associated with this resource.
     */
    tenants?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The token of this resource.
     */
    token?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TokenAccount resource.
 */
export interface TokenAccountArgs {
    /**
     * The description of this token account.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of environment IDs associated with this resource.
     */
    environments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this resource.
     */
    spaceId?: pulumi.Input<string>;
    /**
     * A list of tenant tags associated with this resource.
     */
    tenantTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
     */
    tenantedDeploymentParticipation?: pulumi.Input<string>;
    /**
     * A list of tenant IDs associated with this resource.
     */
    tenants?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The token of this resource.
     */
    token: pulumi.Input<string>;
}
