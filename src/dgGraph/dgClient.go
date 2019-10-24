package dgGraph

type DGClient struct {
	MessagesNumber      uint64
	inManagementChannel *chan ManagementMessage
}

func NewDGClient() DGClient {
	chanIn := make(chan ManagementMessage, 30)
	return DGClient{
		MessagesNumber:      0,
		inManagementChannel: &chanIn,
	}
}

func (client DGClient) Run(graph *dgGraph, preset []*DGRequest) {
	go ReaderChan(&client, graph)
	for i := range preset {
		request := preset[i]
		graph.add(request, client.inManagementChannel)
		i++
	}
}

func ReaderChan(client *DGClient, graph *dgGraph) {
	for {
		message := <-*client.inManagementChannel
		newNode := message.parameter.(*dgNode)
		switch message.messageType {
		case leavingNode:
			*graph.lastNodeInManagementChannel <- NewManagementMessage(leavingNode, newNode)
		}
	}
}