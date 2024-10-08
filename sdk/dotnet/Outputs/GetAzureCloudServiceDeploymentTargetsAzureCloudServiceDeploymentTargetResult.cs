// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetResult
    {
        public readonly string AccountId;
        public readonly string CloudServiceName;
        public readonly string DefaultWorkerPoolId;
        public readonly ImmutableArray<Outputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointResult> Endpoints;
        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> Environments;
        public readonly bool HasLatestCalamari;
        /// <summary>
        /// Represents the health status of this deployment target. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        /// </summary>
        public readonly string HealthStatus;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsDisabled;
        public readonly bool IsInProcess;
        public readonly string MachinePolicyId;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        public readonly string OperatingSystem;
        public readonly ImmutableArray<string> Roles;
        public readonly string ShellName;
        public readonly string ShellVersion;
        public readonly string Slot;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;
        /// <summary>
        /// The status of this resource. Valid statuses are `CalamariNeedsUpgrade`, `Disabled`, `NeedsUpgrade`, `Offline`, `Online`, or `Unknown`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A summary elaborating on the status of this resource.
        /// </summary>
        public readonly string StatusSummary;
        public readonly string StorageAccountName;
        public readonly bool SwapIfPossible;
        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> TenantTags;
        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        public readonly string TenantedDeploymentParticipation;
        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public readonly ImmutableArray<string> Tenants;
        public readonly string Thumbprint;
        public readonly string Uri;
        public readonly bool UseCurrentInstanceCount;

        [OutputConstructor]
        private GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetResult(
            string accountId,

            string cloudServiceName,

            string defaultWorkerPoolId,

            ImmutableArray<Outputs.GetAzureCloudServiceDeploymentTargetsAzureCloudServiceDeploymentTargetEndpointResult> endpoints,

            ImmutableArray<string> environments,

            bool hasLatestCalamari,

            string healthStatus,

            string id,

            bool isDisabled,

            bool isInProcess,

            string machinePolicyId,

            string name,

            string operatingSystem,

            ImmutableArray<string> roles,

            string shellName,

            string shellVersion,

            string slot,

            string spaceId,

            string status,

            string statusSummary,

            string storageAccountName,

            bool swapIfPossible,

            ImmutableArray<string> tenantTags,

            string tenantedDeploymentParticipation,

            ImmutableArray<string> tenants,

            string thumbprint,

            string uri,

            bool useCurrentInstanceCount)
        {
            AccountId = accountId;
            CloudServiceName = cloudServiceName;
            DefaultWorkerPoolId = defaultWorkerPoolId;
            Endpoints = endpoints;
            Environments = environments;
            HasLatestCalamari = hasLatestCalamari;
            HealthStatus = healthStatus;
            Id = id;
            IsDisabled = isDisabled;
            IsInProcess = isInProcess;
            MachinePolicyId = machinePolicyId;
            Name = name;
            OperatingSystem = operatingSystem;
            Roles = roles;
            ShellName = shellName;
            ShellVersion = shellVersion;
            Slot = slot;
            SpaceId = spaceId;
            Status = status;
            StatusSummary = statusSummary;
            StorageAccountName = storageAccountName;
            SwapIfPossible = swapIfPossible;
            TenantTags = tenantTags;
            TenantedDeploymentParticipation = tenantedDeploymentParticipation;
            Tenants = tenants;
            Thumbprint = thumbprint;
            Uri = uri;
            UseCurrentInstanceCount = useCurrentInstanceCount;
        }
    }
}
