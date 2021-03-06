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

package sites

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/cache/interface"
	"path/filepath"
	"time"

	"k8s.io/kubernetes/resourcecollector/pkg/collector/cache"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/client"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/internalinterfaces"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/typed"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/types"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/utils"
)

// InformerFlavor provides access to a shared informer and lister for Sites.
type InformerSite interface {
	Informer() cacheinterface.SharedInformer
}

type informerSite struct {
	factory internalinterfaces.SharedInformerFactory
	name    string
	key     string
	period  time.Duration
}

// New initial the informerSite
func New(f internalinterfaces.SharedInformerFactory, name string, key string, period time.Duration) InformerSite {
	return &informerSite{factory: f, name: name, key: key, period: period}
}

// NewSiteInformer constructs a new informer for flavor.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSiteInformer(client client.Interface, resyncPeriod time.Duration, name string, key string) cacheinterface.SharedInformer {
	return cacheinterface.NewSharedInformer(
		&cache.Lister{ListFunc: func(options interface{}) ([]interface{}, error) {
			// read site init data
			confLocation := utils.GetConfigDirectory()
			confFilePath := filepath.Join(confLocation, "sites.json")
			file, err := ioutil.ReadFile(confFilePath)
			if err != nil {
				return nil, fmt.Errorf("read site data error:%v", err)
			}

			var sites typed.Sites
			err = json.Unmarshal(file, &sites)
			if err != nil {
				return nil, fmt.Errorf("unmarshal site data error:%v", err)
			}

			var interfaceSlice []interface{}
			for _, site := range sites.Sites {
				interfaceSlice = append(interfaceSlice, site)
			}

			return interfaceSlice, nil
		}}, resyncPeriod, name, key, types.ListSiteOpts{})
}

func (f *informerSite) defaultInformer(client client.Interface, resyncPeriod time.Duration, name string, key string) cacheinterface.SharedInformer {
	if f.period > 0 {
		resyncPeriod = f.period
	}
	return NewSiteInformer(client, resyncPeriod, name, key)
}

func (f *informerSite) Informer() cacheinterface.SharedInformer {
	return f.factory.InformerFor(f.name, f.key, f.defaultInformer)
}
