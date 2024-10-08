# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserRoleArgs', 'UserRole']

@pulumi.input_type
class UserRoleArgs:
    def __init__(__self__, *,
                 can_be_deleted: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 granted_space_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 granted_system_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 supported_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 system_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a UserRole resource.
        :param pulumi.Input[str] description: The description of this user role.
        :param pulumi.Input[str] name: The name of this resource.
        """
        if can_be_deleted is not None:
            pulumi.set(__self__, "can_be_deleted", can_be_deleted)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if granted_space_permissions is not None:
            pulumi.set(__self__, "granted_space_permissions", granted_space_permissions)
        if granted_system_permissions is not None:
            pulumi.set(__self__, "granted_system_permissions", granted_system_permissions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_permission_descriptions is not None:
            pulumi.set(__self__, "space_permission_descriptions", space_permission_descriptions)
        if supported_restrictions is not None:
            pulumi.set(__self__, "supported_restrictions", supported_restrictions)
        if system_permission_descriptions is not None:
            pulumi.set(__self__, "system_permission_descriptions", system_permission_descriptions)

    @property
    @pulumi.getter(name="canBeDeleted")
    def can_be_deleted(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "can_be_deleted")

    @can_be_deleted.setter
    def can_be_deleted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_be_deleted", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this user role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="grantedSpacePermissions")
    def granted_space_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "granted_space_permissions")

    @granted_space_permissions.setter
    def granted_space_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "granted_space_permissions", value)

    @property
    @pulumi.getter(name="grantedSystemPermissions")
    def granted_system_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "granted_system_permissions")

    @granted_system_permissions.setter
    def granted_system_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "granted_system_permissions", value)

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
    @pulumi.getter(name="spacePermissionDescriptions")
    def space_permission_descriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "space_permission_descriptions")

    @space_permission_descriptions.setter
    def space_permission_descriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "space_permission_descriptions", value)

    @property
    @pulumi.getter(name="supportedRestrictions")
    def supported_restrictions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "supported_restrictions")

    @supported_restrictions.setter
    def supported_restrictions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "supported_restrictions", value)

    @property
    @pulumi.getter(name="systemPermissionDescriptions")
    def system_permission_descriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "system_permission_descriptions")

    @system_permission_descriptions.setter
    def system_permission_descriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "system_permission_descriptions", value)


@pulumi.input_type
class _UserRoleState:
    def __init__(__self__, *,
                 can_be_deleted: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 granted_space_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 granted_system_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 supported_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 system_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering UserRole resources.
        :param pulumi.Input[str] description: The description of this user role.
        :param pulumi.Input[str] name: The name of this resource.
        """
        if can_be_deleted is not None:
            pulumi.set(__self__, "can_be_deleted", can_be_deleted)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if granted_space_permissions is not None:
            pulumi.set(__self__, "granted_space_permissions", granted_space_permissions)
        if granted_system_permissions is not None:
            pulumi.set(__self__, "granted_system_permissions", granted_system_permissions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_permission_descriptions is not None:
            pulumi.set(__self__, "space_permission_descriptions", space_permission_descriptions)
        if supported_restrictions is not None:
            pulumi.set(__self__, "supported_restrictions", supported_restrictions)
        if system_permission_descriptions is not None:
            pulumi.set(__self__, "system_permission_descriptions", system_permission_descriptions)

    @property
    @pulumi.getter(name="canBeDeleted")
    def can_be_deleted(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "can_be_deleted")

    @can_be_deleted.setter
    def can_be_deleted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_be_deleted", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this user role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="grantedSpacePermissions")
    def granted_space_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "granted_space_permissions")

    @granted_space_permissions.setter
    def granted_space_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "granted_space_permissions", value)

    @property
    @pulumi.getter(name="grantedSystemPermissions")
    def granted_system_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "granted_system_permissions")

    @granted_system_permissions.setter
    def granted_system_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "granted_system_permissions", value)

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
    @pulumi.getter(name="spacePermissionDescriptions")
    def space_permission_descriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "space_permission_descriptions")

    @space_permission_descriptions.setter
    def space_permission_descriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "space_permission_descriptions", value)

    @property
    @pulumi.getter(name="supportedRestrictions")
    def supported_restrictions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "supported_restrictions")

    @supported_restrictions.setter
    def supported_restrictions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "supported_restrictions", value)

    @property
    @pulumi.getter(name="systemPermissionDescriptions")
    def system_permission_descriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "system_permission_descriptions")

    @system_permission_descriptions.setter
    def system_permission_descriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "system_permission_descriptions", value)


class UserRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_be_deleted: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 granted_space_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 granted_system_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 supported_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 system_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resource manages user roles in Octopus Deploy.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.UserRole("example",
            can_be_deleted=True,
            description="Responsible for all development-related operations.",
            granted_space_permissions=[
                "DeploymentCreate",
                "DeploymentDelete",
                "DeploymentView",
            ],
            granted_system_permissions=["SpaceCreate"],
            space_permission_descriptions=[
                "Delete deployments (restrictable to Environments, Projects, Tenants)",
                "Deploy releases to target environments (restrictable to Environments, Projects, Tenants)",
                "View deployments (restrictable to Environments, Projects, Tenants)",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/userRole:UserRole [options] octopusdeploy_user_role.<name> <user-role-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this user role.
        :param pulumi.Input[str] name: The name of this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserRoleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages user roles in Octopus Deploy.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.UserRole("example",
            can_be_deleted=True,
            description="Responsible for all development-related operations.",
            granted_space_permissions=[
                "DeploymentCreate",
                "DeploymentDelete",
                "DeploymentView",
            ],
            granted_system_permissions=["SpaceCreate"],
            space_permission_descriptions=[
                "Delete deployments (restrictable to Environments, Projects, Tenants)",
                "Deploy releases to target environments (restrictable to Environments, Projects, Tenants)",
                "View deployments (restrictable to Environments, Projects, Tenants)",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/userRole:UserRole [options] octopusdeploy_user_role.<name> <user-role-id>
        ```

        :param str resource_name: The name of the resource.
        :param UserRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_be_deleted: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 granted_space_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 granted_system_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 supported_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 system_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserRoleArgs.__new__(UserRoleArgs)

            __props__.__dict__["can_be_deleted"] = can_be_deleted
            __props__.__dict__["description"] = description
            __props__.__dict__["granted_space_permissions"] = granted_space_permissions
            __props__.__dict__["granted_system_permissions"] = granted_system_permissions
            __props__.__dict__["name"] = name
            __props__.__dict__["space_permission_descriptions"] = space_permission_descriptions
            __props__.__dict__["supported_restrictions"] = supported_restrictions
            __props__.__dict__["system_permission_descriptions"] = system_permission_descriptions
        super(UserRole, __self__).__init__(
            'octopusdeploy:index/userRole:UserRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            can_be_deleted: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            granted_space_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            granted_system_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            space_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            supported_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            system_permission_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'UserRole':
        """
        Get an existing UserRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this user role.
        :param pulumi.Input[str] name: The name of this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserRoleState.__new__(_UserRoleState)

        __props__.__dict__["can_be_deleted"] = can_be_deleted
        __props__.__dict__["description"] = description
        __props__.__dict__["granted_space_permissions"] = granted_space_permissions
        __props__.__dict__["granted_system_permissions"] = granted_system_permissions
        __props__.__dict__["name"] = name
        __props__.__dict__["space_permission_descriptions"] = space_permission_descriptions
        __props__.__dict__["supported_restrictions"] = supported_restrictions
        __props__.__dict__["system_permission_descriptions"] = system_permission_descriptions
        return UserRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canBeDeleted")
    def can_be_deleted(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "can_be_deleted")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this user role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="grantedSpacePermissions")
    def granted_space_permissions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "granted_space_permissions")

    @property
    @pulumi.getter(name="grantedSystemPermissions")
    def granted_system_permissions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "granted_system_permissions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="spacePermissionDescriptions")
    def space_permission_descriptions(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "space_permission_descriptions")

    @property
    @pulumi.getter(name="supportedRestrictions")
    def supported_restrictions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "supported_restrictions")

    @property
    @pulumi.getter(name="systemPermissionDescriptions")
    def system_permission_descriptions(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "system_permission_descriptions")

