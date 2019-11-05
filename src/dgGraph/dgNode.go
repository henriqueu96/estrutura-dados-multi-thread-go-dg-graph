package dgGraph

import (
	"strconv"
	"sync"
	"sync/atomic"
)

var idIndex uint64 = 0;

type dgNode struct {
	id                          uint64
	status                      dgNodeStatus
	request                     *DGRequest
	dependenciesNumber          uint64
	solvedDependenciesNumber    uint64
	dependentsChannelsList      *[]*chan ManagementMessage
	inManagementChannel         *chan ManagementMessage
	answersManagementChannel    *chan ManagementMessage
	leavingNodeAnswerChannel    *chan ManagementMessage
	NextNodeInManagementChannel *chan ManagementMessage
	GraphManagementChannel      *chan ManagementMessage
	isOn                        bool
	graph                       *dgGraph
}

func newNode(request *DGRequest, graph *dgGraph) dgNode {
	idIndex++;
	chanIn := make(chan ManagementMessage, 10)
	chanOut := make(chan ManagementMessage, 10)
	chanWant := make(chan ManagementMessage)
	return dgNode{
		id:                          idIndex,
		request:                     request,
		dependenciesNumber:          0,
		solvedDependenciesNumber:    0,
		dependentsChannelsList:      &[]*chan ManagementMessage{},
		inManagementChannel:         &chanIn,
		answersManagementChannel:    &chanOut,
		NextNodeInManagementChannel: graph.lastNodeInManagementChannel,
		status:                      entering,
		GraphManagementChannel:      graph.addAndDeleteChannel,
		isOn:                        true,
		graph:                       graph,
		leavingNodeAnswerChannel:    &chanWant,
	}
}

func (node *dgNode) toString() string {
	return "(Node) Id:" + strconv.FormatUint(node.id, 10)
}

func (node *dgNode) start() {
	go node.inManagementChannelReader()
	go node.answersManagementChannelReader()
}

var addRemoveSequencialy = sync.Mutex{}

func (node *dgNode) IsNextNode(nextNode *dgNode) bool {
	return node.NextNodeInManagementChannel == nextNode.inManagementChannel
}

func (node *dgNode) IsRedyToGo() bool {
	return node.status == waiting && node.dependenciesNumber == node.solvedDependenciesNumber
}

func (node *dgNode) inManagementChannelReader() {
	for node.isOn {
		message := <-*node.inManagementChannel
		inManagementChannelReader(node, message)
	}
}

func inManagementChannelReader(node *dgNode, message ManagementMessage) {

	if message.parameter != nil && message.messageType != wantToDelete {
		var parameter = message.parameter.(Printer)
		PrintMessage(message, node, parameter)
	} else {
		PrintMessageWithoutSender(message, node)
	}

	switch message.messageType {
	case enterNewNode:
		if node.status == entering {
			if node.NextNodeInManagementChannel == nil {
				node.status = ready
				go Work(node)
			} else {
				*node.NextNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, node)
			}
		}
	case newNodeAppeared:
		newNode := message.parameter.(*dgNode)

		if node.request.isDependent(newNode.request) && node.status != leaving {

			newlist := append(*node.dependentsChannelsList, newNode.answersManagementChannel)
			node.dependentsChannelsList = &newlist
			*newNode.answersManagementChannel <- NewManagementMessage(hasConflictMessage, nil)
		}

		if node.NextNodeInManagementChannel == nil {
			*newNode.answersManagementChannel <- NewManagementMessage(endsConflictMessage, nil)
		} else {
			*node.NextNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, newNode)
		}

	case endFunc:
		node.status = leaving
		for _, e := range *node.dependentsChannelsList {
			*e <- NewManagementMessage(decreaseConflict, nil)
		}
		*node.GraphManagementChannel <- NewManagementMessage(leavingNode, node)

	case leavingNode:
		leavingNodeHandle(node, message)

	case wantToDelete:
		wantManagementChannel := message.parameter.(*chan ManagementMessage)
		*wantManagementChannel <- NewManagementMessage(wantToDelete, node)
	}
}

func leavingNodeHandle(node *dgNode, message ManagementMessage) {

	theLeavingNode := message.parameter.(*dgNode)

	if theLeavingNode == node {
		nodeShouldLeaveHandle(theLeavingNode)
	} else if node.IsNextNode(theLeavingNode) {
		isTheNextNodeLeavingHandle(node, theLeavingNode)
	} else {
		*node.NextNodeInManagementChannel <- NewManagementMessage(leavingNode, theLeavingNode)
	}
}

func nodeShouldLeaveHandle(node *dgNode) {
	node.isOn = false
	*node.answersManagementChannel <- NewManagementMessage(leavingNode, nil)
}

func isTheNextNodeLeavingHandle(node, theLeavingNode *dgNode) {
	*theLeavingNode.inManagementChannel <- NewManagementMessage(wantToDelete, node.leavingNodeAnswerChannel)
	message := <-*node.leavingNodeAnswerChannel
	if message.parameter != nil {
		nodeDelete := message.parameter.(*dgNode)
		node.NextNodeInManagementChannel = nodeDelete.NextNodeInManagementChannel
		*nodeDelete.inManagementChannel <- NewManagementMessage(leavingNode, nodeDelete)
	}
}

func (node *dgNode) answersManagementChannelReader() {
	for node.isOn {
		message := <-*node.answersManagementChannel
		answersManagementChannelReader(node, message)
	}

	if node.graph.lastNodeInManagementChannel == node.inManagementChannel {
		*node.graph.addAndDeleteChannel <- NewManagementMessage(leavingNode, node.NextNodeInManagementChannel)
	}

	close(*node.inManagementChannel)
	close(*node.answersManagementChannel)
	close(*node.leavingNodeAnswerChannel)

	cond.Signal()
	node.graph.length--
}

func answersManagementChannelReader(node *dgNode, message ManagementMessage) {

	if message.parameter != nil {
		var parameter = message.parameter.(Printer)
		PrintMessage(message, node, parameter)
	} else {
		PrintMessageWithoutSender(message, node)
	}

	switch message.messageType {
	case hasConflictMessage:
		atomic.AddUint64(&node.dependenciesNumber, 1)

	case endsConflictMessage:
		node.status = waiting
		if atomic.LoadUint64(&node.dependenciesNumber) == atomic.LoadUint64(&node.solvedDependenciesNumber) {
			node.status = ready
			go Work(node)
		}

	case decreaseConflict:
		atomic.AddUint64(&node.solvedDependenciesNumber, 1)
		if node.IsRedyToGo() {
			node.status = ready
			go Work(node)
		}

	case leavingNode:
		return
	}
}
//a