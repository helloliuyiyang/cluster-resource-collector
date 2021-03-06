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

package cache

import (
	"fmt"

	schedulerlisters "k8s.io/kubernetes/resourcecollector/pkg/collector/listers"
	schedulernodeinfo "k8s.io/kubernetes/resourcecollector/pkg/collector/nodeinfo"
	"k8s.io/kubernetes/resourcecollector/pkg/collector/types"
)

// Snapshot is a snapshot of cache NodeInfo and NodeTree order. The scheduler takes a
// snapshot at the beginning of each scheduling cycle and uses it for its operations in that cycle.
type Snapshot struct {
	// NodeInfoMap a map of node name to a snapshot of its NodeInfo.
	NodeInfoMap map[string]*schedulernodeinfo.NodeInfo
	// NodeInfoList is the list of nodes as ordered in the cache's nodeTree.
	NodeInfoList []*schedulernodeinfo.NodeInfo
	// HavePodsWithAffinityNodeInfoList is the list of nodes with at least one pod declaring affinity terms.
	HavePodsWithAffinityNodeInfoList []*schedulernodeinfo.NodeInfo
	generation                       int64
}

var _ schedulerlisters.SharedLister = &Snapshot{}

// NewEmptySnapshot initializes a Snapshot struct and returns it.
func NewEmptySnapshot() *Snapshot {
	return &Snapshot{
		NodeInfoMap: make(map[string]*schedulernodeinfo.NodeInfo),
	}
}

// NewSnapshot initializes a Snapshot struct and returns it.
func NewSnapshot(pods []*types.Stack, nodes []*types.SiteNode) *Snapshot {
	nodeInfoMap := createNodeInfoMap(pods, nodes)
	nodeInfoList := make([]*schedulernodeinfo.NodeInfo, 0, len(nodeInfoMap))
	havePodsWithAffinityNodeInfoList := make([]*schedulernodeinfo.NodeInfo, 0, len(nodeInfoMap))
	for _, v := range nodeInfoMap {
		nodeInfoList = append(nodeInfoList, v)
		if len(v.StackWithAffinity()) > 0 {
			havePodsWithAffinityNodeInfoList = append(havePodsWithAffinityNodeInfoList, v)
		}
	}

	s := NewEmptySnapshot()
	s.NodeInfoMap = nodeInfoMap
	s.NodeInfoList = nodeInfoList
	s.HavePodsWithAffinityNodeInfoList = havePodsWithAffinityNodeInfoList

	return s
}

// createNodeInfoMap obtains a list of pods and pivots that list into a map
// where the keys are node names and the values are the aggregated information
// for that node.
func createNodeInfoMap(stacks []*types.Stack, nodes []*types.SiteNode) map[string]*schedulernodeinfo.NodeInfo {
	nodeNameToInfo := make(map[string]*schedulernodeinfo.NodeInfo)
	for _, stack := range stacks {
		nodeName := stack.Selected.NodeID
		if _, ok := nodeNameToInfo[nodeName]; !ok {
			nodeNameToInfo[nodeName] = schedulernodeinfo.NewNodeInfo()
		}
		nodeNameToInfo[nodeName].AddStack(stack)
	}

	for _, node := range nodes {
		if _, ok := nodeNameToInfo[node.SiteID]; !ok {
			nodeNameToInfo[node.SiteID] = schedulernodeinfo.NewNodeInfo()
		}
		nodeInfo := nodeNameToInfo[node.SiteID]
		nodeInfo.SetNode(node)
	}
	return nodeNameToInfo
}

// Stacks returns a StackLister
func (s *Snapshot) Stacks() schedulerlisters.StackLister {
	return stackLister(s.NodeInfoList)
}

// NodeInfos returns a NodeInfoLister.
func (s *Snapshot) NodeInfos() schedulerlisters.NodeInfoLister {
	return s
}

// NumNodes returns the number of nodes in the snapshot.
func (s *Snapshot) NumNodes() int {
	return len(s.NodeInfoList)
}

type stackLister []*schedulernodeinfo.NodeInfo

// List returns the list of stacks in the snapshot.
func (p stackLister) List() ([]*types.Stack, error) {
	alwaysTrue := func(*types.Stack) bool { return true }
	return p.FilteredList(alwaysTrue)
}

// FilteredList returns a filtered list of stacks in the snapshot.
func (p stackLister) FilteredList(filter schedulerlisters.StackFilter) ([]*types.Stack, error) {
	// stackFilter is expected to return true for most or all of the stacks. We
	// can avoid expensive array growth without wasting too much memory by
	// pre-allocating capacity.
	maxSize := 0
	for _, n := range p {
		maxSize += len(n.Stacks())
	}
	stacks := make([]*types.Stack, 0, maxSize)
	for _, n := range p {
		for _, stack := range n.Stacks() {
			if filter(stack) {
				stacks = append(stacks, stack)
			}
		}
	}
	return stacks, nil
}

// List returns the list of nodes in the snapshot.
func (s *Snapshot) List() ([]*schedulernodeinfo.NodeInfo, error) {
	return s.NodeInfoList, nil
}

// HavePodsWithAffinityList returns the list of nodes with at least one pods with inter-pod affinity
func (s *Snapshot) HavePodsWithAffinityList() ([]*schedulernodeinfo.NodeInfo, error) {
	return s.HavePodsWithAffinityNodeInfoList, nil
}

// Get returns the NodeInfo of the given node name.
func (s *Snapshot) Get(nodeID string) (*schedulernodeinfo.NodeInfo, error) {
	if v, ok := s.NodeInfoMap[nodeID]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("nodeinfo not found for node name %q", nodeID)
}
