package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FooSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

type FooStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// TestType is a top-level type. A client is created for it.
type Foo struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec FooSpec `json:"spec"`
	// +optional
	Status FooStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestTypeList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Foo `json:"items"`
}
