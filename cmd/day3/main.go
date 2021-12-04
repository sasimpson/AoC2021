package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type power struct {
	gamma   uint
	epsilon uint
}

func main() {

	var p power

	data, err := readData("data/diagnostics.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := 11; i >= 0; i-- {
		var notbit uint
		bit := mostCommonBit(data, i)
		if bit == 0 {
			notbit = 1
		}
		fmt.Println("place:", i, " bit:", bit, " not bit:", notbit)
		p.gamma = p.gamma << 1
		p.gamma = p.gamma + bit
		p.epsilon = p.epsilon << 1
		p.epsilon = p.epsilon + notbit
	}

	fmt.Printf("gamma result: \t%012b\n", p.gamma)
	fmt.Printf("epsilon result: \t%012b\n", p.epsilon)
	fmt.Printf("power consumption: %d", p.gamma*p.epsilon)
}

func mostCommonBit(data []int64, pos int) uint {
	var ones int
	for i := 0; i < len(data); i++ {
		if getBit(data[i], pos) == 1 {
			ones++
		}
	}
	if ones >= len(data)-ones {
		return 1
	}
	return 0
}

func getBit(data int64, pos int) int64 {
	shifted := data >> pos
	return shifted & 1
}

func readData(filename string) (data []int64, err error) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.ParseInt(scanner.Text(), 2, 16)
		if err != nil {
			return data, err
		}
		data = append(data, line)
	}
	if err := scanner.Err(); err != nil {
		return data, err
	}
	return
}
