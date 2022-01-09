package exercises

import (
	"errors"
	"fmt"
)

type Account struct {
	LastName  string
	FirstName string
}

type Employee struct {
	accountInfo Account
	load        float64
}

func (a *Account) ChangeName(name string) {
	a.LastName = name
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	if len(firstName) == 0 || len(lastName) == 0 {
		return nil, errors.New("error name when create employee")
	}
	acInfo := Account{FirstName: firstName, LastName: lastName}
	return &Employee{load: credits, accountInfo: acInfo}, nil
}

func (e *Employee) AddCredits(amount float64) (*Employee, error) {
	if amount < 0 {
		return nil, errors.New("error credit amount when add credits")
	}

	e.load += amount
	return e, nil
}

func (e *Employee) RemoveCredits(amount float64) (*Employee, error) {
	if amount < 0 {
		return nil, errors.New("error credit amount when remove credits")
	}

	e.load -= amount
	if e.load < 0 {
		return nil, errors.New("error credit amount more than loaded")
	}
	return e, nil
}

func (e *Employee) CheckCredits() (float64, error) {
	if e.load < 0 {
		return 0, errors.New("error credit amount when check credits")
	}

	return e.load, nil
}

func (e *Employee) String() {
	fmt.Printf("%s·%s have %f credits\n", e.accountInfo.LastName, e.accountInfo.FirstName, e.load)
}

func Employees() {
	fmt.Println("employee examples...")
	e, err := CreateEmployee("Smith", "Fung", 1000)
	if err != nil {
		fmt.Println(e)
		return
	}

	e.accountInfo.ChangeName("Feng")
	e.String()
	e, err = e.AddCredits(5000)
	if err != nil {
		fmt.Println(err)
		return
	}
	e, err = e.RemoveCredits(1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	load, err1 := e.CheckCredits()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("End, acount load is %f\n", load)
}
