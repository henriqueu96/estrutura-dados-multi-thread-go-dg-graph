package main

import (
	"fmt"
	"strconv"
	"tcc/requests"
)

func main() {
	parallelizer := requests.NewParallelizer(10)

	for i := 200; i > 0; i-- {
		curr := strconv.Itoa(i)
		parallelizer.Add("Request: " + curr)
		fmt.Println("Entrou")
	}
	fmt.Println("O Loko bixo")
}
