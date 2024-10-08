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
    public sealed class GetUsersUserResult
    {
        public readonly bool CanPasswordBeEdited;
        /// <summary>
        /// The display name of this resource.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The email address of this resource.
        /// </summary>
        public readonly string EmailAddress;
        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetUsersUserIdentityResult> Identities;
        public readonly bool IsActive;
        public readonly bool IsRequestor;
        public readonly bool IsService;
        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The username associated with this resource.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetUsersUserResult(
            bool canPasswordBeEdited,

            string displayName,

            string emailAddress,

            string id,

            ImmutableArray<Outputs.GetUsersUserIdentityResult> identities,

            bool isActive,

            bool isRequestor,

            bool isService,

            string password,

            string username)
        {
            CanPasswordBeEdited = canPasswordBeEdited;
            DisplayName = displayName;
            EmailAddress = emailAddress;
            Id = id;
            Identities = identities;
            IsActive = isActive;
            IsRequestor = isRequestor;
            IsService = isService;
            Password = password;
            Username = username;
        }
    }
}
