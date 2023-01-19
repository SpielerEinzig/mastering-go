package main

import (
	"fmt")

func main () {

	var name string = "Ryan"
	var age = 1000

	//print line prints in separate lines as
	//opposed to the print
	fmt.Println(name, age)

	//print
	fmt.Print(name)
	fmt.Print(age)

	//formatting strings
	//print f is used to print formatted outputs to the console
	//sprint f saves formatted outputs to a variable
	//%v for varables, %q to quote the variable (only works on strings)
	//%T for the type of variable
	fmt.Printf("Hello, my name is %q and I am %v years old", name, age)
	var data = fmt.Sprintf("Hello, my name is %q and I am %v years old", name, age)
	fmt.Println(data)
} 