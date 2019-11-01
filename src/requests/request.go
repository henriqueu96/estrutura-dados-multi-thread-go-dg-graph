package requests

type Request struct {
	Value        int
	ExecState    ExecState
	RequestType  RequestType
	dependencies *[]*Request
	dependents   *[]*Request
}

func NewRequest(value int, RequestType RequestType) Request {
	return Request{
		Value:        value,
		ExecState:    Ready,
		RequestType:  RequestType,
		dependencies: []*Request{},
		dependents:   []*Request{},
	}
}

func (request *Request) addDependency(dependency *Request) {
	newList := append(*request.dependencies, dependency)
	request.dependencies = &newList
}

func (request *Request) removeDependency(dependency *Request) {
	request.dependencies = removeRequest(request.dependencies, dependency)
}

func removeRequest(requests *[]*Request, requestToRemove *Request) (requestsResult *[]*Request) {
	requestsResult = & []*Request{}
	for _, request := range *requests {
		var isToRemove = isEqual(request, requestToRemove)
		if !isToRemove {
			var obj = &request
			newList := append(*requestsResult, *obj)
			requestsResult = &newList
		}
	}
	return
}

func isEqual(request *Request, request2 *Request) bool {
	return request == request2
}

func (request Request) hasDependency() bool {
	if request.dependencies == nil{
		return false
	}
	return len(*request.dependencies) > 0
}

func (request *Request) addDependent(dependent *Request) {
	newList := append(*request.dependents, dependent)
	request.dependents = &newList
}

func (request *Request) hasDependent() bool {
	return len(*request.dependents) > 0;
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