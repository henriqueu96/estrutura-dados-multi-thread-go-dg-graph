#!/bin/bash

declare -a ARRAY
declare -a repetition

export GOROOT='/usr/local/go'
export GOPATH='/Users/henriqueurruth/Projects/tcc/'

maximunOfThreadsIteration=$1
parallelizerLimit=$2
dependencyOdds=$3
mylistLimit=$4
maxRepetitionNumber=$5
fileName=$6
index=0
numberOfThreads=1

while [ $numberOfThreads -le $maximunOfThreadsIteration ]; do
	repetitionNumber=1
	while [ $repetitionNumber -lt $maxRepetitionNumber ]; do
		# shellcheck disable=SC2006
		repetition[$repetitionNumber]=`go run main.go $numberOfThreads $parallelizerLimit $dependencyOdds $mylistLimit`
		repetitionNumber=$repetitionNumber+1
	done

	SIZER=${#repetition[@]}

	let maximum=0

	for(( j=1;j<=$SIZER;j++)); do
		maximum=$(awk "BEGIN {printf \"%.3f\",$maximum+${repetition[${j}]}}")
	done

	value=$(awk "BEGIN {printf \"%.3f\",$maximum/$SIZER}")

	ARRAY[$index]=$value

	if [ $numberOfThreads -eq 1 ]
	then
		let numberOfThreads=$numberOfThreads+1
	else
		let numberOfThreads=$numberOfThreads+2
	fi

	let index=$index+1
done

ELEMENTS=${#ARRAY[@]}

for (( i=0;i<$ELEMENTS;i++)); do
    printf "${ARRAY[${i}]}" >> $fileName
    printf "\n" >> $fileName
done
#print variable on a screen