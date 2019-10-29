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
	go printPopulation(graph)
	for i := range preset {
		request := preset[i]
		if graph.GetPopulationAdd() {
			*graph.AddChannel <- NewManagementMessage(enterNewNode, request)
			i++
		}
	}
}

func printPopulation(graph *dgGraph) {
	for {
		fmt.Println(graph.Population)
		time.Sleep(time.Second)
	}

}
func ReaderChan(client *DGClient, graph *dgGraph) {
	for {
		message := <-*client.inManagementChannel
		ExitVerification(message, graph)
	}
}

/*
	- Cliente se o leavingNode é o próximo nodo
	- Caso não seja, só passa a mensagem pra frente

*/
func ExitVerification(message ManagementMessage, graph *dgGraph) {
	theLeavingNode := message.parameter.(*dgNode)
	switch message.messageType {
	case leavingNode:
		if graph.lastNodeInManagementChannel == theLeavingNode.inManagementChannel {
			*theLeavingNode.inManagementChannel <- NewManagementMessage(wantToDelete, graph.WantDeleteChannel)
			WaitingForWantToDeleteResponse(graph)
		} else {
			// NÃO PODE SER NULL - se ele não é o próximo, então ele tem que ter um anterior
			*graph.lastNodeInManagementChannel <- NewManagementMessage(leavingNode, theLeavingNode)
		}
	}
}

func WaitingForWantToDeleteResponse(graph *dgGraph) {
	message := <-*graph.WantDeleteChannel
	nodeThatWantDelete := message.parameter.(*dgNode)
	*graph.UpdateLastNodeInManagementChannel <- NewManagementMessage(irrelevant, nodeThatWantDelete.NextNodeInManagementChannel)
	*nodeThatWantDelete.inManagementChannel <- NewManagementMessage(leavingNode, nodeThatWantDelete)
}
