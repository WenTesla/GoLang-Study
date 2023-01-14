package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)
}
