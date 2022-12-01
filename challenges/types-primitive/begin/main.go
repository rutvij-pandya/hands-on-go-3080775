// challenges/types-primitive/begin/main.go
package main

import (
	"fmt"
	"time"
)

// use type inference to create a package-level variable "val" and assign it a value of "global"

func main() {
	// Defining duration parameter of AfterFunc() method
	DurationOfTime := time.Duration(3) * time.Second

	fmt.Println("====== Begin ======")
	// Defining function parameter of AfterFunc() method
	f := func() {
		// Printed when its called by the AfterFunc() method in the time stated above
		fmt.Println("Function called by AfterFunc() after 3 seconds")
	}

	// Calling AfterFunc() method with its parameter
	Timer1 := time.AfterFunc(DurationOfTime, f)

	// Calling stop method w.r.to Timer1
	defer Timer1.Stop()

	// Calling sleep method
	time.Sleep(10 * time.Second)
	fmt.Println("====== End ======")
}
