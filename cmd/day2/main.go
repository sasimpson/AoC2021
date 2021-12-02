package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type move struct {
	Direction string
	Value     int
}

type position struct {
	distance int
	depth    int
}

func main() {
	data, err := readData("data/course.txt")
	if err != nil {
		log.Fatal(err)
	}

	pos := position{}

	for _, m := range data {
		switch m.Direction {
		case "forward":
			pos.distance += m.Value
		case "up":
			pos.depth -= m.Value
		case "down":
			pos.depth += m.Value
		}
	}

	fmt.Println("distance * depth:", pos.distance*pos.depth)

}

func readData(filename string) (data []move, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		value, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			return []move{}, err
		}
		motion := move{
			Direction: line[0],
			Value:     int(value),
		}

		data = append(data, motion)
	}

	if err := scanner.Err(); err != nil {
		return []move{}, err
	}
	return
}
