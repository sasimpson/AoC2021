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
	numbers []int
	cards   [][][]int
	calls   [][][]bool
	lookup  map[int][]numberCoords
	winners []int
	scores  map[int]int
}

type numberCoords struct {
	card, row, col int
}

func (bss *bingoSubSystem) analyze() {
	for _, num := range bss.numbers {
		if _, ok := bss.lookup[num]; ok {
			for _, l := range bss.lookup[num] {
				//check if card already won, if so, we don't need to do the rest
				if _, ok := bss.scores[l.card]; !ok {
					bss.calls[l.card][l.row][l.col] = true
					winner := bss.checkCard(l)
					if winner {
						fmt.Println("card ", l.card, " won on number", num)
						//append winning card, need to not check cards that have already won
						bss.scores[l.card] = bss.calcCard(l, num)
						bss.winners = append(bss.winners, l.card)
					}
				}
			}
		}
	}
}

func (bss *bingoSubSystem) calcCard(l numberCoords, num int) int {
	var sum int
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if bss.calls[l.card][row][col] == false {
				sum += bss.cards[l.card][row][col]
			}
		}
	}
	return sum * num
}
func (bss *bingoSubSystem) checkCard(lookup numberCoords) bool {
	calls := bss.calls[lookup.card]
	//horizontal
	var horzTotal int
	for i := 0; i < len(calls[lookup.row]); i++ {
		if calls[lookup.row][i] == true {
			horzTotal++
		} else {
			continue
		}
	}

	//vertical
	var vertTotal int
	for i := 0; i < len(calls); i++ {
		if calls[i][lookup.col] == true {
			vertTotal++
		} else {
			continue
		}
	}

	if vertTotal == 5 || horzTotal == 5 {
		return true
	}
	return false
}

func (bss *bingoSubSystem) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //get first line

	bss.lookup = make(map[int][]numberCoords)

	for _, x := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(x)
		if err != nil {
			return err
		}
		bss.numbers = append(bss.numbers, num)
	}
	scanner.Scan()    //blank line between numbers and puzzles
	var board [][]int //holder for the boards generated in the for loop
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
			var boardLine []int
			var callLine []bool
			if re.MatchString(line) {
				sm := re.FindStringSubmatch(line)
				for i, x := range sm[1:] {
					number, err := strconv.Atoi(x)
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
	bss.cards = append(bss.cards, board)
	bss.calls = append(bss.calls, calls)
	bss.scores = make(map[int]int, len(bss.cards))

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
	fmt.Println("first winner: ", bss.winners[0], " with a score of ", bss.scores[bss.winners[0]])
	fmt.Println("last winner:", bss.winners[len(bss.winners)-1], " with a score of ", bss.scores[bss.winners[len(bss.winners)-1]])
	//}
}
