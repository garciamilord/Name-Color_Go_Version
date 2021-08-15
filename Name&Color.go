package main

import "fmt"

func red() {
	name, color := "", ""

	fmt.Printf("\nEnter a name and ")
	fmt.Printf("your favorite color: ")
	fmt.Scan(&name)
	fmt.Scan(&color)
	if color == "blue" {
		fmt.Printf("My name is %s and my favorite color %s", name, color)
	} else if color == "red" {
		fmt.Printf("My name is %s and my favorite color %s", name, color)
	} else if color == "yellow" {
		fmt.Printf("My name is %s and my favorite color %s", name, color)
	} else if color == "green" {
		fmt.Printf("My name is %s and my favorite color %s", name, color)
	} else {
		fmt.Printf("wrong answer")
		red()

	}

}

func main() {
	red()

}
