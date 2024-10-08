// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class UserIdentityGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("claims")]
        private InputList<Inputs.UserIdentityClaimGetArgs>? _claims;
        public InputList<Inputs.UserIdentityClaimGetArgs> Claims
        {
            get => _claims ?? (_claims = new InputList<Inputs.UserIdentityClaimGetArgs>());
            set => _claims = value;
        }

        [Input("provider")]
        public Input<string>? Provider { get; set; }

        public UserIdentityGetArgs()
        {
        }
        public static new UserIdentityGetArgs Empty => new UserIdentityGetArgs();
    }
}
