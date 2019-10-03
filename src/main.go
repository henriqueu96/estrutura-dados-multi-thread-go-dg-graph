package main

import (
	"fmt"
	"math/rand"
	"os"
	"requests"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

var args = os.Args
var mylist = requests.NewMyList()

func main() {
	//fmt.Println("-- Inicio --")
	runtime.GOMAXPROCS(8);
	time := time.Now()
	rand.Seed(time.UnixNano())

	//go showActivitie()

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
		worker.Id = i
		workers = append(workers, &worker)
		go worker.Run(&parallelizer, &mylist)
	}

	measureMetrics(&client, &workers)
	//println("-- Final --")
}

func showActivitie(){
	for{
		time.Sleep(2.85 * 1000000000)
		fmt.Print(".");
	}
}

func measureMetrics(client *requests.Client, workers *[]*requests.Worker) {

	var population uint64 = 0;
	// var populationAvarage float64 = 0;
	for i := 0; i < 10; i++ {
		workerProcessesNumber := uint64(0);
		time.Sleep(3 * 1000000000);
		messagesNumber := client.GetMessagesNumber()
		for _, worker := range *workers {
			workerProcessesNumber += atomic.LoadUint64(&worker.ProcessNumber);
		}
		population += messagesNumber - workerProcessesNumber;
	}

	//populationAvarage = float64(population) / 10

	processed := uint64(0);
	for _, worker := range *workers {
		processed += atomic.LoadUint64(&worker.ProcessNumber)
	}
	processedPerSecound := float64(processed) / 30

	fmt.Print(floatToString(processedPerSecound)) //+ " " +  floatToString(populationAvarage)
}

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}


func generatePreset(dependencyOdds float64, myListLimit int) (requests []*requests.Request) {
	for i := 0; i < 16777216; i++ {
		requests = append(requests, generateRequest(getRandonInt(myListLimit), dependencyOdds))
	}
	return
}

func generateRequest(value int, dependencyOdds float64) *requests.Request {
	requestType := requests.Read
	if  getRandonFloat() < dependencyOdds {
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