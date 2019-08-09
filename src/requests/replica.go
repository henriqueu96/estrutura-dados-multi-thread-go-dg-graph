package requests

type Replica struct {
	paralizer *Parallelizer
}

func NewReplica(parallizer *Parallelizer) Replica{
	return Replica{
		paralizer: parallizer,
	}
}

func (replica *Replica) Run(){
	var request *Request;
	for{
		request = replica.paralizer.NextRequest()
		replica.paralizer.Remove(request)
	}
}
