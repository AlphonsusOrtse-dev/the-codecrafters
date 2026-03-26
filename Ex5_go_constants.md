# Go Constants

* Go constants are values that can not be changed after been declared. the keyword (const) is used for declaring constants.

#### for example:
package main
import "fmt"

const Name = "Alphons"
Name = "john"

func main() {
  fmt.Println(Name)
} 
* here, you will have an error "expected declaration, found Name" because can not reasign another value to the variable, Name.

### types of constants


1. typed constanats; they are declared with a particular data type. indicating the type of data passed in. either string, int,  float or bool.

#### for example:

package main
import ("fmt")

const A string = "Good"

func main() {
  fmt.Println(A)
}

2. untyped constants: this are declared without indicating the data type, the compiler has to choose basse on the type of value passed in.

#### for example:

 package main
import ("fmt")

const A = 4

func main() {
  fmt.Println(A)
} 
* the compiler automatically prints the value (4) because it was able to read the intiger type of data passed in.