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

package utils

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
)

const (
	defaultInstanceName = "default"
)

// GetLogSetting gets default LogSetting instance
func (a *k8sAccess) GetLogSetting(
	ctx context.Context,
) (*v1alpha1.LogSetting, error) {

	req := &v1alpha1.LogSetting{}

	reqName := client.ObjectKey{
		Name: defaultInstanceName,
	}

	if err := a.client.Get(ctx, reqName, req); err != nil {
		return nil, err
	}

	return req, nil
}

// UpdateLogSetting creates, if not existing already, default LogSetting. Otherwise
// updates it.
func (a *k8sAccess) UpdateLogSetting(
	ctx context.Context,
	dc *v1alpha1.LogSetting,
) error {

	reqName := client.ObjectKey{
		Name: defaultInstanceName,
	}

	tmp := &v1alpha1.LogSetting{}

	err := a.client.Get(ctx, reqName, tmp)
	if err != nil {
		if apierrors.IsNotFound(err) {
			err = a.client.Create(ctx, dc)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	err = a.client.Update(ctx, dc)
	if err != nil {
		return err
	}

	return nil
}
