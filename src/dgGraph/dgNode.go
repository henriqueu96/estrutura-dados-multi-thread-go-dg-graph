package dgGraph

import "errors"

type dgNode struct {
	status                   dgNodeStatus
	request                  *DGRequest
	dependenciesNumber       int
	solvedDependenciesNumber int
	// used to free the dependents
	dependentsChannelsList      []chan ManagementMessage
	inManagementChannel         chan ManagementMessage
	NextNodeInManagementChannel chan ManagementMessage
}

func (node *dgNode) start() error {
	for {
		message := <-node.inManagementChannel
		switch message.messageType {

		case enterNewNode:
			if node.status == entering {
				if node.NextNodeInManagementChannel == nil {
					node.status = ready
					go Work(node.request)
				} else {
					node.NextNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, &node)
				}

			} else {
				return errors.New("Status incorreto para mensagem enterNewNode")
			}

		case newNodeAppeared:
			newNode := message.parameter.(*dgNode)
			if node.status != (entering | waiting | ready) {
				return errors.New("Status incorreto para mensagem newNodeAppeared")
			} else {
				if node.request.isDependent(newNode.request) {
					node.dependentsChannelsList = append(node.dependentsChannelsList, newNode.inManagementChannel)
					newNode.inManagementChannel <- NewManagementMessage(hasConflictMessage, nil)
				}
			}

			if node.status != (entering | waiting | ready | leaving) {
				return errors.New("Status incorreto para mensagem newNodeAppeared")
			} else {
				if node.NextNodeInManagementChannel == nil {
					newNode.inManagementChannel <- NewManagementMessage(endsConflictMessage, nil)
				} else {
					node.NextNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, &node)
				}
			}
			/*
				if DELETED nao
				deve
				fazer
				nada ...
			*/
		case hasConflictMessage:
			if node.status != entering {
				return errors.New("Status incorreto para mensagem haConflictMessage")
			} else {
				node.dependenciesNumber++
			}

		case endsConflictMessage:
			if node.status != entering {
				return errors.New("Status incorreto para mensagem endsConflictMessage ")
			}
			node.status = waiting
			if node.dependenciesNumber == node.solvedDependenciesNumber {
				node.status = ready
				go Work(node.request)
			}

		case decreaseConflict:
			node.solvedDependenciesNumber++
			if node.status == waiting && node.dependenciesNumber == node.solvedDependenciesNumber {
				node.status = ready
				go Work(node.request)
			}

		case endFunc:
			node.status = leaving
			for _, e := range node.dependentsChannelsList {
				e <- NewManagementMessage(decreaseConflict, nil)
			}

		}
	}
}

func newNode(request *DGRequest) dgNode {
	return dgNode{
		request:                     request,
		dependenciesNumber:          0,
		solvedDependenciesNumber:    0,
		dependentsChannelsList:      []chan ManagementMessage{},
		inManagementChannel:         make(chan ManagementMessage),
		NextNodeInManagementChannel: nil,

	}
}
