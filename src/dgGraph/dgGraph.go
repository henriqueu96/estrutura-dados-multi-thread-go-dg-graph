package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel *chan ManagementMessage
	//lastNodeOutManagementChannel *chan ManagementMessage
	GraphLimit int
}

func (dgGraph *dgGraph) add(request *DGRequest, clientManagementChannel *chan ManagementMessage) {
	node := newNode(request, dgGraph.lastNodeInManagementChannel, clientManagementChannel)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	go node.start()
	*node.inManagementChannel <- NewManagementMessage(enterNewNode, &node)
}

func NewGraph(graphLimit int) dgGraph {
	return dgGraph{
		GraphLimit: graphLimit,
		lastNodeInManagementChannel: nil,
	}
}
