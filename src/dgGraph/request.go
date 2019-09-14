package dgGraph

type request struct {
	Value       int
	requestType requestType
}

func NewRequest(value int, requestType requestType) request {
	return request{
		Value:       value,
		requestType: requestType,
	}
}

func (request *request) isDependent(possibleDependent *request) bool {
	return request.requestType == Write || possibleDependent.requestType == Write
}

// func (request *request) isDependent(request1 *request) bool {
// 	if request1.Value >= request.Conflict {
// 		return false
// 	} else {
// 		return true
// 	}
// }

func (request *request) Execute(myList *MyList) bool {
	if request.requestType == Write {
		return myList.add(request.Value)
	} else {
		return myList.get(request.Value)
	}
}
