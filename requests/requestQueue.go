package requests

import (
	"sync"
	"tcc/boolGenerator"
	"time"
)

type RequestQueue struct {
	pendingRequests []*Request
	limit           int
	Mutex           *sync.Mutex
	NotFull         *sync.Cond
	HasReady        *sync.Cond
}

func newRequestQueue(limit int) (queue RequestQueue) {
	mutex := &sync.Mutex{}
	notFull := sync.NewCond(mutex)
	hasReady := sync.NewCond(mutex)

	queue = RequestQueue{
		limit:    limit,
		Mutex:    mutex,
		NotFull:  notFull,
		HasReady: hasReady,
	}
	return
}

func isDependent() bool {
	return boolGenerator.New(time.Now())
}

func (queue *RequestQueue) add(request *Request) {
	queue.Mutex.Lock()
	for queue.requestCount() >= queue.limit {
		queue.NotFull.Wait()
	}
	for _, current := range queue.pendingRequests {
		if isDependent() {
			request.addDependency(current)
			current.addDependent(request)
			request.ExecState = Blocked
		}
	}
	queue.pendingRequests = append(queue.pendingRequests, request)
	if request.ExecState == Ready {
		queue.HasReady.Signal()
	}
	queue.Mutex.Unlock()
}

func (queue *RequestQueue) hasRequest() bool {
	return len(queue.pendingRequests) > 0
}

func (queue *RequestQueue) remove(request *Request) {
	queue.Mutex.Lock()

	for _, current := range request.dependents {
		if current != nil {
			current.removeDependency(request)
			if !current.hasDependency() {
				current.ExecState = Ready
				queue.HasReady.Signal()
			}
		}
	}
	queue.pendingRequests = removeRequest(queue.pendingRequests, request)
	queue.NotFull.Signal()
	queue.Mutex.Unlock()
}

func (queue *RequestQueue) nextRequest() *Request {
	queue.Mutex.Lock()
	for true {
		for _, request := range queue.pendingRequests {
			if request.ExecState == Ready {
				request.ExecState = Running
				queue.Mutex.Unlock()
				return request
			}
		}
		queue.HasReady.Wait()
	}
	queue.Mutex.Unlock()
	return nil
}

func (queue *RequestQueue) clear() {
	queue.Mutex.Lock()
	queue.pendingRequests = []*Request{}
	queue.Mutex.Unlock()
}

func (queue *RequestQueue) requestCount() int {
	return len(queue.pendingRequests)
}
