/*
Copyright 2023

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

package loglevel_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
)

func TestShow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LogLevel Suite")
}

func getLogSetting() *v1alpha1.LogSetting {
	return &v1alpha1.LogSetting{
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
		},
	}
}
