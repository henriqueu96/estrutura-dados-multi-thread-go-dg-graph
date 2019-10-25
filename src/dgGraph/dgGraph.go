package dgGraph

type dgGraph struct {
	AddChannel                  *chan ManagementMessage
	lastNodeInManagementChannel *chan ManagementMessage
	RemoveChannel               *chan ManagementMessage
	GraphLimit                  int
	Client DGClient
}

func NewGraph(graphLimit int) dgGraph {
	add := make(chan ManagementMessage)
	remove := make(chan ManagementMessage)
	return dgGraph{
		GraphLimit:                  graphLimit,
		lastNodeInManagementChannel: nil,
		AddChannel:                  &add,
		RemoveChannel:                  &remove,
	}
}

func (dgGraph *dgGraph) add(request *DGRequest, clientManagementChannel *chan ManagementMessage) {
	node := newNode(request, dgGraph.lastNodeInManagementChannel, clientManagementChannel, dgGraph)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	go node.start()
	*node.inManagementChannel <- NewManagementMessage(enterNewNode, &node)
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
