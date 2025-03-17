package main //a package is a way to group functions, and it's made up of all the files in the same directory

import "fmt" // popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.

func main() {
	fmt.Println("Hello, World!") // A main function executes by default when you run the main package.
}

//variables - var keyword then name then type = value

var num int = 41 // int8 , int16 , int32 , int64
var name string = "Gopher"
var isTrue bool = true
var float float64 = 3.14 //float32 , float64

//short hand declaration
number := 41
name := "Gopher"	//type inference
isTrue := true
float := 3.14

//constants - const keyword
const number int = 41	//cannot be changed	
const name string = "Gopher"
const isTrue bool = true
const float float64 = 3.14

//multiple variables
var (
	number int = 41		
	name string = "Gopher"
	isTrue bool = true
	float float64 = 3.14
)
//multiple variables short hand
number, name, isTrue, float := 41, "Gopher", true, 3.14

//zero values
var num int //0
var name string //""		
var isTrue bool //false
var float float64 //0

//type conversion	
var num int = 41
var float float64 = float64(num)

//type inference
var num = 41
var float = 3.14

//type alias
type myInt int
var num myInt = 41

//string concatenation
var name string = "Gopher"
var message string = "Hello, " + name

//string interpolation
var name string = "Gopher"
var message string = fmt.Sprintf("Hello, %v", name)

//string formatting
var name string = "Gopher"
var message string = fmt.Sprintf("Hello, %q", name)

//string length
var name string = "Gopher"
var length int = len(name)

//string indexing
var name string = "Gopher"
var firstChar byte = name[0]


//string slicing
var name string = "Gopher"
var slice string = name[0:3] //Gop

//runes
var char rune = 'A' //rune is an alias for int32

//****************************************************************************

//Functions
func main() {
	fmt.Println("Hello, World!") // A main function executes by default when you run the main package.
	add(1,2)
}

func add(a int, b int) int //return type is mentioned after the parameters
{
	return a + b
}








