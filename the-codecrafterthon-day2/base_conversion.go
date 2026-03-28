package main

import (
	"fmt"
	"strconv"
)

var hex string
var num string

func hexToDex(hex string) (int64, error) {
	value, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return value, nil

}

func main() {
	fmt.Println("enter number")
	fmt.Scan(&num)
	fmt.Println(hexToDex(num))

}
