package requests

type Request struct {
	Id           int
	Value        int
	ExecState    ExecState
	RequestType  RequestType
	dependencies []*Request
	dependents   []*Request
}

func NewRequest(id int, value int, state ExecState, requestType RequestType) Request {
	return Request{
		Id:           id,
		Value:        value,
		ExecState:    state,
		RequestType:  requestType,
		dependencies: []*Request{},
		dependents:   []*Request{},
	}
}

func (request *Request) addDependency(dependency *Request) {
	request.dependencies = append(request.dependencies, dependency)
}

func (request *Request) removeDependency(dependency *Request) {
	request.dependencies = removeRequest(request.dependencies, dependency)
}

func removeRequest(requests []*Request, requestToRemove *Request) (requestsResult []*Request) {
	requestsResult = []*Request{}
	for _, request := range requests {
		var isToRemove = isEqual(request, requestToRemove)
		if !isToRemove {
			var obj = &request
			requestsResult = append(requestsResult, *obj)
		}
	}
	return
}

func isEqual(request *Request, request2 *Request) bool {
	return request == request2
}

func (request Request) hasDependency() bool {
	return len(request.dependencies) > 0
}

func (request *Request) addDependent(dependent *Request) {
	request.dependents = append(request.dependents, dependent)
}

func (request *Request) hasDependent() bool {
	return len(request.dependents) > 0;
}

func (request *Request) removeDependent(dependent *Request) {
	request.dependents = removeRequest(request.dependents, dependent)
}

func (request *Request) Execute(myList *MyList) bool{
	if(request.RequestType == Write){
		return myList.add(request.Value)
	} else{
		return myList.get(request.Value)
	}
}

func (request *Request) isDependent(possibleDependent *Request) bool {
	return request.RequestType == Write || possibleDependent.RequestType == Write
}