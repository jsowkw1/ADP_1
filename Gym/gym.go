package Gym

import "fmt"

type Gym struct {
	Members map[int]string
}

func NewGym() Gym {
	return Gym{
		Members: make(map[int]string),
	}
}

func (g *Gym) AddMember(id int, name string) {
	g.Members[id] = name
	fmt.Println("Member added to gym.")
}

func (g Gym) Info() {
	fmt.Println("Gym Members:")
	for id, name := range g.Members {
		fmt.Println(id, "-", name)
	}
}
