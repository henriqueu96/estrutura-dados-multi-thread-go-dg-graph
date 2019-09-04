package dgGraph

import "errors"

type dgNode struct {
	status                   dgNodeStatus
	request                  *request
	dependenciesNumber       int
	solvedDependenciesNumber int
	// used to free the dependents
	dependentsChannelsList      [] chan interface{}
	inManagementChannel         chan ManagementMessage
	NextNodeInManagementChannel chan ManagementMessage
}

func (node *dgNode) start() error {
	for {
		message := <-node.inManagementChannel
		switch message.messageType {
		case enterNewNode:
			// todo: fazer a mensagem de novo node chegando
		case newNodeAppeared:
			if node.status == entering {
				if node.NextNodeInManagementChannel == nil {
					node.status = ready
					go node.start()
				} else {
					node.NextNodeInManagementChannel <- NewManagementMessage(newNodeAppeared, node)
					// todo: quando receber o nodo, pegar o request e ver se conflita, após, atribuir o next chanel para o  node.inManagementChannel
				}

			} else {
				return errors.New("") // todo: definir texto do erro
			}
		case hasConflictMessage:
			if node.status != entering {
				return errors.New("") // todo: definir texto do erro
			} else {
				node.dependenciesNumber++
			}
		case endsConflictMessage:
			if node.status != entering {
				return errors.New("") // todo: definir texto do erro
			}
			node.status = waiting

			if node.dependenciesNumber == node.solvedDependenciesNumber {
				node.status = ready
				go node.start()
			}
		case decreaseConflict:
			node.solvedDependenciesNumber++ // mais uma dep resolvida
			if node.status == waiting && node.dependenciesNumber == node.solvedDependenciesNumber { // era a ultima ?
				node.status = ready
				go node.start()
			}
		case endFunc:
			node.status = leaving // marca como acabou de executar, saindo */
			for _, e := range node.dependentsChannelsList {
				e <- decreaseConflict
			}

			//mchi <- freeMSG -- todo: perguntar dotti
		}
	}
}

func newNode() dgNode {
	return dgNode{}
}
