package main

import (
	"dgGraph"
	"fmt"
	"runtime"

	//	"requests"
//	"time"
)

func main() {

	if len(args) < 2 {
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
