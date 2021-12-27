package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type crab struct {
	position int
}

type swarm struct {
	crabs        []crab
	min, max     int
	bestPosition struct {
		position, fuel int
	}
}

func (s *swarm) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var min, max int
	min = 1000

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		for _, d := range data {
			crabPos, err := strconv.Atoi(d)
			if err != nil {
				return err
			}
			s.crabs = append(s.crabs, crab{position: crabPos})
			if crabPos > max {
				max = crabPos
			}
			if crabPos < min {
				min = crabPos
			}
		}
	}
	s.min = min
	s.max = max
	return nil
}

func (s *swarm) analyze() {
	for i := s.min; i <= s.max; i++ {
		var fuel int
		for _, c := range s.crabs {
			fuel += c.fuelToPos(i)
		}

		if s.bestPosition.fuel == 0 {
			s.bestPosition.fuel = 10000000
		}
		if fuel < s.bestPosition.fuel {
			s.bestPosition.fuel = fuel
			s.bestPosition.position = i
		}
		fmt.Println("Pos", i, "fuel", fuel)
	}
}

func (c *crab) fuelToPos(pos int) int {
	var fuel int
	steps := c.position - pos
	if steps < 0 {
		return -steps
	}
	for i := 1; i <= steps; i++ {
		fuel += i
	}
	fmt.Println("c pos:", c.position, "fuel:", fuel)
	return fuel
}

func main() {
	var s swarm
	err := s.load("data/crabs_sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	s.analyze()
	//fmt.Println("Pos", s.bestPosition.position, "fuel", s.bestPosition.fuel)
}
