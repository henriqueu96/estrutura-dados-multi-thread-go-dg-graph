package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel *chan ManagementMessage
	addAndDeleteChannel         *chan ManagementMessage
	length                      int
	GraphLimit                  int
	Client                      DGClient
}

func NewGraph(graphLimit int) dgGraph {
	addAndDeleteChannel := make(chan ManagementMessage, graphLimit)
	return dgGraph{
		GraphLimit:                  graphLimit,
		lastNodeInManagementChannel: nil,
		addAndDeleteChannel:         &addAndDeleteChannel,
		length:                      0,
	}
}

func (dgGraph *dgGraph) toString() string {
	return "Graph";
}

func (dgGraph *dgGraph) add(request *DGRequest) {
	node := newNode(request, dgGraph)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	go node.start()
	*node.inManagementChannel <- NewManagementMessage(enterNewNode, &node)
}

func (graph *dgGraph) Start() {
	for {
		message := <-*graph.addAndDeleteChannel;
		var printer Printer = nil;
		switch message.messageType {
		case AddRequest:
			printer = graph.Client
			newRequest := message.parameter.(*DGRequest)
			graph.add(newRequest)

		case leavingNode:
			theLeavingNode := message.parameter.(*dgNode)

			if theLeavingNode.inManagementChannel == graph.lastNodeInManagementChannel {
				theLeavingNodeIsTheFirstHandle(graph, theLeavingNode)
			} else {
				*graph.lastNodeInManagementChannel <- NewManagementMessage(leavingNode, theLeavingNode)
			}
		}

		if printer != nil {
			PrintMessage(message, graph, printer)
		} else {
			PrintMessageWithoutSender(message, graph)
		}
	}
}

func (dgGraph *dgGraph) isFull() bool {
	return dgGraph.length >= dgGraph.GraphLimit-1;
}

func theLeavingNodeIsTheFirstHandle(graph *dgGraph, theLeavingNode *dgNode) {
	channel := make(chan ManagementMessage)
	*theLeavingNode.inManagementChannel <- NewManagementMessage(wantToDelete, &channel)
	message := <- channel
	close(channel)
	wantToDeleteAnswerChannelHandle(graph, &message)
}

func wantToDeleteAnswerChannelHandle(graph *dgGraph, message *ManagementMessage) {
	nodeDelete := message.parameter.(*dgNode)

	graph.lastNodeInManagementChannel = nodeDelete.NextNodeInManagementChannel
	*nodeDelete.inManagementChannel <- NewManagementMessage(leavingNode, nodeDelete)
}
