// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "global"
var pointer *int
func main() {
	// create a local variable "val" and assign it an int value
	val := 42
	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v \n", val, val)
	// print the value of the package-level variable "val"
	printVariable()
	// update the package-level variable "val" and print it again
	updateVariable("updated global")
	printVariable()
	// print the pointer address of the local variable "val"
	fmt.Printf("local value = %v, local type = %T \n", val, val)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	pointer = &val
	fmt.Printf("%T, local &val = %v \n", pointer, pointer)
	*pointer = 100
	fmt.Printf("local val = %v \n", *pointer)

}

func updateVariable(str string) {
	val = str
}

func printVariable() {
	fmt.Printf("global val = %v \n", val, val)
}
