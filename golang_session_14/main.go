package main

import "fmt"

type IntContainer struct {
	value int
}

type FloatContainer struct {
	value float64
}

type Container[T any] struct {
	value T
}

func main() {

	fmt.Println(test(5))
	fmt.Println(test1(5.5))

	c := IntContainer{5}
	fmt.Println(c)

	c1 := FloatContainer{5.5}
	fmt.Println(c1)

	c2 := Container[int]{5}
	fmt.Println(c2)

	c3 := Container[float64]{5.5}
	fmt.Println(c3)

	c4 := Container[string]{"Hello"}
	fmt.Println(c4)

	c5 := Container[bool]{true}
	fmt.Println(c5)

	s := sumInt(5, 5)
	fmt.Println(s)

	s1 := sumFloat(5.5, 5.3)
	fmt.Println(s1)

	s2 := sum(5, 5)
	fmt.Println(s2)

	s3 := sum(5.5, 5.3)
	fmt.Println(s3)

	s4 := minu(5, 5)
	fmt.Println(s4)

	s5 := product(5, 5)
	fmt.Println(s5)

	m := multiply(5, 5)
	fmt.Println(m)

	d := divide(5, 5)
	fmt.Println(d)

	printMessage1("Hello")
	//printMessage1(45) // Error

	printMessage("Hello")
	printMessage(5)
	printMessage(5.5)
	printMessage(true)

	printMessage2(5)
	printMessage2(5.5)
	//printMessage2(true) // Error

	printMessageR(5)

	printMessageRS("Hello")
	printMessageRS(5.5)
	//printMessageRS(5) // Error

}

func printMessage2[T Number](msg T) {
	fmt.Println(msg)
}

func printMessage[T any](msg T) {
	fmt.Println(msg)
}

func printMessageR[R any](msg R) {
	fmt.Println(msg)
}

func printMessageRS[R string | float64](msg R) {
	fmt.Println(msg)
}

func printMessage1(msg string) {
	fmt.Println(msg)
}

type Number interface {
	int | float64 | int32 | int64 | float32 | uint | uint32 | uint64 | uint8 | uint16 | int8 | int16 | uintptr
}

func divide[N Number](a, b N) N {
	return a / b
}

func multiply[N Number](a, b N) N {
	return a * b
}

func product[N Number](a, b N) N {
	return a * b
}

func sum[N int | float64](a, b N) N {
	return a + b
}

func minu[N int | float64](a, b N) N {
	return a - b
}

func sumInt(a, b int) int {
	return a + b
}
func sumFloat(a, b float64) float64 {
	return a + b
}

func test(a int) int {
	return a * a
}
func test1(a float64) float64 {
	return a * a
}
