package main

import "fmt"

func main() {

	var a int // Declare a variable
	a = 10    // Initialize the variable
	a = a + 5 // Use of the variable
	fmt.Println(a)

	// Array
	var arrayOfInt [5]int // Declare an array
	fmt.Println(arrayOfInt)
	arrayOfInt[0] = 10 // Initialize the array
	arrayOfInt[1] = 20 // Initialize the array
	arrayOfInt[2] = 30 // Initialize the array
	arrayOfInt[3] = 14 // Initialize the array
	arrayOfInt[4] = 51 // Initialize the array

	fmt.Println(arrayOfInt) // Use of the array

	fmt.Println("Input Array = ", arrayOfInt, ", len = ", len(arrayOfInt), "cap = ", cap(arrayOfInt))

	x := 5
	fmt.Println(x)

	// Array
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y, len(y), cap(y))

	z := y[3]
	fmt.Println(z)

	y[3] = 44
	fmt.Println(y)

	// y[5] = 44 invalid argument: index 5 out of bounds

	fmt.Println("Iterate array element Using for loop")

	for i := 0; i < len(y); i++ {
		v := y[i]
		fmt.Println("Index = ", i, ", Value = ", v, ", Value * 2 = ", v*2)
	}

	fmt.Println("Iterate array element Using range")

	for i, v := range y {
		fmt.Println("Index = ", i, ", Value = ", v, ", Value * 2 = ", v*2)
	}

	fmt.Println("Iterate array element Using range without index")
	for _, v := range y {
		fmt.Println("Value = ", v, ", Value * 2 = ", v*2)
	}

	fmt.Println("Iterate array element Using range without value")
	for i, _ := range y {
		fmt.Println("Index = ", i, ", Value = ", y[i], ", Value * 2 = ", y[i]*2)
	}

	// Slice
	sliceOfInt := []int{1, 2, 3, 4, 5}
	fmt.Println("sliceOfInt = ", sliceOfInt, ", len = ", len(sliceOfInt), ", cap = ", cap(sliceOfInt))

	sliceOfInt = append(sliceOfInt, 6)

	fmt.Println("sliceOfInt = ", sliceOfInt, ", len = ", len(sliceOfInt), ", cap = ", cap(sliceOfInt))

	sliceOfInt = append(sliceOfInt, 7, 8, 9)

	fmt.Println("sliceOfInt = ", sliceOfInt, ", len = ", len(sliceOfInt), ", cap = ", cap(sliceOfInt))

	sliceOfInt = append(sliceOfInt, 10)

	fmt.Println("sliceOfInt = ", sliceOfInt, ", len = ", len(sliceOfInt), ", cap = ", cap(sliceOfInt))

	sliceOfInt = append(sliceOfInt, 11)

	fmt.Println("sliceOfInt = ", sliceOfInt, ", len = ", len(sliceOfInt), ", cap = ", cap(sliceOfInt))
	

	// Slice
	sliceOfIntUsingMake := make([]int, 5, 10)
	fmt.Println("sliceOfIntUsingMake = ", sliceOfIntUsingMake, ", len = ", len(sliceOfIntUsingMake), ", cap = ", cap(sliceOfIntUsingMake))

	sliceOfIntUsingMakeWithoutCapacity := make([]int, 5)
	fmt.Println("sliceOfIntUsingMakeWithoutCapacity = ", sliceOfIntUsingMakeWithoutCapacity, ", len = ", len(sliceOfIntUsingMakeWithoutCapacity), ", cap = ", cap(sliceOfIntUsingMakeWithoutCapacity))

	// Slice
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("s1 = ", s1, ", len = ", len(s1), ", cap = ", cap(s1))

	s2 := s1[1:3]
	fmt.Println("s1[1:3] --> s2 = ", s2, ", len = ", len(s2), ", cap = ", cap(s2))

	s3 := s1[3:]
	fmt.Println("s1[3:] --> s3 = ", s3, ", len = ", len(s3), ", cap = ", cap(s3))

	s4 := s1[:3]
	fmt.Println("s1[:3] --> s4 = ", s4, ", len = ", len(s4), ", cap = ", cap(s4))

	s5 := s1[:]
	fmt.Println("s1[:] --> s5 = ", s5, ", len = ", len(s5), ", cap = ", cap(s5))


	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	s2[0] = 22
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)

}
