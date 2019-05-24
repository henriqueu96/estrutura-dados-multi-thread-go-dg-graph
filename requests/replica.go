package requests

import (
	"fmt"
	"strconv"
)

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

		message := request.Name + "" + strconv.Itoa(request.Id)
		fmt.Println(message)

		replica.paralizer.Remove(request)
	}
}
