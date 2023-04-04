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

package lib_test

import (
	"flag"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
	"github.com/gianlucam76/pod-log-level/lib"
)

var _ = Describe("LogSetting", func() {
	component := v1alpha1.Component{Namespace: componentNamespace, Identifier: componentIdentifier}

	It("change klog level appropriately", func() {
		conf := &v1alpha1.LogSetting{
			ObjectMeta: metav1.ObjectMeta{
				Name: "default",
			},
			Spec: v1alpha1.LogSettingSpec{
				Configuration: []v1alpha1.ComponentConfiguration{
					{Component: component, LogLevel: v1alpha1.LogLevelDebug},
				},
			},
		}

		lib.UpdateLogLevel(conf)
		f := flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(lib.LogDebug)))

		conf.Spec.Configuration = []v1alpha1.ComponentConfiguration{
			{Component: component, LogLevel: v1alpha1.LogLevelInfo},
		}

		lib.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(lib.LogInfo)))

		conf.Spec.Configuration = []v1alpha1.ComponentConfiguration{
			{Component: component, LogLevel: v1alpha1.LogLevelVerbose},
		}

		lib.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(lib.LogVerbose)))

		newDebugValue := 8
		instance.SetDebugValue(newDebugValue)
		conf.Spec.Configuration = []v1alpha1.ComponentConfiguration{
			{Component: component, LogLevel: v1alpha1.LogLevelDebug},
		}

		lib.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(newDebugValue)))

		newInfoValue := 5
		instance.SetInfoValue(newInfoValue)
		conf.Spec.Configuration = []v1alpha1.ComponentConfiguration{
			{Component: component, LogLevel: v1alpha1.LogLevelInfo},
		}

		lib.UpdateLogLevel(conf)
		f = flag.Lookup("v")
		Expect(f).ToNot(BeNil())
		Expect(f.Value.String()).To(Equal(strconv.Itoa(newInfoValue)))
	})
})
