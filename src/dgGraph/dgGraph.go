package dgGraph

import (
	"fmt"
)

var shouldPrint = true;

type dgGraph struct {
	AddChannel                        *chan ManagementMessage
	lastNodeInManagementChannel       *chan ManagementMessage
	UpdateLastNodeInManagementChannel *chan ManagementMessage
	WantDeleteChannel                 *chan ManagementMessage
	GraphLimit                        uint64
	Client                            DGClient
	Population                        uint64
}

func NewGraph(graphLimit uint64) dgGraph {
	add := make(chan ManagementMessage)
	remove := make(chan ManagementMessage)
	delete := make(chan ManagementMessage)
	return dgGraph{
		GraphLimit:                        graphLimit,
		lastNodeInManagementChannel:       nil,
		AddChannel:                        &add,
		UpdateLastNodeInManagementChannel: &remove,
		WantDeleteChannel:                 &delete,
		Population:                        0,
	}
}

func (dgGraph *dgGraph) add(request *DGRequest, clientManagementChannel *chan ManagementMessage) {
	node := newNode(request, dgGraph.lastNodeInManagementChannel, clientManagementChannel, dgGraph)
	node.start()
	if (shouldPrint) {
		fmt.Println("Event:" + "enterNewNode" + " " + node.ToString())
	}
	*dgGraph.UpdateLastNodeInManagementChannel <- NewManagementMessage(irrelevant, node.inManagementChannel)
	if (node.NextNodeInManagementChannel == nil) {
		node.status = ready
		go Work(&node)
	} else {
		*dgGraph.lastNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, &node)
	}
}

func (graph *dgGraph) Start(client *DGClient) {
	for {
		var message ManagementMessage;
		select {
		case message = <-*graph.AddChannel:
			newRequest := message.parameter.(*DGRequest)
			graph.add(newRequest, client.inManagementChannel)
			graph.Population++
			fmt.Println("adicionou no graph")

		case message = <-*graph.UpdateLastNodeInManagementChannel:
			updatedLastInManagementChannel := message.parameter.(*chan ManagementMessage)
			if updatedLastInManagementChannel ==  nil {
				fmt.Println("se fosse em java tava pronto!" )
			}
			graph.lastNodeInManagementChannel = updatedLastInManagementChannel
			graph.Population--
			fmt.Println("removeu do graph")
		}
	}
}

func (graph *dgGraph) GetPopulationAdd() bool {
	if (graph.Population < graph.GraphLimit) {
		return true
	}
	return false
}
