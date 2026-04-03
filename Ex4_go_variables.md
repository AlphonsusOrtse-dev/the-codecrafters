# variables
* A variable is a box or a placeholder for data values.

### there are different types of variable types in go:
* int: stores whole numbers e.g 1,2,3... or -1, 2, 3, 4
* float32: this stores decimal numbers e.g 3.95 or - 3.954
* string: stores text e.g "good"
* bool: this stores conditional statements e.g true or false

### Declaring a variable in Go:
* the (var) keyword is used for declaring variable in Go. A variable name and a type of data is followed when declarating your variable. e.g:
1. var name string "Alphons"
2. name := "Alphons"---this ia called short variable declaration in Go.

#### for example: 
```go
package main

import "fmt"
func main() {
    var name string = "Alphons"
    fmt.Println(name)
}
```


### Variable Declaration Without Initial Value
* In Go, variables declared without asigning initial values will return their  default values.

#### for example:
```go
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a) ---output = (empty string)
  fmt.Println(b) ---output = (0)
  fmt.Println(c) ---output = (false)
}
```