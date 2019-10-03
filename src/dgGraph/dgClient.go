package dgGraph

type DGClient struct {
	MessagesNumber uint64
}

func NewDGClient() DGClient {
	return DGClient{
		MessagesNumber: 0,
	}
}

func (client DGClient) Run(parallelizer *dgGraph, preset []*DGRequest) {
	for i := range preset{
		request := preset[i]
		parallelizer.add(request)
		i++
	}

}
