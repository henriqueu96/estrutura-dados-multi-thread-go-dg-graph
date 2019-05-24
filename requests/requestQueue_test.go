package requests

import "testing"

func newTestQueue() RequestQueue {
	return newRequestQueue(10)
}

func TestRequestQueue_hasRequest_shouldReturnFalse(t *testing.T) {
	queue := newTestQueue()
	if (queue.hasRequest()) {
		t.Fail()
	}
}

func TestRequestQueue_hasRequest_shouldReturnTrue(t *testing.T) {
	queue := newTestQueue()
	request := newTestRequest()

	queue.add(&request)

	if (!queue.hasRequest()) {
		t.Fail()
	}
}

func TestRequestQueue_add_ShouldInsertRequest(t *testing.T) {
	queue := newTestQueue()
	request := newTestRequest()

	queue.add(&request)

	if (!queue.hasRequest()) {
		t.Fail()
	}
}

func TestRequestQueue_Remove_ShouldRemoveRequest(t *testing.T) {
	queue := newTestQueue()
	request := newTestRequest()

	queue.add(&request)
	queue.remove(&request)

	if (queue.hasRequest()) {
		t.Fail()
	}
}

func TestRequestQueue_nextQueue_shouldRetunrTheFirstReady(t *testing.T) {
	request := newTestRequest()
	request2 := newTestRequest()
	request2.ExecState = Ready

	queue := newTestQueue()
	queue.add(&request)
	queue.add(&request2)

	result := queue.nextRequest();
	if (!isEqual(result, &request2)) {
		t.Fail()
	}
}

func TestRequestQueue_nextQueue_shouldRetunNilWhenNoOneIsReady(t *testing.T) {
	request := newTestRequest()

	queue := newTestQueue()
	queue.add(&request)

	result := queue.nextRequest();
	if (result != nil) {
		t.Fail()
	}
}

func TestRequestQueue_clear_shouldClearTheList(t *testing.T) {
	request := newTestRequest()
	request2 := newTestRequest()

	queue := newTestQueue()
	queue.add(&request)
	queue.add(&request2)

	queue.clear()

	if (len(queue.pendingRequests) > 0) {
		t.Fail()
	}
}

func TestRequestQueue_count_shouldReturn0(t *testing.T) {
	queue := newTestQueue()

	if (queue.requestCount() != 0) {
		t.Fail()
	}
}
