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
	prev, increases int64

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

func sum(data *[]int64) int64 {
	var sum int64
	for _, d := range *data {
		sum += d
	}
	return sum
}

func getWindow(data *[]int64, pos int, window int) ([]int64, error) {
	//index issues
	if pos < 0 || pos > len(*data) || pos+window > len(*data) {
		return nil, ErrInvalidIndex
	}

	return (*data)[pos : pos+window], nil

}

func readData(filename string) (data []int64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
	return
}
