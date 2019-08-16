package requests

type Client struct {
	MessagesNumber int
}

func NewClient() Client{
	return Client{
		MessagesNumber: 0,
	}
}

func (client Client) Run(parallelizer *Parallelizer, preset []*Request) {
	for {
		i:= 0
		request := preset[i]
		parallelizer.Add(request)
		i++
	}
}