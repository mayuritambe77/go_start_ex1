package main //a package is a way to group functions, and it's made up of all the files in the same directory

import "fmt" // popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.

func main() {
	fmt.Println("Hello, World!") // A main function executes by default when you run the main package.
}

//variables - var keyword then name then type = value

var number int = 41
var name string = "Gopher"
var isTrue bool = true
var float float64 = 3.14
