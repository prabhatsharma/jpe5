/*
Copyright 2021 Prabhat Sharma.

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

// AdmissionPolicySpec defines the desired state of AdmissionPolicy
type AdmissionPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Rules is a collection of one or more independent rules against which resources will be validated against.
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name                    string `json:"name"`
	ResourceKind            string `json:"resourceKind"`
	ValidationFailureAction string `json:"validationFailureAction"`
	Rule                    string `json:"rule"`
	Description             string `json:"description"`
	Message                 string `json:"message"`
}

type RuleResponse struct {
	Allowed bool   `json:"allowed"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// AdmissionPolicyStatus defines the observed state of AdmissionPolicy
type AdmissionPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AdmissionPolicy is the Schema for the admissionpolicies API
type AdmissionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdmissionPolicySpec   `json:"spec,omitempty"`
	Status AdmissionPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AdmissionPolicyList contains a list of AdmissionPolicy
type AdmissionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmissionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmissionPolicy{}, &AdmissionPolicyList{})
}
