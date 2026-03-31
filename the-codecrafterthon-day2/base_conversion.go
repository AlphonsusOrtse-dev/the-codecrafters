// THE BASE CONVERTER
//   Concept: Number Systems & strconv

//  Rules:
//   → DO NOT USE GOOGLE OR ANY AI TOOL; DO IT YOURSELF - you can share ideas with others, but not copy code.
//   → Write everything in Go. Standard library only.
//   → Must compile and run without errors.
//   → Push to your the-codecrafters repo in a folder
//     named: thecodecrafterthon-day2/
//   → Include a README.md explaining how to run it

//   go-reloaded needs to convert hex and binary
//   strings into decimal numbers. This project
//   teaches you exactly that — and makes you think
//   about what happens when the input is garbage.

//   Build a CLI tool that converts numbers between
//   bases. It runs from the terminal like this:

//      > convert 1E hex
//        ✦ Decimal: 30

//      > convert 10 bin
//        ✦ Decimal: 2

//      > convert 255 dec
//        ✦ Binary:  11111111
//        ✦ Hex:     FF

//   Requirements:

//   → Support three input bases: hex, bin, dec.
//   → For dec input, output both binary and hex.
//   → For hex and bin input, output only decimal.
//   → All hex output must be uppercase.
//   → The program keeps running until: quit

//   Validation — handle all of these cleanly:
//   → "1G" is not valid hex.
//   → "10201" is not valid binary.
//   → "abc" is not a valid decimal.
//   → Negative numbers must be supported for dec.
//   → Empty input must not crash the program.

//   Think about:
//   → Which strconv functions handle base
//     conversions for you?
//   → How do you detect which characters are
//     valid for a given base?
//   → What is the difference between ParseInt
//     and ParseUint — and does it matter here?

//   ✦ Nail this and (hex) / (bin) in go-reloaded
//     becomes a 5-minute task.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hexToDec(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)
}

func binToDec(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 64)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

start:
	fmt.Print("> ")

	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)

	if input == "" {
		goto start
	}

	command := strings.ToLower(input)

	if command == "quit" {
		fmt.Println("Goodbye!")
		return
	}

	if command == "help" {
		fmt.Println("=== Number Base Converter ===")

		fmt.Println("Command:")
		fmt.Println("  convert <number> <base>")
		fmt.Println("Supported bases:")
		fmt.Println("  hex  → converts to decimal")
		fmt.Println("  bin  → converts to decimal")
		fmt.Println("  dec  → converts to binary and hex")
		fmt.Println("Examples:")
		fmt.Println("  > convert 1E hex")
		fmt.Println("     Decimal: 30")
		fmt.Println("  > convert 255 dec")
		fmt.Println("     Binary: 11111111")
		fmt.Println("     Hex:    FF")
		fmt.Println("Other commands:")
		fmt.Println("  help  → show this message")
		fmt.Println("  quit  → exit program")
		goto start
	}

	word := strings.Fields(input)

	if len(word) != 3 {
		fmt.Println("Invalid number of arguments")
		goto start
	}

	if word[0] != "convert" {
		fmt.Println("Invalid command")
		goto start
	}

	if word[2] != "hex" && word[2] != "bin" && word[2] != "dec" {
		fmt.Println("Invalid base (use hex/bin/dec)")
		goto start
	}

	switch word[2] {

	case "hex":
		val, err := hexToDec(word[1])
		if err != nil {
			fmt.Println("Invalid hex input")
			goto start
		}
		fmt.Println(" Decimal:", val)
		goto start

	case "bin":
		val, err := binToDec(word[1])
		if err != nil {
			fmt.Println("Invalid binary input")
			goto start
		}
		fmt.Println(" Decimal:", val)
		goto start

	case "dec":
		val, err := strconv.ParseInt(word[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid decimal input")
			goto start
		}

		bin := strconv.FormatInt(val, 2)
		hex := strings.ToUpper(strconv.FormatInt(val, 16))

		fmt.Println(" Binary: ", bin)
		fmt.Println(" Hex:    ", hex)
		goto start
	}
}
