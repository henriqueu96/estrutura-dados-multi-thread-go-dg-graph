#!/bin/bash

# 1 - quant threads
# 2 - graph size limit
# 3 - from 0 to 1 - prob of pairwise conflict
# 5 - arbitrary processing time
# 6 - repetitions
# 7 - archive where you store results

export GOPATH="/Users/henriqueurruth/Projects/tcc"

declare -a ARRAY
declare -a repetition

count=1
pcount=1
let range=$1+1
let range2=$5+1
index=0

while [ $count -le $range ]; do

	let count2=1

	while [ $count2 -lt $range2 ]; do
		repetition[$count2]=`go run main.go $count $2 $3 $4`
		let count2=$count2+1
	done

	SIZER=${#repetition[@]}

	let maximum=0

	for(( j=1;j<=$SIZER;j++)); do
		maximum=$(awk "BEGIN {printf \"%.3f\",$maximum+${repetition[${j}]}}")
	done

	value=$(awk "BEGIN {printf \"%.3f\",$maximum/$SIZER}")

	ARRAY[$index]=$value


	if [ $count -eq 1 ]
	then
		let pcount=1
	else
		let pcount=2
	fi

	if [ $count -ge 12 ]
	then
		let pcount=4
	fi

	if [ $count -ge 16 ]
	then
		let pcount=8
	fi

	let count=$count+$pcount

	let index=$index+1
done

ELEMENTS=${#ARRAY[@]}


for (( i=0;i<$ELEMENTS;i++)); do
    printf "${ARRAY[${i}]}" >> $6
    printf "\n" >> $6
done
#print variable on a screen
