package requests

type MyList struct {
	firstNode *Node
}

func NewMyList() MyList{
	return MyList{
	}
}

func (myList *MyList) add(value int) (exists bool) {
	node := newNode(value)
	if (myList.firstNode == nil) {
		myList.firstNode = &node
	} else{
		lastNode := myList.getLastNode()
		lastNode.next = &node
	}

	return
}

func (myList *MyList) getLastNode() (lastNode *Node){
	lastNode = myList.firstNode
	for true  {
		if(lastNode.getNext() != nil){
			lastNode = lastNode.next
		} else{
			break
		}
	}
	return
}

func (myList MyList) get(value int) (exists bool) {

	if myList.firstNode == nil{
		return
	}

	var lastNode = myList.firstNode

	for{
		if (lastNode != nil) {
			if(lastNode.getValue() == value){
				exists = true
				return
			}
			lastNode = lastNode.getNext()
		} else{
			break
		}
	}
	return
}
