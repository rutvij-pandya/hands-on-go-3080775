// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors = map[string]author{
		"rp": {name: "rutvij"},
		"sa": {name: "shubhangi"},
	}

	authors["rp"] = author{name: "r. pandya"}

	delete(authors, "sa")

	// initialize the map with make
	// authors = make(map[string]author)

	// add authors to the map
	// authors["rp"] = author{name: "rutvij"}
	// authors["sa"] = author{name: "shubhangi"}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	fmt.Println(authors["rp"])
}
