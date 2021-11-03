// Lesson 20
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	width     = 80
	height    = 25
	seed      = 35
	millisec  = 300
	max_phase = 100
)

type Universe [][]bool

func NewUniverse() Universe {
	univ := make([][]bool, height)
	for i := range univ {
		univ[i] = make([]bool, width)
	}
	return univ
}

func Show(univ Universe) {
	fmt.Println(strings.Repeat("-", width))
	for _, line := range univ {
		for _, u := range line {
			if u {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println(strings.Repeat("-", width))
}

func (u Universe) Seed() {
	for _, line := range u {
		for i := range line {
			if rand.Intn(seed) == 0 {
				line[i] = true
			}
		}
	}
}

func (u Universe) CheckAlive(x, y int) bool {
	if x < 0 || x >= width {
		x = (x + width) % width
	}
	if y < 0 || y >= height {
		y = (y + height) % height
	}
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	var alive int
	if u.CheckAlive(x, y+1) {
		alive++
	}
	if u.CheckAlive(x, y-1) {
		alive++
	}
	if u.CheckAlive(x-1, y) {
		alive++
	}
	if u.CheckAlive(x+1, y) {
		alive++
	}
	if u.CheckAlive(x+1, y+1) {
		alive++
	}
	if u.CheckAlive(x-1, y-1) {
		alive++
	}
	if u.CheckAlive(x+1, y-1) {
		alive++
	}
	if u.CheckAlive(x-1, y+1) {
		alive++
	}
	return alive
}

func (u Universe) Next(x, y int) bool {
	neighbors := u.Neighbors(x, y)
	if neighbors < 2 {
		return false
	}
	if neighbors > 3 {
		return false
	}
	return true
}

func Step(a, b Universe) {
	work := height
	time.Sleep(millisec * time.Millisecond)
	for work > 0 {
		fmt.Println(" ")
		work--
	}
	a, b = b, a
}

func main() {
	univ := NewUniverse()
	univ2 := NewUniverse()
	univ.Seed()
	phase := 1
	for phase < max_phase+1 {
		fmt.Printf("Phase: %v / %v\n", phase, max_phase)
		Show(univ)
		Step(univ, univ2)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if univ.Next(i, j) {
					univ[i][j] = true
				} else {
					univ[i][j] = false
				}
			}
		}
		phase++
	}
}
