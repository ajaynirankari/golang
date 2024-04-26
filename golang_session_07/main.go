package main

import "fmt"

func main() {

	input := []int{1, 2, 3, 4, 5}

	result := sumOfInputSlice(input)
	fmt.Println("Sum of input slice is: ", result)

	result = sumOfInputSliceUsingRange(input)
	fmt.Println("Sum of input slice using range is: ", result)

	product := multiplyOnInputSlice([]int{2, 5, 3})
	fmt.Println("Product of input slice is: ", product)

	intputSlice := []int{}
	product = multiplyOnInputSlice(intputSlice)
	fmt.Println("Product of input slice is: ", product)

	intputSlice = append(intputSlice, 2, 6, 3)
	product = multiplyOnInputSlice(intputSlice)
	fmt.Println("Product of input slice is: ", product)

	product = multiplyOnInputSliceUsingRange(intputSlice)
	fmt.Println("Product of input slice is: ", product)

	intputSliceForMax := []int{7, 4, 8, 6, 2, 4}
	maxElement := findMaxElementInInputSlice(intputSliceForMax)
	fmt.Println("Max element in input slice is: ", maxElement, ", intputSliceForMax: ", intputSliceForMax)

	minElement := findMinElementInInputSlice(intputSliceForMax)

	fmt.Println("Min element in input slice is: ", minElement, ", intputSlice: ", intputSliceForMax)

}

func findMinElementInInputSlice(slice []int) int {
	min := slice[0]

	for i := 1; i < len(slice); i++ {
		// fmt.Println("element: ", slice[i], ", min: ", min, ", if < : ", slice[i] < min)
		if slice[i] < min {
			min = slice[i]
		}
	}

	return min
}

func findMaxElementInInputSlice(slice []int) int {
	max := slice[0]

	for i := 1; i < len(slice); i++ {
		// fmt.Println("element: ", slice[i], ", max: ", max, ", if > : ", slice[i] > max)
		if slice[i] > max {
			max = slice[i]
		}
	}

	return max
}

func multiplyOnInputSliceUsingRange(slice []int) int {
	mul := 1

	for _, value := range slice {
		mul = mul * value
	}

	return mul
}

func multiplyOnInputSlice(slice []int) int {
	lenghtOfSlice := len(slice)
	if lenghtOfSlice == 0 {
		return 0
	}

	mul := slice[0]
	for i := 1; i < lenghtOfSlice; i++ {
		mul = mul * slice[i]
	}
	return mul
}

func sumOfInputSliceUsingRange(slice []int) int {
	sum := 0

	for _, value := range slice {
		sum = sum + value
	}

	return sum
}

func sumOfInputSlice(slice []int) int {
	sum := 0

	for i := 0; i < len(slice); i++ {
		value := slice[i]
		sum = sum + value
	}

	return sum
}
