/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MickaDeploymentSpec defines the desired state of MickaDeployment
type MickaDeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of MickaDeployment. Edit mickadeployment_types.go to remove/update
	Foo        string `json:"foo,omitempty"`
	NotChanged string `json:"notChanged,omitempty"`
}

// MickaDeploymentStatus defines the observed state of MickaDeployment
type MickaDeploymentStatus struct {
	BeforeFoo string `json:"beforeFoo,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MickaDeployment is the Schema for the mickadeployments API
type MickaDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MickaDeploymentSpec   `json:"spec,omitempty"`
	Status MickaDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MickaDeploymentList contains a list of MickaDeployment
type MickaDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MickaDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MickaDeployment{}, &MickaDeploymentList{})
}
