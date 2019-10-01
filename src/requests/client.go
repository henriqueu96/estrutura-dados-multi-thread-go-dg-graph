package requests

import "sync/atomic"

type Client struct {
	messagesNumber uint64
}

func NewClient() Client{
	return Client{
		messagesNumber: 0,
	}
}

func (client *Client) Run(parallelizer *Parallelizer, preset []*Request) {
	for _, message := range preset {
		parallelizer.Add(message)
		client.incrementMessagesNumber()
	}
}

func (client *Client) GetMessagesNumber() uint64{
	return atomic.LoadUint64(&client.messagesNumber)
}

func (client *Client) incrementMessagesNumber(){
	atomic.AddUint64(&client.messagesNumber, 1)
}