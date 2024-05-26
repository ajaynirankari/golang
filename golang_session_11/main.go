package main

import "fmt"

func main() {

	f := fact(5)
	fmt.Println("Factorial of 5 = ", f)

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-2) + fib(n-1)
	}

	output := fib(7)
	fmt.Println("fib output = ", output)

	v := 5
	fmt.Println("main v = ", v)
	change(v)
	fmt.Println("main v = ", v)

	anotherV := 5
	fmt.Println("main anotherV = ", anotherV, "&anotherV = ", &anotherV)
	shareValueViaAddress(&anotherV)
	fmt.Println("main anotherV = ", anotherV)

}
func change(v int) {
	fmt.Println("change v = ", v)
	v = 0
	fmt.Println("change v = ", v)
}

func shareValueViaAddress(v *int) {
	fmt.Println("before shareValueViaAddress v = ", v, " *v = ", *v)
	*v = 0
	fmt.Println("after shareValueViaAddress v = ", v, " *v = ", *v)
}

func fact(n int) int {
	// fmt.Println("Call n = ",n)
	// if n < 2 {
	// 	fmt.Println("return n < 2 = ",n)
	// 	return n
	// }
	// r := fact(n-1)
	// fmt.Println("n = ",n,"call fact(",(n-1),")")
	// f := n * r
	// fmt.Println("return n = ",n,"call fact(",(n-1),") return = ",r, " (",n, " * ",r,") = ",f)
	// return f

	if n < 2 {
		return n
	}
	return n * fact(n-1)
}
