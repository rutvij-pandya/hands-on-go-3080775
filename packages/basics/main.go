// packages/basics/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func greet() string {
	return "Hey!"
}

func greetWithName(name string) string {
	return "My name is " + name + "!"
}

func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "My name is " + name + " and I'm " + strconv.Itoa(age) + " years old."
	return
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return b, errors.New("Cannot divide by 0!")
	}
	return a / b, nil
}

func main() {
	// fmt.Println(greet())
	// fmt.Println(greetWithName("Adam"))
	// fmt.Println(greetWithNameAndAge("Adam", 25))

	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))

	// fmt.Println("Hi there! This is GOlang.")
	fmt.Printf("Today is %s \n", time.Now().Weekday())
}
