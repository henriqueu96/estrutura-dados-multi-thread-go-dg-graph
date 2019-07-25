package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"tcc/requests"
	"time"
)

var args = os.Args
var mylist = requests.NewMyList()
var requestNumber = 0

func main(){
	fmt.Println("Inicio")
	var final = make(chan int)

	if(len(args) < 4){
		println("Seems like you are missing args!")
		return
	}

	threadsNumber, _ := getIntArgument(0)
	parallelizerLimit, _ := getIntArgument(1)
	dependencyOdds, _ := getFloatArgument(2)
	mylistLimit, _ := getIntArgument(3)

	parallelizer := requests.NewParallelizer(parallelizerLimit)
	preset := generatePreset( dependencyOdds , mylistLimit)

	producer(&parallelizer, preset)
	for i:= 0; i < threadsNumber; i++ {
		go consumer(&parallelizer)
	}

	measureMetrics()

	<- final
}


func measureMetrics(){
	for i := 0 ; i < 10 ; i++  {
		time.Sleep(3 * 1000000000);

	}
}

func generatePreset(dependencyOdds float64, myListLimit int) (requests []*requests.Request){
	for i := 0 ; i < 16777216 ; i++  {
		requests = append(requests, generateRequest(getRandonInt(myListLimit), dependencyOdds))
	}
	return
}

func generateRequest(value int, dependencyOdds float64) *requests.Request{
	requestType := requests.Read
	if dependencyOdds < getRandonFloat(){
		requestType = requests.Write
	}
	request := requests.NewRequest(value,requestType)
	return &request
}

func getRandonInt(limit int) int{
	return rand.Intn(limit)
}

func getRandonFloat() float64{
	return rand.Float64()
}

func getStringArgument(index int) string{
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

func consumer(parallelizer *requests.Parallelizer) {
	for {
		request := parallelizer.NextRequest()
		find := request.Execute(&mylist)
		message := ""
		if(find){
			message = "found!"
		} else {
			message = "not found!"
		}
		fmt.Println(message)
		parallelizer.Remove(request)
		requestNumber++
	}
}

func producer(parallelizer *requests.Parallelizer, preset []*requests.Request) {
	for {
		i:= 0
		request := preset[i]
		parallelizer.Add(request)
		i++
	}
}