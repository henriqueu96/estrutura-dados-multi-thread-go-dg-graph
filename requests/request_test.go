package requests

import "testing"

func newTestRequest() Request {
	return NewRequest(10, 10, Blocked, Read)
}

func TestRequest_AddDependency_ShouldAddAnDependencyToRequest(t *testing.T) {
	request := newTestRequest()
	dependency := newTestRequest()
	request.addDependency(&dependency)
	if(len(request.dependencies) == 0){
		t.Fail()
	}
}

func TestRequest_HasDependency_ShouldReturnFalse(t *testing.T){
	request := newTestRequest()

	if(request.hasDependency()){
		t.Fail()
	}
}

func TestRequest_HasDependency_ShouldReturnTrue(t *testing.T){
	request := newTestRequest()
	dependency := newTestRequest()
	request.addDependency(&dependency)

	if(!request.hasDependency()){
		t.Fail()
	}
}

func Test_isEqual_ShouldReturnTrue(t *testing.T) {
	var request = newTestRequest()
	var request2 = &request

	if(!isEqual(&request, request2)){
		t.Fail()
	}
}

func TestRequest_RemoveDependency_ShouldRemoveDependency(t *testing.T){
	request := newTestRequest()
	dependency := newTestRequest()

	request.addDependency(&dependency)
	request.removeDependency(&dependency)

	if(request.hasDependency()){
		t.Fail()
	}
}

func TestRequest_AddDependent_ShouldAddDependent(t *testing.T) {
	request := newTestRequest()
	dependent := newTestRequest()

	request.addDependent(&dependent)

	if(!request.hasDependent()){
		t.Fail()
	}
}

func TestRequest_RemoveDependent_ShouldRemoveDependent(t *testing.T) {
	request := newTestRequest()
	dependent := newTestRequest()

	request.addDependent(&dependent)
	request.removeDependent(&dependent)

	if(request.hasDependent()){
		t.Fail()
	}
}

func TestRequest_HasDependent_ShouldReturnFalse(t *testing.T){
	request := newTestRequest()

	if(request.hasDependent()){
		t.Fail()
	}
}

func TestRequest_HasDependent_ShouldReturnTrue(t *testing.T){
	request := newTestRequest()
	dependent := newTestRequest()

	request.addDependent(&dependent)

	if(!request.hasDependent()){
		t.Fail()
	}
}

func Test_RemoveRequest_ShouldRemoveRequest(t *testing.T){
	request := newTestRequest()
	list := []*Request{ &request }
	list = removeRequest(list, &request)
	if(len(list) > 0){
		t.Fail()
	}

}

func Test_RemoveRequest_ShouldNotRemoveRequest(t *testing.T){
	request := newTestRequest()
	request2 := newTestRequest()

	list := []*Request{ &request }

	list = removeRequest(list, &request2)

	if(len(list) == 0){
		t.Fail()
	}

}

func Test_IsDependent_ShouldNotReturnTrue_WhenNoOneOfRequestsIsWrite(t *testing.T) {
	request := newTestRequest()
	possibleDependent := newTestRequest()

	response := request.isDependent(&possibleDependent)

	if(response){
		t.Fail()
	}
}

func Test_IsDependent_ShouldReturnTrue_WhenOneOfRequestsIsWrite(t *testing.T) {
	request := newTestRequest()
	possibleDependent := newTestRequest()

	request.RequestType = Write

	response := request.isDependent(&possibleDependent)

	if(!response){
		t.Fail()
	}
}