package dgGraph

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

func (node *dgNode) start() {
	for {
		message := <-node.inManagementChannel
		switch message.messageType {
		case newNodeAppeared:
			// some code (async to maintain node receiving messages)....
		case incrementConflict:
			// some code (async to maintain node receiving messages)....
		}
	}
}

func newNode() dgNode {
	return dgNode{}
}
