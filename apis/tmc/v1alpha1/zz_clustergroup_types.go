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

type ClusterGroupMetaObservation struct {

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

type ClusterGroupMetaParameters struct {

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

type ClusterGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata for the resource
	Meta []ClusterGroupMetaObservation `json:"meta,omitempty" tf:"meta,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ClusterGroupParameters struct {

	// Metadata for the resource
	// +kubebuilder:validation:Optional
	Meta []ClusterGroupMetaParameters `json:"meta,omitempty" tf:"meta,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ClusterGroupSpec defines the desired state of ClusterGroup
type ClusterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterGroupParameters `json:"forProvider"`
}

// ClusterGroupStatus defines the observed state of ClusterGroup.
type ClusterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterGroup is the Schema for the ClusterGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tmc}
type ClusterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ClusterGroupSpec   `json:"spec"`
	Status ClusterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterGroupList contains a list of ClusterGroups
type ClusterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterGroup `json:"items"`
}

// Repository type metadata.
var (
	ClusterGroup_Kind             = "ClusterGroup"
	ClusterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterGroup_Kind}.String()
	ClusterGroup_KindAPIVersion   = ClusterGroup_Kind + "." + CRDGroupVersion.String()
	ClusterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ClusterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterGroup{}, &ClusterGroupList{})
}