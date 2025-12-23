package Gym

import "fmt"

type Gym struct {
	Name      string
	Members   map[string]bool
	Equipment []string
}

func (g *Gym) AddMember(name string) {
	g.Members[name] = true
	fmt.Printf("%s joined %s gym!\n", name, g.Name)
}

func (g *Gym) ListEquipment() {
	fmt.Println("Equipment available:", g.Equipment)
}
