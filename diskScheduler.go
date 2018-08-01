package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func fcfs(inputsMap map[string]int, requests []int) (string, int) {
	count := 0
	for _, request := range requests {
		fmt.Printf("Servicing %5d\n", request)
	}
	return "FCFS traversal count ", count
}

//Starting point for the program
func diskSeekingAlgorithms(filename string) error {
	if algoName, inputsMap, requets, err := readFile(filename); err != nil {
		return err
	} else {
		fmt.Printf("Seek algorithm: %s\n", strings.ToUpper(algoName))
		fmt.Printf("\tLower cylinder: %5d\n", inputsMap["lowerCYL"])
		fmt.Printf("\tUpper cylinder: %5d\n", inputsMap["upperCYL"])
		fmt.Printf("\tUpper cylinder: %5d\n", inputsMap["initCYL"])
		fmt.Printf("\tCylinder requests:\n")
		for _, request := range requets {
			fmt.Printf("\t\tCylinder %5d\n", request)
		}
		if algoName == "fcfs" {
			if title, count := fcfs(inputsMap, requets); err != nil {
				return err
			}
		} else if algoName == "sstf" {
			if err := fcfs(inputsMap, requets); err != nil {
				return err
			}
		} else if algoName == "scan" {
			if err := fcfs(inputsMap, requets); err != nil {
				return err
			}
		} else if algoName == "c-scan" {
			if err := fcfs(inputsMap, requets); err != nil {
				return err
			}
		} else if algoName == "look" {
			if err := fcfs(inputsMap, requets); err != nil {
				return err
			}
		} else if algoName == "c-look" {
			if err := fcfs(inputsMap, requets); err != nil {
				return err
			}
		}
	}
	return nil
}
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Filename is missing")
	}
	if err := diskSeekingAlgorithms(os.Args[1]); err != nil {
		log.Fatal(err)
	}

}
