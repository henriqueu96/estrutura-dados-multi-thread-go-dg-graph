package requests

type Worker struct {
	ProcessNumber int
}

func NewWorker() Worker {
	return Worker{
	}
}

func (worker Worker) Run(parallelizer *Parallelizer, myList *MyList, requestNumber *int) {
	for {
		request := parallelizer.NextRequest()
		request.Execute(myList)
		parallelizer.Remove(request)
		*requestNumber++
	}
}
