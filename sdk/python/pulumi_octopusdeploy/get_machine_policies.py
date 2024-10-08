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
    'GetMachinePoliciesResult',
    'AwaitableGetMachinePoliciesResult',
    'get_machine_policies',
    'get_machine_policies_output',
]

@pulumi.output_type
class GetMachinePoliciesResult:
    """
    A collection of values returned by getMachinePolicies.
    """
    def __init__(__self__, id=None, ids=None, machine_policies=None, partial_name=None, skip=None, space_id=None, take=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if machine_policies and not isinstance(machine_policies, list):
            raise TypeError("Expected argument 'machine_policies' to be a list")
        pulumi.set(__self__, "machine_policies", machine_policies)
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
    @pulumi.getter(name="machinePolicies")
    def machine_policies(self) -> Sequence['outputs.GetMachinePoliciesMachinePolicyResult']:
        """
        A list of machine policies that match the filter(s).
        """
        return pulumi.get(self, "machine_policies")

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


class AwaitableGetMachinePoliciesResult(GetMachinePoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMachinePoliciesResult(
            id=self.id,
            ids=self.ids,
            machine_policies=self.machine_policies,
            partial_name=self.partial_name,
            skip=self.skip,
            space_id=self.space_id,
            take=self.take)


def get_machine_policies(ids: Optional[Sequence[str]] = None,
                         partial_name: Optional[str] = None,
                         skip: Optional[int] = None,
                         space_id: Optional[str] = None,
                         take: Optional[int] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMachinePoliciesResult:
    """
    Provides information about existing machine policies.


    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['partialName'] = partial_name
    __args__['skip'] = skip
    __args__['spaceId'] = space_id
    __args__['take'] = take
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('octopusdeploy:index/getMachinePolicies:getMachinePolicies', __args__, opts=opts, typ=GetMachinePoliciesResult).value

    return AwaitableGetMachinePoliciesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        machine_policies=pulumi.get(__ret__, 'machine_policies'),
        partial_name=pulumi.get(__ret__, 'partial_name'),
        skip=pulumi.get(__ret__, 'skip'),
        space_id=pulumi.get(__ret__, 'space_id'),
        take=pulumi.get(__ret__, 'take'))


@_utilities.lift_output_func(get_machine_policies)
def get_machine_policies_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                partial_name: Optional[pulumi.Input[Optional[str]]] = None,
                                skip: Optional[pulumi.Input[Optional[int]]] = None,
                                space_id: Optional[pulumi.Input[Optional[str]]] = None,
                                take: Optional[pulumi.Input[Optional[int]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMachinePoliciesResult]:
    """
    Provides information about existing machine policies.


    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    ...
