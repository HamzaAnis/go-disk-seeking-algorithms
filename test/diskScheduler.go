package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) (string, map[string]int, []int, error) {
	algoName := ""
	// to store the data
	command := map[string]int{}
	// to store the cylinder requests
	requests := []int{}
	// reading file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", nil, nil, err
	}
	// splitting lines
	lines := strings.Split(string(data), "\n")

	// storing data
	for i := 0; i < 4; i++ {
		line := strings.Split(lines[i], " ")
		if i == 0 {
			algoName = line[1]
			continue
		}
		if number, err := strconv.Atoi(line[1]); err != nil {
			return "", nil, nil, err
		} else {
			command[line[0]] = number
		}
	}

	for i := 4; i < len(lines)-1; i++ {
		line := strings.Split(lines[i], " ")
		if line[0] == "end" {
			break
		} else if request, err := strconv.Atoi(line[1]); err != nil {
			return "", nil, nil, err
		} else {
			requests = append(requests, request)
		}
	}
	return algoName, command, requests, nil
}

// FCFC implements First come first serve
func FCFS(inputsMap map[string]int, requests []int) int {
	count := 0
	initPosition := inputsMap["initCYL"]
	for _, request := range requests {
		if request < inputsMap["upperCYL"] && request > inputsMap["lowerCYL"] {
			count = count + int(math.Abs(float64(initPosition-request)))
			fmt.Printf("Servicing %5d\n", request)
			initPosition = request
		} else {
			log.Printf("%d out of bound for %d - %d", request, inputsMap["lowerCYL"], inputsMap["upperCYL"])
		}

	}
	return count
}

// SSTF implement SSTF algo
func SSTF(inputsMap map[string]int, requests []int) int {
	visited := map[int]bool{}
	for _, request := range requests {
		visited[request] = false
	}
	count := 0
	initPosition := inputsMap["initCYL"]
	for {
		difference, idx := findShortestCylinder(requests, visited, initPosition)
		if idx == -1 {
			break
		}
		visited[requests[idx]] = true
		count = count + difference
		initPosition = requests[idx]
		fmt.Printf("Servicing %d\n", requests[idx])
	}
	return count
}

func findShortestCylinder(requests []int, visited map[int]bool, initPosition int) (int, int) {
	idx := -1
	difference := 99999999
	for i, request := range requests {
		test := int(math.Abs(float64(request - initPosition)))
		if !visited[request] {
			if test < difference {
				difference = test
				idx = i
			}
		}
	}
	return difference, idx
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Filename is missing")
	}
	if algoName, inputsMap, requests, err := readFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Seek algorithm: %s\n", strings.ToUpper(algoName))
		fmt.Printf("\tLower cylinder: %5d\n", inputsMap["lowerCYL"])
		fmt.Printf("\tUpper cylinder: %5d\n", inputsMap["upperCYL"])
		fmt.Printf("\tInit cylinder: %5d\n", inputsMap["initCYL"])
		fmt.Printf("\tCylinder requests:\n")
		for _, request := range requests {
			fmt.Printf("\t\tCylinder %5d\n", request)
		}
		if algoName == "fcfs" {
			fmt.Printf("FCFS traversal count = %d\n", FCFS(inputsMap, requests))
		} else if algoName == "sstf" {
			fmt.Printf("SSTF traversal count = %d\n", SSTF(inputsMap, requests))
		} else if algoName == "scan" {
			// fmt.Printf("C-LOCK traversal count = %d\n", C_LOOK(inputsMap, requests))

		} else if algoName == "c-scan" {
			fmt.Printf("FCFS traversal count = %d\n", FCFS(inputsMap, requests))
		} else if algoName == "look" {
			fmt.Printf("FCFS traversal count = %d\n", FCFS(inputsMap, requests))

		} else if algoName == "c-look" {
			fmt.Printf("C-LOCK traversal count = %d\n", C_LOOK(inputsMap, requests))

		}
	}
}
