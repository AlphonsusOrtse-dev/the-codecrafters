package main

import (
	"fmt"
	"strconv"
)

var hex string

func hexToDex() {
	fmt.Scan(&hex)
	value, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println("Choose a valid hexadecimal!")
	} else {
		fmt.Println(value)
	}
	
}

func main() {
	hexToDex()
}