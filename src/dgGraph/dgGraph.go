package dgGraph

type dgGraph struct {
	lastNodeInManagementChannel *chan ManagementMessage
	addAndUpdateLastChannel     *chan ManagementMessage
	leavingNodeChannel          *chan ManagementMessage
	outManagementChannel        *chan ManagementMessage
	GraphLimit                  int
	Client                      DGClient
}

func NewGraph(graphLimit int) dgGraph {
	addAndUpdateLastChannel := make(chan ManagementMessage)
	leavingNodeChannel := make(chan ManagementMessage)
	return dgGraph{
		GraphLimit:                  graphLimit,
		lastNodeInManagementChannel: nil,
		addAndUpdateLastChannel:     &addAndUpdateLastChannel,
		leavingNodeChannel: &leavingNodeChannel,
	}
}

func (dgGraph *dgGraph) toString() string{
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
			printer = graph.Client
			newRequest := message.parameter.(*DGRequest)
			graph.add(newRequest, graph.leavingNodeChannel)

		case UpdateLastInManagementChannel:
			printer = graph
			updatedLastInManagementChannel := message.parameter.(*chan ManagementMessage)
			graph.lastNodeInManagementChannel = updatedLastInManagementChannel
		}

		PrintMessage(message, graph, printer)
	}
}

func (graph *dgGraph) StartLeavingNodeChannelReader() {
	for {
		message := <-*graph.leavingNodeChannel;
		theLeavingNode := message.parameter.(*dgNode)

		PrintMessage(message, graph, theLeavingNode)

		switch message.messageType {
		case leavingNode:
			
		}
	}

}
