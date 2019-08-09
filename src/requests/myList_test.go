package requests

import "testing"

func newTestMyList()  MyList{
	return NewMyList()
}

func newTestNode(value int) Node{
	return Node{
		value: value,
	}
}

func TestMyList_add_should_set_first_node(t *testing.T) {
	myList := newTestMyList()
	myList.add(1)
	if myList.firstNode == nil {
		t.Fail()
	}
}

func TestMyList_getLastNode_should_return_first_node(t *testing.T) {
	myList := newTestMyList()
	myList.add(1)

	if myList.firstNode.value != 1 {
		t.Fail()
	}
}

func TestMyList_getLastNode_should_return_secound_node(t *testing.T) {
	myList := newTestMyList()
	myList.add(1)
	nodeTwo := newTestNode(2)
	myList.firstNode.setNext(&nodeTwo)

	if myList.firstNode.next.value != 2 {
		t.Fail()
	}
}

func TestMyList_getLastNode_should_return_third_node(t *testing.T) {
	myList := newTestMyList()
	myList.add(1)
	nodeTwo := newTestNode(2)
	nodeThree := newTestNode(3)
	nodeTwo.next = &nodeThree
	myList.firstNode.setNext(&nodeTwo)


	if myList.firstNode.next.next.value != 3 {
		t.Fail()
	}
}

func TestMyList_add_should_set_secound_node(t *testing.T) {
	myList := newTestMyList()
	myList.add(1)
	myList.add(2)

	if myList.firstNode.getNext() == nil {
		t.Fail()
	}
}

func TestMyList_add(t *testing.T) {
	myList := newTestMyList()
	myList.add(10)
	myList.add(5)

	ten := myList.get(10)
	five := myList.get(5)

	if !ten || !five {
		t.Fail()
	}
}