package store

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Account
	Credits float64
}

func (a *Account) ChangeName(name string) {
	a.FirstName = name
}

func (e *Employee) String() string {
	return fmt.Sprintf("Name: %s %s\n Credits: %.2f \n", e.FirstName, e.LastName, e.Credits)
}

func CreateEmployee(fistName, LastName string, credits float64) (*Employee, error) {
	return &Employee{Account{fistName, LastName}, credits}, nil
}

func (e *Employee) AddCredits(credits float64) (float64, error) {
	if credits > 0.0 {
		e.Credits += credits
		return e.Credits, nil
	}
	return 0.0, errors.New("参数错误")
}
func (e Employee) RemoveCredit(credit float64) (float64, error) {
	if credit > 0 {
		if e.Credits < credit {
			return 0.0, errors.New("账户金额小于0")
		} else {
			e.Credits -= credit
			return credit, nil
		}
	}

	return 0.0, errors.New("参数错误")
}

func (e Employee) CheckCredits() float64 {
	return e.Credits
}
