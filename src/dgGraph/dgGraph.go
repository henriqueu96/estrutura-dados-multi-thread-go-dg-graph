package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel *chan ManagementMessage
	lastNode                    *dgNode
	addAndUpdateLastChannel     *chan ManagementMessage
	leavingNodeChannel          *chan ManagementMessage
	outManagementChannel        *chan ManagementMessage
	wantToDeleteAnswaerChannel  *chan ManagementMessage
	GraphLimit                  int
	Client                      DGClient
	length                      int
}

func NewGraph(graphLimit int) dgGraph {
	addAndUpdateLastChannel := make(chan ManagementMessage)
	leavingNodeChannel := make(chan ManagementMessage)
	wantToDeleteAnswaerChannel := make(chan ManagementMessage)
	return dgGraph{
		GraphLimit:                  graphLimit,
		lastNodeInManagementChannel: nil,
		addAndUpdateLastChannel:     &addAndUpdateLastChannel,
		leavingNodeChannel:          &leavingNodeChannel,
		wantToDeleteAnswaerChannel:  &wantToDeleteAnswaerChannel,
		length:                      0,
	}
}

func (dgGraph *dgGraph) toString() string {
	return "Graph";
}

func (dgGraph *dgGraph) add(request *DGRequest, clientManagementChannel *chan ManagementMessage) {
	node := newNode(request, dgGraph.lastNodeInManagementChannel, clientManagementChannel, dgGraph)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	go node.start()
	*node.inManagementChannel <- NewManagementMessage(enterNewNode, &node)
}

func (graph *dgGraph) Start() {
	go graph.StartAddAndUpdateLastChannelChannelReader()
	go graph.StartLeavingNodeChannelReader()
}

func (graph *dgGraph) StartAddAndUpdateLastChannelChannelReader() {
	for {
		message := <-*graph.addAndUpdateLastChannel;
		var printer Printer = nil;
		switch message.messageType {
		case AddRequest:
			graph.length++
			printer = graph.Client
			newRequest := message.parameter.(*DGRequest)
			graph.add(newRequest, graph.leavingNodeChannel)

		case UpdateLastInManagementChannel:
			printer = graph
			updatedLastInManagementChannel := message.parameter.(*chan ManagementMessage)
			graph.lastNodeInManagementChannel = updatedLastInManagementChannel
		}

		if printer != nil {
			PrintMessage(message, graph, printer)
		} else {
			PrintMessageWithoutSender(message, graph)
		}
	}
}

func (graph *dgGraph) StartLeavingNodeChannelReader() {
	for {
		message := <-*graph.leavingNodeChannel;
		theLeavingNode := message.parameter.(*dgNode)

		if message.parameter != nil {
			PrintMessage(message, graph, theLeavingNode)
		} else {
			PrintMessageWithoutSender(message, graph)
		}

		switch message.messageType {
		case leavingNode:
			if theLeavingNode.inManagementChannel == graph.lastNodeInManagementChannel {
				go theLeavingNodeIsTheFirstHandle(graph, theLeavingNode)
			} else {
				*graph.lastNodeInManagementChannel <- NewManagementMessage(leavingNode, theLeavingNode)
			}
		}
	}
}

func (dgGraph *dgGraph) isFull() bool {
	return dgGraph.length >= dgGraph.GraphLimit-1;
}

func theLeavingNodeIsTheFirstHandle(graph *dgGraph, theLeavingNode *dgNode) {
	*theLeavingNode.inManagementChannel <- NewManagementMessage(wantToDelete, graph.wantToDeleteAnswaerChannel)
	message := <-*graph.wantToDeleteAnswaerChannel
	wantToDeleteAnswaerChannelHandle(graph, &message)
}

func wantToDeleteAnswaerChannelHandle(graph *dgGraph, message *ManagementMessage) {
	nodeDelete := message.parameter.(*dgNode)
	*graph.addAndUpdateLastChannel <- NewManagementMessage(UpdateLastInManagementChannel, nodeDelete.NextNodeInManagementChannel)
	*nodeDelete.inManagementChannel <- NewManagementMessage(leavingNode, nodeDelete)
}
