package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel *chan ManagementMessage
	lastNodeOutManagementChannel *chan ManagementMessage
	GraphLimit int
}

func (dgGraph *dgGraph) add(request *DGRequest) {
	node := newNode(request, dgGraph.lastNodeInManagementChannel, dgGraph.lastNodeOutManagementChannel)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	dgGraph.lastNodeOutManagementChannel = node.outManagementChannel
	go node.start()
	*node.inManagementChannel <- NewManagementMessage(enterNewNode, &node)
}

func NewGraph(graphLimit int) dgGraph {
	return dgGraph{
		GraphLimit: graphLimit,
		lastNodeInManagementChannel: nil,
	}
}
