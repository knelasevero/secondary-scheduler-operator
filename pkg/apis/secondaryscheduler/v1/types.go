package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecondaryScheduler is the Schema for the secondaryscheduler API
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
type SecondaryScheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +required
	Spec SecondarySchedulerSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status SecondarySchedulerStatus `json:"status"`
}

// SecondarySchedulerSpec defines the desired state of SecondaryScheduler
type SecondarySchedulerSpec struct {
	operatorv1.OperatorSpec `json:",inline"`

	// SchedulerConfig allows configuring the customized scheduler plugin configuration for the secondaryscheduler.
	SchedulerConfig string `json:"schedulerConfig"`

	// SchedulerImage sets the container image url to be pulled for the custom scheduler that's deployed
	SchedulerImage string `json:"schedulerImage"`
}

// SecondarySchedulerStatus defines the observed state of SecondaryScheduler
type SecondarySchedulerStatus struct {
	operatorv1.OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecondarySchedulerList contains a list of SecondaryScheduler
type SecondarySchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecondaryScheduler `json:"items"`
}
