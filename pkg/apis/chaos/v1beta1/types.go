package v1beta1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodChaosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []PodChaos `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              PodChaosSpec   `json:"spec"`
	Status            PodChaosStatus `json:"status,omitempty"`
}

type PodChaosSpec struct {
	Selector         map[string]string `json:"selector"`
	FrequencySeconds int               `json:"frequencySeconds"`
}
type PodChaosStatus struct {
	LastRun time.Time `json:"lastRun"`
}
