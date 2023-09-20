/*
Copyright 2023 zhaoxin.

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

// PvcExportSpec defines the desired state of PvcExport
type PvcExportSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PvcExport. Edit pvcexport_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// PvcExportStatus defines the observed state of PvcExport
type PvcExportStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PvcExport is the Schema for the pvcexports API
type PvcExport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PvcExportSpec   `json:"spec,omitempty"`
	Status PvcExportStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PvcExportList contains a list of PvcExport
type PvcExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PvcExport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PvcExport{}, &PvcExportList{})
}
