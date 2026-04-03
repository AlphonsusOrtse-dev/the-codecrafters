# Go Data Types

Go Data Type is a concept in Golang that tell the computer what kind of data a variable holds. And because Go is statically typed, any asigned data value can not be changed.

### Go basic data types:

1. Bool: prints true or false
2. Numeric: prints whole numbers, both positive and negative integers
3. string: prints text.
4. float64: prints decimal values.

#### For example:
```go
package main

import "fmt"

func main() {
	var isTrue bool = true
	var name string = "car"
	var num int = 2

	fmt.Printf("Is it %t that the %s has ony %d sits in the front?\n", isTrue, name, num)

}
```