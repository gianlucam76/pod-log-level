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

package utils

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
)

// k8sAccess is used to access resources in the management cluster.
type k8sAccess struct {
	client     client.Client
	restConfig *rest.Config
	clientset  *kubernetes.Clientset
	scheme     *runtime.Scheme
}

var (
	accessInstance *k8sAccess
)

// GetAccessInstance return k8sAccess instance used to access resources in the
// management cluster.
func GetAccessInstance() *k8sAccess {
	return accessInstance
}

// Following method could have been called directly by GetAccessInstance is accessInstance was
// nil. Doing this way though it makes it possible to run uts against each of the implemented
// command.

// InitalizeManagementClusterAcces initializes k8sAccess singleton
func InitalizeManagementClusterAcces(scheme *runtime.Scheme, restConfig *rest.Config,
	cs *kubernetes.Clientset, c client.Client) {

	accessInstance = &k8sAccess{
		scheme:     scheme,
		client:     c,
		clientset:  cs,
		restConfig: restConfig,
	}
}

func GetScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	if err := addToScheme(scheme); err != nil {
		return nil, err
	}
	return scheme, nil
}

func addToScheme(scheme *runtime.Scheme) error {
	if err := v1alpha1.AddToScheme(scheme); err != nil {
		return err
	}
	return nil
}

// GetScheme returns scheme
func (a *k8sAccess) GetScheme() *runtime.Scheme {
	return a.scheme
}
