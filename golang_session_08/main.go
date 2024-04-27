package main

import "fmt"

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("input slice: ", input)
	result := isAnyElementsAreDuplicateInInputSlice(input)
	fmt.Println("isAnyElementsAreDuplicateInInputSlice: ", result)

	input = []int{1, 2, 3, 4, 5, 6, 4, 8, 9}
	fmt.Println("input slice: ", input)
	result = isAnyElementsAreDuplicateInInputSlice(input)
	fmt.Println("isAnyElementsAreDuplicateInInputSlice: ", result)

	input = []int{1, 2, 3, 4, 5, 6, 4, 8, 3}
	fmt.Println("input slice: ", input)
	result = isAnyElementsAreDuplicateInInputSlice(input)
	fmt.Println("isAnyElementsAreDuplicateInInputSlice: ", result)

	input = []int{1, 2, 3, 4, 5, 6, 4, 8, 3, 8}
	fmt.Println("input slice: ", input)
	results := findAllDuplicatesInInputSlice(input)
	fmt.Println("findAllDuplicatesInInputSlice: ", results)

	swap()

	inputSlice := []int{5, 3, 8, 6, 2}
	fmt.Println("input slice: ", inputSlice)
	assendingOrderUsingbubbleSort(inputSlice)
	fmt.Println("assendingOrderUsingbubbleSort: ", inputSlice)

	inputSlice = []int{5, 3, 8, 6, 2}
	fmt.Println("input slice: ", inputSlice)
	descendingOrderUsingbubbleSort(inputSlice)
	fmt.Println("descendingOrderUsingbubbleSort: ", inputSlice)

}

func descendingOrderUsingbubbleSort(inputSlice []int) {
	for i := 0; i < len(inputSlice); i++ {
		for j := 0; j < len(inputSlice)-1-i; j++ {
			if inputSlice[j] < inputSlice[j+1] {
				inputSlice[j], inputSlice[j+1] = inputSlice[j+1], inputSlice[j]
			}
		}
	}
}
func assendingOrderUsingbubbleSort(inputSlice []int) {

	for i := 0; i < len(inputSlice); i++ {

		for j := 0; j < len(inputSlice)-1-i; j++ {
			if inputSlice[j] > inputSlice[j+1] {
				inputSlice[j], inputSlice[j+1] = inputSlice[j+1], inputSlice[j]
			}
		}
		//fmt.Println(inputSlice)
	}
	// fmt.Println(inputSlice)
	// for j := 0; j < len(inputSlice)-1; j++ {
	// 	if inputSlice[j] > inputSlice[j+1] {
	// 		inputSlice[j], inputSlice[j+1] = inputSlice[j+1], inputSlice[j]
	// 	}
	// }
	// fmt.Println(inputSlice)
	// for j := 0; j < len(inputSlice)-1; j++ {
	// 	if inputSlice[j] > inputSlice[j+1] {
	// 		inputSlice[j], inputSlice[j+1] = inputSlice[j+1], inputSlice[j]
	// 	}
	// }
	// fmt.Println(inputSlice)
	// for j := 0; j < len(inputSlice)-1; j++ {
	// 	if inputSlice[j] > inputSlice[j+1] {
	// 		inputSlice[j], inputSlice[j+1] = inputSlice[j+1], inputSlice[j]
	// 	}
	// }
	//}
}

func swap() {
	x := 10
	y := 20
	fmt.Println("Before swap x: ", x, " y: ", y)

	t := x
	x = y
	y = t

	fmt.Println("After swap x: ", x, " y: ", y)

	a := 10
	b := 20
	fmt.Println("Before swap a: ", a, " b: ", b)

	a, b = b, a

	fmt.Println("After swap a: ", a, " b: ", b)

}

func findAllDuplicatesInInputSlice(inputSlice []int) []int {
	duplicates := []int{}
	for i := 0; i < len(inputSlice); i++ {
		number := inputSlice[i]

		isAlreadyChecked := checkInputNumberExistInInputSlice(duplicates, number)

		if isAlreadyChecked {
			continue
		}

		isDpulicate := checkInputNumberInInputSliceIsDuplicate(inputSlice, number)
		//fmt.Println("number: ", number, " isDpulicate: ", isDpulicate)

		if isDpulicate {
			duplicates = append(duplicates, inputSlice[i])
			//fmt.Println("duplicates: ", duplicates)
		}
	}

	return duplicates
}
func checkInputNumberExistInInputSlice(inputSlice []int, number int) bool {
	for i := 0; i < len(inputSlice); i++ {
		if number == inputSlice[i] {
			return true
		}
	}
	return false
}
func checkInputNumberInInputSliceIsDuplicate(inputSlice []int, number int) bool {
	count := 0
	for i := 0; i < len(inputSlice); i++ {
		if number == inputSlice[i] {
			count = count + 1
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func isAnyElementsAreDuplicateInInputSlice(inputSlice []int) bool {
	// Write your code here

	for t := 0; t < len(inputSlice); t++ {
		element := inputSlice[t]
		count := 0
		for i := 0; i < len(inputSlice); i++ {
			if element == inputSlice[i] {
				count = count + 1
			}
			//fmt.Println("element: ", element, " count: ", count)
			if count > 1 {
				return true
			}
		}
	}
	return false

	// first = inputSlice[1]
	// for i := 1; i < len(inputSlice); i++ {
	// 	if first == inputSlice[i] {
	// 		return true
	// 	}
	// }

	// first = inputSlice[2]
	// for i := 1; i < len(inputSlice); i++ {
	// 	if first == inputSlice[i] {
	// 		return true
	// 	}
	// }
	//return false
}
