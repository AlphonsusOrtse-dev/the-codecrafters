package main

import (
	"fmt"
	"strconv"
)


func hexToDex(hex string) (int64, error) {
	value, err := strconv.ParseInt(hex, 16, 64)
	return value, err
}

func main() {
	fmt.Println(hexToDex("1E"))
}