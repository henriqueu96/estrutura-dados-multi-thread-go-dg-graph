package main

import (
	"fmt"
		"math/rand"
		"strconv"
		"time"

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
	fmt.Println("Inicio")
	var final = make(chan int)

	if (len(args) < 2) {
		println("Seems like you are missing args!")
		return
		dgGraph.NewRequest()

	}
	func getRandonInt(limit int) int {
		return rand.Intn(limit)
	}

	
			threadsNumber, _ := getIntArgument(1)
			parallelizerLimit, _ := getIntArgument(2)
			dependencyOdds, _ := getFloatArgument(1)
			mylistLimit, _ := getIntArgument(2)

			parallelizer := requests.NewParallelizer(parallelizerLimit)
			preset := generatePreset(dependencyOdds, mylistLimit)
			fmt.Println(preset)
			client := requests.NewClient()
			go client.Run(&parallelizer, preset)

			workers := []*requests.Worker{}

			for i := 0; i < threadsNumber; i++ {
				worker := requests.NewWorker()
				workers = append(workers, &worker)
				go worker.Run(&parallelizer, &mylist, &requestNumber)
			}

			measureMetrics(&client, workers)

			<-final
		}

		func measureMetrics(client *requests.Client, workers []*requests.Worker) {
			metric := 0;
			for i := 0; i < 10; i++ {
				workerProcessesNumber := 0;
				time.Sleep(3 * 1000000000);
				messagesNumber := client.MessagesNumber;
				for _, worker := range workers {
					workerProcessesNumber += worker.ProcessNumber;
				}
				metric += messagesNumber - workerProcessesNumber;
			}
			metric = metric / 10;
			for _, worker := range workers {
				metric += worker.ProcessNumber
			}

			commands := metric / 30;
			fmt.Print(commands);
		}

		func generatePreset(dependencyOdds float64, myListLimit int) (requests []*requests.request) {
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
	
