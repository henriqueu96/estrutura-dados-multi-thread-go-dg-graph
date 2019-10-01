package requests

type MyList struct {
	nodes *[]*Node
}

func NewMyList() MyList{
	return MyList{
	}
}

func (myList *MyList) add(value int) bool {
	node := NewNode(value)
	newList := append(*myList.nodes, &node)
	myList.nodes = &newList

	for _, node := range *myList.nodes{
		if node.value == value {
			return true
		}
	}
	return false
}

func (myList *MyList) get(value int)  bool {
	for _, currentNode := range *myList.nodes{
		if currentNode.value == value{
			myList.nodes = removeNode(myList.nodes, currentNode)
			return true
		}
	}
	return false
}

func removeNode(nodes *[]*Node, node *Node) *[]*Node {
	list := []*Node{}
	for _, currentNode := range *nodes{
		if currentNode != node {
			list = append(list, currentNode)
		}
	}
	return &list
}
