package dgGraph

type DGRequest struct {
	Value       int
	requestType requestType
}

func NewRequest(value int, requestType requestType) DGRequest {
	return DGRequest{
		Value:       value,
		requestType: requestType,
	}
}

func (request *DGRequest) isDependent(possibleDependent *DGRequest) bool {
	return request.requestType == Write || possibleDependent.requestType == Write
}

func (request *DGRequest) Execute(myList *MyList) bool {
	if request.requestType == Write {
		return myList.add(request.Value)
	} else {
		return myList.get(request.Value)
	}
}
