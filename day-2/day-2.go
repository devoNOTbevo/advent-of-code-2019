package main

import (
	"advent-of-code-2019/common"
	"fmt"
)

func main() {
	inputs := common.GetNumInputs(",")
	fmt.Printf("%v", inputs)
	readOpscode(inputs)
}

func readOpscode(arr []int) {
	var newArray [][]int
	fmt.Println("old length: ", len(arr))
	fmt.Println("new length: ", len(newArray))

	for i := 0; i < len(arr); i += 4 {
		chunk := arr[i:min(i+4, len(arr))]
		fmt.Println(i, "----", chunk)
		newArray = append(newArray, chunk)
		chunk = chunk[:0]
	}
	fmt.Printf("%v", newArray)

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

