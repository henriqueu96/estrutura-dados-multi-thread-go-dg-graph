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
var freeToAdd = make(chan int, 1000)


func (client DGClient) Run(graph *dgGraph, preset []*DGRequest) {
	go graph.Start()
	go freeToAddReader(graph)
	for _, request := range preset {
		mut.Lock()
		if graph.isFull() {
			cond.Wait()
		}
		*graph.addAndDeleteChannel <- NewManagementMessage(AddRequest, request)
		mut.Unlock()
	}
}

var ExitedNodes uint64 = 0
func freeToAddReader(graph *dgGraph) {
	for true{
		_ = <-freeToAdd
		ExitedNodes++
		graph.length--
		cond.Signal()
	}
}

func (client DGClient) toString() string {
	return "Client";
}
