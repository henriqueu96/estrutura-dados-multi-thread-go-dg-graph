package dgGraph

import (
	"fmt"
	"strconv"
	"sync"
)

var idIndex uint64 = 0;

type dgNode struct {
	id                          uint64
	status                      dgNodeStatus
	request                     *DGRequest
	dependenciesNumber          int
	solvedDependenciesNumber    int
	dependentsChannelsList      *[]*chan ManagementMessage
	inManagementChannel         *chan ManagementMessage
	outManagementChannel        *chan ManagementMessage
	WantManagementChannel       *chan ManagementMessage
	NextNodeInManagementChannel *chan ManagementMessage
	ClientManagementChannel     *chan ManagementMessage
	ShouldContinue              bool
	graph                       *dgGraph
}

func newNode(request *DGRequest, nextNodeInManagementChannel *chan ManagementMessage, clientManagementChannel *chan ManagementMessage, graph *dgGraph) dgNode {
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
		outManagementChannel:        &chanOut,
		NextNodeInManagementChannel: nextNodeInManagementChannel,
		status:                      entering,
		ClientManagementChannel:     clientManagementChannel,
		ShouldContinue:              true,
		graph:                       graph,
		WantManagementChannel:       &chanWant,
	}
}

func (node *dgNode) ToString() string {
	return "(Node) Id:" + strconv.FormatUint(node.id, 10)
}

func (node *dgNode) start() {
	go node.StartIn()
	go node.StartOut()
}

func newMethodOut(node *dgNode, message ManagementMessage) {
	if(GetPrint()){
		messageType := MessageTypes[message.messageType]
		fmt.Println("Event:" + messageType + " " + node.ToString())
	}
	switch message.messageType {
	case hasConflictMessage:
		node.dependenciesNumber++

	case endsConflictMessage:
		node.status = waiting
		if node.dependenciesNumber == node.solvedDependenciesNumber {
			node.status = ready
			go Work(node)
		}

	case decreaseConflict:
		node.solvedDependenciesNumber++
		if node.IsRedyToGo() {
			node.status = ready
			go Work(node)
		}

	case leavingNode:
		return
	}
}

var addRemoveSequencialy = sync.Mutex{}

func newMethodIn(node *dgNode, message ManagementMessage) {

	if(GetPrint()){
		messageType := MessageTypes[message.messageType]
		fmt.Println("Event:" + messageType + " " + node.ToString())
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

			newlist := append(*node.dependentsChannelsList, newNode.outManagementChannel)
			node.dependentsChannelsList = &newlist
			*newNode.outManagementChannel <- NewManagementMessage(hasConflictMessage, nil)
		}

		if node.NextNodeInManagementChannel == nil {
			*newNode.outManagementChannel <- NewManagementMessage(endsConflictMessage, nil)
		} else {
			*node.NextNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, newNode)
		}

	case endFunc:
		node.status = leaving
		for _, e := range *node.dependentsChannelsList {
			*e <- NewManagementMessage(decreaseConflict, nil)
		}
		*node.ClientManagementChannel <- NewManagementMessage(leavingNode, node)

	case leavingNode:
		theLeavingNode := message.parameter.(*dgNode)
		if theLeavingNode == node {
			node.ShouldContinue = false
			*node.outManagementChannel <- NewManagementMessage(leavingNode, nil)
			return
		}
		if node.IsNextNode(theLeavingNode) {
			*node.NextNodeInManagementChannel <- NewManagementMessage(wantToDelete, node.WantManagementChannel)
			for node.ShouldContinue {
				message := <- *node.WantManagementChannel
				if message.parameter != nil{
					nodeDelete := message.parameter.(*dgNode)
					node.NextNodeInManagementChannel = nodeDelete.NextNodeInManagementChannel
					*nodeDelete.inManagementChannel <- NewManagementMessage(leavingNode, nodeDelete)
				}
				return
			}
		} else {
			if (node.NextNodeInManagementChannel != nil) {
				*node.NextNodeInManagementChannel <- NewManagementMessage(leavingNode, theLeavingNode)
			}
		}
	case wantToDelete:
		wantManagementChannel := message.parameter.(*chan ManagementMessage)
		*wantManagementChannel <- NewManagementMessage(wantToDelete, node)
	}

}

func (node *dgNode) IsNextNode(nextNode *dgNode) bool {
	return node.NextNodeInManagementChannel == nextNode.inManagementChannel
}

func (node *dgNode) IsRedyToGo() bool {
	return node.status == waiting && node.dependenciesNumber == node.solvedDependenciesNumber
}

func (node *dgNode) StartIn() {
	for node.ShouldContinue {
		message := <-*node.inManagementChannel
		newMethodIn(node, message)
	}
}

func (node *dgNode) StartOut() {
	for node.ShouldContinue {
		message := <-*node.outManagementChannel
		newMethodOut(node, message)
	}

	//if node.graph.lastNodeInManagementChannel == node.inManagementChannel {
	//	*node.graph.RemoveChannel <- NewManagementMessage(leavingNode, node.NextNodeInManagementChannel)
	//}
	if(GetPrint()){
	fmt.Println("Leaving" + node.ToString())}
	incrementProcessNumber()
}
