package requests

type Parallelizer struct {
	queue *RequestQueue
	count int
}

func NewParallelizer(limit int) Parallelizer {
	queue := newRequestQueue(limit)
	return Parallelizer{
		queue: &queue,
		count: 0,
	}
}

func (parallelizer *Parallelizer) Add(value int, requestType RequestType) *Request {
	request := NewRequest(parallelizer.count, value, Ready, requestType)
	parallelizer.queue.add(&request)
	parallelizer.count++
	return &request
}

func (parallelizer *Parallelizer) Remove(request *Request) {
	parallelizer.queue.remove(request)
	parallelizer.count--
}

func (parallelizer *Parallelizer) NextRequest() *Request {
	return parallelizer.queue.nextRequest()
}

func (parallelizer *Parallelizer) clear() {
	parallelizer.queue.clear()
}

func (parallelizer *Parallelizer) requestCount() int {
	return parallelizer.queue.requestCount()
}
