package collector

import (
	"k8s.io/kubernetes/resourcecollector/pkg/collector/cache"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/common/constants"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/common/logger"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/informers"
	cache2 "k8s.io/kubernetes/resourcecollector/pkg/collector/internal/cache"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/typed"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/types"
	"time"
)

var collector *Collector

// InitScheduler
func InitCollector(stopCh <-chan struct{}) error {
	var err error
	collector, err = NewCollector(stopCh)
	return err
}

// GetScheduler gets single scheduler instance. New Scheduler will only run once,
// if it runs failed, nil will be return.
func GetCollector() *Collector {
	if collector == nil {
		logger.Errorf("Collector need to be init correctly")
		return collector
	}

	return collector
}

type Collector struct {
	CollectorCache   cache2.Cache
	nodeInfoSnapshot *cache2.Snapshot
	RegionAZCache    *cache.RegionAZCache

	//InformerFactory informers.SharedInformerFactory
}

func NewCollector(stopCh <-chan struct{}) (*Collector, error) {
	c := &Collector{
		CollectorCache:   cache2.New(30*time.Second, stopCh),
		nodeInfoSnapshot: cache2.NewEmptySnapshot(),
		RegionAZCache:    cache.NewRegionAZCache(),
	}
	return c, nil
}

// snapshot snapshots scheduler cache and node infos for all fit and priority
// functions.
func (c *Collector) snapshot() error {
	// Used for all fit and priority funcs.
	return c.Cache().UpdateSnapshot(c.nodeInfoSnapshot)
}

func (c *Collector) GetSnapshot() (*cache2.Snapshot, error) {
	err := c.snapshot()
	if err != nil {
		return nil, err
	}
	return c.nodeInfoSnapshot, nil
}

// start resource cache informer and run
func (c *Collector) StartInformersAndRun(stopCh <-chan struct{}) {
	go func(stopCh2 <-chan struct{}) {
		interval := 10

		// init informer
		informers.InformerFac = informers.NewSharedInformerFactory(nil, 60*time.Second)

		// init volume type informer
		volumetypeInterval := interval
		informers.InformerFac.VolumeType(informers.VOLUMETYPE, "ID",
			time.Duration(volumetypeInterval)*time.Second).Informer()

		// init site informer
		siteInterval := interval
		informers.InformerFac.Sites(informers.SITES, "ID",
			time.Duration(siteInterval)*time.Second).Informer()

		// init flavor informer
		flavorInterval := 60
		informers.InformerFac.Flavor(informers.FLAVOR, "RegionFlavorID",
			time.Duration(flavorInterval)*time.Second).Informer()

		// init eip pool informer
		eipPoolInterval := 60
		eipPoolInformer := informers.InformerFac.EipPools(informers.EIPPOOLS, "Region",
			time.Duration(eipPoolInterval)*time.Second).Informer()
		eipPoolInformer.AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				ListFunc: updateEipPools,
			})

		// init volume pool informer
		volumePoolInterval := 60
		volumePoolInformer := informers.InformerFac.VolumePools(informers.VOLUMEPOOLS, "Region",
			time.Duration(volumePoolInterval)*time.Second).Informer()
		volumePoolInformer.AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				ListFunc: updateVolumePools,
			})

		// init node informer
		nodeInterval := 60
		nodeInformer := informers.InformerFac.Nodes(informers.NODES, "EdgeSiteID",
			time.Duration(nodeInterval)*time.Second).Informer()
		nodeInformer.AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				ListFunc: addSiteNodesToCache,
			})

		informers.InformerFac.Start(stopCh2)

	}(stopCh)
}

// Cache returns the cache in scheduler for test to check the data in scheduler.
func (c *Collector) Cache() cache2.Cache {
	return c.CollectorCache
}

// update EipPools with sched cache
func updateEipPools(obj []interface{}) {
	if obj == nil {
		return
	}

	for _, eipPoolObj := range obj {
		eipPool, ok := eipPoolObj.(typed.EipPool)
		if !ok {
			logger.Warnf("convert interface to (typed.EipPool) failed.")
			continue
		}

		err := collector.Cache().UpdateEipPool(&eipPool)
		if err != nil {
			logger.Infof("UpdateEipPool failed! err: %s", err)
		}
	}
}

// update VolumePools with sched cache
func updateVolumePools(obj []interface{}) {
	if obj == nil {
		return
	}

	for _, volumePoolObj := range obj {
		volumePool, ok := volumePoolObj.(typed.RegionVolumePool)
		if !ok {
			logger.Warnf("convert interface to (typed.VolumePools) failed.")
			continue
		}

		err := collector.Cache().UpdateVolumePool(&volumePool)
		if err != nil {
			logger.Infof("updateVolumePools failed! err: %s", err)
		}
	}
}

// add site node to cache
func addSiteNodesToCache(obj []interface{}) {
	if obj == nil {
		return
	}

	sites := informers.InformerFac.GetInformer(informers.SITES).GetStore().List()

	for _, sn := range obj {
		siteNode, ok := sn.(typed.SiteNode)
		if !ok {
			logger.Warnf("convert interface to (typed.SiteNode) failed.")
			continue
		}

		var isFind = false
		for _, site := range sites {
			siteInfo, ok := site.(typed.Site)
			if !ok {
				continue
			}

			if siteInfo.Region == siteNode.Region && siteInfo.Az == siteNode.AvailabilityZone {
				info := convertToSiteNode(siteInfo, siteNode)
				err := collector.Cache().AddNode(info)
				if err != nil {
					logger.Infof("add node to cache failed! err: %s", err)
				}

				isFind = true
				break
			}
		}

		if !isFind {
			site := &types.SiteNode{
				SiteID: siteNode.Region + "--" + siteNode.AvailabilityZone,
				RegionAzMap: types.RegionAzMap{
					Region:           siteNode.Region,
					AvailabilityZone: siteNode.AvailabilityZone,
				},
				Status: constants.SiteStatusNormal,
			}

			site.Nodes = append(site.Nodes, siteNode.Nodes...)
			err := collector.Cache().AddNode(site)
			if err != nil {
				logger.Infof("add node to cache failed! err: %s", err)
			}
		}
	}

	collector.Cache().PrintString()
}

func convertToSiteNode(site typed.Site, node typed.SiteNode) *types.SiteNode {
	siteNode := &types.SiteNode{
		SiteID: site.ID,
		GeoLocation: types.GeoLocation{
			Country:  site.Country,
			Area:     site.Area,
			Province: site.Province,
			City:     site.City,
		},
		RegionAzMap: types.RegionAzMap{
			Region:           site.Region,
			AvailabilityZone: site.Az,
		},
		Operator:      site.Operator.Name,
		EipTypeName:   site.EipTypeName,
		Status:        site.Status,
		SiteAttribute: site.SiteAttributes,
	}

	siteNode.Nodes = append(siteNode.Nodes, node.Nodes...)
	return siteNode
}

// initPodInformers init collector with podInformer
//func (c *Collector) initPodInformers(stopCh <-chan struct{}) error {
//	masterURL := config.DefaultString("master", "127.0.0.1:8080")
//	kubeconfig := config.DefaultString("kubeconfig", "/var/run/kubernetes/admin.kubeconfig")
//
//	// init client
//	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
//	if err != nil {
//		return err
//	}
//
//	client, err := clientset.NewForConfig(cfg)
//	if err != nil {
//		return err
//	}
//
//
//	c.InformerFactory = informers.NewSharedInformerFactory(client, 0)
//	c.PodInformer = factory.NewPodInformer(client, 0)
//	c.NextStack = internalqueue.MakeNextStackFunc(c.StackQueue)
//	c.Client = client
//	return nil
//}
