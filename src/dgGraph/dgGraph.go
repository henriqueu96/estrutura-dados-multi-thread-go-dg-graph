package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel *chan ManagementMessage
	//lastNodeOutManagementChannel *chan ManagementMessage
	GraphLimit int
}

func (graph *dgGraph) setNextNodeInManagmentChanel(newChandel *chan ManagementMessage){
	addRemoveSequencialy.Lock()

	graph.lastNodeInManagementChannel = newChandel

	addRemoveSequencialy.Unlock()
}

func (dgGraph *dgGraph) add(request *DGRequest, clientManagementChannel *chan ManagementMessage) {
	addRemoveSequencialy.Lock()

	node := newNode(request, dgGraph.lastNodeInManagementChannel, clientManagementChannel, dgGraph)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	go node.start()
	*node.inManagementChannel <- NewManagementMessage(enterNewNode, &node)

	addRemoveSequencialy.Unlock()
}

func NewGraph(graphLimit int) dgGraph {
	return dgGraph{
		GraphLimit: graphLimit,
		lastNodeInManagementChannel: nil,
	}
}