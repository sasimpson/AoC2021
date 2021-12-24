package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	coordRegExp = `^(\d+),(\d+) -> (\d+),(\d+)`
)

type point struct {
	x, y int
}

type line struct {
	start, end point
}

type vents struct {
	data     []line
	chart    [1000][1000]int
	overlaps map[point]int
}

func (p point) String() string {
	return fmt.Sprintf("point: (%d,%d)", p.x, p.y)
}

func (l line) String() string {
	return fmt.Sprintf("line: (%d,%d) -> (%d,%d) length: %d", l.start.x, l.start.y, l.end.x, l.end.y, l.Length())
}

func (l line) IsHorizontal() bool {
	if l.start.y == l.end.y {
		return true
	}
	return false
}

func (l line) IsVertical() bool {
	if l.start.x == l.end.x {
		return true
	}
	return false
}

func (l line) IsDiagonal() bool {
	if (l.start.x != l.end.x) && (l.start.y != l.end.y) {
		return true
	}
	return false
}

func (l line) Length() int {
	var length int
	switch {
	//if x is the same, range on y
	case l.start.x == l.end.x:
		length = l.start.y - l.end.y
	case l.start.y == l.end.y:
		length = l.start.x - l.end.x
	}
	if length < 0 {
		return -length
	}
	return length
}

func (l line) Iterator() ([]int, []int) {
	var returnX, returnY []int
	switch {
	case l.IsHorizontal():
		if l.start.x > l.end.x {
			for i := l.start.x; i >= l.end.x; i-- {
				returnX = append(returnX, i)
			}
		} else {
			for i := l.start.x; i <= l.end.x; i++ {
				returnX = append(returnX, i)
			}
		}
	case l.IsVertical():
		if l.start.y > l.end.y {
			for i := l.start.y; i >= l.end.y; i-- {
				returnY = append(returnY, i)
			}
		} else {
			for i := l.start.y; i <= l.end.y; i++ {
				returnY = append(returnY, i)
			}
		}
	default:
		if l.start.x > l.end.x {
			for i := l.start.y; i >= l.end.y; i-- {
				returnX = append(returnX, i)
			}
		} else {
			for i := l.start.x; i <= l.end.x; i++ {
				returnX = append(returnX, i)
			}
		}
		if l.start.y > l.end.y {
			for i := l.start.y; i >= l.end.y; i-- {
				returnY = append(returnY, i)
			}
		} else {
			for i := l.start.y; i <= l.end.y; i++ {
				returnY = append(returnY, i)
			}
		}
	}
	return returnX, returnY
}

func (v *vents) analyze() {
	fmt.Println("v.overlaps length", len(v.overlaps))
	count := 0
	for _, v := range v.overlaps {
		count += v
	}

	fmt.Println("v.overlaps total", count)
}

func (v *vents) display() {
	for y := range v.chart {
		for x := range v.chart[y] {
			if v.chart[y][x] > 0 {
				fmt.Printf("%d", v.chart[y][x])
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println()
	}
}

func (v *vents) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(coordRegExp)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		startX, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return err
		}
		startY, err := strconv.Atoi(matches[0][2])
		if err != nil {
			return err
		}
		endX, err := strconv.Atoi(matches[0][3])
		if err != nil {
			return err
		}
		endY, err := strconv.Atoi(matches[0][4])
		if err != nil {
			return err
		}
		lp := line{
			start: point{
				x: startX,
				y: startY,
			},
			end: point{
				x: endX,
				y: endY,
			},
		}
		v.data = append(v.data, lp)
		//fmt.Println(lp)
	}
	v.overlaps = make(map[point]int, 0)

	for _, l := range v.data {
		switch {
		//if x is the same, range on y
		case l.IsVertical():
			_, iterY := l.Iterator()
			for _, y := range iterY {
				p := point{x: l.start.x, y: y}
				v.chart[y][l.start.x]++
				if v.chart[y][l.start.x] >= 2 {
					if _, ok := v.overlaps[p]; !ok {
						v.overlaps[p] = 1
					} else {
						v.overlaps[p]++
					}
				}
			}

		//if y is the same, range on x
		case l.IsHorizontal():
			iterX, _ := l.Iterator()
			for _, x := range iterX {
				p := point{x: x, y: l.start.y}
				v.chart[l.start.y][x]++
				if v.chart[l.start.y][x] >= 2 {
					if _, ok := v.overlaps[p]; !ok {
						v.overlaps[p] = 1
					} else {
						v.overlaps[p]++
					}
				}
			}
		}
	}
	return nil
}

func main() {
	var v vents
	err := v.load("data/vents_sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	//v.display()
	v.analyze()
	fmt.Println()
}
