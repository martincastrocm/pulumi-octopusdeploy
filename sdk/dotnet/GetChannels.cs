// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetChannels
    {
        /// <summary>
        /// Provides information about existing channels.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetChannels.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "Channels-123",
        ///             "Channels-321",
        ///         },
        ///         PartialName = "Defau",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetChannelsResult> InvokeAsync(GetChannelsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetChannelsResult>("octopusdeploy:index/getChannels:getChannels", args ?? new GetChannelsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing channels.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetChannels.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "Channels-123",
        ///             "Channels-321",
        ///         },
        ///         PartialName = "Defau",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetChannelsResult> Invoke(GetChannelsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetChannelsResult>("octopusdeploy:index/getChannels:getChannels", args ?? new GetChannelsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChannelsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public int? Skip { get; set; }

        [Input("spaceId")]
        public string? SpaceId { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public int? Take { get; set; }

        public GetChannelsArgs()
        {
        }
        public static new GetChannelsArgs Empty => new GetChannelsArgs();
    }

    public sealed class GetChannelsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public Input<int>? Skip { get; set; }

        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public Input<int>? Take { get; set; }

        public GetChannelsInvokeArgs()
        {
        }
        public static new GetChannelsInvokeArgs Empty => new GetChannelsInvokeArgs();
    }


    [OutputType]
    public sealed class GetChannelsResult
    {
        /// <summary>
        /// A channel that matches the specified filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetChannelsChannelResult> Channels;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;

        [OutputConstructor]
        private GetChannelsResult(
            ImmutableArray<Outputs.GetChannelsChannelResult> channels,

            string id,

            ImmutableArray<string> ids,

            string? partialName,

            int? skip,

            string spaceId,

            int? take)
        {
            Channels = channels;
            Id = id;
            Ids = ids;
            PartialName = partialName;
            Skip = skip;
            SpaceId = spaceId;
            Take = take;
        }
    }
}
