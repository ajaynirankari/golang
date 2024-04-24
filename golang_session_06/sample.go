package main

import "fmt"

func main() {
	
	var x int // Declare a variable x of type int
	x = 10    // Assign or initialize a value to x

	inc := x + 5 // Use the value of x in an expression

	fmt.Println(x, inc) // Print the values of x and inc

	var arrayOfIntValues [5]int // Declare an array of 5 integers

	fmt.Println(arrayOfIntValues)

	arrayOfIntValues[2] = 20 // Assign a value to the first element of the array

	fmt.Println(arrayOfIntValues)

	v := arrayOfIntValues[2]
	fmt.Println(v)

	arrayOfIntValues[0] = 9
	arrayOfIntValues[1] = 10
	arrayOfIntValues[2] = 20
	arrayOfIntValues[3] = 30
	arrayOfIntValues[4] = 30

	fmt.Println("arrayOfIntValues = ", arrayOfIntValues, "Length = ", len(arrayOfIntValues))

	//v = arrayOfIntValues[6]   //invalid argument: index 6 out of bounds
	//v = arrayOfIntValues[-2]  // index -2 (constant of type int) must not be negative

	// Variable declaration and initialization
	s := 10
	fmt.Println("s = ", s)

	// Array declaration and initialization
	a := [5]int{11, 12, 13, 14, 15} // Declare and initialize an array of 5 integers

	fmt.Println("array = ", a, "Length = ", len(a))

	a[2] = 22
	fmt.Println("array = ", a, "Length = ", len(a))

	fmt.Println("a[0] = ", a[0])
	fmt.Println("a[2] = ", a[2])
	fmt.Println("a[4] = ", a[4])

	// iterate over the array
	for i := 0; i < len(a); i++ {
		v = a[i]
		v = v * 2
		fmt.Println("a[", i, "] = ", a[i], ", v*2 = ", v)
	}

	fmt.Println("Iterating over the array using range")
	for i, v := range a {
		fmt.Println("a[", i, "] = ", a[i], ", v*2 = ", v*2)
	}

	fmt.Println("Iterating over the array using range without using the value")
	for i, _ := range a {
		v = a[i]
		v = v * 2
		fmt.Println("a[", i, "] = ", a[i], ", v*2 = ", v)
	}

	fmt.Println("Iterating over the array using range without using the index")
	for _, v := range a {
		fmt.Println("v = ", v, "v*2 = ", v*2)
	}

	//Create Slice
	slice := []int{21, 22, 23} // Declare and initialize a slice of integers

	fmt.Println("slice = ", slice, ", len = ", len(slice), " cap = ", cap(slice))

	slice = append(slice, 24) // Append an element to the slice

	fmt.Println("slice = ", slice, ", len = ", len(slice), " cap = ", cap(slice))

	slice = append(slice, 25, 26) // Append an element to the slice

	fmt.Println("slice = ", slice, ", len = ", len(slice), " cap = ", cap(slice))

	slice = append(slice, 27) // Append an element to the slice

	fmt.Println("slice = ", slice, ", len = ", len(slice), " cap = ", cap(slice))

	//Create Slice from built-in function make
	p := make([]int, 4, 6)
	fmt.Println("p = ", p, ", len = ", len(p), " cap = ", cap(p))

	p1 := make([]int, 3)
	fmt.Println("p1 = ", p1, ", len = ", len(p1), " cap = ", cap(p1))

	ar := [7]int{31, 32, 33, 34, 35, 36, 37}
	fmt.Println("ar = ", ar)

	sl := []int{41, 42, 43, 44, 45, 46, 47, 48, 49}
	fmt.Println("sl = ", sl)

	//Create Slice from Slice Expression - ar[low:high] 
	sl1 := ar[1:4]
	fmt.Println("ar[1:4] sl1 = ", sl1, "len = ", len(sl1), ",cap = ", cap(sl1))

	sl2 := sl[1:4]
	fmt.Println("sl[1:4] sl2 = ", sl2, "len = ", len(sl2), ",cap = ", cap(sl2))

	sl2 = sl[:4]     // same like sl[0:4]
	fmt.Println("sl[:4] sl2 = ", sl2, "len = ", len(sl2), ",cap = ", cap(sl2))

	sl2 = sl[1:]      // same like sl[1:len(sl)]
	fmt.Println("sl[1:] sl2 = ", sl2, "len = ", len(sl2), ",cap = ", cap(sl2))

	sl2 = sl[:]       // same like sl[0:len(sl)]
	fmt.Println("sl[:] sl2 = ", sl2, "len = ", len(sl2), ",cap = ", cap(sl2))

	z1 := []int{1, 2, 3, 4, 5, 6}
	z2 := z1[2:5]
	fmt.Println("Beofre: ", "z1 = ", z1, ", slice z2 of z1[2:5] = ", z2, ", z2 len = ", len(z2), ", z2 cap = ", cap(z2))

	z2[1] = 111
	fmt.Println("After: ", "z1 = ", z1, ", slice z2 of z1[2:5] = ", z2, ", z2 len = ", len(z2), ", z2 cap = ", cap(z2))

}
