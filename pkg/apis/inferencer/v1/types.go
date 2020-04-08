package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Inferencer struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec InferencerSpec `json:"spec"`
}

type InferencerSpec struct {
	 ModelName   string `json:"model_name"`
	 // init pod number
	 Number      int    `json:"number"`
	 // framework like tf-serving、caffe-serving etc...
	 Framework   string `json:"framework"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList is a list of Network resources
type InferencerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Inferencer `json:"items"`
}