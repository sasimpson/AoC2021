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
	crabs []crab
}

func (s *swarm) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		for _, d := range data {
			crabPos, err := strconv.Atoi(d)
			if err != nil {
				return err
			}
			s.crabs = append(s.crabs, crab{position: crabPos})
		}
	}
	return nil
}

func main() {

	var s swarm
	err := s.load("data/crabs_sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", s)

}
