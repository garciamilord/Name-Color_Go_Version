package main

import (
	"fmt"
)

func main() {

	name, color := "", ""

	fmt.Printf("\nEnter a name and ")
	fmt.Printf("your favorite color: ")
	fmt.Scan(&name)
	fmt.Scan(&color)

	fmt.Printf("My name is %s and my favorite color %s", name, color)
}
