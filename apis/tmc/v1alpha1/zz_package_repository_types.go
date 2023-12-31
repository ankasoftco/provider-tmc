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

type ImgpkgBundleObservation struct {

	// image url string.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`
}

type ImgpkgBundleParameters struct {

	// image url string.
	// +kubebuilder:validation:Required
	Image *string `json:"image" tf:"image,omitempty"`
}

type Package_RepositoryMetaObservation struct {

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

type Package_RepositoryMetaParameters struct {

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

type Package_RepositoryObservation struct {

	// If true, Package Repository is disabled for cluster.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata for the resource
	Meta []Package_RepositoryMetaObservation `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the package repository resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of Namespace where package repository will be created.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Scope for the package repository, having one of the valid scopes: cluster.
	Scope []Package_RepositoryScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// spec for package repository.
	Spec []Package_RepositorySpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// status for package repository.
	State []StateObservation `json:"state,omitempty" tf:"state,omitempty"`
}

type Package_RepositoryParameters struct {

	// If true, Package Repository is disabled for cluster.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Metadata for the resource
	// +kubebuilder:validation:Optional
	Meta []Package_RepositoryMetaParameters `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the package repository resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Scope for the package repository, having one of the valid scopes: cluster.
	// +kubebuilder:validation:Optional
	Scope []Package_RepositoryScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// spec for package repository.
	// +kubebuilder:validation:Optional
	Spec []Package_RepositorySpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type Package_RepositoryScopeClusterObservation struct {

	// Name of the management cluster
	ManagementClusterName *string `json:"managementClusterName,omitempty" tf:"management_cluster_name,omitempty"`

	// Name of this cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Provisioner of the cluster
	ProvisionerName *string `json:"provisionerName,omitempty" tf:"provisioner_name,omitempty"`
}

type Package_RepositoryScopeClusterParameters struct {

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

type Package_RepositoryScopeObservation struct {

	// The schema for cluster full name
	Cluster []Package_RepositoryScopeClusterObservation `json:"cluster,omitempty" tf:"cluster,omitempty"`
}

type Package_RepositoryScopeParameters struct {

	// The schema for cluster full name
	// +kubebuilder:validation:Optional
	Cluster []Package_RepositoryScopeClusterParameters `json:"cluster,omitempty" tf:"cluster,omitempty"`
}

type Package_RepositorySpecObservation struct {

	// Docker image url; unqualified, tagged, or digest references supported.
	ImgpkgBundle []ImgpkgBundleObservation `json:"imgpkgBundle,omitempty" tf:"imgpkg_bundle,omitempty"`
}

type Package_RepositorySpecParameters struct {

	// Docker image url; unqualified, tagged, or digest references supported.
	// +kubebuilder:validation:Optional
	ImgpkgBundle []ImgpkgBundleParameters `json:"imgpkgBundle,omitempty" tf:"imgpkg_bundle,omitempty"`
}

type StateObservation struct {
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	PackageRepositoryPhase *string `json:"packageRepositoryPhase,omitempty" tf:"package_repository_phase,omitempty"`

	Subscribed *bool `json:"subscribed,omitempty" tf:"subscribed,omitempty"`
}

type StateParameters struct {
}

// Package_RepositorySpec defines the desired state of Package_Repository
type Package_RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Package_RepositoryParameters `json:"forProvider"`
}

// Package_RepositoryStatus defines the observed state of Package_Repository.
type Package_RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Package_RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Package_Repository is the Schema for the Package_Repositorys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tmc}
type Package_Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scope)",message="scope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.spec)",message="spec is a required parameter"
	Spec   Package_RepositorySpec   `json:"spec"`
	Status Package_RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Package_RepositoryList contains a list of Package_Repositorys
type Package_RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Package_Repository `json:"items"`
}

// Repository type metadata.
var (
	Package_Repository_Kind             = "Package_Repository"
	Package_Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Package_Repository_Kind}.String()
	Package_Repository_KindAPIVersion   = Package_Repository_Kind + "." + CRDGroupVersion.String()
	Package_Repository_GroupVersionKind = CRDGroupVersion.WithKind(Package_Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Package_Repository{}, &Package_RepositoryList{})
}
