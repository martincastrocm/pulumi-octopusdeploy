// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetFeeds
    {
        /// <summary>
        /// Provides information about existing feeds.
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
        ///     var example = Octopusdeploy.GetFeeds.Invoke(new()
        ///     {
        ///         FeedType = "NuGet",
        ///         Ids = new[]
        ///         {
        ///             "Feeds-123",
        ///             "Feeds-321",
        ///         },
        ///         PartialName = "Develop",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFeedsResult> InvokeAsync(GetFeedsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeedsResult>("octopusdeploy:index/getFeeds:getFeeds", args ?? new GetFeedsArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing feeds.
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
        ///     var example = Octopusdeploy.GetFeeds.Invoke(new()
        ///     {
        ///         FeedType = "NuGet",
        ///         Ids = new[]
        ///         {
        ///             "Feeds-123",
        ///             "Feeds-321",
        ///         },
        ///         PartialName = "Develop",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFeedsResult> Invoke(GetFeedsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeedsResult>("octopusdeploy:index/getFeeds:getFeeds", args ?? new GetFeedsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeedsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
        /// </summary>
        [Input("feedType")]
        public string? FeedType { get; set; }

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
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A filter to search by a partial name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public int? Skip { get; set; }

        /// <summary>
        /// The space ID associated with this feeds.
        /// </summary>
        [Input("spaceId")]
        public string? SpaceId { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public int? Take { get; set; }

        public GetFeedsArgs()
        {
        }
        public static new GetFeedsArgs Empty => new GetFeedsArgs();
    }

    public sealed class GetFeedsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
        /// </summary>
        [Input("feedType")]
        public Input<string>? FeedType { get; set; }

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
        /// The name of this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A filter to search by a partial name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public Input<int>? Skip { get; set; }

        /// <summary>
        /// The space ID associated with this feeds.
        /// </summary>
        [Input("spaceId")]
        public Input<string>? SpaceId { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public Input<int>? Take { get; set; }

        public GetFeedsInvokeArgs()
        {
        }
        public static new GetFeedsInvokeArgs Empty => new GetFeedsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeedsResult
    {
        /// <summary>
        /// A filter to search by feed type. Valid feed types are `AwsElasticContainerRegistry`, `BuiltIn`, `Docker`, `GitHub`, `Helm`, `Maven`, `NuGet`, or `OctopusProject`.
        /// </summary>
        public readonly string? FeedType;
        public readonly ImmutableArray<Outputs.GetFeedsFeedResult> Feeds;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A filter to search by a partial name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// The space ID associated with this feeds.
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;

        [OutputConstructor]
        private GetFeedsResult(
            string? feedType,

            ImmutableArray<Outputs.GetFeedsFeedResult> feeds,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? partialName,

            int? skip,

            string spaceId,

            int? take)
        {
            FeedType = feedType;
            Feeds = feeds;
            Id = id;
            Ids = ids;
            Name = name;
            PartialName = partialName;
            Skip = skip;
            SpaceId = spaceId;
            Take = take;
        }
    }
}
