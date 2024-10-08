# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['KubernetesAgentWorkerArgs', 'KubernetesAgentWorker']

@pulumi.input_type
class KubernetesAgentWorkerArgs:
    def __init__(__self__, *,
                 thumbprint: pulumi.Input[str],
                 uri: pulumi.Input[str],
                 worker_pool_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 communication_mode: Optional[pulumi.Input[str]] = None,
                 is_disabled: Optional[pulumi.Input[bool]] = None,
                 machine_policy_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 upgrade_locked: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a KubernetesAgentWorker resource.
        :param pulumi.Input[str] thumbprint: The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
        :param pulumi.Input[str] uri: The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] worker_pool_ids: A list of worker pool Ids specifying the pools in which this worker belongs
        :param pulumi.Input[str] communication_mode: The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
        :param pulumi.Input[bool] is_disabled: Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
        :param pulumi.Input[str] machine_policy_id: Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[bool] upgrade_locked: If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
        """
        pulumi.set(__self__, "thumbprint", thumbprint)
        pulumi.set(__self__, "uri", uri)
        pulumi.set(__self__, "worker_pool_ids", worker_pool_ids)
        if communication_mode is not None:
            pulumi.set(__self__, "communication_mode", communication_mode)
        if is_disabled is not None:
            pulumi.set(__self__, "is_disabled", is_disabled)
        if machine_policy_id is not None:
            pulumi.set(__self__, "machine_policy_id", machine_policy_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if upgrade_locked is not None:
            pulumi.set(__self__, "upgrade_locked", upgrade_locked)

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Input[str]:
        """
        The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: pulumi.Input[str]):
        pulumi.set(self, "thumbprint", value)

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Input[str]:
        """
        The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter(name="workerPoolIds")
    def worker_pool_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of worker pool Ids specifying the pools in which this worker belongs
        """
        return pulumi.get(self, "worker_pool_ids")

    @worker_pool_ids.setter
    def worker_pool_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "worker_pool_ids", value)

    @property
    @pulumi.getter(name="communicationMode")
    def communication_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
        """
        return pulumi.get(self, "communication_mode")

    @communication_mode.setter
    def communication_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "communication_mode", value)

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
        """
        return pulumi.get(self, "is_disabled")

    @is_disabled.setter
    def is_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_disabled", value)

    @property
    @pulumi.getter(name="machinePolicyId")
    def machine_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
        """
        return pulumi.get(self, "machine_policy_id")

    @machine_policy_id.setter
    def machine_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_policy_id", value)

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
    @pulumi.getter(name="upgradeLocked")
    def upgrade_locked(self) -> Optional[pulumi.Input[bool]]:
        """
        If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
        """
        return pulumi.get(self, "upgrade_locked")

    @upgrade_locked.setter
    def upgrade_locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "upgrade_locked", value)


@pulumi.input_type
class _KubernetesAgentWorkerState:
    def __init__(__self__, *,
                 agent_helm_release_name: Optional[pulumi.Input[str]] = None,
                 agent_kubernetes_namespace: Optional[pulumi.Input[str]] = None,
                 agent_tentacle_version: Optional[pulumi.Input[str]] = None,
                 agent_upgrade_status: Optional[pulumi.Input[str]] = None,
                 agent_version: Optional[pulumi.Input[str]] = None,
                 communication_mode: Optional[pulumi.Input[str]] = None,
                 is_disabled: Optional[pulumi.Input[bool]] = None,
                 machine_policy_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 upgrade_locked: Optional[pulumi.Input[bool]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 worker_pool_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering KubernetesAgentWorker resources.
        :param pulumi.Input[str] agent_helm_release_name: Name of the Helm release that the agent belongs to.
        :param pulumi.Input[str] agent_kubernetes_namespace: Name of the Kubernetes namespace where the agent is installed.
        :param pulumi.Input[str] agent_tentacle_version: Current Tentacle version of the agent
        :param pulumi.Input[str] agent_upgrade_status: Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
        :param pulumi.Input[str] agent_version: Current Helm chart version of the agent.
        :param pulumi.Input[str] communication_mode: The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
        :param pulumi.Input[bool] is_disabled: Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
        :param pulumi.Input[str] machine_policy_id: Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[str] thumbprint: The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
        :param pulumi.Input[bool] upgrade_locked: If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
        :param pulumi.Input[str] uri: The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] worker_pool_ids: A list of worker pool Ids specifying the pools in which this worker belongs
        """
        if agent_helm_release_name is not None:
            pulumi.set(__self__, "agent_helm_release_name", agent_helm_release_name)
        if agent_kubernetes_namespace is not None:
            pulumi.set(__self__, "agent_kubernetes_namespace", agent_kubernetes_namespace)
        if agent_tentacle_version is not None:
            pulumi.set(__self__, "agent_tentacle_version", agent_tentacle_version)
        if agent_upgrade_status is not None:
            pulumi.set(__self__, "agent_upgrade_status", agent_upgrade_status)
        if agent_version is not None:
            pulumi.set(__self__, "agent_version", agent_version)
        if communication_mode is not None:
            pulumi.set(__self__, "communication_mode", communication_mode)
        if is_disabled is not None:
            pulumi.set(__self__, "is_disabled", is_disabled)
        if machine_policy_id is not None:
            pulumi.set(__self__, "machine_policy_id", machine_policy_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if thumbprint is not None:
            pulumi.set(__self__, "thumbprint", thumbprint)
        if upgrade_locked is not None:
            pulumi.set(__self__, "upgrade_locked", upgrade_locked)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)
        if worker_pool_ids is not None:
            pulumi.set(__self__, "worker_pool_ids", worker_pool_ids)

    @property
    @pulumi.getter(name="agentHelmReleaseName")
    def agent_helm_release_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Helm release that the agent belongs to.
        """
        return pulumi.get(self, "agent_helm_release_name")

    @agent_helm_release_name.setter
    def agent_helm_release_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_helm_release_name", value)

    @property
    @pulumi.getter(name="agentKubernetesNamespace")
    def agent_kubernetes_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Kubernetes namespace where the agent is installed.
        """
        return pulumi.get(self, "agent_kubernetes_namespace")

    @agent_kubernetes_namespace.setter
    def agent_kubernetes_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_kubernetes_namespace", value)

    @property
    @pulumi.getter(name="agentTentacleVersion")
    def agent_tentacle_version(self) -> Optional[pulumi.Input[str]]:
        """
        Current Tentacle version of the agent
        """
        return pulumi.get(self, "agent_tentacle_version")

    @agent_tentacle_version.setter
    def agent_tentacle_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_tentacle_version", value)

    @property
    @pulumi.getter(name="agentUpgradeStatus")
    def agent_upgrade_status(self) -> Optional[pulumi.Input[str]]:
        """
        Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
        """
        return pulumi.get(self, "agent_upgrade_status")

    @agent_upgrade_status.setter
    def agent_upgrade_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_upgrade_status", value)

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> Optional[pulumi.Input[str]]:
        """
        Current Helm chart version of the agent.
        """
        return pulumi.get(self, "agent_version")

    @agent_version.setter
    def agent_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_version", value)

    @property
    @pulumi.getter(name="communicationMode")
    def communication_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
        """
        return pulumi.get(self, "communication_mode")

    @communication_mode.setter
    def communication_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "communication_mode", value)

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
        """
        return pulumi.get(self, "is_disabled")

    @is_disabled.setter
    def is_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_disabled", value)

    @property
    @pulumi.getter(name="machinePolicyId")
    def machine_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
        """
        return pulumi.get(self, "machine_policy_id")

    @machine_policy_id.setter
    def machine_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_policy_id", value)

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
    @pulumi.getter
    def thumbprint(self) -> Optional[pulumi.Input[str]]:
        """
        The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thumbprint", value)

    @property
    @pulumi.getter(name="upgradeLocked")
    def upgrade_locked(self) -> Optional[pulumi.Input[bool]]:
        """
        If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
        """
        return pulumi.get(self, "upgrade_locked")

    @upgrade_locked.setter
    def upgrade_locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "upgrade_locked", value)

    @property
    @pulumi.getter
    def uri(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter(name="workerPoolIds")
    def worker_pool_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of worker pool Ids specifying the pools in which this worker belongs
        """
        return pulumi.get(self, "worker_pool_ids")

    @worker_pool_ids.setter
    def worker_pool_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "worker_pool_ids", value)


class KubernetesAgentWorker(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 communication_mode: Optional[pulumi.Input[str]] = None,
                 is_disabled: Optional[pulumi.Input[bool]] = None,
                 machine_policy_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 upgrade_locked: Optional[pulumi.Input[bool]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 worker_pool_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resource manages Kubernetes agent workers in Octopus Deploy.

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/kubernetesAgentWorker:KubernetesAgentWorker [options] octopusdeploy_kubernetes_agent_worker.<name> <machine-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] communication_mode: The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
        :param pulumi.Input[bool] is_disabled: Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
        :param pulumi.Input[str] machine_policy_id: Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[str] thumbprint: The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
        :param pulumi.Input[bool] upgrade_locked: If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
        :param pulumi.Input[str] uri: The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] worker_pool_ids: A list of worker pool Ids specifying the pools in which this worker belongs
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubernetesAgentWorkerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages Kubernetes agent workers in Octopus Deploy.

        ## Import

        ```sh
        $ pulumi import octopusdeploy:index/kubernetesAgentWorker:KubernetesAgentWorker [options] octopusdeploy_kubernetes_agent_worker.<name> <machine-id>
        ```

        :param str resource_name: The name of the resource.
        :param KubernetesAgentWorkerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubernetesAgentWorkerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 communication_mode: Optional[pulumi.Input[str]] = None,
                 is_disabled: Optional[pulumi.Input[bool]] = None,
                 machine_policy_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 upgrade_locked: Optional[pulumi.Input[bool]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 worker_pool_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubernetesAgentWorkerArgs.__new__(KubernetesAgentWorkerArgs)

            __props__.__dict__["communication_mode"] = communication_mode
            __props__.__dict__["is_disabled"] = is_disabled
            __props__.__dict__["machine_policy_id"] = machine_policy_id
            __props__.__dict__["name"] = name
            __props__.__dict__["space_id"] = space_id
            if thumbprint is None and not opts.urn:
                raise TypeError("Missing required property 'thumbprint'")
            __props__.__dict__["thumbprint"] = thumbprint
            __props__.__dict__["upgrade_locked"] = upgrade_locked
            if uri is None and not opts.urn:
                raise TypeError("Missing required property 'uri'")
            __props__.__dict__["uri"] = uri
            if worker_pool_ids is None and not opts.urn:
                raise TypeError("Missing required property 'worker_pool_ids'")
            __props__.__dict__["worker_pool_ids"] = worker_pool_ids
            __props__.__dict__["agent_helm_release_name"] = None
            __props__.__dict__["agent_kubernetes_namespace"] = None
            __props__.__dict__["agent_tentacle_version"] = None
            __props__.__dict__["agent_upgrade_status"] = None
            __props__.__dict__["agent_version"] = None
        super(KubernetesAgentWorker, __self__).__init__(
            'octopusdeploy:index/kubernetesAgentWorker:KubernetesAgentWorker',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_helm_release_name: Optional[pulumi.Input[str]] = None,
            agent_kubernetes_namespace: Optional[pulumi.Input[str]] = None,
            agent_tentacle_version: Optional[pulumi.Input[str]] = None,
            agent_upgrade_status: Optional[pulumi.Input[str]] = None,
            agent_version: Optional[pulumi.Input[str]] = None,
            communication_mode: Optional[pulumi.Input[str]] = None,
            is_disabled: Optional[pulumi.Input[bool]] = None,
            machine_policy_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None,
            thumbprint: Optional[pulumi.Input[str]] = None,
            upgrade_locked: Optional[pulumi.Input[bool]] = None,
            uri: Optional[pulumi.Input[str]] = None,
            worker_pool_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'KubernetesAgentWorker':
        """
        Get an existing KubernetesAgentWorker resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_helm_release_name: Name of the Helm release that the agent belongs to.
        :param pulumi.Input[str] agent_kubernetes_namespace: Name of the Kubernetes namespace where the agent is installed.
        :param pulumi.Input[str] agent_tentacle_version: Current Tentacle version of the agent
        :param pulumi.Input[str] agent_upgrade_status: Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
        :param pulumi.Input[str] agent_version: Current Helm chart version of the agent.
        :param pulumi.Input[str] communication_mode: The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
        :param pulumi.Input[bool] is_disabled: Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
        :param pulumi.Input[str] machine_policy_id: Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[str] thumbprint: The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
        :param pulumi.Input[bool] upgrade_locked: If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
        :param pulumi.Input[str] uri: The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] worker_pool_ids: A list of worker pool Ids specifying the pools in which this worker belongs
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KubernetesAgentWorkerState.__new__(_KubernetesAgentWorkerState)

        __props__.__dict__["agent_helm_release_name"] = agent_helm_release_name
        __props__.__dict__["agent_kubernetes_namespace"] = agent_kubernetes_namespace
        __props__.__dict__["agent_tentacle_version"] = agent_tentacle_version
        __props__.__dict__["agent_upgrade_status"] = agent_upgrade_status
        __props__.__dict__["agent_version"] = agent_version
        __props__.__dict__["communication_mode"] = communication_mode
        __props__.__dict__["is_disabled"] = is_disabled
        __props__.__dict__["machine_policy_id"] = machine_policy_id
        __props__.__dict__["name"] = name
        __props__.__dict__["space_id"] = space_id
        __props__.__dict__["thumbprint"] = thumbprint
        __props__.__dict__["upgrade_locked"] = upgrade_locked
        __props__.__dict__["uri"] = uri
        __props__.__dict__["worker_pool_ids"] = worker_pool_ids
        return KubernetesAgentWorker(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentHelmReleaseName")
    def agent_helm_release_name(self) -> pulumi.Output[str]:
        """
        Name of the Helm release that the agent belongs to.
        """
        return pulumi.get(self, "agent_helm_release_name")

    @property
    @pulumi.getter(name="agentKubernetesNamespace")
    def agent_kubernetes_namespace(self) -> pulumi.Output[str]:
        """
        Name of the Kubernetes namespace where the agent is installed.
        """
        return pulumi.get(self, "agent_kubernetes_namespace")

    @property
    @pulumi.getter(name="agentTentacleVersion")
    def agent_tentacle_version(self) -> pulumi.Output[str]:
        """
        Current Tentacle version of the agent
        """
        return pulumi.get(self, "agent_tentacle_version")

    @property
    @pulumi.getter(name="agentUpgradeStatus")
    def agent_upgrade_status(self) -> pulumi.Output[str]:
        """
        Current upgrade availability status of the agent. One of 'NoUpgrades', 'UpgradeAvailable', 'UpgradeSuggested', 'UpgradeRequired'
        """
        return pulumi.get(self, "agent_upgrade_status")

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> pulumi.Output[str]:
        """
        Current Helm chart version of the agent.
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter(name="communicationMode")
    def communication_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The communication mode used by the Kubernetes agent to communicate with Octopus Server. Currently, the only supported value is 'Polling'.
        """
        return pulumi.get(self, "communication_mode")

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the Kubernetes agent is disabled. If the agent is disabled, it will not be included in any deployments.
        """
        return pulumi.get(self, "is_disabled")

    @property
    @pulumi.getter(name="machinePolicyId")
    def machine_policy_id(self) -> pulumi.Output[str]:
        """
        Optional ID of the machine policy that the Kubernetes agent will use. If not provided the default machine policy will be used.
        """
        return pulumi.get(self, "machine_policy_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Output[str]:
        """
        The thumbprint of the Kubernetes agent's certificate used by server to verify the identity of the agent. This is the same thumbprint that was used when installing the agent.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter(name="upgradeLocked")
    def upgrade_locked(self) -> pulumi.Output[Optional[bool]]:
        """
        If enabled the Kubernetes agent will not automatically upgrade and will stay on the currently installed version, even if the associated machine policy is configured to automatically upgrade.
        """
        return pulumi.get(self, "upgrade_locked")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        The URI of the Kubernetes agent's used by the server to queue messages. This is the same subscription uri that was used when installing the agent.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter(name="workerPoolIds")
    def worker_pool_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of worker pool Ids specifying the pools in which this worker belongs
        """
        return pulumi.get(self, "worker_pool_ids")

