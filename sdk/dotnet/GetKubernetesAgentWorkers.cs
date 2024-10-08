// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetKubernetesAgentWorkers
    {
        /// <summary>
        /// Provides information about existing kubernetes agent workers.
        /// </summary>
        public static Task<GetKubernetesAgentWorkersResult> InvokeAsync(GetKubernetesAgentWorkersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesAgentWorkersResult>("octopusdeploy:index/getKubernetesAgentWorkers:getKubernetesAgentWorkers", args ?? new GetKubernetesAgentWorkersArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing kubernetes agent workers.
        /// </summary>
        public static Output<GetKubernetesAgentWorkersResult> Invoke(GetKubernetesAgentWorkersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubernetesAgentWorkersResult>("octopusdeploy:index/getKubernetesAgentWorkers:getKubernetesAgentWorkers", args ?? new GetKubernetesAgentWorkersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubernetesAgentWorkersArgs : global::Pulumi.InvokeArgs
    {
        [Input("healthStatuses")]
        private List<string>? _healthStatuses;

        /// <summary>
        /// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public List<string> HealthStatuses
        {
            get => _healthStatuses ?? (_healthStatuses = new List<string>());
            set => _healthStatuses = value;
        }

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

        [Input("isDisabled")]
        public bool? IsDisabled { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        [Input("roles")]
        private List<string>? _roles;

        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public List<string> Roles
        {
            get => _roles ?? (_roles = new List<string>());
            set => _roles = value;
        }

        [Input("shellNames")]
        private List<string>? _shellNames;

        /// <summary>
        /// A list of shell names to match in the query and/or search
        /// </summary>
        public List<string> ShellNames
        {
            get => _shellNames ?? (_shellNames = new List<string>());
            set => _shellNames = value;
        }

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

        [Input("thumbprint")]
        public string? Thumbprint { get; set; }

        public GetKubernetesAgentWorkersArgs()
        {
        }
        public static new GetKubernetesAgentWorkersArgs Empty => new GetKubernetesAgentWorkersArgs();
    }

    public sealed class GetKubernetesAgentWorkersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("healthStatuses")]
        private InputList<string>? _healthStatuses;

        /// <summary>
        /// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public InputList<string> HealthStatuses
        {
            get => _healthStatuses ?? (_healthStatuses = new InputList<string>());
            set => _healthStatuses = value;
        }

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

        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("shellNames")]
        private InputList<string>? _shellNames;

        /// <summary>
        /// A list of shell names to match in the query and/or search
        /// </summary>
        public InputList<string> ShellNames
        {
            get => _shellNames ?? (_shellNames = new InputList<string>());
            set => _shellNames = value;
        }

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

        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        public GetKubernetesAgentWorkersInvokeArgs()
        {
        }
        public static new GetKubernetesAgentWorkersInvokeArgs Empty => new GetKubernetesAgentWorkersInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubernetesAgentWorkersResult
    {
        /// <summary>
        /// A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public readonly ImmutableArray<string> HealthStatuses;
        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A filter to search by the disabled status of a resource.
        /// </summary>
        public readonly bool? IsDisabled;
        /// <summary>
        /// A list of kubernetes agent workers that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesAgentWorkersKubernetesAgentWorkerResult> KubernetesAgentWorkers;
        /// <summary>
        /// A filter to search by name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A filter to search by a list of role IDs.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// A list of shell names to match in the query and/or search
        /// </summary>
        public readonly ImmutableArray<string> ShellNames;
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
        /// <summary>
        /// The thumbprint of the deployment target to match in the query and/or search
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private GetKubernetesAgentWorkersResult(
            ImmutableArray<string> healthStatuses,

            string id,

            ImmutableArray<string> ids,

            bool? isDisabled,

            ImmutableArray<Outputs.GetKubernetesAgentWorkersKubernetesAgentWorkerResult> kubernetesAgentWorkers,

            string? name,

            string? partialName,

            ImmutableArray<string> roles,

            ImmutableArray<string> shellNames,

            int? skip,

            string spaceId,

            int? take,

            string? thumbprint)
        {
            HealthStatuses = healthStatuses;
            Id = id;
            Ids = ids;
            IsDisabled = isDisabled;
            KubernetesAgentWorkers = kubernetesAgentWorkers;
            Name = name;
            PartialName = partialName;
            Roles = roles;
            ShellNames = shellNames;
            Skip = skip;
            SpaceId = spaceId;
            Take = take;
            Thumbprint = thumbprint;
        }
    }
}
