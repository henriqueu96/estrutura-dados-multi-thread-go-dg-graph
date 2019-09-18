package dgGraph

import (
	"fmt"
	"sync/atomic"
)

var myList = NewMyList()
var processNumber uint64 = 0;

func incrementProcessNumber(){
	atomic.AddUint64(&processNumber, 1);
}

func getProcessNumber() uint64{
	return atomic.LoadUint64(&processNumber)
}


func Work(request *DGRequest){
	request.Execute(&myList)
	incrementProcessNumber()
	fmt.Print("o loko bixo" )
}

