package dgGraph

import (
	"fmt"
	"strconv"
)

var idIndex uint64 = 0;

type dgNode struct {
	id                       uint64
	status                   dgNodeStatus
	request                  *DGRequest
	dependenciesNumber       int
	solvedDependenciesNumber int
	// used to free the dependents
	dependentsChannelsList       *[]*chan ManagementMessage
	inManagementChannel          *chan ManagementMessage
	outManagementChannel         *chan ManagementMessage
	NextNodeInManagementChannel  *chan ManagementMessage
	NextNodeOutManagementChannel *chan ManagementMessage
	ClientManagementChannel      *chan ManagementMessage
}

func newNode(request *DGRequest, nextNodeInManagementChannel *chan ManagementMessage, nextNodeOutManagementChannel *chan ManagementMessage, clientManagementChannel *chan ManagementMessage) dgNode {
	idIndex++;
	chanIn := make(chan ManagementMessage, 30)
	chanOut := make(chan ManagementMessage, 30)
	return dgNode{
		id:                           idIndex,
		request:                      request,
		dependenciesNumber:           0,
		solvedDependenciesNumber:     0,
		dependentsChannelsList:       &[]*chan ManagementMessage{},
		inManagementChannel:          &chanIn,
		outManagementChannel:         &chanOut,
		NextNodeInManagementChannel:  nextNodeInManagementChannel,
		NextNodeOutManagementChannel: nextNodeOutManagementChannel,
		status:                       entering,
		ClientManagementChannel:      clientManagementChannel,
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

	/*messageType := MessageTypes[message.messageType]
	fmt.Println("Event:" + messageType + " " + node.ToString())*/
	switch message.messageType {
	case hasConflictMessage:
		node.dependenciesNumber++

	case endsConflictMessage:
		node.status = waiting
		if node.dependenciesNumber == node.solvedDependenciesNumber {
			node.status = ready
			//fmt.Println("finish wainting, go work")
			go Work(node)
		}

	case decreaseConflict:
		node.solvedDependenciesNumber++
		if node.IsRedyToGo() {
			node.status = ready
			go Work(node)
			//fmt.Println("decrease go work")
		}

	}
}

func newMethodIn(node *dgNode, message ManagementMessage) {

	/*messageType := MessageTypes[message.messageType]
	fmt.Println("Event:" + messageType + " " + node.ToString())*/
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

		/*
			if DELETED nao
			deve
			fazer
			nada ...
		*/

	case endFunc:
		node.status = leaving

		/*depNubmer := len(*node.dependentsChannelsList)
		fmt.Println("Event Free: " + strconv.Itoa(depNubmer))*/
		for _, e := range *node.dependentsChannelsList {
			*e <- NewManagementMessage(decreaseConflict, nil)
		}
		fmt.Println(node.id)
	}
}

func (node *dgNode) IsRedyToGo() bool {
	return node.status == waiting && node.dependenciesNumber == node.solvedDependenciesNumber
}

func (node *dgNode) StartIn() {
	for {
		message := <-*node.inManagementChannel
		newMethodIn(node, message)
	}
}

func (node *dgNode) StartOut() {
	for {
		message := <-*node.outManagementChannel
		newMethodOut(node, message)
	}
}
