package dgGraph

type Node struct {
	value int;
	next *Node;
}

func NewNode(value int) Node{
	return Node{
		value: value,
	}
}

func(node Node) getValue() int{
	return node.value
}

func(node Node) getNext() *Node{
	return node.next
}

func(node *Node) setNext(next *Node){
	node.next = next
}