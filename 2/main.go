package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	horizontal := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " ")
		value, _ := strconv.Atoi(parts[1])
		if parts[0] == "forward" {
			horizontal += value
			depth += aim * value
		} else if parts[0] == "up" {
			aim -= value
		} else {
			aim += value
		}
	}
	fmt.Println("Horizontal: ", horizontal)
	fmt.Println("Depth: ", depth)

	fmt.Println("Multiplied: ", horizontal*depth)
}
