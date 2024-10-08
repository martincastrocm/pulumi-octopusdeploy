# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetCloudRegionDeploymentTargetsResult',
    'AwaitableGetCloudRegionDeploymentTargetsResult',
    'get_cloud_region_deployment_targets',
    'get_cloud_region_deployment_targets_output',
]

@pulumi.output_type
class GetCloudRegionDeploymentTargetsResult:
    """
    A collection of values returned by getCloudRegionDeploymentTargets.
    """
    def __init__(__self__, cloud_region_deployment_targets=None, deployment_id=None, environments=None, health_statuses=None, id=None, ids=None, is_disabled=None, name=None, partial_name=None, roles=None, shell_names=None, skip=None, space_id=None, take=None, tenant_tags=None, tenants=None, thumbprint=None):
        if cloud_region_deployment_targets and not isinstance(cloud_region_deployment_targets, list):
            raise TypeError("Expected argument 'cloud_region_deployment_targets' to be a list")
        pulumi.set(__self__, "cloud_region_deployment_targets", cloud_region_deployment_targets)
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if environments and not isinstance(environments, list):
            raise TypeError("Expected argument 'environments' to be a list")
        pulumi.set(__self__, "environments", environments)
        if health_statuses and not isinstance(health_statuses, list):
            raise TypeError("Expected argument 'health_statuses' to be a list")
        pulumi.set(__self__, "health_statuses", health_statuses)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if is_disabled and not isinstance(is_disabled, bool):
            raise TypeError("Expected argument 'is_disabled' to be a bool")
        pulumi.set(__self__, "is_disabled", is_disabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partial_name and not isinstance(partial_name, str):
            raise TypeError("Expected argument 'partial_name' to be a str")
        pulumi.set(__self__, "partial_name", partial_name)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if shell_names and not isinstance(shell_names, list):
            raise TypeError("Expected argument 'shell_names' to be a list")
        pulumi.set(__self__, "shell_names", shell_names)
        if skip and not isinstance(skip, int):
            raise TypeError("Expected argument 'skip' to be a int")
        pulumi.set(__self__, "skip", skip)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)
        if take and not isinstance(take, int):
            raise TypeError("Expected argument 'take' to be a int")
        pulumi.set(__self__, "take", take)
        if tenant_tags and not isinstance(tenant_tags, list):
            raise TypeError("Expected argument 'tenant_tags' to be a list")
        pulumi.set(__self__, "tenant_tags", tenant_tags)
        if tenants and not isinstance(tenants, list):
            raise TypeError("Expected argument 'tenants' to be a list")
        pulumi.set(__self__, "tenants", tenants)
        if thumbprint and not isinstance(thumbprint, str):
            raise TypeError("Expected argument 'thumbprint' to be a str")
        pulumi.set(__self__, "thumbprint", thumbprint)

    @property
    @pulumi.getter(name="cloudRegionDeploymentTargets")
    def cloud_region_deployment_targets(self) -> Sequence['outputs.GetCloudRegionDeploymentTargetsCloudRegionDeploymentTargetResult']:
        """
        A list of cloud region deployment targets that match the filter(s).
        """
        return pulumi.get(self, "cloud_region_deployment_targets")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[str]:
        """
        A filter to search by deployment ID.
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def environments(self) -> Optional[Sequence[str]]:
        """
        A filter to search by a list of environment IDs.
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter(name="healthStatuses")
    def health_statuses(self) -> Optional[Sequence[str]]:
        """
        A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
        """
        return pulumi.get(self, "health_statuses")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        An auto-generated identifier that includes the timestamp when this data source was last modified.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        """
        A filter to search by a list of IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> Optional[bool]:
        """
        A filter to search by the disabled status of a resource.
        """
        return pulumi.get(self, "is_disabled")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        A filter to search by name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partialName")
    def partial_name(self) -> Optional[str]:
        """
        A filter to search by the partial match of a name.
        """
        return pulumi.get(self, "partial_name")

    @property
    @pulumi.getter
    def roles(self) -> Optional[Sequence[str]]:
        """
        A filter to search by a list of role IDs.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="shellNames")
    def shell_names(self) -> Optional[Sequence[str]]:
        """
        A list of shell names to match in the query and/or search
        """
        return pulumi.get(self, "shell_names")

    @property
    @pulumi.getter
    def skip(self) -> Optional[int]:
        """
        A filter to specify the number of items to skip in the response.
        """
        return pulumi.get(self, "skip")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> str:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter
    def take(self) -> Optional[int]:
        """
        A filter to specify the number of items to take (or return) in the response.
        """
        return pulumi.get(self, "take")

    @property
    @pulumi.getter(name="tenantTags")
    def tenant_tags(self) -> Optional[Sequence[str]]:
        """
        A filter to search by a list of tenant tags.
        """
        return pulumi.get(self, "tenant_tags")

    @property
    @pulumi.getter
    def tenants(self) -> Optional[Sequence[str]]:
        """
        A filter to search by a list of tenant IDs.
        """
        return pulumi.get(self, "tenants")

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[str]:
        """
        The thumbprint of the deployment target to match in the query and/or search
        """
        return pulumi.get(self, "thumbprint")


class AwaitableGetCloudRegionDeploymentTargetsResult(GetCloudRegionDeploymentTargetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudRegionDeploymentTargetsResult(
            cloud_region_deployment_targets=self.cloud_region_deployment_targets,
            deployment_id=self.deployment_id,
            environments=self.environments,
            health_statuses=self.health_statuses,
            id=self.id,
            ids=self.ids,
            is_disabled=self.is_disabled,
            name=self.name,
            partial_name=self.partial_name,
            roles=self.roles,
            shell_names=self.shell_names,
            skip=self.skip,
            space_id=self.space_id,
            take=self.take,
            tenant_tags=self.tenant_tags,
            tenants=self.tenants,
            thumbprint=self.thumbprint)


def get_cloud_region_deployment_targets(deployment_id: Optional[str] = None,
                                        environments: Optional[Sequence[str]] = None,
                                        health_statuses: Optional[Sequence[str]] = None,
                                        ids: Optional[Sequence[str]] = None,
                                        is_disabled: Optional[bool] = None,
                                        name: Optional[str] = None,
                                        partial_name: Optional[str] = None,
                                        roles: Optional[Sequence[str]] = None,
                                        shell_names: Optional[Sequence[str]] = None,
                                        skip: Optional[int] = None,
                                        space_id: Optional[str] = None,
                                        take: Optional[int] = None,
                                        tenant_tags: Optional[Sequence[str]] = None,
                                        tenants: Optional[Sequence[str]] = None,
                                        thumbprint: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudRegionDeploymentTargetsResult:
    """
    Provides information about existing cloud region deployment targets.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_octopusdeploy as octopusdeploy

    example = octopusdeploy.get_cloud_region_deployment_targets(environments=[
            "Environments-123",
            "Environments-321",
        ],
        ids=["Machines-123"],
        name="Azure North America",
        partial_name="Azure Nor",
        skip=5,
        take=100)
    ```
    <!--End PulumiCodeChooser -->


    :param str deployment_id: A filter to search by deployment ID.
    :param Sequence[str] health_statuses: A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param Sequence[str] shell_names: A list of shell names to match in the query and/or search
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    __args__ = dict()
    __args__['deploymentId'] = deployment_id
    __args__['environments'] = environments
    __args__['healthStatuses'] = health_statuses
    __args__['ids'] = ids
    __args__['isDisabled'] = is_disabled
    __args__['name'] = name
    __args__['partialName'] = partial_name
    __args__['roles'] = roles
    __args__['shellNames'] = shell_names
    __args__['skip'] = skip
    __args__['spaceId'] = space_id
    __args__['take'] = take
    __args__['tenantTags'] = tenant_tags
    __args__['tenants'] = tenants
    __args__['thumbprint'] = thumbprint
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('octopusdeploy:index/getCloudRegionDeploymentTargets:getCloudRegionDeploymentTargets', __args__, opts=opts, typ=GetCloudRegionDeploymentTargetsResult).value

    return AwaitableGetCloudRegionDeploymentTargetsResult(
        cloud_region_deployment_targets=pulumi.get(__ret__, 'cloud_region_deployment_targets'),
        deployment_id=pulumi.get(__ret__, 'deployment_id'),
        environments=pulumi.get(__ret__, 'environments'),
        health_statuses=pulumi.get(__ret__, 'health_statuses'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        is_disabled=pulumi.get(__ret__, 'is_disabled'),
        name=pulumi.get(__ret__, 'name'),
        partial_name=pulumi.get(__ret__, 'partial_name'),
        roles=pulumi.get(__ret__, 'roles'),
        shell_names=pulumi.get(__ret__, 'shell_names'),
        skip=pulumi.get(__ret__, 'skip'),
        space_id=pulumi.get(__ret__, 'space_id'),
        take=pulumi.get(__ret__, 'take'),
        tenant_tags=pulumi.get(__ret__, 'tenant_tags'),
        tenants=pulumi.get(__ret__, 'tenants'),
        thumbprint=pulumi.get(__ret__, 'thumbprint'))


@_utilities.lift_output_func(get_cloud_region_deployment_targets)
def get_cloud_region_deployment_targets_output(deployment_id: Optional[pulumi.Input[Optional[str]]] = None,
                                               environments: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                               health_statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                               ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                               is_disabled: Optional[pulumi.Input[Optional[bool]]] = None,
                                               name: Optional[pulumi.Input[Optional[str]]] = None,
                                               partial_name: Optional[pulumi.Input[Optional[str]]] = None,
                                               roles: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                               shell_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                               skip: Optional[pulumi.Input[Optional[int]]] = None,
                                               space_id: Optional[pulumi.Input[Optional[str]]] = None,
                                               take: Optional[pulumi.Input[Optional[int]]] = None,
                                               tenant_tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                               tenants: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                               thumbprint: Optional[pulumi.Input[Optional[str]]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCloudRegionDeploymentTargetsResult]:
    """
    Provides information about existing cloud region deployment targets.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_octopusdeploy as octopusdeploy

    example = octopusdeploy.get_cloud_region_deployment_targets(environments=[
            "Environments-123",
            "Environments-321",
        ],
        ids=["Machines-123"],
        name="Azure North America",
        partial_name="Azure Nor",
        skip=5,
        take=100)
    ```
    <!--End PulumiCodeChooser -->


    :param str deployment_id: A filter to search by deployment ID.
    :param Sequence[str] health_statuses: A filter to search by a list of health statuses of resources. Valid health statuses are `HasWarnings`, `Healthy`, `Unavailable`, `Unhealthy`, or `Unknown`.
    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param Sequence[str] shell_names: A list of shell names to match in the query and/or search
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    ...
