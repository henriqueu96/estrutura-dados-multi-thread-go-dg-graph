package requests

type Worker struct {
	ProcessNumber int
}

func NewWorker() Worker {
	return Worker{
		ProcessNumber: 0,
	}
}

func (worker *Worker) Run(parallelizer *Parallelizer, myList *MyList) {
	for {
		request := parallelizer.NextRequest()
		request.Execute(myList)
		parallelizer.Remove(request)
		worker.ProcessNumber++
	}
}
