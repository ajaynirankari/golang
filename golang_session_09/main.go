package main

import (
	"fmt"
	"strings"
)

func main() {

	inputSlice := []int{5, 3, 6, 2, 10}
	fmt.Println("inputSlice: ", inputSlice)
	selectionSort(inputSlice)
	fmt.Println("Sorted Slice: ", inputSlice)

	name := "golang"
	fmt.Println("name: ", name)

	for i, c := range name {
		fmt.Println("Index = ", i, "ASCII = ", c, " Character In Lower case = ", string(c), " Convert to Upper Case = ", string(c-32))
	}

	toLower := "golang" // string is []byte  or  // []rune
	fmt.Println("input toLower: ", toLower)
	toUpper := toUpperCase(toLower)
	fmt.Println("output toUpper: ", toUpper, ", From library: ", strings.ToUpper(toLower))

	toLower = toLowerCase(toUpper)
	fmt.Println("input toUpper: ", toUpper)
	fmt.Println("out toLower: ", toLower, ", From library: ", strings.ToLower(toUpper))

	sum := sumOfNumbers(1, 2, 3, 4, 5)
	fmt.Println("Sum of 1, 2, 3, 4, 5: ", sum)

	sum = sumOfNumbers(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("Sum of 1, 2, 3, 4, 5,6,7,8: ", sum)

	sum = sumOfNumbers(6, 7, 8)
	fmt.Println("Sum of 6,7,8: ", sum)

	sum = sumInutSlice([]int{1, 2, 3, 4, 5})
	fmt.Println("Sum of 1, 2, 3, 4, 5: ", sum)

	fmt.Println(findMaxAndMin([]int{1, 2, 3, 4, 5}))

	fmt.Println(findAllEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

}

func findAllEvenNumbers(input []int) []int {
	var result []int
	for _, v := range input {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result

}

func findMaxAndMin(input []int) (int, int) {
	var max int
	max = input[0]
	for i := 1; i < len(input); i++ {
		if max < input[i] {
			max = input[i]
		}
	}

	var min int
	min = input[0]
	for i := 1; i < len(input); i++ {
		if min > input[i] {
			min = input[i]
		}
	}

	return max, min
}

func sumOfNumbers(a ...int) int {
	var sum int
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func sumInutSlice(a []int) int {
	var sum int
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func sum2Digits(a, b int) int {
	return a + b
}
func sum3Digits(a, b, c int) int {
	return a + b + c
}
func sum4Digits(a, b, c, d int) int {
	return a + b + c + d
}

func toLowerCase(input string) string {
	var result string
	for _, c := range input {
		result = result + string(c+32)
	}
	return result
}

func toUpperCase(input string) string {
	var result string
	for _, c := range input {
		result = result + string(c-32)
	}
	return result
}

func selectionSort(slice []int) {

	for i := 0; i < len(slice); i++ {
		minIndex := i
		minValue := slice[i]

		for j := i + 1; j < len(slice); j++ {
			if slice[j] < minValue {
				minValue = slice[j]
				minIndex = j
			}
		}

		if minIndex != i {
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
		}
	}

}
