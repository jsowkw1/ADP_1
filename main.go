package main

import (
	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Hotel"
	"Assignment1/Wallet"
)

func main() {

	emp := Employee.Employee{Name: "John", Age: 30, Salary: 5000}
	emp.Work()

	hotel := Hotel.Hotel{Name: "Grand Hotel", Rooms: 5}
	hotel.BookRoom()

	gym := Gym.Gym{
		Name:      "FitLife",
		Members:   make(map[string]bool),
		Equipment: []string{"Treadmill", "Dumbbells", "Stationary Bike"},
	}
	gym.AddMember("Alice")
	gym.ListEquipment()

	// Wallet example
	wallet := Wallet.Wallet{Owner: "John", Balance: 100}
	wallet.AddMoney(50)
	wallet.SpendMoney(30)
}
