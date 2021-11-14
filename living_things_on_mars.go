// Lesson 25
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (a Animal) String() string {
	return fmt.Sprintf("%v", a.name)
}

type Animal struct {
	name string
}

func (a Animal) move(s string) string {
	rand := rand.Intn(3)
	d := ""
	switch rand {
	case 0:
		d = "North"
	case 1:
		d = "South"
	case 2:
		d = "East"
	case 3:
		d = "West"

	}
	fmt.Printf("%v is heading to %v.\n", s, d)
	return strings.ToLower(d)
}
func (a Animal) eat(s string) string {
	rand := rand.Intn(3)
	d := ""
	switch rand {
	case 0:
		d = "Feds"
	case 1:
		d = "Meats"
	case 2:
		d = "Fish"
	case 3:
		d = "Insects"

	}
	fmt.Printf("%v is eating %v.\n", s, d)
	return strings.ToLower(d)
}
func timesleep(i int) {
	t := time.Duration(i)
	time.Sleep(time.Second * t)
}
func main() {
	for day := 0; day < 3; day++ {
		lifeLog := make(map[string]string, 5)
		fmt.Println("A new day has begun.")
		timesleep(3)
		for i := 0; i < 12; i++ {
			r := rand.Intn(5)
			name := ""
			switch r {
			case 0:
				name = "Rabbit"
			case 1:
				name = "Bear"
			case 2:
				name = "Duck"
			case 3:
				name = "Dog"
			case 4:
				name = "Cat"

			}
			lifeLog[name] = ""
			animal := Animal{name}
			a := rand.Intn(2)
			if a == 0 {
				place := animal.move(animal.String())
				lifeLog[name] = fmt.Sprintf("%v is somewhere in %v.", name, place)
			} else {
				food := animal.eat(animal.String())
				lifeLog[name] = fmt.Sprintf("%v is just finished to eat %v.", name, food)
			}
			timesleep(1)
		}
		fmt.Printf("\nEnd of Day...\n")
		timesleep(3)
		fmt.Println("--last activity--")
		timesleep(2)
		for _, a := range lifeLog {
			fmt.Printf("%v\n", a)
			timesleep(2)
		}
		if day != 2 {
			fmt.Printf("\nAll living things on Mars is sleeping for 12 hours...\n\n")
			timesleep(5)
		} else {
			fmt.Printf("\n\nEnd of observation.")
		}
	}
}
