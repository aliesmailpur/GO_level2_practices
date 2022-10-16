package main

import "fmt"

// write a program that
// assigns an int to a variable
// prints that int in decimal, binary, and hex
// shifts the bits of that int over 1 position to the left
// prints that variable in decimal, binary, and hex

var a int = 800

func main(){

	fmt.Printf("%d\t%b\t%#x", a , a , a)

	b:= a << 1

	fmt.Printf("%d\t%b\t%#x", b , b , b)

}