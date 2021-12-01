package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
	fmt.Println(counter)
}
