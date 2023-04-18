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

type GpuResource struct {
	GpuName         string
	Uuid            string
	Node            string
	CoreCapacity    int64
	CoreAllocated   int64
	MemoryCapacity  int64
	MemoryAllocated int64
}

type NodeResource struct {
	NodeName        string
	HasDevicePlugin bool
	GpusByUuid      map[string]*GpuResource
}

type GpuManager struct {
	queue *util.LinkedQueue
	mu    sync.Mutex
}

func NewGpuManager(queue *util.LinkedQueue) *GpuManager {
	return &GpuManager{queue: queue}
}

func (gpuMgr *GpuManager) DoAdded(obj map[string]interface{}) {
	bytes, _ := json.Marshal(obj)
	gpuMgr.mu.Lock()
	gpuMgr.queue.Add(kubesys.ToJsonObject(bytes))
	gpuMgr.mu.Unlock()
}

func (gpuMgr *GpuManager) DoModified(obj map[string]interface{}) {
}

func (gpuMgr *GpuManager) DoDeleted(obj map[string]interface{}) {
}
