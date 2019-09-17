package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel chan ManagementMessage
	GraphLimit int
}

func (dgGraph *dgGraph) add(request *request) {
	node := newNode()
	node.status = entering
	node.request = request
	node.NextNodeInManagementChannel = dgGraph.lastNodeInManagementChannel
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	node.start()
	node.inManagementChannel <- NewManagementMessage(enterNewNode, nil)
}

func NewGraph(graphLimit int) dgGraph {
	return dgGraph{
		GraphLimit: graphLimit,
		lastNodeInManagementChannel: make(chan ManagementMessage),
	}
}
