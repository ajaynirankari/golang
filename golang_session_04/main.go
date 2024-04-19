package main

import "fmt"

func main() {

	num := 40
	result := isEvenNumber(num)
	fmt.Println(num,"Is the number even? ", result)

	num = 35
	result = isEvenNumber(num)
	fmt.Println(num,"Is the number even? ", result)

	num = 50
	result = isEvenNumber(num)
	fmt.Println(num,"Is the number even? ", result)


	number := 1
	if isEvenNumber(number) {
		fmt.Print(number," ")
	}

	number = 2
	if isEvenNumber(number) {
		fmt.Print(number," ")
	}

	number = 3
	if isEvenNumber(number) {
		fmt.Print(number," ")
	}

	number = 4
	if isEvenNumber(number) {
		fmt.Print(number," ")
	}

	number = 5
	if isEvenNumber(number) {
		fmt.Print(number," ")
	}

	fmt.Println()

	for i := 1; i <= 50; i++ {
		if isEvenNumber(i) {
			fmt.Print(i," ")
		}
	}

	fmt.Println()

	j := 1;
	for  j <= 50 {
		if isEvenNumber(j) {
			fmt.Print(j," ")
		}
		j++;
	}


	fmt.Println()

	k := 1;
	for {
		if isEvenNumber(k) {
			fmt.Print(k," ")
		}
		k++;
		if k > 50 {
			break
		}
	}
	fmt.Println("---------------------------------------------------")

	m:=0
	for {
		m = m+3
		if m % 5 == 0 {
			fmt.Println("Found (",m,") m % 5 == 0 true, going to continue the loop")
			continue
		}
		if m > 50 {
			fmt.Println("Found (",m,") m > 50 true, going to break the loop")
			break
		}
		fmt.Print(m," ")
		
	}

	fmt.Println("--Outside of the loop------------------------------")

	for p := range 20 {
        fmt.Print(p, " ")
    }
	
 fmt.Println()

 checkIfElseCondition(30)
 checkIfElseCondition(7)
 checkIfElseCondition(3)
 checkIfElseCondition(0)

 checkIfElseCondition(80)
 fmt.Println()

 sum := sumOfFirstNNumbers(100)
 fmt.Println("sumOfFirstNNumbers Sum of first 100 numbers = ",sum)

 sum = sumOfFirstNNumbersUsingRange(100)
 fmt.Println("sumOfFirstNNumbersUsingRange Sum of first 100 numbers = ",sum)

 sum = sumOfFirstNEvenNumbers(100)
 fmt.Println("sumOfFirstNEvenNumbers Sum of first 100 numbers = ",sum)

 sum = sumOfFirstNOddNumbers(100)
 fmt.Println("sumOfFirstNOddNumbers Sum of first 100 numbers = ",sum)

 mul := multiplicationOfFirstNNumbers(10)
 fmt.Println("Multiplication of first 10 numbers = ",mul)

}

func multiplicationOfFirstNNumbers(n int) int {
	mul := 1;
	for i := 1; i <= n; i++ {
		fmt.Println("mul = ",mul,", * i = ",i)
		mul = mul * i
	}
	return mul;
}

func sumOfFirstNOddNumbers(n int) int {
	sum := 0;
	for i := 1; i <= n; i++ {
		if isOddNumber(i) {
			sum = sum + i
		}
	}
	return sum;
}


func sumOfFirstNEvenNumbers(n int) int {
	sum := 0;
	for i := 1; i <= n; i++ {
		if isEvenNumber(i) {
			sum = sum + i
		}
	}
	return sum;
}


func sumOfFirstNNumbersUsingRange(n int) int {
	sum := 0;
	for i := range n {
		sum = sum + i		
	}
	return sum;
}

func sumOfFirstNNumbers(n int) int {

	sum := 0;

	for i := 1; i < n; i++ {
		fmt.Println("sum = ",sum,", + i = ",i)
		sum = sum + i
		
	}
	
	return sum;
}

func checkIfElseCondition(n int) {

	if n > 50 {
		fmt.Print("n > 50")
	}


	if n > 30 {
		fmt.Print("n > 30")
	} else {
		fmt.Print("n <= 30")
	}

	if n > 10 {
		fmt.Print("n > 10")
	} else if n > 5 {
		fmt.Print("n > 5")
	} else if n > 2 {
		fmt.Print("n > 2")
	} else {
		fmt.Print("no if match")
	}

	fmt.Println("out side of if block")

	
}

func isEvenNumber(number int) bool {
	// rem := number % 2
	// if rem == 0 {
	// 	return true
	// } else {
	// 	return false
	// }
	return number % 2 == 0
}

func isOddNumber(number int) bool {
	return number % 2 != 0
}