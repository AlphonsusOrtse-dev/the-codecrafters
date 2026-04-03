# Go for loop

In Go, loop is a way of telling the computer to repeating a block of code until a certain condition is met.

#### for example:

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}

### nested loop:
A loop in side another loop.

#### for example:
for i := 0; i < 3; i++ {        
    for j := 0; j < 3; j++ {    
        fmt.Println(i, j)
    }
}
