package main

import (
	"dgGraph"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"

	//	"requests"
	//	"time"
)

var args = os.Args

func main() {

	if len(args) < 4 {
		println("Seems like you are missing args!")
		return
	}


	threadsNumber, _ := getIntArgument(1)
	graphLimit, _ := getIntArgument(2)
	dependencyOdds, _ := getFloatArgument(3)
	mylistLimit, _ := getIntArgument(4)
	runtime.GOMAXPROCS(threadsNumber)

	graph := dgGraph.NewGraph(graphLimit)
	preset := generatePreset(dependencyOdds, mylistLimit)
	fmt.Println(preset)
	client := dgGraph.NewDGClient()// vai criar os nodos
	go client.Run(&graph, preset)

	/*workers := []*requests.Worker{} // vai receber os resultados

	for i := 0; i < threadsNumber; i++ {
		worker := requests.NewWorker()
		workers = append(workers, &worker)
		go worker.Run(&graph, &mylist, &requestNumber)
	}

	measureMetrics(&client, workers)
}*/




}


func getRandonInt(limit int) int {
	return rand.Intn(limit)
}

func getRandonFloat() float64 {
	return rand.Float64()
}

func getStringArgument(index int) string {
	return args[index];
}

func getIntArgument(index int) (int, error) {
	argument := getStringArgument(index)
	return strconv.Atoi(argument)
}

func getFloatArgument(index int) (float64, error) {
	argument := getStringArgument(index)
	return strconv.ParseFloat(argument, 64)
}

func generatePreset(dependencyOdds float64, myListLimit int) (requests []*dgGraph.DGRequest) {
	for i := 0; i < 16777216; i++ {
		requests = append(requests, generateRequest(getRandonInt(myListLimit), dependencyOdds))
	}
	return
}

func generateRequest(value int, dependencyOdds float64) *dgGraph.DGRequest {
	requestType := dgGraph.Read
	if dependencyOdds < getRandonFloat() {
		requestType = dgGraph.Write
	}
	request := dgGraph.NewRequest(value, requestType)
	return &request
}