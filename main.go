package main

import (
	"./greeting"
	"fmt"
)

func main() {
	fmt.Println(greeting.Get())
	fmt.Println(greeting.Hello())
}
