/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package output

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	glog "github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/app/output/action"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/metrics"
)

const (
	// defaultHandlerQueueSize is default queue size of Handler.
	defaultHandlerQueueSize = 1024

	// defaultHandleInterval is default interval of handle.
	defaultHandleInterval = 500 * time.Millisecond

	// defaultHandlerReportPeriod report queue length for handler dataType
	defaultHandlerReportPeriod = 5 * time.Second
)

// Action handles the metadata in ADD/DEL/UPDATE methods.
type Action interface {
	// Add adds new resource metadata.
	Add(syncData *action.SyncData)

	// Delete deletes target resource metadata.
	Delete(syncData *action.SyncData)

	// Update updates target resource metadata.
	Update(syncData *action.SyncData)
}

// Handler is resource handler, consumes metadata distributed from
// Writer, and handles data with the action.
type Handler struct {
	// clusterID
	clusterID string
	// resource metadata type.
	dataType string

	// distributed metadata queue.
	queue chan *action.SyncData

	// metadata action.
	act Action
}

// NewHandler creates a new resource Handler instance with the action.
func NewHandler(clusterID string, dataType string, act Action) *Handler {
	h := &Handler{
		dataType:  dataType,
		queue:     make(chan *action.SyncData, defaultHandlerQueueSize),
		act:       act,
		clusterID: clusterID,
	}
	return h
}

// Handle sends the metadata into handler queue.
func (h *Handler) Handle(data *action.SyncData) {
	h.queue <- data
}

// HandleWithTimeout sends the metadata into handler queue with timeout.
func (h *Handler) HandleWithTimeout(data *action.SyncData, timeout time.Duration) {
	select {
	case h.queue <- data:
		metrics.ReportK8sWatchHandlerQueueLengthInc(h.clusterID, h.dataType)
	case <-time.After(timeout):
		metrics.ReportK8sWatchHandlerDiscardEvents(h.clusterID, h.dataType)
		glog.Warn("can't handle data, queue timeout")
	}
}

// debugs here.
func (h *Handler) debug() {
	for {
		time.Sleep(debugInterval)
		glog.Infof("Handler[%+v] debug: QueueLen[%d]", h.dataType, len(h.queue))
	}
}

// reportQueueLength report datatype length to prometheus metrics
func (h *Handler) reportHandlerQueueLength() {
	metrics.ReportK8sWatchHandlerQueueLength(h.clusterID, h.dataType, float64(len(h.queue)))
}

// handle func is invoked by wait.NonSlidingUntil with a stop channel, do not block
// to recv the queue here in order to make it have runtime to handle the stop channel.
func (h *Handler) handle() {
	// try to keep reading from queue until there is no more data every period.
	for {
		select {
		case data := <-h.queue:
			metrics.ReportK8sWatchHandlerQueueLengthDec(h.clusterID, h.dataType)
			switch data.Action {
			case action.SyncDataActionAdd:
				h.act.Add(data)

			case action.SyncDataActionDelete:
				h.act.Delete(data)

			case action.SyncDataActionUpdate:
				h.act.Update(data)

			default:
				glog.Errorf("can't handle metadata, unknown action type[%+v]", data.Action)
			}

		case <-time.After(defaultQueueTimeout):
			// no more data, break loop.
			return
		}
	}
}

// Run starts the handler.
func (h *Handler) Run(stopCh <-chan struct{}) {
	glog.Infof("%+v resource handler is starting now", h.dataType)
	go wait.NonSlidingUntil(h.handle, defaultHandleInterval, stopCh)
	go wait.Until(h.reportHandlerQueueLength, defaultHandlerReportPeriod, stopCh)

	// setup debug.
	//go h.debug()
}
