package dgGraph

type dgClient struct {
	MessagesNumber int
}

func NewDGClient() dgClient{
	return dgClient{
		MessagesNumber: 0,
	}
}

func (client dgClient) Run(parallelizer *dgGraph, preset []*DGRequest) {
	for {
		i:= 0
		request := preset[i]
		parallelizer.add(request)
		i++
	}
}