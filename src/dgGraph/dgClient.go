package dgGraph

import (
	"time"
)

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

	go graph.Start(&client)
	go ReaderChan(&client, graph)
	for i := range preset {
		request := preset[i]
		time.Sleep(time.Nanosecond * 5)
		*graph.AddChannel <- NewManagementMessage(enterNewNode, request)
		i++
	}
}

func ReaderChan(client *DGClient, graph *dgGraph) {
	for {
		message := <-*client.inManagementChannel
		verificacaoSaida(message, graph)
	}
}
func verificacaoSaida(message ManagementMessage, graph *dgGraph) {
	theLeavingNode := message.parameter.(*dgNode)
	switch message.messageType {
	case leavingNode:
		if (graph.lastNodeInManagementChannel == theLeavingNode.inManagementChannel) {
			*theLeavingNode.inManagementChannel <- NewManagementMessage(wantToDelete, graph.WantDeleteChannel)
			for {
				message := <-*graph.WantDeleteChannel
				if message.parameter != nil {
					nodeDelete := message.parameter.(*dgNode)
					graph.lastNodeInManagementChannel = nodeDelete.NextNodeInManagementChannel
					*nodeDelete.inManagementChannel <- NewManagementMessage(leavingNode, nodeDelete)
					return
				}

				*graph.RemoveChannel <- NewManagementMessage(leavingNode, theLeavingNode.NextNodeInManagementChannel)
			}
		} else {

			*graph.lastNodeInManagementChannel <- NewManagementMessage(leavingNode, theLeavingNode) // explodiu erro aqui - chan = nil
		}

	}

}
