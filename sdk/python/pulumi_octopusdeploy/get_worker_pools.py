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
    'GetWorkerPoolsResult',
    'AwaitableGetWorkerPoolsResult',
    'get_worker_pools',
    'get_worker_pools_output',
]

@pulumi.output_type
class GetWorkerPoolsResult:
    """
    A collection of values returned by getWorkerPools.
    """
    def __init__(__self__, id=None, ids=None, name=None, partial_name=None, skip=None, space_id=None, take=None, worker_pools=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partial_name and not isinstance(partial_name, str):
            raise TypeError("Expected argument 'partial_name' to be a str")
        pulumi.set(__self__, "partial_name", partial_name)
        if skip and not isinstance(skip, int):
            raise TypeError("Expected argument 'skip' to be a int")
        pulumi.set(__self__, "skip", skip)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)
        if take and not isinstance(take, int):
            raise TypeError("Expected argument 'take' to be a int")
        pulumi.set(__self__, "take", take)
        if worker_pools and not isinstance(worker_pools, list):
            raise TypeError("Expected argument 'worker_pools' to be a list")
        pulumi.set(__self__, "worker_pools", worker_pools)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
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
    @pulumi.getter(name="workerPools")
    def worker_pools(self) -> Sequence['outputs.GetWorkerPoolsWorkerPoolResult']:
        """
        A list of worker pools that match the filter(s).
        """
        return pulumi.get(self, "worker_pools")


class AwaitableGetWorkerPoolsResult(GetWorkerPoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkerPoolsResult(
            id=self.id,
            ids=self.ids,
            name=self.name,
            partial_name=self.partial_name,
            skip=self.skip,
            space_id=self.space_id,
            take=self.take,
            worker_pools=self.worker_pools)


def get_worker_pools(ids: Optional[Sequence[str]] = None,
                     name: Optional[str] = None,
                     partial_name: Optional[str] = None,
                     skip: Optional[int] = None,
                     space_id: Optional[str] = None,
                     take: Optional[int] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkerPoolsResult:
    """
    Provides information about existing worker pools.


    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['name'] = name
    __args__['partialName'] = partial_name
    __args__['skip'] = skip
    __args__['spaceId'] = space_id
    __args__['take'] = take
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('octopusdeploy:index/getWorkerPools:getWorkerPools', __args__, opts=opts, typ=GetWorkerPoolsResult).value

    return AwaitableGetWorkerPoolsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name=pulumi.get(__ret__, 'name'),
        partial_name=pulumi.get(__ret__, 'partial_name'),
        skip=pulumi.get(__ret__, 'skip'),
        space_id=pulumi.get(__ret__, 'space_id'),
        take=pulumi.get(__ret__, 'take'),
        worker_pools=pulumi.get(__ret__, 'worker_pools'))


@_utilities.lift_output_func(get_worker_pools)
def get_worker_pools_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            name: Optional[pulumi.Input[Optional[str]]] = None,
                            partial_name: Optional[pulumi.Input[Optional[str]]] = None,
                            skip: Optional[pulumi.Input[Optional[int]]] = None,
                            space_id: Optional[pulumi.Input[Optional[str]]] = None,
                            take: Optional[pulumi.Input[Optional[int]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkerPoolsResult]:
    """
    Provides information about existing worker pools.


    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    ...
