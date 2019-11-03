package dgGraph

import (
	"sync"
)

type DGClient struct {
	//MessagesNumber uint64
}

func NewDGClient() DGClient {
	return DGClient{
	//	MessagesNumber: 0,
	}
}

var mut = sync.Mutex{}
var cond = sync.NewCond(&mut)


func (client DGClient) Run(graph *dgGraph, preset []*DGRequest) {
	go graph.Start()

	for _, request := range preset {
		mut.Lock()
		if graph.isFull() {
			cond.Wait()
		}
		*graph.addAndDeleteChannel <- NewManagementMessage(AddRequest, request)
		mut.Unlock()
	}
}

func (client DGClient) toString() string {
	return "Client";
}
