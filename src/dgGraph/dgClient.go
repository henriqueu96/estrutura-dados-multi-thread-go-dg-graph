package dgGraph

type DGClient struct {
	MessagesNumber      uint64
}

func NewDGClient() DGClient {
	return DGClient{
		MessagesNumber:      0,
	}
}

func (client DGClient) Run(graph *dgGraph, preset []*DGRequest) {
	go graph.Start()

	for _, request := range preset {
		*graph.addAndUpdateLastChannel <- NewManagementMessage(AddRequest, request)
	}
}

func (client DGClient) toString() string{
	return "Client";
}