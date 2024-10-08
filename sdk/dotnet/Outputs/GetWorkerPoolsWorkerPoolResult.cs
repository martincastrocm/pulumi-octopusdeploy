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
    public sealed class GetWorkerPoolsWorkerPoolResult
    {
        public readonly bool CanAddWorkers;
        /// <summary>
        /// The description of this worker pool.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsDefault;
        /// <summary>
        /// The name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The order number to sort a dynamic worker pool.
        /// </summary>
        public readonly int SortOrder;
        /// <summary>
        /// The space ID associated with this resource.
        /// </summary>
        public readonly string SpaceId;
        public readonly string WorkerPoolType;
        public readonly string WorkerType;

        [OutputConstructor]
        private GetWorkerPoolsWorkerPoolResult(
            bool canAddWorkers,

            string description,

            string id,

            bool isDefault,

            string name,

            int sortOrder,

            string spaceId,

            string workerPoolType,

            string workerType)
        {
            CanAddWorkers = canAddWorkers;
            Description = description;
            Id = id;
            IsDefault = isDefault;
            Name = name;
            SortOrder = sortOrder;
            SpaceId = spaceId;
            WorkerPoolType = workerPoolType;
            WorkerType = workerType;
        }
    }
}
