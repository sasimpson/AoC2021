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
	generations [8]generation
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
	for i := 0; i < 8; i++ {
		s.generations[i].id = i
	}
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
			s.generations[timer].count++
			//s.fish = append(s.fish, fish{timer: timer})
		}
	}
	return nil
}

func (s *school) incrementDay() {
	//for i, g := range s.generations {
	//
	//}
	//var newFish []fish
	//for i, f := range s.fish {
	//	if f.timer == 0 {
	//		nf := f.spawn()
	//		if nf != nil {
	//			newFish = append(newFish, *nf)
	//		}
	//		s.fish[i].timer = 6
	//	} else {
	//		s.fish[i].timer--
	//	}
	//}
	//s.fish = append(s.fish, newFish...)
	//s.day++
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

	//for i := 0; i < 256; i++ {
	//	s.incrementDay()
	//	fmt.Println("day", i, "len", len(s.fish))
	//	//fmt.Println(s)
	//}
	//fmt.Printf("total fish: %d", len(s.fish))
}
