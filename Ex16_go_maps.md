# Go maps
* Maps are used to store data values in key:value pairs.
* Every item in a map has a unique key. You use that key to quickly retrieve, update, or delete the corresponding value.
* The length of a map is the number of its elements. You can find it using the len() function.
* The default value of a map is nil.

### how to create maps in go:
contacts := map[string]string{
    "alice": "08012345678",
    "bob":   "08098765432",
    "carol": "08055544433",
}

map[string]string simply means:

* the key is a string (a name)
* the value is a string (a phone number)

### creating an empty map
contacts := make(map[string]string)

### Adding something to the map
contacts["david"] = "08011122233"
* Now "david" is in the map with his number.

### Updating a value
* Like crossing out an old number and writing a new one
contacts["david"] = "08099999999"
* here, david's old number is gone, replaced with the new one.

### Deleting from the map
delete(contacts, "bob")
* Bob is gone from the map forever.

### Looping Through Maps with range
* basic syntax:
for key, value := range myMap {
}

#### example:
```go
package main

import "fmt"

func main() {
	ages := map[string]int{
		"ada":  25,
		"john": 30,
		"ben":  17,
	}

	for name, age := range ages {
		fmt.Println(name, "is", age, "years old")
	}
}
```
you use this when you want to ignore index
for name, _ := range ages {
    fmt.Println(name)
}

