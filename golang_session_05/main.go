package main

import "fmt"

func main() {

	i := 3
	fmt.Print("Write ", i, " as ")
	switchTest(i)

	i = 1
	fmt.Print("Write ", i, " as ")
	switchTest(i)

	i = 12
	fmt.Print("Write ", i, " as ")
	switchTest(i)

	fmt.Printf("grade(95): %v\n", grade(95))
	fmt.Printf("grade(85): %v\n", grade(85))
	fmt.Printf("grade(75): %v\n", grade(75))
	fmt.Printf("grade(65): %v\n", grade(65))
	fmt.Printf("grade(55): %v\n", grade(55))

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		case bool:
			fmt.Println("I'm a bool")
		default:
			fmt.Println("I'm another type")
		}
	}

	whoAmI(1)
	whoAmI("1")
	whoAmI(true)
	whoAmI(1.0)

	fmt.Printf("factorial(5): %v\n", factorial(5))
	fmt.Printf("factorial(4): %v\n", factorial(4))
	fmt.Printf("factorial(7): %v\n", factorial(7))

	fmt.Printf("isPrimeNumber(7): %v\n", isPrimeNumber(7))
	fmt.Printf("isPrimeNumber(8): %v\n", isPrimeNumber(8))
	fmt.Printf("isPrimeNumber(9): %v\n", isPrimeNumber(9))
	fmt.Printf("isPrimeNumber(10): %v\n", isPrimeNumber(10))
	fmt.Printf("isPrimeNumber(11): %v\n", isPrimeNumber(11))
	fmt.Printf("isPrimeNumber(12): %v\n", isPrimeNumber(12))

	for i := range 100 {
		fmt.Printf("isPrimeNumber(%v): %v\n", i, isPrimeNumber(i))
	}
	printOnlyPrimeNumbers(100)
	fmt.Println()
	printOnlyNonPrimeNumbers(100)

	fmt.Printf("digitSum(12345): %v\n", digitSum(12345))
	fmt.Printf("digitSum(1234): %v\n", digitSum(1234))
	fmt.Printf("digitSum(123): %v\n", digitSum(123))
	fmt.Printf("digitSum(12): %v\n", digitSum(12))

	fmt.Printf("digitCount(12345): %v\n", digitCount(12345))
	fmt.Printf("digitCount(1234): %v\n", digitCount(1234))
	fmt.Printf("digitCount(123): %v\n", digitCount(123))
}

func digitCount(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}

func digitSum(n int) int {
	sum := 0
	rem := 0

	// rem = n % 10
	// n = n / 10
	// fmt.Println("rem = ",rem, " n = ",n)

	// rem = n % 10
	// n = n / 10
	// fmt.Println("rem = ",rem, " n = ",n)

	// rem = n % 10
	// n = n / 10
	// fmt.Println("rem = ",rem, " n = ",n)

	// rem = n % 10
	// n = n / 10
	// fmt.Println("rem = ",rem, " n = ",n)

	for n > 0 {
		rem = n % 10
		n = n / 10
		sum = sum + rem
		fmt.Println("rem = ", rem, " n = ", n, " sum = ", sum)
	}

	return sum
}

func printOnlyNonPrimeNumbers(n int) {
	for i := 1; i <= n; i++ {
		if !isPrimeNumber(i) {
			fmt.Print(i, ", ")
		}
	}
}

func printOnlyPrimeNumbers(n int) {
	for i := 1; i <= n; i++ {
		if isPrimeNumber(i) {
			fmt.Print(i, ", ")
		}
	}
}

func isPrimeNumber(n int) bool {
	if n <= 0 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		//fmt.Println("i = ",i," n = ",n," n % i = ",n % i)
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	fact := 1
	fmt.Println("fact = ", fact)

	for i := n; i > 1; i-- {
		fact = fact * i
		fmt.Println("i = ", i, " * fact = ", fact)
	}

	return fact
}

func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func switchTest(i int) {
	switch i {

	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")

	default:
		fmt.Println("default case")
	}
}
