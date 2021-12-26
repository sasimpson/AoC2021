package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stringer interface {
	String() string
}

type school struct {
	day         int
	fish        []fish
	generations [][]generation
}

type generation struct {
	id    int
	count int
}

type fish struct {
	timer int
}

//join - join a stringer into 1 string.  playing with generics :D
func join[T Stringer](t []T, sep string) string {
	// if the len is not >= 2 then we have nothing to join.
	if len(t) < 2 {
		return fmt.Sprintf("%s", t[0])
	}

	retValue := fmt.Sprintf("%s", t[0])
	for _, v := range t[1:] {
		retValue = fmt.Sprintf("%s%s%s", retValue, sep, v)
	}
	return retValue
}

func (s school) String() string {
	data := join(s.fish, ",")
	return fmt.Sprintf("After %2d days: %s", s.day, data)
}

func (s *school) init() {
	initGenerations := make([]generation, 9)
	for i := 0; i <= 8; i++ {
		initGenerations[i].id = i
	}
	s.generations = append(s.generations, initGenerations)
	fmt.Printf("%#v\n", s.generations)
}

func (s *school) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		for _, d := range data {
			timer, err := strconv.Atoi(d)
			if err != nil {
				return err
			}
			s.generations[0][timer].count++
		}
	}
	return nil
}

func (s *school) incrementDay() []generation {
	currentGen := s.generations[len(s.generations)-1]
	nextGen := make([]generation, 9)
	for i := 0; i <= 8; i++ {
		switch {
		//if current gen is 0, the next gen will spawn count of current gen
		case i == 0:
			nextGen[8].count += currentGen[0].count
			nextGen[6].count = currentGen[0].count
		case i < 8:
			nextGen[i].count = currentGen[i+1].count
		}
	}
	return nextGen
}

func (f fish) String() string {
	return strconv.Itoa(f.timer)
}

func (f *fish) spawn() *fish {
	if f.timer == 0 {
		f.timer = 6
		return &fish{timer: 8}
	}
	return nil
}

func main() {
	var s school

	s.init()
	err := s.load("data/lanternfish_sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", s.generations)

	for i := 0; i < 12; i++ {
		s.generations = append(s.generations, s.incrementDay())
		fmt.Println("day", i, "len", len(s.fish))
		fmt.Println(s)
	}
	fmt.Printf("total fish: %d", len(s.fish))
}
