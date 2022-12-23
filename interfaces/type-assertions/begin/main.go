// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var i any = "Hello"
	// var i any = 10
	// fmt.Println(i.(string))

	// perform a type assertion and handle the error
	if _, ok := i.(int); !ok {
		fmt.Printf("%v is not an integer \n", i)
	}
}
