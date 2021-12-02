package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(secondPart())
}

func secondPart() int {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	differencesCounter := 0

	recorder := make([]int, 0, 4)

	for scanner.Scan() {
		text := scanner.Text()
		currentElement, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(recorder) > 3 {
			differencesCounter += compareSums(recorder[:3], recorder[1:])
			recorder = recorder[1:]
		}
		recorder = append(recorder, currentElement)
	}
	return differencesCounter
}

func compareSums(base, next []int) int {
	sumBase := 0
	sumNext := 0
	for _, v := range base {
		sumBase += v
	}

	for _, v := range next {
		sumNext += v
	}
	if sumNext > sumBase {
		return 1
	}
	return 0
}

func firstPart() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	mem := 0
	counter := 0

	for scanner.Scan() {
		text := scanner.Text()
		current, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if mem == 0 {
			mem = current
			continue
		}
		if current > mem {
			counter += 1
		}
		mem = current
	}
	return counter
}
