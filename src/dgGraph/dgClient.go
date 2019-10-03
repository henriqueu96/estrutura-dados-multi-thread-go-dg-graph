package dgGraph

type DGClient struct {
	MessagesNumber uint64
}

func NewDGClient() DGClient {
	return DGClient{
		MessagesNumber: 0,
	}
}


func (client DGClient) Run(graph *dgGraph, preset []*DGRequest) {
	for i := range preset{
		request := preset[i]
		graph.add(request)
		i++
	}

}
