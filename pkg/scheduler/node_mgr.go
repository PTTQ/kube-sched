/**
 * Copyright (2021, ) Institute of Software, Chinese Academy of Sciences
 **/

package scheduler

import (
	"encoding/json"
	"github.com/kubesys/client-go/pkg/kubesys"
	"github.com/pttq/kube-sched/pkg/util"
	"sync"
)

type NodeManager struct {
	queue *util.LinkedQueue
	mu    sync.Mutex
}

func NewNodeManager(queue *util.LinkedQueue) *NodeManager {
	return &NodeManager{queue: queue}
}

func (nodeMgr *NodeManager) DoAdded(obj map[string]interface{}) {
}

func (nodeMgr *NodeManager) DoModified(obj map[string]interface{}) {
	bytes, _ := json.Marshal(obj)
	nodeMgr.mu.Lock()
	nodeMgr.queue.Add(kubesys.ToJsonObject(bytes))
	nodeMgr.mu.Unlock()
}

func (nodeMgr *NodeManager) DoDeleted(obj map[string]interface{}) {
}
