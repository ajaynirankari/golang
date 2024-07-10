package main

import (
	"fmt"
	"sort"
)

func main() {

	x := []int{41, 2, 23, 14, 5}
	fmt.Println(x)

	sort.Ints(x)
	fmt.Println(x)

	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})

	fmt.Println(x)

	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)

	sort.Slice(names, func(i, j int) bool {
		return names[i] > names[j]
	})
	fmt.Println(names)

	floatNumbers := []float64{4.2, 7.1, 1.3, 3.7}
	fmt.Println(floatNumbers)
	sort.Float64s(floatNumbers)
	fmt.Println(floatNumbers)

	sort.Slice(floatNumbers, func(i, j int) bool {
		return floatNumbers[i] > floatNumbers[j]
	})
	fmt.Println(floatNumbers)

	type Person struct {
		Name string
		Age  int
	}

	listOfPeople := []Person{
		{"John", 50},
		{"Paul", 35},
		{"George", 40},
		{"Ringo", 45},
	}
	fmt.Println(listOfPeople)

	sort.Slice(listOfPeople, func(i, j int) bool {
		return listOfPeople[i].Age < listOfPeople[j].Age
	})
	fmt.Println(listOfPeople)

	sort.Slice(listOfPeople, func(i, j int) bool {
		return listOfPeople[i].Name < listOfPeople[j].Name
	})
	fmt.Println(listOfPeople)

	sort.Slice(listOfPeople, func(i, j int) bool {
		return listOfPeople[i].Name > listOfPeople[j].Name
	})
	fmt.Println("Sort by name: ", listOfPeople)

}
