package cache

import cacheinterface "k8s.io/kubernetes/resourcecollector/pkg/collector/cache/interface"

// Lister is defined to invoke ListFunc
type Lister struct {
	ListFunc cacheinterface.ListFunc
}

// List a set of apiServer resources
func (lw *Lister) List(options interface{}) ([]interface{}, error) {
	return lw.ListFunc(options)
}
