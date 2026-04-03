# Go functions

functions are reusable block of code or set of instructions to achieve a particular task.

## creating a Go function
* Also caled function declaration. this Uses func keyword, followed by name, parameters, return types, and the body in {}.


#### for example:
```go
package main
import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println(result) 
}
```
