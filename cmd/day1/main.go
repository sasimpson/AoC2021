package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	prev, increases int

	ErrInvalidIndex = errors.New("invalid index for array given (over or negative)")
)

func main() {

	data, err := readData("data/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(data)-2; i++ {
		window, err := getWindow(&data, i, 3)
		if err != nil {
			log.Fatal(err)
		}

		total := sum(&window)

		if prev != 0 {
			if total > prev {
				increases++
			}
		}
		prev = total
	}

	fmt.Println("total increases:", increases)

}

func sum(data *[]int) int {
	var sum int
	for _, d := range *data {
		sum += d
	}
	return sum
}

func getWindow(data *[]int, pos int, window int) ([]int, error) {
	//index issues
	if pos < 0 || pos > len(*data) || pos+window > len(*data) {
		return nil, ErrInvalidIndex
	}

	return (*data)[pos : pos+window], nil

}

func readData(filename string) (data []int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return []int{}, err
		}

		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}
	return
}
