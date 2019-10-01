package requests

type Node struct {
	value int;
}

func NewNode(value int) Node{
	return Node{
		value: value,
	}
}

func(node Node) getValue() int{
	return node.value
}