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

package loglevel

import (
	"context"
	"sort"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
	"github.com/gianlucam76/pod-log-level/internal/utils"
)

type componentConfiguration struct {
	component   v1alpha1.Component
	logSeverity v1alpha1.LogLevel
}

// byComponent sorts componentConfiguration by name.
type byComponent []*componentConfiguration

func (c byComponent) Len() int      { return len(c) }
func (c byComponent) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c byComponent) Less(i, j int) bool {
	if c[i].component.Namespace == c[j].component.Namespace {
		return c[i].component.Identifier < c[j].component.Identifier
	}
	return c[i].component.Namespace < c[j].component.Namespace
}

func collectLogLevelConfiguration(ctx context.Context) ([]*componentConfiguration, error) {
	instance := utils.GetAccessInstance()

	dc, err := instance.GetLogSetting(ctx)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return make([]*componentConfiguration, 0), nil
		}
		return nil, err
	}

	configurationSettings := make([]*componentConfiguration, len(dc.Spec.Configuration))

	for i, c := range dc.Spec.Configuration {
		configurationSettings[i] = &componentConfiguration{
			component:   c.Component,
			logSeverity: c.LogLevel,
		}
	}

	// Sort this by component name first. Component/node is higher priority than Component
	sort.Sort(byComponent(configurationSettings))

	return configurationSettings, nil
}

func updateLogLevelConfiguration(
	ctx context.Context,
	spec []v1alpha1.ComponentConfiguration,
) error {

	instance := utils.GetAccessInstance()

	dc, err := instance.GetLogSetting(ctx)
	if err != nil {
		if apierrors.IsNotFound(err) {
			dc = &v1alpha1.LogSetting{
				ObjectMeta: metav1.ObjectMeta{
					Name: "default",
				},
			}
		} else {
			return err
		}
	}

	dc.Spec = v1alpha1.LogSettingSpec{
		Configuration: spec,
	}

	return instance.UpdateLogSetting(ctx, dc)
}
