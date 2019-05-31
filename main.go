package main

import (
	"fmt"
	"strconv"
	"tcc/requests"
)

func main() {
	parallelizer := requests.NewParallelizer(100)

	fmt.Println("Inicio")

	go consumer(parallelizer)
	go consumer(parallelizer)
	go consumer(parallelizer)

	go producer(parallelizer)
	blq := make(chan int)
	<-blq
}

func consumer(parallelizer requests.Parallelizer) {
	for {
		request := parallelizer.NextRequest()
		fmt.Println("Request " + request.Name)
		parallelizer.Remove(request)
	}
}

func producer(parallelizer requests.Parallelizer) {
	i := 0
	for {
		curr := strconv.Itoa(i)
		parallelizer.Add("Request: " + curr)
		i++
	}
}
