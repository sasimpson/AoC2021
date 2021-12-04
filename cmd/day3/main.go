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

func (p *power) consumption() uint {
	return p.gamma * p.epsilon
}

func (p *power) analyze(d diagnostics) {
	for i := 11; i >= 0; i-- {
		var notbit uint
		bit := mostCommonBit(d.Readings, i)
		if bit == 0 {
			notbit = 1
		}
		p.gamma = p.gamma << 1
		p.gamma = p.gamma + bit
		p.epsilon = p.epsilon << 1
		p.epsilon = p.epsilon + notbit
	}
}

type lifeSupport struct {
	oxygen uint
	co2    uint
}

func (ls *lifeSupport) analyzeOxygen(d diagnostics) {
	data := new(diagnostics)
	for _, val := range d.Readings {
		data.Readings = append(data.Readings, val)
	}
	readings(data, 11, true)
	ls.oxygen = uint(data.Readings[0])
}

func (ls *lifeSupport) analyzeCO2(d diagnostics) {
	data := new(diagnostics)
	for _, val := range d.Readings {
		data.Readings = append(data.Readings, val)
	}
	readings(data, 11, false)
	ls.co2 = uint(data.Readings[0])
}

func (ls *lifeSupport) rating() uint {
	return ls.oxygen * ls.co2
}

func readings(d *diagnostics, level int, most bool) int {
	var deleteList []int
	var compBit uint
	if most {
		compBit = mostCommonBit(d.Readings, level)
	} else {
		compBit = leastCommonBit(d.Readings, level)
	}
	for i, r := range d.Readings {
		b := getBit(r, level)
		if b != int64(compBit) {
			deleteList = append(deleteList, i)
		}
	}
	for i, dl := range deleteList {
		d.delete(dl - i)
	}
	if level <= 0 || len(d.Readings) == 1 {
		return 0
	}
	level--
	return readings(d, level, most)
}

type diagnostics struct {
	Readings []int64
}

func (d *diagnostics) load(filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.ParseInt(scanner.Text(), 2, 16)
		if err != nil {
			return err
		}
		d.Readings = append(d.Readings, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (d *diagnostics) delete(i int) int64 {
	x := d.Readings[i]
	d.Readings = append(d.Readings[:i], d.Readings[i+1:]...)
	return x
}

func main() {
	var (
		diags diagnostics
		p     power
		ls    lifeSupport
	)

	err := diags.load("data/diagnostics.txt")
	if err != nil {
		log.Fatal(err)
	}

	p.analyze(diags)
	fmt.Printf("power consumption: %d\n", p.consumption())

	ls.analyzeOxygen(diags)
	ls.analyzeCO2(diags)
	fmt.Printf("life support rating: %d\n", ls.rating())
}

// mostCommonBit - get most common bit for a position in a byte.
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

func leastCommonBit(data []int64, pos int) uint {
	var ones int
	for i := 0; i < len(data); i++ {
		if getBit(data[i], pos) == 1 {
			ones++
		}
	}
	if ones < len(data)-ones {
		return 1
	}
	return 0
}

func getBit(data int64, pos int) int64 {
	shifted := data >> pos
	return shifted & 1
}
