package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (fullTimeEmployee FullTimeEmployee) GetMessage() string {
	return "Full time employee"
}

func (halfTimeEmployee HalfTimeEmployee) GetMessage() string {
	return "Half time employee"
}

type HalfTimeEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	GetMessage() string
}

func GetMessage(p PrintInfo){
	fmt.Println(p.GetMessage())
}

func main() {

	empleado := FullTimeEmployee{}

	empleado.name = "Juan"
	empleado.age = 20
	empleado.id = 1

	fmt.Println(empleado)
	GetMessage(empleado)

	temporal := HalfTimeEmployee{}

	temporal.name = "Maria"
	temporal.age = 25
	temporal.id = 2

	fmt.Println(temporal)
	GetMessage(temporal)
}
