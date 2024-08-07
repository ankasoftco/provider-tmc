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

type IamPolicyMetaObservation struct {

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

type IamPolicyMetaParameters struct {

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

type IamPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata for the resource
	Meta []IamPolicyMetaObservation `json:"meta,omitempty" tf:"meta,omitempty"`

	// List of role bindings associated with the policy
	RoleBindings []RoleBindingsObservation `json:"roleBindings,omitempty" tf:"role_bindings,omitempty"`

	// Scope of the resource on which the rolebinding has to be added, having one of the valid scopes: organization, cluster_group, cluster, workspace or namespace.
	Scope []IamPolicyScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`
}

type IamPolicyParameters struct {

	// Metadata for the resource
	// +kubebuilder:validation:Optional
	Meta []IamPolicyMetaParameters `json:"meta,omitempty" tf:"meta,omitempty"`

	// List of role bindings associated with the policy
	// +kubebuilder:validation:Optional
	RoleBindings []RoleBindingsParameters `json:"roleBindings,omitempty" tf:"role_bindings,omitempty"`

	// Scope of the resource on which the rolebinding has to be added, having one of the valid scopes: organization, cluster_group, cluster, workspace or namespace.
	// +kubebuilder:validation:Optional
	Scope []IamPolicyScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`
}

type IamPolicyScopeClusterGroupObservation struct {

	// Name of the cluster group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IamPolicyScopeClusterGroupParameters struct {

	// Name of the cluster group
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type IamPolicyScopeClusterObservation struct {

	// Name of the management cluster
	ManagementClusterName *string `json:"managementClusterName,omitempty" tf:"management_cluster_name,omitempty"`

	// Name of this cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Provisioner of the cluster
	ProvisionerName *string `json:"provisionerName,omitempty" tf:"provisioner_name,omitempty"`
}

type IamPolicyScopeClusterParameters struct {

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

type IamPolicyScopeObservation struct {

	// The schema for cluster full name
	Cluster []IamPolicyScopeClusterObservation `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The schema for cluster group full name
	ClusterGroup []IamPolicyScopeClusterGroupObservation `json:"clusterGroup,omitempty" tf:"cluster_group,omitempty"`

	// The schema for namespace iam policy full name
	Namespace []NamespaceObservation `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The schema for organization iam policy full name
	Organization []ScopeOrganizationObservation `json:"organization,omitempty" tf:"organization,omitempty"`

	// The schema for workspace iam policy full name
	Workspace []IamPolicyScopeWorkspaceObservation `json:"workspace,omitempty" tf:"workspace,omitempty"`
}

type IamPolicyScopeParameters struct {

	// The schema for cluster full name
	// +kubebuilder:validation:Optional
	Cluster []IamPolicyScopeClusterParameters `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The schema for cluster group full name
	// +kubebuilder:validation:Optional
	ClusterGroup []IamPolicyScopeClusterGroupParameters `json:"clusterGroup,omitempty" tf:"cluster_group,omitempty"`

	// The schema for namespace iam policy full name
	// +kubebuilder:validation:Optional
	Namespace []NamespaceParameters `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The schema for organization iam policy full name
	// +kubebuilder:validation:Optional
	Organization []ScopeOrganizationParameters `json:"organization,omitempty" tf:"organization,omitempty"`

	// The schema for workspace iam policy full name
	// +kubebuilder:validation:Optional
	Workspace []IamPolicyScopeWorkspaceParameters `json:"workspace,omitempty" tf:"workspace,omitempty"`
}

type IamPolicyScopeWorkspaceObservation struct {

	// Name of the workspace
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IamPolicyScopeWorkspaceParameters struct {

	// Name of the workspace
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type NamespaceObservation struct {

	// Name of Cluster
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Name of ManagementCluster
	ManagementClusterName *string `json:"managementClusterName,omitempty" tf:"management_cluster_name,omitempty"`

	// Name of the Namespace
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of Provisioner
	ProvisionerName *string `json:"provisionerName,omitempty" tf:"provisioner_name,omitempty"`
}

type NamespaceParameters struct {

	// Name of Cluster
	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// Name of ManagementCluster
	// +kubebuilder:validation:Optional
	ManagementClusterName *string `json:"managementClusterName,omitempty" tf:"management_cluster_name,omitempty"`

	// Name of the Namespace
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Name of Provisioner
	// +kubebuilder:validation:Optional
	ProvisionerName *string `json:"provisionerName,omitempty" tf:"provisioner_name,omitempty"`
}

type RoleBindingsObservation struct {

	// Role for this rolebinding: max length for a role is 126 characters.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Subject for this rolebinding.
	Subjects []SubjectsObservation `json:"subjects,omitempty" tf:"subjects,omitempty"`
}

type RoleBindingsParameters struct {

	// Role for this rolebinding: max length for a role is 126 characters.
	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// Subject for this rolebinding.
	// +kubebuilder:validation:Required
	Subjects []SubjectsParameters `json:"subjects" tf:"subjects,omitempty"`
}

type ScopeOrganizationObservation struct {

	// ID of the Organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type ScopeOrganizationParameters struct {

	// ID of the Organization
	// +kubebuilder:validation:Required
	OrgID *string `json:"orgId" tf:"org_id,omitempty"`
}

type SubjectsObservation struct {

	// Subject type, having one of the subject types: USER or GROUP or K8S_SERVICEACCOUNT
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Subject name: allow max characters for email - 320 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SubjectsParameters struct {

	// Subject type, having one of the subject types: USER or GROUP or K8S_SERVICEACCOUNT
	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// Subject name: allow max characters for email - 320 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// IamPolicySpec defines the desired state of IamPolicy
type IamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IamPolicyParameters `json:"forProvider"`
}

// IamPolicyStatus defines the observed state of IamPolicy.
type IamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamPolicy is the Schema for the IamPolicys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tmc}
type IamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.roleBindings)",message="roleBindings is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scope)",message="scope is a required parameter"
	Spec   IamPolicySpec   `json:"spec"`
	Status IamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamPolicyList contains a list of IamPolicys
type IamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamPolicy `json:"items"`
}

// Repository type metadata.
var (
	IamPolicy_Kind             = "IamPolicy"
	IamPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IamPolicy_Kind}.String()
	IamPolicy_KindAPIVersion   = IamPolicy_Kind + "." + CRDGroupVersion.String()
	IamPolicy_GroupVersionKind = CRDGroupVersion.WithKind(IamPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&IamPolicy{}, &IamPolicyList{})
}
