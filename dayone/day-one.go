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
	sum := getSumFromRange(inputs)
	fmt.Println("SUM:  ", sum)
}

//divide by 3, round down, subtract 2
func calcFuelRequirements(mass int) int {
	return int(mass/3) - 2
}

func getSumFromRange(arr []int) int {
	sum := 0
	for index, element := range arr {
		fmt.Println(index, "----", element)
		sum += element
	}
	return sum
}

func getInputs() []int {
	dat, err := ioutil.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(dat), "\n")
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
