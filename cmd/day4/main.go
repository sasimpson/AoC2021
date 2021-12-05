package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bingoSubSystem struct {
	numbers []uint64
	cards   [][][]uint64
}

func (bss *bingoSubSystem) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //get first line

	for _, x := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return err
		}
		bss.numbers = append(bss.numbers, num)
	}

	scanner.Scan() //blank line between numbers and puzzles

	var board [][]uint64 //holder for the boards generated in the for loop
	re := regexp.MustCompile("\\s+")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			bss.cards = append(bss.cards, board)
			board = nil
		} else {
			var boardLine []uint64
			if re.MatchString(line) {
				for _, x := range re.FindStringSubmatch(line)[1:] {
					number, err := strconv.ParseUint(x, 10, 64)
					if err != nil {
						return err
					}
					boardLine = append(boardLine, number)
				}
				board = append(board, boardLine)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("bingo subsystem loaded")
	return nil
}

func main() {

	bss := bingoSubSystem{}
	bss.load("data/bingo.txt")

}
