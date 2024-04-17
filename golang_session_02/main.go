package main

import "fmt"

func main() {

	var name string   						   // Declaring a variable
	name="Tim"        						// Assigning a value to the variable

	fmt.Printf("Type = %T\n", name)               // use the variable
	name="John"  
	
	fmt.Printf("Type = %T, Value = %v\n", name, name)  

	var x int
	x=7
	fmt.Println(x)

	var isMale bool

	isMale=true

	fmt.Println(isMale)

	var a float64 = 56.89
	fmt.Println(a)

	
	age := 70   // Declaring and assigning a value to a variable

	fmt.Printf("Type = %T \n", age)

	var age1 uint = 70
	fmt.Println(age1)

	var test byte = 255
	fmt.Println(test)

	var test1 rune = 2550
	fmt.Println(test1)

	rune1 := 'A'
	fmt.Println(rune1)
	fmt.Printf("Type = %T ", rune1)

	//rune1 := 'A'      // Duplicate variable is not allowed

	test1 = 'B'
	fmt.Println(rune1)
	fmt.Printf("Type = %T \n", rune1)


	//Data types

	// byte, rune, 
	// int, int8, int16, int32, int64, 
	// uint, uint8, uint16, uint32, uint64, uintptr
	// float32, float64, 
	// complex64, complex128, 
	// string, 
	// bool

	const pi = 3.14159
	fmt.Println(pi)

	// pi = 4.14159   // Error: cannot assign to cosnt variable pi
}
