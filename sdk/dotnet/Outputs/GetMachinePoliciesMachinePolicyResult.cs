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
    public sealed class GetMachinePoliciesMachinePolicyResult
    {
        /// <summary>
        /// In nanoseconds. Minimum value: 10000000000 (10 seconds).
        /// </summary>
        public readonly int ConnectionConnectTimeout;
        public readonly int ConnectionRetryCountLimit;
        /// <summary>
        /// In nanoseconds.
        /// </summary>
        public readonly int ConnectionRetrySleepInterval;
        /// <summary>
        /// In nanoseconds.
        /// </summary>
        public readonly int ConnectionRetryTimeLimit;
        /// <summary>
        /// The description of this machine policy.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsDefault;
        public readonly ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineCleanupPolicyResult> MachineCleanupPolicies;
        public readonly ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineConnectivityPolicyResult> MachineConnectivityPolicies;
        public readonly ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyResult> MachineHealthCheckPolicies;
        public readonly ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineUpdatePolicyResult> MachineUpdatePolicies;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// In nanoseconds.
        /// </summary>
        public readonly int PollingRequestQueueTimeout;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;

        [OutputConstructor]
        private GetMachinePoliciesMachinePolicyResult(
            int connectionConnectTimeout,

            int connectionRetryCountLimit,

            int connectionRetrySleepInterval,

            int connectionRetryTimeLimit,

            string description,

            string id,

            bool isDefault,

            ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineCleanupPolicyResult> machineCleanupPolicies,

            ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineConnectivityPolicyResult> machineConnectivityPolicies,

            ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineHealthCheckPolicyResult> machineHealthCheckPolicies,

            ImmutableArray<Outputs.GetMachinePoliciesMachinePolicyMachineUpdatePolicyResult> machineUpdatePolicies,

            string name,

            int pollingRequestQueueTimeout,

            string spaceId)
        {
            ConnectionConnectTimeout = connectionConnectTimeout;
            ConnectionRetryCountLimit = connectionRetryCountLimit;
            ConnectionRetrySleepInterval = connectionRetrySleepInterval;
            ConnectionRetryTimeLimit = connectionRetryTimeLimit;
            Description = description;
            Id = id;
            IsDefault = isDefault;
            MachineCleanupPolicies = machineCleanupPolicies;
            MachineConnectivityPolicies = machineConnectivityPolicies;
            MachineHealthCheckPolicies = machineHealthCheckPolicies;
            MachineUpdatePolicies = machineUpdatePolicies;
            Name = name;
            PollingRequestQueueTimeout = pollingRequestQueueTimeout;
            SpaceId = spaceId;
        }
    }
}
