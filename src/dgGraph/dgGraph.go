package dgGraph

import "sync/atomic"

type DGGraph struct {
	lastNodeInManagementChannel *chan ManagementMessage
	addAndDeleteChannel         *chan ManagementMessage
	GraphLimit                  int64
	Client                      DGClient
	Length                      int64
	AddedNodes                  int64
}

func NewGraph(graphLimit int64) DGGraph {
	addAndDeleteChannel := make(chan ManagementMessage, graphLimit)
	return DGGraph{
		GraphLimit:                  graphLimit,
		lastNodeInManagementChannel: nil,
		addAndDeleteChannel:         &addAndDeleteChannel,
		Length:                      0,
	}
}

func (dgGraph *DGGraph) toString() string {
	return "Graph";
}

func (dgGraph *DGGraph) add(request *DGRequest) {
	node := newNode(request, dgGraph)
	dgGraph.lastNodeInManagementChannel = node.inManagementChannel
	go node.start()
	*node.inManagementChannel <- NewManagementMessage(enterNewNode, &node)
}

func (graph *DGGraph) Start() {
	for {
		message := <-*graph.addAndDeleteChannel;
		var printer Printer = nil;
		switch message.messageType {
		case AddRequest:
			atomic.AddInt64(&graph.AddedNodes, 1)
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

func (dgGraph *DGGraph) isFull() bool {
	return dgGraph.Length >= dgGraph.GraphLimit-1;
}

func theLeavingNodeIsTheFirstHandle(graph *DGGraph, theLeavingNode *dgNode) {
	channel := make(chan ManagementMessage)
	*theLeavingNode.inManagementChannel <- NewManagementMessage(wantToDelete, &channel)
	message := <- channel
	close(channel)
	wantToDeleteAnswerChannelHandle(graph, &message)
}

func wantToDeleteAnswerChannelHandle(graph *DGGraph, message *ManagementMessage) {
	nodeDelete := message.parameter.(*dgNode)

	graph.lastNodeInManagementChannel = nodeDelete.NextNodeInManagementChannel
	*nodeDelete.inManagementChannel <- NewManagementMessage(leavingNode, nodeDelete)
}
