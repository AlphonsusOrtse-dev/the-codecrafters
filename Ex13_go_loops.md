# Go for loop

In Go, loop is a way of telling the computer to repeating a block of code until a certain condition is met.

#### for example:
```go
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

### nested loop:
A loop in side another loop.

#### for example:
```go
for i := 0; i < 3; i++ {        
    for j := 0; j < 3; j++ {    
        fmt.Println(i, j)
    }
}
```


### The Range Keyword
* The range keyword is used to iterate through the elements of an array, slice or map and returns both the index and the value.

#### for example:
```go
package main

import "fmt"

func main() {
	s := []string{"welcome", "to", "the", "world", "where", "problem", "is", "your", "friend"}
	for i, v := range s {
		fmt.Println(i, v)

	}
}
```