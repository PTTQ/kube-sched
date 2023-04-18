/**
 * Copyright (2021, ) Institute of Software, Chinese Academy of Sciences
 **/

package scheduler

type Result struct {
	NodeName string
	GpuUuid    []string
}

type Algorithm interface {
	Name() string
	// Schedule selects a node with gpus for a pod, returns nil if there is no suitable resource.
	Schedule(requestMemory, requestCore int64, availableNode []string, resourceOnNode map[string]*NodeResource) *Result
}
