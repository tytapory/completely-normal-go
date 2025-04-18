// Arthouse of my work

package main

import (
	"completely_normal_go/fibonacci"
	"completely_normal_go/simpleprintablestring"
	"fmt"
)

func main() {
	fmt.Printf("Presentation of my simple printable string implementation\n")
	mySimpleString := simpleprintablestring.New("Hello world!")
	mySimpleString.Print()

	fmt.Printf("\n\nThis is my fibonacci\n")
	fmt.Printf("%d", fibonacci.Fibonacci(20))
}
