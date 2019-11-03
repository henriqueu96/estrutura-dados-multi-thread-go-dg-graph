package dgGraph

import (
	"fmt"
	"sync/atomic"
)

var myList = NewMyList()
var processNumber uint64 = 0;

func incrementProcessNumber() {
	atomic.AddUint64(&processNumber, 1);
}

func GetProcessNumber() uint64 {
	return atomic.LoadUint64(&processNumber)
}

func Work(node *dgNode) {
	request := node.request
	request.Execute(&myList)
	incrementProcessNumber()
	fmt.Println(node.id)
	*node.inManagementChannel <- NewManagementMessage(endFunc, nil);
}

