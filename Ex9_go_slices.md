# Go Slices

slices are like modified arrays, but the squre brackets [] are empty. unlike arrays the length of a slice can be grown or shrinked. that is items can be added or reduced.

### Various ways of creating a slice

1. Using the []datatype{values} Format (Slice Literal)
* this is when you have a specific number of values you want to create a slice with

#### for example:
package main

import "fmt"

func main() {
	mySlice := []string{"work", "hard", "Alphonsus"}
	fmt.Printf("sclice: %v\n", mySlice)
	fmt.Printf("capacity: %v\n", cap(mySlice))
	fmt.Printf("length: %v\n", len(mySlice))

}

2. creating a slice fron an array
* this is when you have a fix len of an array but still want to work with a specific number. you can use the Append keyword to add values to a slice.
#### for example:
package main

import "fmt"

func main() {
	mision := []string{"work", "hard", "Alphonsus"}
	mision = append(mision, "Ortse")
	fmt.Println(mision)
}

* you can also reduce the length of a slice like this:
package main

import "fmt"

func main() {
	s := []string{"come", "to", "my", "table", "immediately"}
	s = s[:len(s)-1]
	fmt.Println(s)
}

3. Using the make() Function
* this is use when you don't have the data you want to use yet but know how much space you wil need work with. probably may want add some data later.

package main

import "fmt"

func main() {
	items := make([]string, 0, 6) // only this wiil print an empty slice []
	items = append(items, "a", "b", "c", "d") //// output = [a b c d]

	fmt.Println(items) 

}