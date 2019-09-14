package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel chan ManagementMessage
}

func (dgGraph *dgGraph) add(request *request) {
	node := newNode()
	node.request = request
	node.NextNodeInManagementChannel = dgGraph.lastNodeInManagementChannel
	dgGraph.lastNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, node)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	node.start()
}

func NewGraph() dgGraph {
	return dgGraph{}
}
