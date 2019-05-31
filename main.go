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
	producer(parallelizer)

	fmt.Println("Fim")
}

func consumer(parallelizer requests.Parallelizer) {
	for {
		request := parallelizer.NextRequest()
		fmt.Println("Request " + request.Name)
		parallelizer.Remove(request)
	}
}

func producer(parallelizer requests.Parallelizer) {
	for i := 0; i < 200; i++ {
		curr := strconv.Itoa(i)
		parallelizer.Add("Request: " + curr)
	}
}
