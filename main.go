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
	data, err := readData(scanner)
	if err != nil {
		log.Fatal(err)
	}

	prev = 0
	current = 0
	increases = 0

	for i := 2; i < len(data); i++ {
		window := data[i-2] + data[i-1] + data[i]

		if prev != 0 {
			if window > prev {
				increases++
			}
		}
		prev = window
	}

	fmt.Println("total increases:", increases)

}

func readData(scanner *bufio.Scanner) ([]int64, error) {
	var data []int64

	for scanner.Scan() {
		line, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return []int64{}, err
		}

		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		return []int64{}, err
	}

	return data, nil
}
