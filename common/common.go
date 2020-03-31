package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("common package initialized.")
}

func GetNumInputs(sep string) []int {
	dat, err := ioutil.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(dat), sep)
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

func GetStringInputs(sep string) []string {
	dat, err := ioutil.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(dat), sep)
	return split
}