package main

import "fmt"

func main() {

	// var a int = 10
	// fmt.Println("Value of a is: ", a)

	/*
	x:= 20
	y:= 40

	sum:=x+y
	fmt.Println("Sum of x and y is: ", sum)

	subtract:=x-y
	fmt.Println("Subtraction of x and y is: ", subtract)

	multiply:=x*y
	fmt.Println("Multiplication of x and y is: ", multiply)

	divide:=y/x
	fmt.Println("Division of x and y is: ", divide)


	a1:= 15
	b2:= 2

	c2:=a1%b2	// Modulus of a and b
	fmt.Println("Modulus of a and b is: ", c2)

	one:= 24
	v1 := one%2
	fmt.Println("Modulus of one and 2 is: ", v1)

	

	a2:= 10
	a2 = a2 + 10
	fmt.Println("Value of a2 is: ", a2)

	a2 = 10
	a2 += 10
	fmt.Println("Value of a2 is: ", a2)

	a1 = 10
	a2 = 20
	b3 := a2==a1
	fmt.Println("a2==a1 is: ", b3)

	b3 = a2 != a1
	fmt.Println("a2 !=a1 is: ", b3)

	fmt.Println("!b3 is: ", !b3)

	VAR1 := 10;
	fmt.Println("Value of VAR1 is: ", VAR1)

	var1 := 10;
	fmt.Println("Value of var1 is: ", var1)	
*/
	sumOfTwoFixNumbers()

	message := sayHello()
	fmt.Println("Message is: ", message)

	sumOfTwoNumbers(10, 20)
	sumOfTwoNumbers(110, 20)
	sumOfTwoNumbers(101, 20)

	result := takeTwoNumbersReturnSum(10, 20)
	fmt.Println("Sum of two numbers is: ", result)

	 
}

func sayHello() string {
	return "Hello"
}

func takeTwoNumbersReturnSum(x int, y int) int {
	sum:=x+y
	return sum
}

func sumOfTwoNumbers(x int, y int) {

	sum:=x+y
	fmt.Println("Sum of x and y is: ", sum)
}

func sumOfTwoFixNumbers(){
	x:= 20
	y:= 40

	//sum:=x+y
	
	//sum:= takeTwoNumbersReturnSum(x, y)
	//fmt.Println("Sum of x and y is: ", sum)

	sumOfTwoNumbers(x, y)
}