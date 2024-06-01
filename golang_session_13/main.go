package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
}

func newEmployee(id int, name string, age int) *Employee {
	return &Employee{id, name, age}
}

type Shape interface {
	area() string
	parameter() string
}

type Rectange struct{}

func (r Rectange) area() string {
	return "Rectange area"
}
func (r Rectange) parameter() string {
	return "Rectange parameter"
}

type Triangle struct{}

func (r Triangle) area() string {
	return "Triangle area"
}
func (r Triangle) parameter() string {
	return "Triangle parameter"
}

func shapeArea(s Shape) {
	a := s.area()
	fmt.Println(a)
	fmt.Println(s.parameter())
}

func shapeForAll(s interface{}) {
	fmt.Println("shapeForAll = ", s)
}

func main() {

	var e1 Employee
	e1 = Employee{1, "John", 45}
	fmt.Println("e1 = ", e1)

	e2 := Employee{2, "John2", 42}
	fmt.Println("e2 = ", e2)

	e3 := Employee{Name: "John2", Age: 42, Id: 3}
	fmt.Println("e3 = ", e3)

	e4 := &Employee{Name: "John3", Age: 43, Id: 4}
	fmt.Println("e4 = ", e4)

	e5 := new(Employee)
	e5.Id = 5
	e5.Name = "Smith"
	e5.Age = 56
	fmt.Println("e5 = ", e5)

	e6 := newEmployee(7, "Tim", 73)
	fmt.Println("e6 = ", e6)

	r := Rectange{}
	t := Triangle{}
	shapeArea(r)
	shapeArea(t)

	shapeForAll(r)
	shapeForAll(t)

	s2 := Super{a: 45}

	shapeForAll(s2)

	s := Sub{
		Super: Super{a: 10},
		b:     20,
	}
	fmt.Println("s= ", s)

	s1 := Sub{
		Super{a: 10},
		20,
	}
	fmt.Println("s1= ", s1)
	fmt.Println("s1.a= ", s1.a)
	fmt.Println("s1.b= ", s1.b)
	fmt.Println("s1.Super.a = ", s1.Super.a)

	fmt.Println("ok")
}

type Super struct {
	a int
}

type Sub struct {
	Super
	b int
}

type Subb struct {
	ss Super
	b  int
}
