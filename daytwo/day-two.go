package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputs := getInputs()
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

func getInputs() []int {
	dat, err := ioutil.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(dat), ",")
	returnArr := []int{}
	for _, element := range split {
		i, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
		}
		returnArr = append(returnArr, i)
	}

	return returnArr
}
