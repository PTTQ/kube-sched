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

type PodManager struct {
	queueOfAdded   *util.LinkedQueue
	queueOfDeleted *util.LinkedQueue
	muOfAdd        sync.Mutex
	muOfDelete     sync.Mutex
}

func NewPodManager(queueOfAdded, queueOfDeleted *util.LinkedQueue) *PodManager {
	return &PodManager{queueOfAdded: queueOfAdded, queueOfDeleted: queueOfDeleted}
}

func (podMgr *PodManager) DoAdded(obj map[string]interface{}) {
	bytes, _ := json.Marshal(obj)
	podMgr.muOfAdd.Lock()
	podMgr.queueOfAdded.Add(kubesys.ToJsonObject(bytes))
	podMgr.muOfAdd.Unlock()
}

func (podMgr *PodManager) DoModified(obj map[string]interface{}) {
}

func (podMgr *PodManager) DoDeleted(obj map[string]interface{}) {
	bytes, _ := json.Marshal(obj)
	podMgr.muOfDelete.Lock()
	podMgr.queueOfDeleted.Add(kubesys.ToJsonObject(bytes))
	podMgr.muOfDelete.Unlock()
}
