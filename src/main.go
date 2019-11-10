package main

import (
	"dgGraph"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

var args = os.Args


var presetLengthNumber uint64 = 10000000;
var graphLimit int64 = 0;

func main() {

	if len(args) < 4 {
		println("Seems like you are missing args!")
		return
	}

	threadsNumber, _ := getIntArgument(1)
	graphLimit, _ = getInt64Argument(2)
	dependencyOdds, _ := getFloatArgument(3)
	mylistLimit, _ := getIntArgument(4)
	runtime.GOMAXPROCS(threadsNumber)

	graph := dgGraph.NewGraph(graphLimit)
	preset := generatePreset(dependencyOdds, mylistLimit)
	client := dgGraph.NewDGClient()
	go client.Run(&graph, preset)

	measureMetrics(&graph);
}

func getRandomInt(limit int) int {
	return rand.Intn(limit)
}

func getRandomFloat() float64 {
	return rand.Float64()
}

func getStringArgument(index int) string {
	return args[index];
}

func getIntArgument(index int) (int, error) {
	argument := getStringArgument(index)
	return strconv.Atoi(argument)
}

func getInt64Argument(index int) (int64, error) {
	argument := getStringArgument(index)
	return strconv.ParseInt(argument, 10,64)
}

func getFloatArgument(index int) (float64, error) {
	argument := getStringArgument(index)
	return strconv.ParseFloat(argument, 64)
}

func generatePreset(dependencyOdds float64, myListLimit int) (requests []*dgGraph.DGRequest) {
	for i := 0; i < int(presetLengthNumber); i++ {
		requests = append(requests, generateRequest(getRandomInt(myListLimit), dependencyOdds))
	}
	return requests
}

func generateRequest(value int, dependencyOdds float64) *dgGraph.DGRequest {
	requestType := dgGraph.Read

	if getRandomFloat() < dependencyOdds {
		requestType = dgGraph.Write
	}
	request := dgGraph.NewRequest(value, requestType)
	return &request
}

func measureMetrics(graph *dgGraph.DGGraph) {

	var population int64 = 0;
	var populationAvarage float64 = 0;
	for i := 0; i < 10; i++ {
		time.Sleep(3 * 1000000000);
		population = graph.Length

		/*
		graphLength := atomic.LoadInt64(&graph.Length)
		addedNodes := atomic.LoadInt64(&graph.AddedNodes)

		fmt.Println("Add messages and nodes: " + int64ToString(graphLength))
		fmt.Println("Nodes: " + int64ToString(addedNodes));*/
	}

	populationAvarage = float64(population) / 10

	processed := dgGraph.GetProcessNumber()
	processedPerSecound := float64(processed) / 30

	fmt.Print(floatToString(processedPerSecound) + " - " +  floatToString(populationAvarage))
}

func floatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func int64ToString(input_num int64) string {
	return strconv.FormatInt(input_num, 10)
}