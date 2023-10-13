/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APIServerAccessConfigObservation struct {

	// IP ranges authorized to access the Kubernetes API server
	AuthorizedIPRanges []*string `json:"authorizedIpRanges,omitempty" tf:"authorized_ip_ranges,omitempty"`

	// Enable Private Cluster
	EnablePrivateCluster *bool `json:"enablePrivateCluster,omitempty" tf:"enable_private_cluster,omitempty"`
}

type APIServerAccessConfigParameters struct {

	// IP ranges authorized to access the Kubernetes API server
	// +kubebuilder:validation:Optional
	AuthorizedIPRanges []*string `json:"authorizedIpRanges,omitempty" tf:"authorized_ip_ranges,omitempty"`

	// Enable Private Cluster
	// +kubebuilder:validation:Required
	EnablePrivateCluster *bool `json:"enablePrivateCluster" tf:"enable_private_cluster,omitempty"`
}

type AadConfigObservation struct {

	// List of AAD group object IDs that will have admin role of the cluster.
	AdminGroupIds []*string `json:"adminGroupIds,omitempty" tf:"admin_group_ids,omitempty"`

	// Enable Azure RBAC for Kubernetes authorization
	EnableAzureRbac *bool `json:"enableAzureRbac,omitempty" tf:"enable_azure_rbac,omitempty"`

	// Enable Managed RBAC
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AadConfigParameters struct {

	// List of AAD group object IDs that will have admin role of the cluster.
	// +kubebuilder:validation:Optional
	AdminGroupIds []*string `json:"adminGroupIds,omitempty" tf:"admin_group_ids,omitempty"`

	// Enable Azure RBAC for Kubernetes authorization
	// +kubebuilder:validation:Optional
	EnableAzureRbac *bool `json:"enableAzureRbac,omitempty" tf:"enable_azure_rbac,omitempty"`

	// Enable Managed RBAC
	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AccessConfigObservation struct {

	// Azure Active Directory config
	AadConfig []AadConfigObservation `json:"aadConfig,omitempty" tf:"aad_config,omitempty"`

	// Disable local accounts
	DisableLocalAccounts *bool `json:"disableLocalAccounts,omitempty" tf:"disable_local_accounts,omitempty"`

	// Enable kubernetes RBAC
	EnableRbac *bool `json:"enableRbac,omitempty" tf:"enable_rbac,omitempty"`
}

type AccessConfigParameters struct {

	// Azure Active Directory config
	// +kubebuilder:validation:Optional
	AadConfig []AadConfigParameters `json:"aadConfig,omitempty" tf:"aad_config,omitempty"`

	// Disable local accounts
	// +kubebuilder:validation:Optional
	DisableLocalAccounts *bool `json:"disableLocalAccounts,omitempty" tf:"disable_local_accounts,omitempty"`

	// Enable kubernetes RBAC
	// +kubebuilder:validation:Optional
	EnableRbac *bool `json:"enableRbac,omitempty" tf:"enable_rbac,omitempty"`
}

type AddonConfigObservation struct {

	// Keyvault secrets provider addon
	AzureKeyvaultSecretsProviderAddonConfig []AzureKeyvaultSecretsProviderAddonConfigObservation `json:"azureKeyvaultSecretsProviderAddonConfig,omitempty" tf:"azure_keyvault_secrets_provider_addon_config,omitempty"`

	// Azure policy addon
	AzurePolicyAddonConfig []AzurePolicyAddonConfigObservation `json:"azurePolicyAddonConfig,omitempty" tf:"azure_policy_addon_config,omitempty"`

	// Monitor addon
	MonitorAddonConfig []MonitorAddonConfigObservation `json:"monitorAddonConfig,omitempty" tf:"monitor_addon_config,omitempty"`
}

type AddonConfigParameters struct {

	// Keyvault secrets provider addon
	// +kubebuilder:validation:Optional
	AzureKeyvaultSecretsProviderAddonConfig []AzureKeyvaultSecretsProviderAddonConfigParameters `json:"azureKeyvaultSecretsProviderAddonConfig,omitempty" tf:"azure_keyvault_secrets_provider_addon_config,omitempty"`

	// Azure policy addon
	// +kubebuilder:validation:Optional
	AzurePolicyAddonConfig []AzurePolicyAddonConfigParameters `json:"azurePolicyAddonConfig,omitempty" tf:"azure_policy_addon_config,omitempty"`

	// Monitor addon
	// +kubebuilder:validation:Optional
	MonitorAddonConfig []MonitorAddonConfigParameters `json:"monitorAddonConfig,omitempty" tf:"monitor_addon_config,omitempty"`
}

type AksClusterObservation struct {

	// Name of the Azure Credential in Tanzu Mission Control
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata for the resource
	Meta []MetaObservation `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of this cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero.  The default duration is 30m
	ReadyWaitTimeout *string `json:"readyWaitTimeout,omitempty" tf:"ready_wait_timeout,omitempty"`

	// Resource group for this cluster
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// Spec for the cluster
	Spec []SpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// Azure Subscription for this cluster
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type AksClusterParameters struct {

	// Name of the Azure Credential in Tanzu Mission Control
	// +kubebuilder:validation:Optional
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	// Metadata for the resource
	// +kubebuilder:validation:Optional
	Meta []MetaParameters `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of this cluster
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero.  The default duration is 30m
	// +kubebuilder:validation:Optional
	ReadyWaitTimeout *string `json:"readyWaitTimeout,omitempty" tf:"ready_wait_timeout,omitempty"`

	// Resource group for this cluster
	// +kubebuilder:validation:Optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// Spec for the cluster
	// +kubebuilder:validation:Optional
	Spec []SpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// Azure Subscription for this cluster
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type AutoScalingConfigObservation struct {

	// Enable auto scaling
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Maximum node count
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// Minimum node count
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
}

type AutoScalingConfigParameters struct {

	// Enable auto scaling
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Maximum node count
	// +kubebuilder:validation:Optional
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// Minimum node count
	// +kubebuilder:validation:Optional
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
}

type AutoUpgradeConfigObservation struct {

	// Upgrade Channel. Allowed values include: NONE, PATCH, STABLE, RAPID or NODE_IMAGE
	UpgradeChannel *string `json:"upgradeChannel,omitempty" tf:"upgrade_channel,omitempty"`
}

type AutoUpgradeConfigParameters struct {

	// Upgrade Channel. Allowed values include: NONE, PATCH, STABLE, RAPID or NODE_IMAGE
	// +kubebuilder:validation:Optional
	UpgradeChannel *string `json:"upgradeChannel,omitempty" tf:"upgrade_channel,omitempty"`
}

type AzureKeyvaultSecretsProviderAddonConfigObservation struct {

	// Enable Azure Key Vault Secrets Provider
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Enable secrets rotation
	EnableSecretRotation *bool `json:"enableSecretRotation,omitempty" tf:"enable_secret_rotation,omitempty"`

	// Secret rotation interval
	RotationPollInterval *string `json:"rotationPollInterval,omitempty" tf:"rotation_poll_interval,omitempty"`
}

type AzureKeyvaultSecretsProviderAddonConfigParameters struct {

	// Enable Azure Key Vault Secrets Provider
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Enable secrets rotation
	// +kubebuilder:validation:Optional
	EnableSecretRotation *bool `json:"enableSecretRotation,omitempty" tf:"enable_secret_rotation,omitempty"`

	// Secret rotation interval
	// +kubebuilder:validation:Optional
	RotationPollInterval *string `json:"rotationPollInterval,omitempty" tf:"rotation_poll_interval,omitempty"`
}

type AzurePolicyAddonConfigObservation struct {

	// Enable policy addon
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type AzurePolicyAddonConfigParameters struct {

	// Enable policy addon
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type ConfigObservation struct {

	// API Server Access Config
	APIServerAccessConfig []APIServerAccessConfigObservation `json:"apiServerAccessConfig,omitempty" tf:"api_server_access_config,omitempty"`

	// Access config
	AccessConfig []AccessConfigObservation `json:"accessConfig,omitempty" tf:"access_config,omitempty"`

	// Addons Config
	AddonConfig []AddonConfigObservation `json:"addonConfig,omitempty" tf:"addon_config,omitempty"`

	// Auto Upgrade Config
	AutoUpgradeConfig []AutoUpgradeConfigObservation `json:"autoUpgradeConfig,omitempty" tf:"auto_upgrade_config,omitempty"`

	// Resource ID of the disk encryption set to use for enabling
	DiskEncryptionSet *string `json:"diskEncryptionSet,omitempty" tf:"disk_encryption_set,omitempty"`

	// Kubernetes version of the cluster
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`

	// Linux Config
	LinuxConfig []LinuxConfigObservation `json:"linuxConfig,omitempty" tf:"linux_config,omitempty"`

	// The geo-location where the resource lives for the cluster.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Network Config
	NetworkConfig []NetworkConfigObservation `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// Name of the resource group containing nodepools.
	NodeResourceGroupName *string `json:"nodeResourceGroupName,omitempty" tf:"node_resource_group_name,omitempty"`

	// Azure Kubernetes Service SKU
	Sku []SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// Storage Config
	StorageConfig []StorageConfigObservation `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`

	// Metadata to apply to the cluster to assist with categorization and organization
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigParameters struct {

	// API Server Access Config
	// +kubebuilder:validation:Optional
	APIServerAccessConfig []APIServerAccessConfigParameters `json:"apiServerAccessConfig,omitempty" tf:"api_server_access_config,omitempty"`

	// Access config
	// +kubebuilder:validation:Optional
	AccessConfig []AccessConfigParameters `json:"accessConfig,omitempty" tf:"access_config,omitempty"`

	// Addons Config
	// +kubebuilder:validation:Optional
	AddonConfig []AddonConfigParameters `json:"addonConfig,omitempty" tf:"addon_config,omitempty"`

	// Auto Upgrade Config
	// +kubebuilder:validation:Optional
	AutoUpgradeConfig []AutoUpgradeConfigParameters `json:"autoUpgradeConfig,omitempty" tf:"auto_upgrade_config,omitempty"`

	// Resource ID of the disk encryption set to use for enabling
	// +kubebuilder:validation:Optional
	DiskEncryptionSet *string `json:"diskEncryptionSet,omitempty" tf:"disk_encryption_set,omitempty"`

	// Kubernetes version of the cluster
	// +kubebuilder:validation:Required
	KubernetesVersion *string `json:"kubernetesVersion" tf:"kubernetes_version,omitempty"`

	// Linux Config
	// +kubebuilder:validation:Optional
	LinuxConfig []LinuxConfigParameters `json:"linuxConfig,omitempty" tf:"linux_config,omitempty"`

	// The geo-location where the resource lives for the cluster.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Network Config
	// +kubebuilder:validation:Required
	NetworkConfig []NetworkConfigParameters `json:"networkConfig" tf:"network_config,omitempty"`

	// Name of the resource group containing nodepools.
	// +kubebuilder:validation:Optional
	NodeResourceGroupName *string `json:"nodeResourceGroupName,omitempty" tf:"node_resource_group_name,omitempty"`

	// Azure Kubernetes Service SKU
	// +kubebuilder:validation:Optional
	Sku []SkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// Storage Config
	// +kubebuilder:validation:Optional
	StorageConfig []StorageConfigParameters `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`

	// Metadata to apply to the cluster to assist with categorization and organization
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LinuxConfigObservation struct {

	// Administrator username to use for Linux VMs
	AdminUsername *string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`

	// Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
}

type LinuxConfigParameters struct {

	// Administrator username to use for Linux VMs
	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers
	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
}

type MetaObservation struct {

	// Annotations for the resource
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Description of the resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels for the resource
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Resource version of the resource
	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`

	// UID of the resource
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type MetaParameters struct {

	// Annotations for the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Description of the resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels for the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type MonitorAddonConfigObservation struct {

	// Enable monitor
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Log analytics workspace ID for the monitoring addon
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`
}

type MonitorAddonConfigParameters struct {

	// Enable monitor
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Log analytics workspace ID for the monitoring addon
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`
}

type NetworkConfigObservation struct {

	// DNS prefix of the cluster
	DNSPrefix *string `json:"dnsPrefix,omitempty" tf:"dns_prefix,omitempty"`

	// IP address assigned to the Kubernetes DNS service
	DNSServiceIP *string `json:"dnsServiceIp,omitempty" tf:"dns_service_ip,omitempty"`

	// A CIDR notation IP range assigned to the Docker bridge network
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty" tf:"docker_bridge_cidr,omitempty"`

	// Load balancer SKU
	LoadBalancerSku *string `json:"loadBalancerSku,omitempty" tf:"load_balancer_sku,omitempty"`

	// Network plugin
	NetworkPlugin *string `json:"networkPlugin,omitempty" tf:"network_plugin,omitempty"`

	// Network policy
	NetworkPolicy *string `json:"networkPolicy,omitempty" tf:"network_policy,omitempty"`

	// CIDR notation IP ranges from which to assign pod IPs
	PodCidr []*string `json:"podCidr,omitempty" tf:"pod_cidr,omitempty"`

	// CIDR notation IP ranges from which to assign service cluster IP
	ServiceCidr []*string `json:"serviceCidr,omitempty" tf:"service_cidr,omitempty"`
}

type NetworkConfigParameters struct {

	// DNS prefix of the cluster
	// +kubebuilder:validation:Required
	DNSPrefix *string `json:"dnsPrefix" tf:"dns_prefix,omitempty"`

	// IP address assigned to the Kubernetes DNS service
	// +kubebuilder:validation:Optional
	DNSServiceIP *string `json:"dnsServiceIp,omitempty" tf:"dns_service_ip,omitempty"`

	// A CIDR notation IP range assigned to the Docker bridge network
	// +kubebuilder:validation:Optional
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty" tf:"docker_bridge_cidr,omitempty"`

	// Load balancer SKU
	// +kubebuilder:validation:Optional
	LoadBalancerSku *string `json:"loadBalancerSku,omitempty" tf:"load_balancer_sku,omitempty"`

	// Network plugin
	// +kubebuilder:validation:Optional
	NetworkPlugin *string `json:"networkPlugin,omitempty" tf:"network_plugin,omitempty"`

	// Network policy
	// +kubebuilder:validation:Optional
	NetworkPolicy *string `json:"networkPolicy,omitempty" tf:"network_policy,omitempty"`

	// CIDR notation IP ranges from which to assign pod IPs
	// +kubebuilder:validation:Optional
	PodCidr []*string `json:"podCidr,omitempty" tf:"pod_cidr,omitempty"`

	// CIDR notation IP ranges from which to assign service cluster IP
	// +kubebuilder:validation:Optional
	ServiceCidr []*string `json:"serviceCidr,omitempty" tf:"service_cidr,omitempty"`
}

type NodepoolObservation struct {

	// Name of the nodepool, immutable
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Spec for the nodepool
	Spec []NodepoolSpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`
}

type NodepoolParameters struct {

	// Name of the nodepool, immutable
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Spec for the nodepool
	// +kubebuilder:validation:Required
	Spec []NodepoolSpecParameters `json:"spec" tf:"spec,omitempty"`
}

type NodepoolSpecObservation struct {

	// Auto scaling config.
	AutoScalingConfig []AutoScalingConfigObservation `json:"autoScalingConfig,omitempty" tf:"auto_scaling_config,omitempty"`

	// The list of Availability zones to use for nodepool. This can only be specified if the type of the nodepool is AvailabilitySet.
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Count is the number of nodes
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Whether each node is allocated its own public IP
	EnableNodePublicIP *bool `json:"enableNodePublicIp,omitempty" tf:"enable_node_public_ip,omitempty"`

	// The maximum number of pods that can run on a node
	MaxPods *float64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// The mode of the nodepool. Allowed values include: SYSTEM or USER. A cluster must have at least one 'SYSTEM' nodepool at all times.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The node image version of the nodepool.
	NodeImageVersion *string `json:"nodeImageVersion,omitempty" tf:"node_image_version,omitempty"`

	// The node labels to be persisted across all nodes in nodepool
	NodeLabels map[string]*string `json:"nodeLabels,omitempty" tf:"node_labels,omitempty"`

	// OS Disk Size in GB to be used to specify the disk size for every machine in the nodepool. If you specify 0, it will apply the default osDisk size according to the vmSize specified
	OsDiskSizeGb *float64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`

	// OS Disk Type. Allowed values include: EPHEMERAL or MANAGED.
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`

	// The OS type of the nodepool. Allowed values include: LINUX.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Scale set eviction policy, Allowed values include: DELETE or DEALLOCATE.
	ScaleSetEvictionPolicy *string `json:"scaleSetEvictionPolicy,omitempty" tf:"scale_set_eviction_policy,omitempty"`

	// Scale set priority. Allowed values include: REGULAR or SPOT.
	ScaleSetPriority *string `json:"scaleSetPriority,omitempty" tf:"scale_set_priority,omitempty"`

	// Max spot price
	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`

	// AKS specific node tags
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The taints added to new nodes during nodepool create and scale
	Taints []TaintsObservation `json:"taints,omitempty" tf:"taints,omitempty"`

	// The Nodepool type. Allowed values include: VIRTUAL_MACHINE_SCALE_SETS or AVAILABILITY_SET.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// upgrade config
	UpgradeConfig []UpgradeConfigObservation `json:"upgradeConfig,omitempty" tf:"upgrade_config,omitempty"`

	// Virtual Machine Size
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`

	// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes
	VnetSubnetID *string `json:"vnetSubnetId,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type NodepoolSpecParameters struct {

	// Auto scaling config.
	// +kubebuilder:validation:Optional
	AutoScalingConfig []AutoScalingConfigParameters `json:"autoScalingConfig,omitempty" tf:"auto_scaling_config,omitempty"`

	// The list of Availability zones to use for nodepool. This can only be specified if the type of the nodepool is AvailabilitySet.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Count is the number of nodes
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// Whether each node is allocated its own public IP
	// +kubebuilder:validation:Optional
	EnableNodePublicIP *bool `json:"enableNodePublicIp,omitempty" tf:"enable_node_public_ip,omitempty"`

	// The maximum number of pods that can run on a node
	// +kubebuilder:validation:Optional
	MaxPods *float64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// The mode of the nodepool. Allowed values include: SYSTEM or USER. A cluster must have at least one 'SYSTEM' nodepool at all times.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The node image version of the nodepool.
	// +kubebuilder:validation:Optional
	NodeImageVersion *string `json:"nodeImageVersion,omitempty" tf:"node_image_version,omitempty"`

	// The node labels to be persisted across all nodes in nodepool
	// +kubebuilder:validation:Optional
	NodeLabels map[string]*string `json:"nodeLabels,omitempty" tf:"node_labels,omitempty"`

	// OS Disk Size in GB to be used to specify the disk size for every machine in the nodepool. If you specify 0, it will apply the default osDisk size according to the vmSize specified
	// +kubebuilder:validation:Optional
	OsDiskSizeGb *float64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`

	// OS Disk Type. Allowed values include: EPHEMERAL or MANAGED.
	// +kubebuilder:validation:Optional
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`

	// The OS type of the nodepool. Allowed values include: LINUX.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Scale set eviction policy, Allowed values include: DELETE or DEALLOCATE.
	// +kubebuilder:validation:Optional
	ScaleSetEvictionPolicy *string `json:"scaleSetEvictionPolicy,omitempty" tf:"scale_set_eviction_policy,omitempty"`

	// Scale set priority. Allowed values include: REGULAR or SPOT.
	// +kubebuilder:validation:Optional
	ScaleSetPriority *string `json:"scaleSetPriority,omitempty" tf:"scale_set_priority,omitempty"`

	// Max spot price
	// +kubebuilder:validation:Optional
	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`

	// AKS specific node tags
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The taints added to new nodes during nodepool create and scale
	// +kubebuilder:validation:Optional
	Taints []TaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// The Nodepool type. Allowed values include: VIRTUAL_MACHINE_SCALE_SETS or AVAILABILITY_SET.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// upgrade config
	// +kubebuilder:validation:Optional
	UpgradeConfig []UpgradeConfigParameters `json:"upgradeConfig,omitempty" tf:"upgrade_config,omitempty"`

	// Virtual Machine Size
	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes
	// +kubebuilder:validation:Optional
	VnetSubnetID *string `json:"vnetSubnetId,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type SkuObservation struct {

	// Name of the cluster SKU. Allowed values include: BASIC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tier of the cluster SKU. Allowed values include: FREE or PAID.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SkuParameters struct {

	// Name of the cluster SKU. Allowed values include: BASIC.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Tier of the cluster SKU. Allowed values include: FREE or PAID.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SpecObservation struct {

	// Name of the cluster in TMC
	AgentName *string `json:"agentName,omitempty" tf:"agent_name,omitempty"`

	// Name of the cluster group to which this cluster belongs
	ClusterGroup *string `json:"clusterGroup,omitempty" tf:"cluster_group,omitempty"`

	// AKS config for the cluster control plane
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// Nodepool definitions for the cluster
	Nodepool []NodepoolObservation `json:"nodepool,omitempty" tf:"nodepool,omitempty"`

	// Optional proxy name is the name of the Proxy Config to be used for the cluster
	Proxy *string `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// Resource ID of the cluster in Azure.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type SpecParameters struct {

	// Name of the cluster in TMC
	// +kubebuilder:validation:Optional
	AgentName *string `json:"agentName,omitempty" tf:"agent_name,omitempty"`

	// Name of the cluster group to which this cluster belongs
	// +kubebuilder:validation:Optional
	ClusterGroup *string `json:"clusterGroup,omitempty" tf:"cluster_group,omitempty"`

	// AKS config for the cluster control plane
	// +kubebuilder:validation:Required
	Config []ConfigParameters `json:"config" tf:"config,omitempty"`

	// Nodepool definitions for the cluster
	// +kubebuilder:validation:Required
	Nodepool []NodepoolParameters `json:"nodepool" tf:"nodepool,omitempty"`

	// Optional proxy name is the name of the Proxy Config to be used for the cluster
	// +kubebuilder:validation:Optional
	Proxy *string `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// Resource ID of the cluster in Azure.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type StorageConfigObservation struct {

	// Enable the azure disk CSI driver for the storage
	EnableDiskCsiDriver *bool `json:"enableDiskCsiDriver,omitempty" tf:"enable_disk_csi_driver,omitempty"`

	// Enable the azure file CSI driver for the storage
	EnableFileCsiDriver *bool `json:"enableFileCsiDriver,omitempty" tf:"enable_file_csi_driver,omitempty"`

	// Enable the snapshot controller for the storage
	EnableSnapshotController *bool `json:"enableSnapshotController,omitempty" tf:"enable_snapshot_controller,omitempty"`
}

type StorageConfigParameters struct {

	// Enable the azure disk CSI driver for the storage
	// +kubebuilder:validation:Optional
	EnableDiskCsiDriver *bool `json:"enableDiskCsiDriver,omitempty" tf:"enable_disk_csi_driver,omitempty"`

	// Enable the azure file CSI driver for the storage
	// +kubebuilder:validation:Optional
	EnableFileCsiDriver *bool `json:"enableFileCsiDriver,omitempty" tf:"enable_file_csi_driver,omitempty"`

	// Enable the snapshot controller for the storage
	// +kubebuilder:validation:Optional
	EnableSnapshotController *bool `json:"enableSnapshotController,omitempty" tf:"enable_snapshot_controller,omitempty"`
}

type TaintsObservation struct {

	// Current effect state of the node pool
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// The taint key to be applied to a node
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The taint value corresponding to the taint key
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TaintsParameters struct {

	// Current effect state of the node pool
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// The taint key to be applied to a node
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The taint value corresponding to the taint key
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UpgradeConfigObservation struct {

	// Max Surge
	MaxSurge *string `json:"maxSurge,omitempty" tf:"max_surge,omitempty"`
}

type UpgradeConfigParameters struct {

	// Max Surge
	// +kubebuilder:validation:Optional
	MaxSurge *string `json:"maxSurge,omitempty" tf:"max_surge,omitempty"`
}

// AksClusterSpec defines the desired state of AksCluster
type AksClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AksClusterParameters `json:"forProvider"`
}

// AksClusterStatus defines the observed state of AksCluster.
type AksClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AksClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AksCluster is the Schema for the AksClusters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tmc}
type AksCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.credentialName)",message="credentialName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.resourceGroup)",message="resourceGroup is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.spec)",message="spec is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.subscriptionId)",message="subscriptionId is a required parameter"
	Spec   AksClusterSpec   `json:"spec"`
	Status AksClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AksClusterList contains a list of AksClusters
type AksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AksCluster `json:"items"`
}

// Repository type metadata.
var (
	AksCluster_Kind             = "AksCluster"
	AksCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AksCluster_Kind}.String()
	AksCluster_KindAPIVersion   = AksCluster_Kind + "." + CRDGroupVersion.String()
	AksCluster_GroupVersionKind = CRDGroupVersion.WithKind(AksCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&AksCluster{}, &AksClusterList{})
}