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
	dependentsChannelsList      *[]*chan ManagementMessage
	inManagementChannel         *chan ManagementMessage
	NextNodeInManagementChannel *chan ManagementMessage
}

func newNode(request *DGRequest, nextNodeInManagementChannel *chan ManagementMessage) dgNode {
	idIndex++;
	chanIn := make(chan ManagementMessage, 30)

	return dgNode{
		id:                          idIndex,
		request:                     request,
		dependenciesNumber:          0,
		solvedDependenciesNumber:    0,
		dependentsChannelsList:      &[]*chan ManagementMessage{},
		inManagementChannel:         &chanIn,
		NextNodeInManagementChannel: nextNodeInManagementChannel,
		status:                      entering,
	}
}

func (node *dgNode) ToString() string {
	return "(Node) Id:" + strconv.FormatUint(node.id, 10)
}

func (node *dgNode) start() error {
	for {
		message := <-*node.inManagementChannel
		newMethod(node, message)
	}
}

func newMethod(node *dgNode, message ManagementMessage) {

	/*	messageType := MessageTypes[message.messageType]
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

			newlist := append(*node.dependentsChannelsList, newNode.inManagementChannel)
			node.dependentsChannelsList = &newlist
			*newNode.inManagementChannel <- NewManagementMessage(hasConflictMessage, nil)
		}

		if node.NextNodeInManagementChannel == nil {
			*newNode.inManagementChannel <- NewManagementMessage(endsConflictMessage, nil)
		} else {
			*node.NextNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, newNode)
		}

		/*
			if DELETED nao
			deve
			fazer
			nada ...
		*/
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
			fmt.Println("decrease go work")
		}

	case endFunc:
		node.status = leaving

		/*depNubmer := len(*node.dependentsChannelsList)
		fmt.Println("Event Free: " + strconv.Itoa(depNubmer))*/
		for _, e := range *node.dependentsChannelsList {
			*e <- NewManagementMessage(decreaseConflict, nil)
		}
	}
}

func (node *dgNode) IsRedyToGo() bool {
	return node.status == waiting && node.dependenciesNumber == node.solvedDependenciesNumber
}
