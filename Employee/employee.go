package Employee

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e Employee) Info() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Salary:", e.Salary)
}
