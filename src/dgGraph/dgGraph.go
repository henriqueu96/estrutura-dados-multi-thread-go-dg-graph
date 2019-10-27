package dgGraph

import "fmt"

type dgGraph struct {
	AddChannel                  *chan ManagementMessage
	lastNodeInManagementChannel *chan ManagementMessage
	RemoveChannel               *chan ManagementMessage
	WantDeleteChannel 			*chan ManagementMessage
	GraphLimit                  int
	Client DGClient
}

func NewGraph(graphLimit int) dgGraph {
	add := make(chan ManagementMessage)
	remove := make(chan ManagementMessage)
	delete := make(chan ManagementMessage)
	return dgGraph{
		GraphLimit:                  graphLimit,
		lastNodeInManagementChannel: nil,
		AddChannel:                  &add,
		RemoveChannel:                  &remove,
		WantDeleteChannel: &delete,
	}
}

func (dgGraph *dgGraph) add(request *DGRequest, clientManagementChannel *chan ManagementMessage) {
	node := newNode(request, dgGraph.lastNodeInManagementChannel, clientManagementChannel, dgGraph)
	node.start()
	if(GetPrint()){
		fmt.Println("Event:" + "enterNewNode" + " " + node.ToString())
	}
	if (node.NextNodeInManagementChannel == nil) {
		node.status = ready
		go Work(&node)
	}else{
		*dgGraph.lastNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, &node)
	}
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
}

func (graph *dgGraph) Start(client *DGClient) {
	for {
		var message ManagementMessage;
		select  {
		case message = <- *graph.AddChannel:
			newRequest := message.parameter.(*DGRequest)
			graph.add(newRequest, client.inManagementChannel)

		case message = <- *graph.RemoveChannel:
			updatedLastInManagementChannel :=  message.parameter.(*chan ManagementMessage)
			graph.lastNodeInManagementChannel = updatedLastInManagementChannel
		}
	}
}

func GetPrint() bool {
	return false
}