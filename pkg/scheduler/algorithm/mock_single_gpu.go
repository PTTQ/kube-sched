/**
 * Copyright (2021, ) Institute of Software, Chinese Academy of Sciences
 **/

package scheduler

import (
	"github.com/pttq/kube-sched/pkg/scheduler"
)

type MockSingleGPU struct {
	scheduler.Algorithm
}

func NewMockSingleGPU() *MockSingleGPU {
	return &MockSingleGPU{}
}

func (alg *MockSingleGPU) Schedule(requestMemory, requestCore int64, availableNode []string, resourceOnNode map[string]*scheduler.NodeResource) *scheduler.Result {
	gpuUuid := ""
	nodeName := ""
	res := int64(0)
	for node, gpus := range resourceOnNode {
		if gpus.HasDevicePlugin {
			for id, gpu := range gpus.GpusByUuid {
				if gpu.CoreCapacity-gpu.CoreAllocated >= requestCore && gpu.MemoryCapacity-gpu.MemoryAllocated >= requestMemory && gpu.MemoryCapacity-gpu.MemoryAllocated > res {
					gpuUuid = id
					nodeName = node
					res = gpu.MemoryCapacity - gpu.MemoryAllocated
				}
			}
		}
	}

	if gpuUuid == "" || nodeName == "" {
		return nil
	}

	result := scheduler.Result{
		NodeName: nodeName,
		GpuUuid:  []string{gpuUuid},
	}
	return &result
}

func (alg *MockSingleGPU) Name() string {
	return "MockSingleGPU"
}
