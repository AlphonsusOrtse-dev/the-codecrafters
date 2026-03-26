* Go syntax is the consistent arrangement of Go files for easy and better readability.

### Every Go program must have the following for it to succesfully run:
1. package main
2. import package
3. the function main and 
4. the your code

### For example:
```
package main

import "fmt"
func main() {
    a := "have "
    a += "you"
    fmt.Println(a)
}
```