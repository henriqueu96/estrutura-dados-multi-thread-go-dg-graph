package requests

import "sync/atomic"

type Worker struct {
	Id int
	ProcessNumber uint64
}

func NewWorker() Worker {
	return Worker{
		ProcessNumber: 0,
	}
}

func (worker *Worker) Run(parallelizer *Parallelizer, myList *MyList) {
	for {
		request := parallelizer.NextRequest()
		//request.Execute(myList)
		parallelizer.Remove(request)
		atomic.AddUint64(&worker.ProcessNumber, 1)
	}
}
