package cache

import "sync"

type RegionAZCache struct {
	mutex       sync.Mutex
	RegionAZMap map[string]string
}

func NewRegionAZCache() *RegionAZCache {
	c := &RegionAZCache{}
	c.RegionAZMap = make(map[string]string)
	return c
}

func (c *RegionAZCache) AddRegionAZ(regionAZ, ip string) {
	c.mutex.Lock()
	c.RegionAZMap[regionAZ] = ip
	c.mutex.Unlock()
}

func (c *RegionAZCache) RemoveRegionAZ(regionAZ string) {
	c.mutex.Lock()
	delete(c.RegionAZMap, regionAZ)
	c.mutex.Unlock()
}
