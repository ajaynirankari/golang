package main

import (
	"fmt"
	"strings"
)

type Customer struct {
	name string
	age  int
}

func (c Customer) displayCustomer() {
	fmt.Println("displayCustomer Befor Name = ", c.name, ", Age = ", c.age)
	c.age = c.age + 1
	fmt.Println("displayCustomer After Name = ", c.name, ", Age = ", c.age)
}

func (c *Customer) updateCustomerAge() {
	fmt.Println("updateCustomerAge Befor Name = ", c.name, ", Age = ", c.age)
	c.age = c.age + 1
	fmt.Println("updateCustomerAge After Name = ", c.name, ", Age = ", c.age)
}

func display(name string) {
	fmt.Println(strings.ToUpper(name))
}

func main() {
	var x int
	x = 10
	fmt.Println("x = ", x)

	y := 20
	fmt.Println("y = ", y)

	display("Sam")

	var customer1 Customer
	customer1 = Customer{name: "Raj", age: 34}
	fmt.Println("main() customer1 = ", customer1)

	customer2 := Customer{name: "Tim", age: 45}
	fmt.Println("main() customer2 = ", customer2)

	fmt.Println("main() customer2.name = ", customer2.name)
	fmt.Println("main() customer2.age = ", customer2.age)

	customer2.displayCustomer()
	customer1.displayCustomer()

	fmt.Println("main() customer2 = ", customer2)
	customer2.updateCustomerAge()
	fmt.Println("main() customer2 = ", customer2)
	customer2.displayCustomer()
	fmt.Println("main() customer2 = ", customer2)

}
