# Go switch statement

This is a cleaner way to write long if/else chains, especially when you are comparing one variable against many specific values.

* The default keyword is optional. It specifies some code to run if there is no case match.

### for example:
```go
package main

import "fmt"

func main() {

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Back to work!")
	case "Friday":
		fmt.Println("Weekend is close!")
	default:
		fmt.Println("Just another day.")
	}
}
```

