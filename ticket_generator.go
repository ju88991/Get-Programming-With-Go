package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var len = 25 + 10 + 20 + 5
	fmt.Printf("%-25v%-10v%-20v%v\n", "SpaceLine", "Days", "Trip type", "Price")
	fmt.Println(strings.Repeat("=", len))
	for i := 0; i < 10; i++ {
		var spaceline = ""
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "Space X"
		case 2:
			spaceline = "Virgin Galactic"
		}
		// var depart_day = "2020-10-13"
		var distance = 62100000
		var speed_s = 36 + rand.Intn(51-36)
		var speed_d = speed_s * 60 * 60 * 24
		var days = distance / speed_d
		var price = speed_s + 20
		var trip_type = ""
		if rand.Intn(2) == 0 {
			trip_type = "One-way"
		} else {
			trip_type = "Round-Trip"
			price *= 2
		}
		fmt.Printf("%-25v%-10v%-20v%v\n", spaceline, days, trip_type, price)
	}
}
