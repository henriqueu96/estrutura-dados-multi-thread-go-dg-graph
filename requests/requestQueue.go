package requests

import (
	"tcc/boolGenerator"
	"time"
)

type RequestQueue struct {
	pendingRequests []*Request
	limit           int
}

func newRequestQueue(limit int) RequestQueue {
	return RequestQueue{
		limit: limit,
	}
}

func isDependent() bool {
	return boolGenerator.New(time.Now())
}

func (queue *RequestQueue) add(request *Request) {
	for _, current := range queue.pendingRequests  {
		if(isDependent()){
			request.addDependency(current)
			current.addDependent(request)
			request.ExecState = Blocked
		}
	}
	queue.pendingRequests = append(queue.pendingRequests, request)
}

func (queue *RequestQueue) hasRequest() bool {
	return len(queue.pendingRequests) > 0
}

func (queue *RequestQueue) remove(request *Request) {
	for _, current := range request.dependents  {
		if(current!= nil){
			current.removeDependency(request)
			if(!current.hasDependency()){
				current.ExecState = Ready
			}
		}
	}
	queue.pendingRequests = removeRequest(queue.pendingRequests, request)
}

func (queue *RequestQueue) nextRequest() *Request {
	for _, request := range queue.pendingRequests {
		if request.ExecState == Ready {
			request.ExecState = Running
			return request
		}
	}
	return nil
}

func (queue *RequestQueue) clear() {
	queue.pendingRequests = []*Request{}
}

func (queue *RequestQueue) requestCount() int{
	return len(queue.pendingRequests)
}
