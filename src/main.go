package main

import (
	"fmt"
	"math/rand"
	"os"
	"requests"
	"strconv"
	"time"
)

var args = os.Args
var mylist = requests.NewMyList()
var requestNumber = 0

func main() {
	fmt.Println("-- Inicio --")

	if (len(args) < 4) {
		println("Seems like you are missing args!")
		return
	}

	threadsNumber, _ := getIntArgument(1)
	parallelizerLimit, _ := getIntArgument(2)
	dependencyOdds, _ := getFloatArgument(3)
	mylistLimit, _ := getIntArgument(4)

	parallelizer := requests.NewParallelizer(parallelizerLimit)
	preset := generatePreset(dependencyOdds, mylistLimit)

	client := requests.NewClient()
	go client.Run(&parallelizer, preset)

	workers := []*requests.Worker{}

	for i := 0; i < threadsNumber; i++ {
		worker := requests.NewWorker()
		workers = append(workers, &worker)
		go worker.Run(&parallelizer, &mylist)
	}

	measureMetrics(&client, workers)
	println("-- Final --")
}

func measureMetrics(client *requests.Client, workers []*requests.Worker) {

	population := 0;
	for i := 0; i < 10; i++ {
		workerProcessesNumber := 0;
		time.Sleep(3 * 1000000000);
		messagesNumber := client.MessagesNumber;
		for _, worker := range workers {
			workerProcessesNumber += worker.ProcessNumber;
		}
		population += messagesNumber - workerProcessesNumber;
	}

	population = population / 10;

	processed := 0;
	for _, worker := range workers {
		processed += worker.ProcessNumber
	}
	processedPerSecound := float64(processed) / 30.0;

	fmt.Print(processedPerSecound);
	fmt.Print(" ")
	fmt.Println(population);
}

func generatePreset(dependencyOdds float64, myListLimit int) (requests []*requests.Request) {
	for i := 0; i < 16777216; i++ {
		requests = append(requests, generateRequest(getRandonInt(myListLimit), dependencyOdds))
	}
	return
}

func generateRequest(value int, dependencyOdds float64) *requests.Request {
	requestType := requests.Read
	if dependencyOdds < getRandonFloat() {
		requestType = requests.Write
	}
	request := requests.NewRequest(value, requestType)
	return &request
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
