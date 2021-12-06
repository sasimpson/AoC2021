package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ttacon/chalk"
)

type bingoSubSystem struct {
	numbers []uint64
	cards   [][][]uint64
	calls   [][][]bool
}

func (bss *bingoSubSystem) validateCoords(card, row, col uint64) bool {
	return (uint64(len(bss.cards[card])) > row && row >= 0) && (uint64(len(bss.cards[card][row])) > col && col >= 0)
}

func (bss *bingoSubSystem) playNumber(num uint64) {
	for i, card := range bss.cards {
		for row := 0; row < len(card); row++ {
			for col := 0; col < len(card[row]); col++ {
				if num == card[row][col] {
					bss.calls[i][row][col] = true
				} else {
					bss.calls[i][row][col] = false
				}
			}
		}
	}
}
func (bss *bingoSubSystem) analyze() {
	for _, num := range bss.numbers {
		for card := range bss.cards {
			for row := range bss.cards[card] {
				for col, cardNum := range bss.cards[card][row] {
					if num == cardNum {
						bss.calls[card][row][col] = true
					} else {
						bss.calls[card][row][col] = false
					}
				}
			}
		}
	}
}

func (bss *bingoSubSystem) displayCard(card uint64) {
	for row := range bss.cards[card] {
		for col := range bss.cards[card][row] {
			if bss.calls[card][row][col] {
				fmt.Printf(chalk.White.Color("[%2d]"), bss.cards[card][row][col])
			} else {
				fmt.Printf(chalk.White.Color("[%2d]"), bss.cards[card][row][col])
			}
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

	for _, x := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return err
		}
		bss.numbers = append(bss.numbers, num)
	}

	scanner.Scan() //blank line between numbers and puzzles

	var board [][]uint64 //holder for the boards generated in the for loop
	re := regexp.MustCompile(`\s*(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			bss.cards = append(bss.cards, board)
			board = nil
		} else {
			var boardLine []uint64
			if re.MatchString(line) {
				sm := re.FindStringSubmatch(line)
				for _, x := range sm[1:] {
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
	_ = bss.load("data/bingo.txt")

	for card := range bss.cards {
		bss.displayCard(uint64(card))
	}
}
