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

type GeneratedResourcesObservation struct {
	ClusterRoleName *string `json:"clusterRoleName,omitempty" tf:"cluster_role_name,omitempty"`

	RoleBindingName *string `json:"roleBindingName,omitempty" tf:"role_binding_name,omitempty"`

	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`
}

type GeneratedResourcesParameters struct {
}

type PackageRefObservation struct {

	// Name of the Package Metadata.
	PackageMetadataName *string `json:"packageMetadataName,omitempty" tf:"package_metadata_name,omitempty"`

	// Version Selection of the Package.
	VersionSelection []VersionSelectionObservation `json:"versionSelection,omitempty" tf:"version_selection,omitempty"`
}

type PackageRefParameters struct {

	// Name of the Package Metadata.
	// +kubebuilder:validation:Required
	PackageMetadataName *string `json:"packageMetadataName" tf:"package_metadata_name,omitempty"`

	// Version Selection of the Package.
	// +kubebuilder:validation:Required
	VersionSelection []VersionSelectionParameters `json:"versionSelection" tf:"version_selection,omitempty"`
}

type Package_InstallMetaObservation struct {

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

type Package_InstallMetaParameters struct {

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

type Package_InstallObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata for the resource
	Meta []Package_InstallMetaObservation `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the package install resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of Namespace where package install will be created.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Scope for the package install, having one of the valid scopes: cluster.
	Scope []Package_InstallScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// spec for package install.
	Spec []Package_InstallSpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// status for package install.
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type Package_InstallParameters struct {

	// Metadata for the resource
	// +kubebuilder:validation:Optional
	Meta []Package_InstallMetaParameters `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the package install resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of Namespace where package install will be created.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Scope for the package install, having one of the valid scopes: cluster.
	// +kubebuilder:validation:Optional
	Scope []Package_InstallScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// spec for package install.
	// +kubebuilder:validation:Optional
	Spec []Package_InstallSpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type Package_InstallScopeClusterObservation struct {

	// Name of the management cluster
	ManagementClusterName *string `json:"managementClusterName,omitempty" tf:"management_cluster_name,omitempty"`

	// Name of this cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Provisioner of the cluster
	ProvisionerName *string `json:"provisionerName,omitempty" tf:"provisioner_name,omitempty"`
}

type Package_InstallScopeClusterParameters struct {

	// Name of the management cluster
	// +kubebuilder:validation:Optional
	ManagementClusterName *string `json:"managementClusterName,omitempty" tf:"management_cluster_name,omitempty"`

	// Name of this cluster
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Provisioner of the cluster
	// +kubebuilder:validation:Optional
	ProvisionerName *string `json:"provisionerName,omitempty" tf:"provisioner_name,omitempty"`
}

type Package_InstallScopeObservation struct {

	// The schema for cluster full name
	Cluster []Package_InstallScopeClusterObservation `json:"cluster,omitempty" tf:"cluster,omitempty"`
}

type Package_InstallScopeParameters struct {

	// The schema for cluster full name
	// +kubebuilder:validation:Optional
	Cluster []Package_InstallScopeClusterParameters `json:"cluster,omitempty" tf:"cluster,omitempty"`
}

type Package_InstallSpecObservation struct {

	// Reference to the Package which will be installed.
	PackageRef []PackageRefObservation `json:"packageRef,omitempty" tf:"package_ref,omitempty"`

	// Role binding scope for service account which will be used by Package Install.
	RoleBindingScope *string `json:"roleBindingScope,omitempty" tf:"role_binding_scope,omitempty"`
}

type Package_InstallSpecParameters struct {

	// Inline values to configure the Package Install.
	// +kubebuilder:validation:Optional
	InlineValuesSecretRef *v1.SecretReference `json:"inlineValuesSecretRef,omitempty" tf:"-"`

	// Reference to the Package which will be installed.
	// +kubebuilder:validation:Required
	PackageRef []PackageRefParameters `json:"packageRef" tf:"package_ref,omitempty"`
}

type StatusObservation struct {
	GeneratedResources []GeneratedResourcesObservation `json:"generatedResources,omitempty" tf:"generated_resources,omitempty"`

	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	PackageInstallPhase *string `json:"packageInstallPhase,omitempty" tf:"package_install_phase,omitempty"`

	ReferredBy []*string `json:"referredBy,omitempty" tf:"referred_by,omitempty"`

	ResolvedVersion *string `json:"resolvedVersion,omitempty" tf:"resolved_version,omitempty"`
}

type StatusParameters struct {
}

type VersionSelectionObservation struct {

	// Constraints to select Package. Example: constraints: 'v1.2.3', constraints: '<v1.4.0' etc.
	Constraints *string `json:"constraints,omitempty" tf:"constraints,omitempty"`
}

type VersionSelectionParameters struct {

	// Constraints to select Package. Example: constraints: 'v1.2.3', constraints: '<v1.4.0' etc.
	// +kubebuilder:validation:Required
	Constraints *string `json:"constraints" tf:"constraints,omitempty"`
}

// Package_InstallSpec defines the desired state of Package_Install
type Package_InstallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Package_InstallParameters `json:"forProvider"`
}

// Package_InstallStatus defines the observed state of Package_Install.
type Package_InstallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Package_InstallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Package_Install is the Schema for the Package_Installs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tmc}
type Package_Install struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.namespace)",message="namespace is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scope)",message="scope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.spec)",message="spec is a required parameter"
	Spec   Package_InstallSpec   `json:"spec"`
	Status Package_InstallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Package_InstallList contains a list of Package_Installs
type Package_InstallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Package_Install `json:"items"`
}

// Repository type metadata.
var (
	Package_Install_Kind             = "Package_Install"
	Package_Install_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Package_Install_Kind}.String()
	Package_Install_KindAPIVersion   = Package_Install_Kind + "." + CRDGroupVersion.String()
	Package_Install_GroupVersionKind = CRDGroupVersion.WithKind(Package_Install_Kind)
)

func init() {
	SchemeBuilder.Register(&Package_Install{}, &Package_InstallList{})
}
