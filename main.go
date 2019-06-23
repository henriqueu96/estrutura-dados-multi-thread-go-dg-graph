package main

import (
	"fmt"
	"math/rand"
	"tcc/boolGenerator"
	"tcc/requests"
	"time"
)

var mylist = requests.NewMyList()

func main(){
	fmt.Println("Inicio")
	parallelizer := requests.NewParallelizer(5)

	var final = make(chan int)

	go consumer(&parallelizer)
	go producer(&parallelizer)

	<- final
}

func consumer(parallelizer *requests.Parallelizer) {
	for {
		request := parallelizer.NextRequest()
		find := request.Execute(&mylist)
		message := ""
		if(find){
			message = "found!"
		} else {
			message = "not found!"
		}
		fmt.Println(message)
		parallelizer.Remove(request)
	}
}

func producer(parallelizer *requests.Parallelizer) {
	for {
		isWrite := boolGenerator.NewByPercent(time.Now(), 10)
		requestType := requests.Read
		if isWrite {
			requestType = requests.Write
		}
		parallelizer.Add(rand.Intn(5), requestType)
	}
}