# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SshKeyAccountArgs', 'SshKeyAccount']

@pulumi.input_type
class SshKeyAccountArgs:
    def __init__(__self__, *,
                 private_key_file: pulumi.Input[str],
                 username: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SshKeyAccount resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        :param pulumi.Input[str] description: The description of this SSH key account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        """
        pulumi.set(__self__, "private_key_file", private_key_file)
        pulumi.set(__self__, "username", username)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environments is not None:
            pulumi.set(__self__, "environments", environments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key_passphrase is not None:
            pulumi.set(__self__, "private_key_passphrase", private_key_passphrase)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if tenant_tags is not None:
            pulumi.set(__self__, "tenant_tags", tenant_tags)
        if tenanted_deployment_participation is not None:
            pulumi.set(__self__, "tenanted_deployment_participation", tenanted_deployment_participation)
        if tenants is not None:
            pulumi.set(__self__, "tenants", tenants)

    @property
    @pulumi.getter(name="privateKeyFile")
    def private_key_file(self) -> pulumi.Input[str]:
        return pulumi.get(self, "private_key_file")

    @private_key_file.setter
    def private_key_file(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key_file", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The username associated with this resource.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this SSH key account.
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
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_key_passphrase")

    @private_key_passphrase.setter
    def private_key_passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_passphrase", value)

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
class _SshKeyAccountState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key_file: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SshKeyAccount resources.
        :param pulumi.Input[str] description: The description of this SSH key account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environments is not None:
            pulumi.set(__self__, "environments", environments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key_file is not None:
            pulumi.set(__self__, "private_key_file", private_key_file)
        if private_key_passphrase is not None:
            pulumi.set(__self__, "private_key_passphrase", private_key_passphrase)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if tenant_tags is not None:
            pulumi.set(__self__, "tenant_tags", tenant_tags)
        if tenanted_deployment_participation is not None:
            pulumi.set(__self__, "tenanted_deployment_participation", tenanted_deployment_participation)
        if tenants is not None:
            pulumi.set(__self__, "tenants", tenants)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this SSH key account.
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
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKeyFile")
    def private_key_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_key_file")

    @private_key_file.setter
    def private_key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_file", value)

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_key_passphrase")

    @private_key_passphrase.setter
    def private_key_passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_passphrase", value)

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

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username associated with this resource.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class SshKeyAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key_file: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource manages SSH key accounts in Octopus Deploy.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.SshKeyAccount("example",
            private_key_file="[private_key_file]",
            username="[username]")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/sshKeyAccount:SshKeyAccount [options] octopusdeploy_ssh_key_account.<name> <account-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this SSH key account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SshKeyAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages SSH key accounts in Octopus Deploy.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.SshKeyAccount("example",
            private_key_file="[private_key_file]",
            username="[username]")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/sshKeyAccount:SshKeyAccount [options] octopusdeploy_ssh_key_account.<name> <account-id>
        ```

        :param str resource_name: The name of the resource.
        :param SshKeyAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SshKeyAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key_file: Optional[pulumi.Input[str]] = None,
                 private_key_passphrase: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
                 tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SshKeyAccountArgs.__new__(SshKeyAccountArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["environments"] = environments
            __props__.__dict__["name"] = name
            if private_key_file is None and not opts.urn:
                raise TypeError("Missing required property 'private_key_file'")
            __props__.__dict__["private_key_file"] = None if private_key_file is None else pulumi.Output.secret(private_key_file)
            __props__.__dict__["private_key_passphrase"] = None if private_key_passphrase is None else pulumi.Output.secret(private_key_passphrase)
            __props__.__dict__["space_id"] = space_id
            __props__.__dict__["tenant_tags"] = tenant_tags
            __props__.__dict__["tenanted_deployment_participation"] = tenanted_deployment_participation
            __props__.__dict__["tenants"] = tenants
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = None if username is None else pulumi.Output.secret(username)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKeyFile", "privateKeyPassphrase", "username"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SshKeyAccount, __self__).__init__(
            'octopusdeploy:index/sshKeyAccount:SshKeyAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_key_file: Optional[pulumi.Input[str]] = None,
            private_key_passphrase: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None,
            tenant_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tenanted_deployment_participation: Optional[pulumi.Input[str]] = None,
            tenants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'SshKeyAccount':
        """
        Get an existing SshKeyAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this SSH key account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A list of environment IDs associated with this resource.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenant_tags: A list of tenant tags associated with this resource.
        :param pulumi.Input[str] tenanted_deployment_participation: The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tenants: A list of tenant IDs associated with this resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SshKeyAccountState.__new__(_SshKeyAccountState)

        __props__.__dict__["description"] = description
        __props__.__dict__["environments"] = environments
        __props__.__dict__["name"] = name
        __props__.__dict__["private_key_file"] = private_key_file
        __props__.__dict__["private_key_passphrase"] = private_key_passphrase
        __props__.__dict__["space_id"] = space_id
        __props__.__dict__["tenant_tags"] = tenant_tags
        __props__.__dict__["tenanted_deployment_participation"] = tenanted_deployment_participation
        __props__.__dict__["tenants"] = tenants
        __props__.__dict__["username"] = username
        return SshKeyAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this SSH key account.
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
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKeyFile")
    def private_key_file(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_key_file")

    @property
    @pulumi.getter(name="privateKeyPassphrase")
    def private_key_passphrase(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "private_key_passphrase")

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

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The username associated with this resource.
        """
        return pulumi.get(self, "username")

