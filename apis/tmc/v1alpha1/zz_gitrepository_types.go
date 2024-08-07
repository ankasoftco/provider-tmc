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

type GitRepositoryMetaObservation struct {

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

type GitRepositoryMetaParameters struct {

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

type GitRepositoryObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata for the resource
	Meta []GitRepositoryMetaObservation `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the Repository.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of Namespace.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Scope for the git repository, having one of the valid scopes: cluster, cluster_group.
	Scope []GitRepositoryScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// Spec for the Repository.
	Spec []GitRepositorySpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// Status for the Repository.
	Status map[string]*string `json:"status,omitempty" tf:"status,omitempty"`
}

type GitRepositoryParameters struct {

	// Metadata for the resource
	// +kubebuilder:validation:Optional
	Meta []GitRepositoryMetaParameters `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the Repository.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of Namespace.
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Scope for the git repository, having one of the valid scopes: cluster, cluster_group.
	// +kubebuilder:validation:Optional
	Scope []GitRepositoryScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// Spec for the Repository.
	// +kubebuilder:validation:Optional
	Spec []GitRepositorySpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type GitRepositoryScopeClusterGroupObservation struct {

	// Name of the cluster group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GitRepositoryScopeClusterGroupParameters struct {

	// Name of the cluster group
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type GitRepositoryScopeClusterObservation struct {

	// Name of the management cluster
	ManagementClusterName *string `json:"managementClusterName,omitempty" tf:"management_cluster_name,omitempty"`

	// Name of this cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Provisioner of the cluster
	ProvisionerName *string `json:"provisionerName,omitempty" tf:"provisioner_name,omitempty"`
}

type GitRepositoryScopeClusterParameters struct {

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

type GitRepositoryScopeObservation struct {

	// The schema for cluster full name
	Cluster []GitRepositoryScopeClusterObservation `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The schema for cluster group full name
	ClusterGroup []GitRepositoryScopeClusterGroupObservation `json:"clusterGroup,omitempty" tf:"cluster_group,omitempty"`
}

type GitRepositoryScopeParameters struct {

	// The schema for cluster full name
	// +kubebuilder:validation:Optional
	Cluster []GitRepositoryScopeClusterParameters `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The schema for cluster group full name
	// +kubebuilder:validation:Optional
	ClusterGroup []GitRepositoryScopeClusterGroupParameters `json:"clusterGroup,omitempty" tf:"cluster_group,omitempty"`
}

type GitRepositorySpecObservation struct {

	// GitImplementation specifies which client library implementation to use. go-git is the default git implementation.
	GitImplementation *string `json:"gitImplementation,omitempty" tf:"git_implementation,omitempty"`

	// Interval at which to check gitrepository for updates. This is the interval at which Tanzu Mission Control will attempt to reconcile changes in the repository to the cluster. A sync interval of 0 would result in no future syncs. If no value is entered, a default interval of 5 minutes will be applied as `5m`.
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// Reference specifies git reference to resolve.
	Ref []RefObservation `json:"ref,omitempty" tf:"ref,omitempty"`

	// Reference to the secret. Repository credential.
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`

	// URL of the git repository. Repository URL should begin with http, https, or ssh
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type GitRepositorySpecParameters struct {

	// GitImplementation specifies which client library implementation to use. go-git is the default git implementation.
	// +kubebuilder:validation:Optional
	GitImplementation *string `json:"gitImplementation,omitempty" tf:"git_implementation,omitempty"`

	// Interval at which to check gitrepository for updates. This is the interval at which Tanzu Mission Control will attempt to reconcile changes in the repository to the cluster. A sync interval of 0 would result in no future syncs. If no value is entered, a default interval of 5 minutes will be applied as `5m`.
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// Reference specifies git reference to resolve.
	// +kubebuilder:validation:Optional
	Ref []RefParameters `json:"ref,omitempty" tf:"ref,omitempty"`

	// Reference to the secret. Repository credential.
	// +kubebuilder:validation:Optional
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`

	// URL of the git repository. Repository URL should begin with http, https, or ssh
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type RefObservation struct {

	// Branch from git to checkout. When branch is given, then that branch from the git repository will be checked out. If the given branch doesn’t exist in the git repository, then adding the git repository will fail. If no branch is given, the `master` branch will be used.
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Commit SHA to checkout. Takes precedence over all other reference fields. When git_implementation is `GO_GIT`, this can be combined with branch to shallow clone branch in which the commit is expected to exist.
	Commit *string `json:"commit,omitempty" tf:"commit,omitempty"`

	// SemVer expression to checkout from git tags. Takes precedence over tag. When semver is given, then the latest tag matching that semver will be checked out from the git repository. If no tag in the git repository matches semver, then adding the git repository will fail. If semver is given, tag and branch will be ignored if they are populated.
	Semver *string `json:"semver,omitempty" tf:"semver,omitempty"`

	// Tag from git to checkout. Takes precedence over branch. When a tag is given, that tag from the git repository will be checked out. If the given tag doesn’t exist in the git repository, then adding the git repository will fail. If both tag and branch are given, tag overrides branch and the branch value will be ignored.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type RefParameters struct {

	// Branch from git to checkout. When branch is given, then that branch from the git repository will be checked out. If the given branch doesn’t exist in the git repository, then adding the git repository will fail. If no branch is given, the `master` branch will be used.
	// +kubebuilder:validation:Optional
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Commit SHA to checkout. Takes precedence over all other reference fields. When git_implementation is `GO_GIT`, this can be combined with branch to shallow clone branch in which the commit is expected to exist.
	// +kubebuilder:validation:Optional
	Commit *string `json:"commit,omitempty" tf:"commit,omitempty"`

	// SemVer expression to checkout from git tags. Takes precedence over tag. When semver is given, then the latest tag matching that semver will be checked out from the git repository. If no tag in the git repository matches semver, then adding the git repository will fail. If semver is given, tag and branch will be ignored if they are populated.
	// +kubebuilder:validation:Optional
	Semver *string `json:"semver,omitempty" tf:"semver,omitempty"`

	// Tag from git to checkout. Takes precedence over branch. When a tag is given, that tag from the git repository will be checked out. If the given tag doesn’t exist in the git repository, then adding the git repository will fail. If both tag and branch are given, tag overrides branch and the branch value will be ignored.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// GitRepositorySpec defines the desired state of GitRepository
type GitRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GitRepositoryParameters `json:"forProvider"`
}

// GitRepositoryStatus defines the observed state of GitRepository.
type GitRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GitRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GitRepository is the Schema for the GitRepositorys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tmc}
type GitRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.namespaceName)",message="namespaceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scope)",message="scope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.spec)",message="spec is a required parameter"
	Spec   GitRepositorySpec   `json:"spec"`
	Status GitRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GitRepositoryList contains a list of GitRepositorys
type GitRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitRepository `json:"items"`
}

// Repository type metadata.
var (
	GitRepository_Kind             = "GitRepository"
	GitRepository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GitRepository_Kind}.String()
	GitRepository_KindAPIVersion   = GitRepository_Kind + "." + CRDGroupVersion.String()
	GitRepository_GroupVersionKind = CRDGroupVersion.WithKind(GitRepository_Kind)
)

func init() {
	SchemeBuilder.Register(&GitRepository{}, &GitRepositoryList{})
}
