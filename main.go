package main

import (
	"flag"
	"fmt"
)

func main() {
	helloName := flag.String("name", "", "name to be called.")
	flag.Parse()
	if *helloName == "" {
		fmt.Println("hello")
	} else {
		fmt.Printf("hello %s\n", *helloName)
	}
}
