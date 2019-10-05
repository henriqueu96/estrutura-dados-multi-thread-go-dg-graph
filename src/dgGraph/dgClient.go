package dgGraph

type DGClient struct {
	MessagesNumber uint64
	inManagementChannel         *chan ManagementMessage
}

func NewDGClient() DGClient {
	chanIn := make(chan ManagementMessage, 30)
	return DGClient{
		MessagesNumber: 0,
		inManagementChannel:         &chanIn,
	}
}


func (client DGClient) Run(graph *dgGraph, preset []*DGRequest) {
	for i := range preset{
		request := preset[i]
		graph.add(request, client.inManagementChannel)
		i++
	}

}
