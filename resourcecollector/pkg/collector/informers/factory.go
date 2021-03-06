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

package informers

import (
	"k8s.io/kubernetes/resourcecollector/pkg/collector/cache/interface"
	"sync"
	"time"

	"k8s.io/kubernetes/resourcecollector/pkg/collector/client"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/eipavailability"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/flavor"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/internalinterfaces"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/nodes"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/sites"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/volumepool"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers/volumetype"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/typed"
)

var (
	// InformerFac is the global param and produce the informers
	InformerFac SharedInformerFactory
)

const (
	// FLAVOR is the symbol of informer flavor
	FLAVOR = "flavor"

	// NODES is the symbol of informer nodes
	NODES = "nodes"

	// VOLUMETYPE is the symbol of informer volumeType
	VOLUMETYPE = "volumeType"

	// Sites is the symbol of informer site
	SITES = "sites"

	// EIPPOOLS eip pool
	EIPPOOLS = "eipools"

	// VOLUMEPOOLS volume pool
	VOLUMEPOOLS = "volumepools"
)

// SharedInformerFactory is used as a factory to produce informers
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	GetInformer(name string) cacheinterface.SharedInformer

	GetFlavor(flavorID string, region string) (typed.Flavor, bool)
	Flavor(name, key string, period time.Duration) flavor.InformerFlavor
	Sites(name, key string, period time.Duration) sites.InformerSite
	Nodes(name, key string, period time.Duration) nodes.InformerNodes
	VolumeType(name, key string, period time.Duration) volumetype.InformerVolumeType
	EipPools(name, key string, period time.Duration) eipavailability.InformerEipAvailability
	VolumePools(name, key string, period time.Duration) volumepool.InformerVolumePool
}

type sharedInformerFactory struct {
	client        client.Interface
	lock          sync.Mutex
	defaultResync time.Duration

	informers map[string]cacheinterface.SharedInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[string]bool
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory
func NewSharedInformerFactory(client client.Interface, defaultResync time.Duration) SharedInformerFactory {
	return NewFilteredSharedInformerFactory(client, defaultResync)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this SharedInformerFactory will be subject to the same filters
// as specified here.
func NewFilteredSharedInformerFactory(client client.Interface, defaultResync time.Duration) SharedInformerFactory {
	return &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[string]cacheinterface.SharedInformer),
		startedInformers: make(map[string]bool),
	}
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) GetInformer(name string) cacheinterface.SharedInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informer, exists := f.informers[name]
	if !exists {
		return cacheinterface.NewEmptyInformer()
	}

	if !informer.HasSynced() {
		informer.SyncOnce()
	}
	return informer

}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(name string, key string, newFunc internalinterfaces.NewInformerFunc) cacheinterface.SharedInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informer, exists := f.informers[name]
	if exists {
		return informer
	}
	informer = newFunc(f.client, f.defaultResync, name, key)
	f.informers[name] = informer

	return informer
}

//GetFlavor get flavor
func (f *sharedInformerFactory) GetFlavor(flavorID string, region string) (typed.Flavor, bool) {
	if region != "" {
		flvInter, exist := f.GetInformer(FLAVOR).GetStore().Get(region + "|" + flavorID)
		var ret = typed.Flavor{}
		if exist {
			ret = flvInter.(typed.RegionFlavor).Flavor
		}
		return ret, exist
	} else {
		flvInters := f.GetInformer(FLAVOR).GetStore().List()
		for _, flvInter := range flvInters {
			ret := flvInter.(typed.RegionFlavor).Flavor
			if ret.ID == flavorID {
				return ret, true
			}
		}
	}

	return typed.Flavor{}, false
}

//Flavor new Flavor informer
func (f *sharedInformerFactory) Flavor(name, key string, period time.Duration) flavor.InformerFlavor {
	return flavor.New(f, name, key, period)
}

//Nodes new nodes informer
func (f *sharedInformerFactory) Nodes(name, key string, period time.Duration) nodes.InformerNodes {
	return nodes.New(f, name, key, period)
}

//VolumeType new volume type informer
func (f *sharedInformerFactory) VolumeType(name, key string, period time.Duration) volumetype.InformerVolumeType {
	return volumetype.New(f, name, key, period)
}

//Sites new site informer
func (f *sharedInformerFactory) Sites(name, key string, period time.Duration) sites.InformerSite {
	return sites.New(f, name, key, period)
}

//EipPools new eip pool informer
func (f *sharedInformerFactory) EipPools(name, key string, period time.Duration) eipavailability.InformerEipAvailability {
	return eipavailability.New(f, name, key, period)
}

//VolumePools new volume pool informer
func (f *sharedInformerFactory) VolumePools(name, key string, period time.Duration) volumepool.InformerVolumePool {
	return volumepool.New(f, name, key, period)
}
