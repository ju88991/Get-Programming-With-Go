// Lesson 15
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var lineWidth = 40
var space = 13

func printHeader(s, d string) {
	fmt.Println(strings.Repeat("=", lineWidth))
	fmt.Printf("|%v%v%v|%v%v%v|\n", s, strings.Repeat(" ", space), "", d, strings.Repeat(" ", space+2), "")
	fmt.Println(strings.Repeat("=", lineWidth))
}

func drawTable(tempreture float64, mode string) {
	var fahren float64
	var cels float64
	var celsLen int
	var fahrenLen int

	if mode == "cels" {
		cels = tempreture
		fahren = celsiusToFahrenheit(cels)
		celsLen = len(strconv.Itoa(int(cels)))
		fahrenLen = len(strconv.Itoa(int(fahren)))
		fmt.Printf("|%6.2f%v|%6.2f%v|\n", cels, strings.Repeat(" ", space-celsLen), fahren, strings.Repeat(" ", space-fahrenLen))
	} else if mode == "fahen" {
		fahren = tempreture
		cels = fahrenheitToCelsius(fahren)
		celsLen = len(strconv.Itoa(int(cels)))
		fahrenLen = len(strconv.Itoa(int(fahren)))
		fmt.Printf("|%6.2f%v|%6.2f%v|\n", fahren, strings.Repeat(" ", space-fahrenLen), cels, strings.Repeat(" ", space-celsLen))
	}
}

func printFooter() {
	fmt.Println(strings.Repeat("=", lineWidth))
}

func celsiusToFahrenheit(cels float64) float64 {
	return (cels * 9 / 5) + 32.0
}
func fahrenheitToCelsius(fahen float64) float64 {
	return (fahen * 5 / 9) - 17.8
}

func main() {

	startTemp := 10.0
	endTemp := 20.0

	nextStartTemp := celsiusToFahrenheit(startTemp)
	nextEndTemp := celsiusToFahrenheit(endTemp)
	nextStep := 2.0

	printHeader("[C]", "[F]")
	for startTemp <= endTemp {
		drawTable(startTemp, "cels")
		startTemp += nextStep
	}
	printFooter()

	startTemp = nextStartTemp
	endTemp = nextEndTemp
	printHeader("[F]", "[C]")
	for startTemp < endTemp {
		drawTable(startTemp, "fahen")
		startTemp += nextStep
	}
	printFooter()
}
