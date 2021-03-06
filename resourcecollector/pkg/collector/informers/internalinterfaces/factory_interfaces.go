/*
Copyright 2019 The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

package internalinterfaces

import (
	"k8s.io/kubernetes/resourcecollector/pkg/collector/cache/interface"
	"time"

	"k8s.io/kubernetes/resourcecollector/pkg/collector/client"
)

// NewInformerFunc provide the function template that informer could set as param
type NewInformerFunc func(client.Interface, time.Duration, string, string) cacheinterface.SharedInformer

// SharedInformerFactory a small interface to allow for adding an informer without an import cycle
type SharedInformerFactory interface {
	Start(stopCh <-chan struct{})
	InformerFor(name string, key string, newFunc NewInformerFunc) cacheinterface.SharedInformer
}
