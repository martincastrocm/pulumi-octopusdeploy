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
from ._inputs import *

__all__ = [
    'GetVariablesResult',
    'AwaitableGetVariablesResult',
    'get_variables',
    'get_variables_output',
]

@pulumi.output_type
class GetVariablesResult:
    """
    A collection of values returned by getVariables.
    """
    def __init__(__self__, description=None, id=None, is_editable=None, is_sensitive=None, name=None, owner_id=None, prompts=None, scopes=None, sensitive_value=None, space_id=None, type=None, value=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_editable and not isinstance(is_editable, bool):
            raise TypeError("Expected argument 'is_editable' to be a bool")
        pulumi.set(__self__, "is_editable", is_editable)
        if is_sensitive and not isinstance(is_sensitive, bool):
            raise TypeError("Expected argument 'is_sensitive' to be a bool")
        pulumi.set(__self__, "is_sensitive", is_sensitive)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if prompts and not isinstance(prompts, list):
            raise TypeError("Expected argument 'prompts' to be a list")
        pulumi.set(__self__, "prompts", prompts)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if sensitive_value and not isinstance(sensitive_value, str):
            raise TypeError("Expected argument 'sensitive_value' to be a str")
        pulumi.set(__self__, "sensitive_value", sensitive_value)
        if space_id and not isinstance(space_id, str):
            raise TypeError("Expected argument 'space_id' to be a str")
        pulumi.set(__self__, "space_id", space_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of this variable.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The identifier of the variable to find.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isEditable")
    def is_editable(self) -> bool:
        """
        Indicates whether or not this variable is considered editable.
        """
        return pulumi.get(self, "is_editable")

    @property
    @pulumi.getter(name="isSensitive")
    def is_sensitive(self) -> bool:
        """
        Indicates whether or not this resource is considered sensitive and should be kept secret.
        """
        return pulumi.get(self, "is_sensitive")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of variable to find.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        Owner ID for the variable to find.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def prompts(self) -> Sequence['outputs.GetVariablesPromptResult']:
        return pulumi.get(self, "prompts")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence['outputs.GetVariablesScopeResult']:
        """
        As variable names can appear more than once under different scopes, a VariableScope must also be provided
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter(name="sensitiveValue")
    def sensitive_value(self) -> str:
        return pulumi.get(self, "sensitive_value")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> str:
        """
        The space ID associated with this variable.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of variable represented by this resource. Valid types are `AmazonWebServicesAccount`, `AzureAccount`, `GoogleCloudAccount`, `UsernamePasswordAccount`, `Certificate`, `Sensitive`, `String`, `WorkerPool`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


class AwaitableGetVariablesResult(GetVariablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVariablesResult(
            description=self.description,
            id=self.id,
            is_editable=self.is_editable,
            is_sensitive=self.is_sensitive,
            name=self.name,
            owner_id=self.owner_id,
            prompts=self.prompts,
            scopes=self.scopes,
            sensitive_value=self.sensitive_value,
            space_id=self.space_id,
            type=self.type,
            value=self.value)


def get_variables(name: Optional[str] = None,
                  owner_id: Optional[str] = None,
                  scopes: Optional[Sequence[Union['GetVariablesScopeArgs', 'GetVariablesScopeArgsDict']]] = None,
                  space_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVariablesResult:
    """
    Provides information about existing variables.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_octopusdeploy as octopusdeploy

    example = octopusdeploy.get_variables()
    ```


    :param str name: The name of variable to find.
    :param str owner_id: Owner ID for the variable to find.
    :param Sequence[Union['GetVariablesScopeArgs', 'GetVariablesScopeArgsDict']] scopes: As variable names can appear more than once under different scopes, a VariableScope must also be provided
    :param str space_id: The space ID associated with this variable.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['ownerId'] = owner_id
    __args__['scopes'] = scopes
    __args__['spaceId'] = space_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('octopusdeploy:index/getVariables:getVariables', __args__, opts=opts, typ=GetVariablesResult).value

    return AwaitableGetVariablesResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        is_editable=pulumi.get(__ret__, 'is_editable'),
        is_sensitive=pulumi.get(__ret__, 'is_sensitive'),
        name=pulumi.get(__ret__, 'name'),
        owner_id=pulumi.get(__ret__, 'owner_id'),
        prompts=pulumi.get(__ret__, 'prompts'),
        scopes=pulumi.get(__ret__, 'scopes'),
        sensitive_value=pulumi.get(__ret__, 'sensitive_value'),
        space_id=pulumi.get(__ret__, 'space_id'),
        type=pulumi.get(__ret__, 'type'),
        value=pulumi.get(__ret__, 'value'))


@_utilities.lift_output_func(get_variables)
def get_variables_output(name: Optional[pulumi.Input[str]] = None,
                         owner_id: Optional[pulumi.Input[str]] = None,
                         scopes: Optional[pulumi.Input[Sequence[Union['GetVariablesScopeArgs', 'GetVariablesScopeArgsDict']]]] = None,
                         space_id: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVariablesResult]:
    """
    Provides information about existing variables.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_octopusdeploy as octopusdeploy

    example = octopusdeploy.get_variables()
    ```


    :param str name: The name of variable to find.
    :param str owner_id: Owner ID for the variable to find.
    :param Sequence[Union['GetVariablesScopeArgs', 'GetVariablesScopeArgsDict']] scopes: As variable names can appear more than once under different scopes, a VariableScope must also be provided
    :param str space_id: The space ID associated with this variable.
    """
    ...
