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

var presetLengthNumber uint64 = 16777216;

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
	client := dgGraph.NewDGClient()
	go client.Run(&graph, preset)
	measureMetrics()
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

func measureMetrics() {
	for i := 0; i < 10; i++ {
		workerProcessesNumber := uint64(0);
		time.Sleep(3 * time.Second);
		workerProcessesNumber += dgGraph.GetProcessNumber();
	}

	processed := uint64(0);
	processed += dgGraph.GetProcessNumber();
	processedPerSecound := float64(processed) / 30

	fmt.Print(floatToString(processedPerSecound))
}

func floatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
