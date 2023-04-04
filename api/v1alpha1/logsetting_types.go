/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum:=LogLevelNotSet;LogLevelInfo;LogLevelDebug;LogLevelVerbose
type LogLevel string

const (
	// LogLevelNotSet indicates log severity is not set. Default configuration will apply.
	LogLevelNotSet = LogLevel("LogLevelNotSet")

	// LogLevelInfo indicates log severity info (default to V(0)) is set
	LogLevelInfo = LogLevel("LogLevelInfo")

	// LogLevelDebug indicates log severity debug (default to V(5)) is set
	LogLevelDebug = LogLevel("LogLevelDebug")

	// LogLevelVerbose indicates log severity debug (default to V(10)) is set
	LogLevelVerbose = LogLevel("LogLevelVerbose")
)

// Component identifies the entity that has registered to have
// log level managed via LogSetting
type Component struct {
	// Namespace is resource namespace
	Namespace string `json:"namespace"`

	// Identifier is an ID that uniquely in a given namespace, identify
	// a resource
	Identifier string `json:"identifier"`
}

// ComponentConfiguration is the debugging configuration to be applied to a Sveltos component.
type ComponentConfiguration struct {
	// Component indicates which component the configuration applies to.
	Component Component `json:"component"`

	// LogLevel is the log severity above which logs are sent to the stdout. [Default: Info]
	LogLevel LogLevel `json:"logLevel,omitempty"`
}

// LogSettingSpec defines the desired state of LogSetting
type LogSettingSpec struct {
	// Configuration contains log level configuration as granular as per component.
	// +listType=atomic
	// +optional
	Configuration []ComponentConfiguration `json:"configuration,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:path=logsettings,scope=Cluster

// LogSetting is the Schema for the logsettings API
type LogSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec LogSettingSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// LogSettingList contains a list of LogSetting
type LogSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogSetting `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogSetting{}, &LogSettingList{})
}
