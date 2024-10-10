// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    /// <summary>
    /// This resource manages variables in Octopus Deploy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Octopusdeploy = Pulumi.Octopusdeploy;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // create an Amazon web services account variable
    ///     var amazonWebServicesAccountVariable = new Octopusdeploy.Variable("amazonWebServicesAccountVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "AmazonWebServicesAccount",
    ///         Value = "Accounts-123",
    ///     });
    /// 
    ///     // create an Azure service principal account variable
    ///     var azureServicePrincipalAccountVariable = new Octopusdeploy.Variable("azureServicePrincipalAccountVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "AzureAccount",
    ///         Value = "Accounts-123",
    ///     });
    /// 
    ///     // create a Google Cloud account variable
    ///     var googleCloudAccountVariable = new Octopusdeploy.Variable("googleCloudAccountVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "GoogleCloudAccount",
    ///         Value = "Accounts-123",
    ///     });
    /// 
    ///     // Create a UsernamePassword account variable
    ///     var usernamepasswordAccountVariable = new Octopusdeploy.Variable("usernamepasswordAccountVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "UsernamePasswordAccount",
    ///         Value = octopusdeploy_username_password_account.Account_user_pass.Id,
    ///     });
    /// 
    ///     // create a Certificate variable
    ///     var certificateVariable = new Octopusdeploy.Variable("certificateVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "Certificate",
    ///         Value = "Certificates-123",
    ///     });
    /// 
    ///     // create a Sensitive variable
    ///     var sensitiveVariable = new Octopusdeploy.Variable("sensitiveVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "Sensitive",
    ///         IsSensitive = true,
    ///         SensitiveValue = "YourSecrets",
    ///     });
    /// 
    ///     // create a String variable
    ///     var stringVariable = new Octopusdeploy.Variable("stringVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "String",
    ///         Value = "PlainText",
    ///     });
    /// 
    ///     // create a WorkerPool variable
    ///     var workerpoolVariable = new Octopusdeploy.Variable("workerpoolVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "WorkerPool",
    ///         Value = "WorkerPools-123",
    ///     });
    /// 
    ///     // create a prompted variable
    ///     var promptedVariable = new Octopusdeploy.Variable("promptedVariable", new()
    ///     {
    ///         OwnerId = "Projects-123",
    ///         Type = "String",
    ///         Prompt = new Octopusdeploy.Inputs.VariablePromptArgs
    ///         {
    ///             Description = "Variable Description",
    ///             IsRequired = true,
    ///             Label = "Variable Label",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import octopusdeploy:index/variable:Variable [options] octopusdeploy_variable.&lt;name&gt; &lt;variable-id&gt;
    /// ```
    /// </summary>
    [OctopusdeployResourceType("octopusdeploy:index/variable:Variable")]
    public partial class Variable : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this variable.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not this variable is considered editable.
        /// </summary>
        [Output("isEditable")]
        public Output<bool> IsEditable { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not this resource is considered sensitive and should be kept secret.
        /// </summary>
        [Output("isSensitive")]
        public Output<bool> IsSensitive { get; private set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("ownerId")]
        public Output<string?> OwnerId { get; private set; } = null!;

        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        [Output("prompt")]
        public Output<Outputs.VariablePrompt?> Prompt { get; private set; } = null!;

        [Output("scope")]
        public Output<Outputs.VariableScope?> Scope { get; private set; } = null!;

        [Output("sensitiveValue")]
        public Output<string?> SensitiveValue { get; private set; } = null!;

        /// <summary>
        /// The space ID associated with this variable.
        /// </summary>
        [Output("spaceId")]
        public Output<string> SpaceId { get; private set; } = null!;

        /// <summary>
        /// The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `UsernamePasswordAccount`, `Certificate`, `Sensitive`, `String`, `WorkerPool`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Variable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Variable(string name, VariableArgs args, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/variable:Variable", name, args ?? new VariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Variable(string name, Input<string> id, VariableState? state = null, CustomResourceOptions? options = null)
            : base("octopusdeploy:index/variable:Variable", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/OctopusDeploy/terraform-provider-octopusdeploy/releases",
                AdditionalSecretOutputs =
                {
                    "sensitiveValue",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Variable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Variable Get(string name, Input<string> id, VariableState? state = null, CustomResourceOptions? options = null)
        {
            return new Variable(name, id, state, options);
        }
    }

    public sealed class VariableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this variable.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether or not this variable is considered editable.
        /// </summary>
        [Input("isEditable")]
        public Input<bool>? IsEditable { get; set; }

        /// <summary>
        /// Indicates whether or not this resource is considered sensitive and should be kept secret.
        /// </summary>
        [Input("isSensitive")]
        public Input<bool>? IsSensitive { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("prompt")]
        public Input<Inputs.VariablePromptArgs>? Prompt { get; set; }

        [Input("scope")]
        public Input<Inputs.VariableScopeArgs>? Scope { get; set; }

        [Input("sensitiveValue")]
        private Input<string>? _sensitiveValue;
        public Input<string>? SensitiveValue
        {
            get => _sensitiveValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sensitiveValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The space ID associated with this variable.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `UsernamePasswordAccount`, `Certificate`, `Sensitive`, `String`, `WorkerPool`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public VariableArgs()
        {
        }
        public static new VariableArgs Empty => new VariableArgs();
    }

    public sealed class VariableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this variable.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether or not this variable is considered editable.
        /// </summary>
        [Input("isEditable")]
        public Input<bool>? IsEditable { get; set; }

        /// <summary>
        /// Indicates whether or not this resource is considered sensitive and should be kept secret.
        /// </summary>
        [Input("isSensitive")]
        public Input<bool>? IsSensitive { get; set; }

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("prompt")]
        public Input<Inputs.VariablePromptGetArgs>? Prompt { get; set; }

        [Input("scope")]
        public Input<Inputs.VariableScopeGetArgs>? Scope { get; set; }

        [Input("sensitiveValue")]
        private Input<string>? _sensitiveValue;
        public Input<string>? SensitiveValue
        {
            get => _sensitiveValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sensitiveValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The space ID associated with this variable.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `UsernamePasswordAccount`, `Certificate`, `Sensitive`, `String`, `WorkerPool`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public VariableState()
        {
        }
        public static new VariableState Empty => new VariableState();
    }
}
