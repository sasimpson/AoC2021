package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	prev, current, increases int64
)

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	prev = 0
	current = 0
	increases = 0
	for scanner.Scan() {
		current, err = strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		if prev != 0 {
			if current > prev {
				increases++
			}
		}

		prev = current
	}

	fmt.Println("total increases:", increases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
