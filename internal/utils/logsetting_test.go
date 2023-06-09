/*
Copyright 2022. projectsveltos.io. All rights reserved.

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

package utils_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
	"github.com/gianlucam76/pod-log-level/internal/utils"
)

var _ = Describe("LogSettings", func() {
	It("GetLogSetting returns the default instance", func() {
		dc := &v1alpha1.LogSetting{
			ObjectMeta: metav1.ObjectMeta{
				Name: utils.DefaultInstanceName,
			},
		}

		initObjects := []client.Object{dc}
		scheme := runtime.NewScheme()
		Expect(utils.AddToScheme(scheme)).To(Succeed())
		c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(initObjects...).Build()

		k8sAccess := utils.GetK8sAccess(scheme, c)
		currentDC, err := k8sAccess.GetLogSetting(context.TODO())
		Expect(err).To(BeNil())
		Expect(currentDC).ToNot(BeNil())
		Expect(currentDC.Name).To(Equal(dc.Name))
	})

	It("UpdateLogSetting updates default LogSetting instance", func() {
		dc := &v1alpha1.LogSetting{
			ObjectMeta: metav1.ObjectMeta{
				Name: utils.DefaultInstanceName,
			},
		}

		scheme := runtime.NewScheme()
		Expect(utils.AddToScheme(scheme)).To(Succeed())
		c := fake.NewClientBuilder().WithScheme(scheme).Build()

		k8sAccess := utils.GetK8sAccess(scheme, c)
		Expect(k8sAccess.UpdateLogSetting(context.TODO(), dc)).To(Succeed())

		currentDC := &v1alpha1.LogSetting{}
		Expect(c.Get(context.TODO(), types.NamespacedName{Name: utils.DefaultInstanceName}, currentDC)).To(Succeed())
		currentDC.Spec.Configuration = []v1alpha1.ComponentConfiguration{
			{Component: v1alpha1.Component{Namespace: "dc", Identifier: "database"}, LogLevel: v1alpha1.LogLevelDebug},
		}

		Expect(k8sAccess.UpdateLogSetting(context.TODO(), currentDC)).To(Succeed())
		Expect(c.Get(context.TODO(), types.NamespacedName{Name: utils.DefaultInstanceName}, currentDC)).To(Succeed())
		Expect(len(currentDC.Spec.Configuration)).To(Equal(1))
	})
})
