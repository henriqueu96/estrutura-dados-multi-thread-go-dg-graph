package main

import (
	"testing"
)

func setupArgs(){
	args = []string{
		"12",
		"10",
		"0.2",
		"100",
	}
}

func TestMain_GetStringArgument_ShouldReturnTheNumberOfThreads_When0IsTheIndex(t *testing.T) {
	setupArgs()
	nThreads := getStringArgument(0)
	if nThreads != "12"{
		t.Fail()
	}
}

func TestMain_GetIntArgument_ShouldReturnTheNumberOfThreads_When0IsTheIndex(t *testing.T) {
	setupArgs()
	nThreads, _ := getIntArgument(0)
	if nThreads != 12{
		t.Fail()
	}
}

func TestMain_GetIntArgument_ShouldReturnTheNumberOfTheParalellizerLimit(t *testing.T) {
	setupArgs()
	parallelizerLimit, _ := getIntArgument(1)
	if parallelizerLimit != 10{
		t.Fail()
	}
}


func TestMain_GetDoubleParamter_ShouldReturnTheNumberOfDependencyOdds(t *testing.T) {
	setupArgs()
	dependencyOdds, _ := getFloatArgument(2)
	if dependencyOdds != 0.2{
		t.Fail()
	}
}

func TestMain_GetIntArgument_ShouldReturnTheMyListLimit(t *testing.T) {
	setupArgs()
	myListLimit, _ := getIntArgument(3)
	if myListLimit != 100{
		t.Fail()
	}
}

func TestMain_getRandonFloat_ShouldReturnANumberGraterThan0AndLessThan1(t *testing.T) {
	randonFloat := getRandomFloat()
	if randonFloat < 0 || randonFloat > 1{
		t.Fail()
	}
}