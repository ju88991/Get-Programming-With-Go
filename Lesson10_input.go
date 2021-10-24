package main

import (
	"fmt"
)

func main() {
	text := "1"
	text_bool := true
	if (text == "true") || (text == "yes") || (text == "1") {
		text_bool = true
	} else if (text == "false") || (text == "no") || (text == "0") {
		text_bool = false
	} else {
		fmt.Println("Error.")
	}
	msg := fmt.Sprintf("%v", text_bool)
	fmt.Println(msg)
}
