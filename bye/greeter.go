package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Bye")
}

var Greeter greeting
