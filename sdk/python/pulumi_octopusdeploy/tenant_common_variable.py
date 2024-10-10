# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TenantCommonVariableArgs', 'TenantCommonVariable']

@pulumi.input_type
class TenantCommonVariableArgs:
    def __init__(__self__, *,
                 library_variable_set_id: pulumi.Input[str],
                 template_id: pulumi.Input[str],
                 tenant_id: pulumi.Input[str],
                 space_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TenantCommonVariable resource.
        :param pulumi.Input[str] library_variable_set_id: The ID of the library variable set.
        :param pulumi.Input[str] template_id: The ID of the variable template.
        :param pulumi.Input[str] tenant_id: The ID of the tenant.
        :param pulumi.Input[str] space_id: The space ID associated with this Tenant Common Variable.
        :param pulumi.Input[str] value: The value of the variable.
        """
        pulumi.set(__self__, "library_variable_set_id", library_variable_set_id)
        pulumi.set(__self__, "template_id", template_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="libraryVariableSetId")
    def library_variable_set_id(self) -> pulumi.Input[str]:
        """
        The ID of the library variable set.
        """
        return pulumi.get(self, "library_variable_set_id")

    @library_variable_set_id.setter
    def library_variable_set_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "library_variable_set_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Input[str]:
        """
        The ID of the variable template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The ID of the tenant.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this Tenant Common Variable.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _TenantCommonVariableState:
    def __init__(__self__, *,
                 library_variable_set_id: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TenantCommonVariable resources.
        :param pulumi.Input[str] library_variable_set_id: The ID of the library variable set.
        :param pulumi.Input[str] space_id: The space ID associated with this Tenant Common Variable.
        :param pulumi.Input[str] template_id: The ID of the variable template.
        :param pulumi.Input[str] tenant_id: The ID of the tenant.
        :param pulumi.Input[str] value: The value of the variable.
        """
        if library_variable_set_id is not None:
            pulumi.set(__self__, "library_variable_set_id", library_variable_set_id)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="libraryVariableSetId")
    def library_variable_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the library variable set.
        """
        return pulumi.get(self, "library_variable_set_id")

    @library_variable_set_id.setter
    def library_variable_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "library_variable_set_id", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this Tenant Common Variable.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the variable template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the tenant.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the variable.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class TenantCommonVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 library_variable_set_id: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a tenant common variable in Octopus Deploy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] library_variable_set_id: The ID of the library variable set.
        :param pulumi.Input[str] space_id: The space ID associated with this Tenant Common Variable.
        :param pulumi.Input[str] template_id: The ID of the variable template.
        :param pulumi.Input[str] tenant_id: The ID of the tenant.
        :param pulumi.Input[str] value: The value of the variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TenantCommonVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a tenant common variable in Octopus Deploy.

        :param str resource_name: The name of the resource.
        :param TenantCommonVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TenantCommonVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 library_variable_set_id: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TenantCommonVariableArgs.__new__(TenantCommonVariableArgs)

            if library_variable_set_id is None and not opts.urn:
                raise TypeError("Missing required property 'library_variable_set_id'")
            __props__.__dict__["library_variable_set_id"] = library_variable_set_id
            __props__.__dict__["space_id"] = space_id
            if template_id is None and not opts.urn:
                raise TypeError("Missing required property 'template_id'")
            __props__.__dict__["template_id"] = template_id
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["value"] = None if value is None else pulumi.Output.secret(value)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["value"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(TenantCommonVariable, __self__).__init__(
            'octopusdeploy:index/tenantCommonVariable:TenantCommonVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            library_variable_set_id: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None,
            template_id: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'TenantCommonVariable':
        """
        Get an existing TenantCommonVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] library_variable_set_id: The ID of the library variable set.
        :param pulumi.Input[str] space_id: The space ID associated with this Tenant Common Variable.
        :param pulumi.Input[str] template_id: The ID of the variable template.
        :param pulumi.Input[str] tenant_id: The ID of the tenant.
        :param pulumi.Input[str] value: The value of the variable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TenantCommonVariableState.__new__(_TenantCommonVariableState)

        __props__.__dict__["library_variable_set_id"] = library_variable_set_id
        __props__.__dict__["space_id"] = space_id
        __props__.__dict__["template_id"] = template_id
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["value"] = value
        return TenantCommonVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="libraryVariableSetId")
    def library_variable_set_id(self) -> pulumi.Output[str]:
        """
        The ID of the library variable set.
        """
        return pulumi.get(self, "library_variable_set_id")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        The space ID associated with this Tenant Common Variable.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        """
        The ID of the variable template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The ID of the tenant.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[str]]:
        """
        The value of the variable.
        """
        return pulumi.get(self, "value")

