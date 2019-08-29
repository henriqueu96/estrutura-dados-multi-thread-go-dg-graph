package requests

type Client struct {
	MessagesNumber int
}

func NewClient() Client{
	return Client{
		MessagesNumber: 0,
	}
}

func (client *Client) Run(parallelizer *Parallelizer, preset []*Request) {

	for _, message := range preset {
		parallelizer.Add(message)
		client.MessagesNumber++
	}
}