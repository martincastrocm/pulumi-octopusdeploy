// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-octopus/sdk/go/octopusdeploy/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The endpoint of the Octopus REST API
func GetAddress(ctx *pulumi.Context) string {
	return config.Get(ctx, "octopusdeploy:address")
}

// The API key to use with the Octopus REST API
func GetApiKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "octopusdeploy:apiKey")
}

// The space ID to target
func GetSpaceId(ctx *pulumi.Context) string {
	return config.Get(ctx, "octopusdeploy:spaceId")
}
