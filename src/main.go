package main

import (
	"dgGraph"
	"fmt"
	"math/rand"
	"os"
	//"requests"
	"runtime"
	"strconv"
	"time"

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
	client := dgGraph.NewDGClient()
	go client.Run(&graph, preset)


	time.Sleep(time.Minute);
	//measureMetrics(&client)
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
	// todo: change to 16777216
	for i := 0; i < 500; i++ {
		requests = append(requests, generateRequest(getRandonInt(myListLimit), dependencyOdds))
	}
	return requests
}

func generateRequest(value int, dependencyOdds float64) *dgGraph.DGRequest {
	requestType := dgGraph.Read
	if getRandonFloat() < dependencyOdds{
		requestType = dgGraph.Write
	}
	request := dgGraph.NewRequest(value, requestType)
	return &request
}

func measureMetrics(client *dgGraph.DGClient) {
	var metric uint64 = 0;
	for i := 0; i < 1; i++ {
		var workerProcessesNumber uint64 = 0;
		time.Sleep(10000000);
		messagesNumber := client.MessagesNumber;
		workerProcessesNumber += dgGraph.GetProcessNumber()
		metric += messagesNumber - workerProcessesNumber;
	}
	metric = metric / 10;
	metric += dgGraph.GetProcessNumber()

	commands := metric / 30;
	fmt.Print(commands);
}
