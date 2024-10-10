// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource manages library variable sets in Octopus Deploy.
 */
export class LibraryVariableSet extends pulumi.CustomResource {
    /**
     * Get an existing LibraryVariableSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LibraryVariableSetState, opts?: pulumi.CustomResourceOptions): LibraryVariableSet {
        return new LibraryVariableSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'octopusdeploy:index/libraryVariableSet:LibraryVariableSet';

    /**
     * Returns true if the given object is an instance of LibraryVariableSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LibraryVariableSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LibraryVariableSet.__pulumiType;
    }

    /**
     * The description of this library variable set.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The space ID associated with this library variable set.
     */
    public readonly spaceId!: pulumi.Output<string>;
    public /*out*/ readonly templateIds!: pulumi.Output<{[key: string]: string}>;
    public readonly templates!: pulumi.Output<outputs.LibraryVariableSetTemplate[] | undefined>;
    public /*out*/ readonly variableSetId!: pulumi.Output<string>;

    /**
     * Create a LibraryVariableSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LibraryVariableSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LibraryVariableSetArgs | LibraryVariableSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LibraryVariableSetState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["spaceId"] = state ? state.spaceId : undefined;
            resourceInputs["templateIds"] = state ? state.templateIds : undefined;
            resourceInputs["templates"] = state ? state.templates : undefined;
            resourceInputs["variableSetId"] = state ? state.variableSetId : undefined;
        } else {
            const args = argsOrState as LibraryVariableSetArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["spaceId"] = args ? args.spaceId : undefined;
            resourceInputs["templates"] = args ? args.templates : undefined;
            resourceInputs["templateIds"] = undefined /*out*/;
            resourceInputs["variableSetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LibraryVariableSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LibraryVariableSet resources.
 */
export interface LibraryVariableSetState {
    /**
     * The description of this library variable set.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this library variable set.
     */
    spaceId?: pulumi.Input<string>;
    templateIds?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    templates?: pulumi.Input<pulumi.Input<inputs.LibraryVariableSetTemplate>[]>;
    variableSetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LibraryVariableSet resource.
 */
export interface LibraryVariableSetArgs {
    /**
     * The description of this library variable set.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of this resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The space ID associated with this library variable set.
     */
    spaceId?: pulumi.Input<string>;
    templates?: pulumi.Input<pulumi.Input<inputs.LibraryVariableSetTemplate>[]>;
}
