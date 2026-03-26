# Go Output Functions

* this are functions that allow Go to print text and format data to the console.

### Go has three functions to output text:

1. Print()- the prints the default value to the console
2. Println() - the prints the result with a new line
3. Printf() - this allows users to print formatted data to the console.

#### for example:
package main

import "fmt"

func main() {
	a := "code"
	b := "crafters"

	fmt.Print(a + b)
	fmt.Println(a + b)
	fmt.Printf("the %s\n", a+b)

}
