package Employee

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

// Method
func (e *Employee) Work() {
	fmt.Printf("%s is working...\n", e.Name)
}

// Interface
type Worker interface {
	Work()
}
