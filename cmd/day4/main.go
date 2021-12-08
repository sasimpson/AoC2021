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
	calls   [][][]bool
	lookup  map[uint64][]numberCoords
}

type numberCoords struct {
	card, row, col int
}

//func (bss *bingoSubSystem) playNumber(num uint64) {
//	for i, card := range bss.cards {
//		for row := 0; row < len(card); row++ {
//			for col := 0; col < len(card[row]); col++ {
//				if num == card[row][col] {
//					fmt.Println("match", i, row, col)
//					bss.calls[i][row][col] = true
//				} else {
//					bss.calls[i][row][col] = false
//				}
//			}
//		}
//	}
//}

func (bss *bingoSubSystem) analyze() {
	for _, num := range bss.numbers {
		if _, ok := bss.lookup[num]; ok {
			for _, l := range bss.lookup[num] {
				fmt.Printf("looking up card %d (%d, %d)\n", l.card, l.row, l.col)
				bss.calls[l.card][l.row][l.col] = true
			}
		}
	}
}

func (bss *bingoSubSystem) displayCard(card uint64) {
	for row := range bss.cards[card] {
		for col := range bss.cards[card][row] {
			fmt.Printf("[%2d]", bss.cards[card][row][col])

		}
		fmt.Printf("\n")
	}
}

func (bss *bingoSubSystem) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //get first line

	bss.lookup = make(map[uint64][]numberCoords)

	for _, x := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return err
		}
		bss.numbers = append(bss.numbers, num)
	}
	scanner.Scan()       //blank line between numbers and puzzles
	var board [][]uint64 //holder for the boards generated in the for loop
	var calls [][]bool
	var cardCounter int
	re := regexp.MustCompile(`\s*(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			bss.cards = append(bss.cards, board)
			board = nil
			bss.calls = append(bss.calls, calls)
			calls = nil
			cardCounter++
		} else {
			var boardLine []uint64
			var callLine []bool
			if re.MatchString(line) {
				sm := re.FindStringSubmatch(line)
				for i, x := range sm[1:] {
					number, err := strconv.ParseUint(x, 10, 64)
					if err != nil {
						return err
					}
					boardLine = append(boardLine, number)
					callLine = append(callLine, false)
					bss.lookup[number] = append(bss.lookup[number], numberCoords{len(bss.cards), len(board), i})
				}
				board = append(board, boardLine)
				calls = append(calls, callLine)
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
	_ = bss.load("data/bingo.txt")
	bss.analyze()

	//for card := range bss.cards {
	//bss.displayCard(uint64(card))
	fmt.Println()
	//}
}
