package dgGraph

import (
	"fmt"
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
	// go printPopulation(graph)
	for i := range preset {
		request := preset[i]
		if graph.GetPopulationAdd(){
			*graph.AddChannel <- NewManagementMessage(enterNewNode, request)
			i++
		}
	}
}

func printPopulation(graph *dgGraph){
	for{
		fmt.Println(graph.Population)
		time.Sleep(time.Second)
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

				*graph.RemoveChannel <- NewManagementMessage(leavingNode, *theLeavingNode.NextNodeInManagementChannel)
			}
		} else {
			*graph.lastNodeInManagementChannel <- NewManagementMessage(leavingNode, theLeavingNode) // explodiu erro aqui - chan = nil
		}

	}

}
