package main

import (
	"fmt"
	collector "k8s.io/kubernetes/resourcecollector/pkg/collector"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/common/apiserver"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/common/logger"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/router"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/rpc"
)

func main() {
	// todo k8s stopCh
	stopCh := make(<-chan struct{})

	commandRun(stopCh)
}

func commandRun(stopCh <-chan struct{}) error {
	// init collector
	collector.InitCollector(stopCh)
	col := collector.GetCollector()
	if col == nil {
		return fmt.Errorf("get new collector failed")
	}

	// start scheduler resource cache informer and run
	col.StartInformersAndRun(stopCh)

	// start the gRPC service to get cluster static info from ClusterController
	go rpc.NewRpcServer()

	// todo  wait until all informer synced
	// start http server to provide resource information to the Scheduler
	router.Register()
	hs, err := apiserver.NewHTTPServer()
	if err != nil {
		logger.Errorf("new http server failed!, err: %s", err.Error())
		return err
	}

	return hs.BlockingRun(stopCh)
}
