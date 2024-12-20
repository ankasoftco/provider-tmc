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

type AwsCredentialObservation struct {

	// Account ID of the AWS credential
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Generic credential
	GenericCredential *string `json:"genericCredential,omitempty" tf:"generic_credential,omitempty"`

	// AWS IAM role ARN and external ID
	IAMRole []IAMRoleObservation `json:"iamRole,omitempty" tf:"iam_role,omitempty"`
}

type AwsCredentialParameters struct {

	// Account ID of the AWS credential
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Generic credential
	// +kubebuilder:validation:Optional
	GenericCredential *string `json:"genericCredential,omitempty" tf:"generic_credential,omitempty"`

	// AWS IAM role ARN and external ID
	// +kubebuilder:validation:Optional
	IAMRole []IAMRoleParameters `json:"iamRole,omitempty" tf:"iam_role,omitempty"`
}

type AzureCredentialObservation struct {

	// Azure service principal
	ServicePrincipal []ServicePrincipalObservation `json:"servicePrincipal,omitempty" tf:"service_principal,omitempty"`

	// Azure service principal
	ServicePrincipalWithCertificate []ServicePrincipalWithCertificateObservation `json:"servicePrincipalWithCertificate,omitempty" tf:"service_principal_with_certificate,omitempty"`
}

type AzureCredentialParameters struct {

	// Azure service principal
	// +kubebuilder:validation:Optional
	ServicePrincipal []ServicePrincipalParameters `json:"servicePrincipal,omitempty" tf:"service_principal,omitempty"`

	// Azure service principal
	// +kubebuilder:validation:Optional
	ServicePrincipalWithCertificate []ServicePrincipalWithCertificateParameters `json:"servicePrincipalWithCertificate,omitempty" tf:"service_principal_with_certificate,omitempty"`
}

type CredentialMetaObservation struct {

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

type CredentialMetaParameters struct {

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

type CredentialObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata for the resource
	Meta []CredentialMetaObservation `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of this credential
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Wait timeout duration until credential resource reaches VALID state. Accepted timeout duration values like 5s, 5m, or 1h, higher than zero.
	ReadyWaitTimeout *string `json:"readyWaitTimeout,omitempty" tf:"ready_wait_timeout,omitempty"`

	// Spec of credential resource
	Spec []CredentialSpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// Status of credential resource
	Status map[string]*string `json:"status,omitempty" tf:"status,omitempty"`
}

type CredentialParameters struct {

	// Metadata for the resource
	// +kubebuilder:validation:Optional
	Meta []CredentialMetaParameters `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of this credential
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Wait timeout duration until credential resource reaches VALID state. Accepted timeout duration values like 5s, 5m, or 1h, higher than zero.
	// +kubebuilder:validation:Optional
	ReadyWaitTimeout *string `json:"readyWaitTimeout,omitempty" tf:"ready_wait_timeout,omitempty"`

	// Spec of credential resource
	// +kubebuilder:validation:Optional
	Spec []CredentialSpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type CredentialSpecObservation struct {

	// The Tanzu capability for which the credential shall be used. Value must be in list [DATA_PROTECTION TANZU_OBSERVABILITY TANZU_SERVICE_MESH PROXY_CONFIG MANAGED_K8S_PROVIDER IMAGE_REGISTRY]
	Capability *string `json:"capability,omitempty" tf:"capability,omitempty"`

	// Holds credentials sensitive data
	Data []DataObservation `json:"data,omitempty" tf:"data,omitempty"`

	// The Tanzu provider for which describes credential data type. Value must be in list [PROVIDER_UNSPECIFIED,AWS_EC2,GENERIC_S3,AZURE_AD,AWS_EKS,AZURE_AKS,GENERIC_KEY_VALUE]
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`
}

type CredentialSpecParameters struct {

	// The Tanzu capability for which the credential shall be used. Value must be in list [DATA_PROTECTION TANZU_OBSERVABILITY TANZU_SERVICE_MESH PROXY_CONFIG MANAGED_K8S_PROVIDER IMAGE_REGISTRY]
	// +kubebuilder:validation:Optional
	Capability *string `json:"capability,omitempty" tf:"capability,omitempty"`

	// Holds credentials sensitive data
	// +kubebuilder:validation:Optional
	Data []DataParameters `json:"data,omitempty" tf:"data,omitempty"`

	// The Tanzu provider for which describes credential data type. Value must be in list [PROVIDER_UNSPECIFIED,AWS_EC2,GENERIC_S3,AZURE_AD,AWS_EKS,AZURE_AKS,GENERIC_KEY_VALUE]
	// +kubebuilder:validation:Optional
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`
}

type DataObservation struct {

	// AWS credential data type
	AwsCredential []AwsCredentialObservation `json:"awsCredential,omitempty" tf:"aws_credential,omitempty"`

	// Azure credential
	AzureCredential []AzureCredentialObservation `json:"azureCredential,omitempty" tf:"azure_credential,omitempty"`

	// Generic credential data type used to hold a blob of data represented as string
	GenericCredential *string `json:"genericCredential,omitempty" tf:"generic_credential,omitempty"`

	// Key Value credential
	KeyValue []KeyValueObservation `json:"keyValue,omitempty" tf:"key_value,omitempty"`
}

type DataParameters struct {

	// AWS credential data type
	// +kubebuilder:validation:Optional
	AwsCredential []AwsCredentialParameters `json:"awsCredential,omitempty" tf:"aws_credential,omitempty"`

	// Azure credential
	// +kubebuilder:validation:Optional
	AzureCredential []AzureCredentialParameters `json:"azureCredential,omitempty" tf:"azure_credential,omitempty"`

	// Generic credential data type used to hold a blob of data represented as string
	// +kubebuilder:validation:Optional
	GenericCredential *string `json:"genericCredential,omitempty" tf:"generic_credential,omitempty"`

	// Key Value credential
	// +kubebuilder:validation:Optional
	KeyValue []KeyValueParameters `json:"keyValue,omitempty" tf:"key_value,omitempty"`
}

type IAMRoleObservation struct {

	// AWS IAM role ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// An external ID used to assume an AWS IAM role
	ExtID *string `json:"extId,omitempty" tf:"ext_id,omitempty"`
}

type IAMRoleParameters struct {

	// AWS IAM role ARN
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// An external ID used to assume an AWS IAM role
	// +kubebuilder:validation:Optional
	ExtID *string `json:"extId,omitempty" tf:"ext_id,omitempty"`
}

type KeyValueObservation struct {

	// Data secret data in the format of key-value pair
	Data map[string]*string `json:"data,omitempty" tf:"data,omitempty"`

	// Type of Secret data, usually mapped to k8s secret type. Supported types: [SECRET_TYPE_UNSPECIFIED,OPAQUE_SECRET_TYPE,DOCKERCONFIGJSON_SECRET_TYPE]
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type KeyValueParameters struct {

	// Data secret data in the format of key-value pair
	// +kubebuilder:validation:Optional
	Data map[string]*string `json:"data,omitempty" tf:"data,omitempty"`

	// Type of Secret data, usually mapped to k8s secret type. Supported types: [SECRET_TYPE_UNSPECIFIED,OPAQUE_SECRET_TYPE,DOCKERCONFIGJSON_SECRET_TYPE]
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServicePrincipalObservation struct {

	// Azure Cloud name
	AzureCloudName *string `json:"azureCloudName,omitempty" tf:"azure_cloud_name,omitempty"`

	// Client ID of the Service Principal
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Client Secret of the Service Principal
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// Resource Group name
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// Subscription ID of the Azure credential
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// Tenant ID of the Azure credential
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ServicePrincipalParameters struct {

	// Azure Cloud name
	// +kubebuilder:validation:Optional
	AzureCloudName *string `json:"azureCloudName,omitempty" tf:"azure_cloud_name,omitempty"`

	// Client ID of the Service Principal
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// Client Secret of the Service Principal
	// +kubebuilder:validation:Optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret,omitempty"`

	// Resource Group name
	// +kubebuilder:validation:Required
	ResourceGroup *string `json:"resourceGroup" tf:"resource_group,omitempty"`

	// Subscription ID of the Azure credential
	// +kubebuilder:validation:Required
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`

	// Tenant ID of the Azure credential
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

type ServicePrincipalWithCertificateObservation struct {

	// Azure Cloud name
	AzureCloudName *string `json:"azureCloudName,omitempty" tf:"azure_cloud_name,omitempty"`

	// Client certificate of the Service Principal
	ClientCertificate *string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`

	// Client ID of the Service Principal
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// IDs of the Azure Subscriptions that the Service Principal can manage
	ManagedSubscriptions []*string `json:"managedSubscriptions,omitempty" tf:"managed_subscriptions,omitempty"`

	// Subscription ID of the Azure credential
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// Tenant ID of the Azure credential
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ServicePrincipalWithCertificateParameters struct {

	// Azure Cloud name
	// +kubebuilder:validation:Optional
	AzureCloudName *string `json:"azureCloudName,omitempty" tf:"azure_cloud_name,omitempty"`

	// Client certificate of the Service Principal
	// +kubebuilder:validation:Required
	ClientCertificate *string `json:"clientCertificate" tf:"client_certificate,omitempty"`

	// Client ID of the Service Principal
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// IDs of the Azure Subscriptions that the Service Principal can manage
	// +kubebuilder:validation:Optional
	ManagedSubscriptions []*string `json:"managedSubscriptions,omitempty" tf:"managed_subscriptions,omitempty"`

	// Subscription ID of the Azure credential
	// +kubebuilder:validation:Required
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`

	// Tenant ID of the Azure credential
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// CredentialSpec defines the desired state of Credential
type CredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CredentialParameters `json:"forProvider"`
}

// CredentialStatus defines the observed state of Credential.
type CredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Credential is the Schema for the Credentials API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tmc}
type Credential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   CredentialSpec   `json:"spec"`
	Status CredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialList contains a list of Credentials
type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Credential `json:"items"`
}

// Repository type metadata.
var (
	Credential_Kind             = "Credential"
	Credential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Credential_Kind}.String()
	Credential_KindAPIVersion   = Credential_Kind + "." + CRDGroupVersion.String()
	Credential_GroupVersionKind = CRDGroupVersion.WithKind(Credential_Kind)
)

func init() {
	SchemeBuilder.Register(&Credential{}, &CredentialList{})
}
