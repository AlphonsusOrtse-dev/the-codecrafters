# Go Conditions
In Go, conditions are how your program maks decisions,They allow you to run specific blocks of code only when certain conditions are met.conditions can  eithe be true or false.


### Go has the following conditional statements:
1. the if statement: It checks if a condition is true.

#### for example:
package main

import "fmt"

func main() {
	b := "see"
	if b == "see" {
		fmt.Println("you can see me")
	}
}

2. the else if: Use else if to specify a new condition to test, if the first condition is false

#### for example:

package main

import "fmt"

func main() {
	b := "see"
	if b != "see" {
		fmt.Println("you get wahala")
	} else {
		fmt.Println("you can see")
	}
}

3. else statement: Use else to specify a block of code to be executed, if the same condition is false

#### for example:

package main

import "fmt"

func main() {
	b := "see"
	if b != "see" {
		fmt.Println("you get wahala")
	} else if b == "code" {
		fmt.Println("you can code anytime of the night")
	} else {
		fmt.Println("I can't come and kill myself")
	}
}

