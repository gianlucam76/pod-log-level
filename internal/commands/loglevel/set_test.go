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
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
	"github.com/gianlucam76/pod-log-level/internal/commands/loglevel"
	"github.com/gianlucam76/pod-log-level/internal/utils"
)

var _ = Describe("Set", func() {
	It("set updates default LogSetting instance", func() {
		component := v1alpha1.Component{Namespace: "foo", Identifier: "bar"}

		scheme, err := utils.GetScheme()
		Expect(err).To(BeNil())
		c := fake.NewClientBuilder().WithScheme(scheme).Build()

		utils.InitalizeManagementClusterAcces(scheme, nil, nil, c)
		Expect(loglevel.UpdateLogSetting(context.TODO(), v1alpha1.LogLevelDebug,
			component)).To(Succeed())

		k8sAccess := utils.GetAccessInstance()

		currentDC, err := k8sAccess.GetLogSetting(context.TODO())
		Expect(err).To(BeNil())
		Expect(currentDC).ToNot(BeNil())
		Expect(currentDC.Spec.Configuration).ToNot(BeNil())
		Expect(len(currentDC.Spec.Configuration)).To(Equal(1))
		Expect(currentDC.Spec.Configuration[0].Component).To(Equal(component))
		Expect(currentDC.Spec.Configuration[0].LogLevel).To(Equal(v1alpha1.LogLevelDebug))

		Expect(loglevel.UpdateLogSetting(context.TODO(), v1alpha1.LogLevelInfo,
			component)).To(Succeed())
		currentDC, err = k8sAccess.GetLogSetting(context.TODO())
		Expect(err).To(BeNil())
		Expect(currentDC).ToNot(BeNil())
		Expect(currentDC.Spec.Configuration).ToNot(BeNil())
		Expect(len(currentDC.Spec.Configuration)).To(Equal(1))
		Expect(currentDC.Spec.Configuration[0].Component).To(Equal(component))
		Expect(currentDC.Spec.Configuration[0].LogLevel).To(Equal(v1alpha1.LogLevelInfo))
	})
})
