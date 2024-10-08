# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AwsAccountArgs', 'AwsAccount']

@pulumi.input_type
class AwsAccountArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AwsAccount resource.
        :param pulumi.Input[str] access_key: The access key associated with this AWS account.
        :param pulumi.Input[str] secret_key: The secret key associated with this resource.
        :param pulumi.Input[str] description: A user-friendly description of this AWS account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this AWS account.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "secret_key", secret_key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environments is not None:
            pulumi.set(__self__, "environments", environments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if tenant_tags is not None:
            pulumi.set(__self__, "tenant_tags", tenant_tags)
        if tenanted_deployment_participation is not None:
            pulumi.set(__self__, "tenanted_deployment_participation", tenanted_deployment_participation)
        if tenants is not None:
            pulumi.set(__self__, "tenants", tenants)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        The access key associated with this AWS account.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        The secret key associated with this resource.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A user-friendly description of this AWS account.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of environment IDs associated with this resource.
        """
        return pulumi.get(self, "environments")

    @environments.setter
    def environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "environments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this AWS account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter(name="tenantTags")
    def tenant_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tenant tags associated with this resource.
        """
        return pulumi.get(self, "tenant_tags")

    @tenant_tags.setter
    def tenant_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tenant_tags", value)

    @property
    @pulumi.getter(name="tenantedDeploymentParticipation")
    def tenanted_deployment_participation(self) -> Optional[pulumi.Input[str]]:
        """
        The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        """
        return pulumi.get(self, "tenanted_deployment_participation")

    @tenanted_deployment_participation.setter
    def tenanted_deployment_participation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenanted_deployment_participation", value)

    @property
    @pulumi.getter
    def tenants(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tenant IDs associated with this resource.
        """
        return pulumi.get(self, "tenants")

    @tenants.setter
    def tenants(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tenants", value)


@pulumi.input_type
class _AwsAccountState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AwsAccount resources.
        :param pulumi.Input[str] access_key: The access key associated with this AWS account.
        :param pulumi.Input[str] description: A user-friendly description of this AWS account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this AWS account.
        :param pulumi.Input[str] secret_key: The secret key associated with this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environments is not None:
            pulumi.set(__self__, "environments", environments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if tenant_tags is not None:
            pulumi.set(__self__, "tenant_tags", tenant_tags)
        if tenanted_deployment_participation is not None:
            pulumi.set(__self__, "tenanted_deployment_participation", tenanted_deployment_participation)
        if tenants is not None:
            pulumi.set(__self__, "tenants", tenants)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The access key associated with this AWS account.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A user-friendly description of this AWS account.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of environment IDs associated with this resource.
        """
        return pulumi.get(self, "environments")

    @environments.setter
    def environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "environments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this AWS account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key associated with this resource.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter(name="tenantTags")
    def tenant_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tenant tags associated with this resource.
        """
        return pulumi.get(self, "tenant_tags")

    @tenant_tags.setter
    def tenant_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tenant_tags", value)

    @property
    @pulumi.getter(name="tenantedDeploymentParticipation")
    def tenanted_deployment_participation(self) -> Optional[pulumi.Input[str]]:
        """
        The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        """
        return pulumi.get(self, "tenanted_deployment_participation")

    @tenanted_deployment_participation.setter
    def tenanted_deployment_participation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenanted_deployment_participation", value)

    @property
    @pulumi.getter
    def tenants(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tenant IDs associated with this resource.
        """
        return pulumi.get(self, "tenants")

    @tenants.setter
    def tenants(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tenants", value)


class AwsAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resource manages AWS accounts in Octopus Deploy.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.AwsAccount("example",
            access_key="access-key",
            secret_key="###########")
        # required; get from secure environment/store
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/awsAccount:AwsAccount [options] octopusdeploy_aws_account.<name> <account-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key associated with this AWS account.
        :param pulumi.Input[str] description: A user-friendly description of this AWS account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this AWS account.
        :param pulumi.Input[str] secret_key: The secret key associated with this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages AWS accounts in Octopus Deploy.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.AwsAccount("example",
            access_key="access-key",
            secret_key="###########")
        # required; get from secure environment/store
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/awsAccount:AwsAccount [options] octopusdeploy_aws_account.<name> <account-id>
        ```

        :param str resource_name: The name of the resource.
        :param AwsAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsAccountArgs.__new__(AwsAccountArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["description"] = description
            __props__.__dict__["environments"] = environments
            __props__.__dict__["name"] = name
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["space_id"] = space_id
            __props__.__dict__["tenant_tags"] = tenant_tags
            __props__.__dict__["tenanted_deployment_participation"] = tenanted_deployment_participation
            __props__.__dict__["tenants"] = tenants
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AwsAccount, __self__).__init__(
            'octopusdeploy:index/awsAccount:AwsAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None,
            tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
            tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AwsAccount':
        """
        Get an existing AwsAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key associated with this AWS account.
        :param pulumi.Input[str] description: A user-friendly description of this AWS account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this AWS account.
        :param pulumi.Input[str] secret_key: The secret key associated with this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsAccountState.__new__(_AwsAccountState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["description"] = description
        __props__.__dict__["environments"] = environments
        __props__.__dict__["name"] = name
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["space_id"] = space_id
        __props__.__dict__["tenant_tags"] = tenant_tags
        __props__.__dict__["tenanted_deployment_participation"] = tenanted_deployment_participation
        __props__.__dict__["tenants"] = tenants
        return AwsAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        The access key associated with this AWS account.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A user-friendly description of this AWS account.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environments(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of environment IDs associated with this resource.
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this AWS account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The secret key associated with this resource.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter(name="tenantTags")
    def tenant_tags(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of tenant tags associated with this resource.
        """
        return pulumi.get(self, "tenant_tags")

    @property
    @pulumi.getter(name="tenantedDeploymentParticipation")
    def tenanted_deployment_participation(self) -> pulumi.Output[str]:
        """
        The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        """
        return pulumi.get(self, "tenanted_deployment_participation")

    @property
    @pulumi.getter
    def tenants(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of tenant IDs associated with this resource.
        """
        return pulumi.get(self, "tenants")

