package main

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Doe"
	fmt.Println(employee.FirstName)
}
