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
	aim      int
	data     []move
}

func (p *position) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		value, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			return err
		}
		motion := move{
			Direction: line[0],
			Value:     int(value),
		}

		p.data = append(p.data, motion)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (p *position) analyze() {
	for _, m := range p.data {
		switch m.Direction {
		case "forward":
			p.distance += m.Value
			p.depth += p.aim * m.Value
		case "up":
			p.aim -= m.Value
		case "down":
			p.aim += m.Value
		}
	}
}

func main() {
	pos := position{}
	err := pos.load("data/course.txt")
	if err != nil {
		log.Fatal(err)
	}

	pos.analyze()
	fmt.Println("distance * depth:", pos.distance*pos.depth)

}
