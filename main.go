package main

import (
	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Hotel"
	"Assignment1/Wallet"
	"fmt"
)

func main() {
	hotel := Hotel.NewHotel("Astana Hotel", 2)
	gym := Gym.NewGym()
	wallet := Wallet.Wallet{Balance: 1000}

	employees := []Employee.Employee{
		{ID: 1, Name: "Ali", Salary: 300000},
		{ID: 2, Name: "Dana", Salary: 280000},
	}

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Hotel")
		fmt.Println("2. Gym")
		fmt.Println("3. Wallet")
		fmt.Println("4. Employees")
		fmt.Println("0. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("1. Add guest")
			fmt.Println("2. Hotel info")
			var c int
			fmt.Scan(&c)

			if c == 1 {
				var name string
				fmt.Print("Guest name: ")
				fmt.Scan(&name)
				hotel.AddGuest(name)
			} else {
				hotel.Info()
			}

		case 2:
			fmt.Println("1. Add member")
			fmt.Println("2. Gym info")
			var c int
			fmt.Scan(&c)

			if c == 1 {
				var id int
				var name string
				fmt.Print("ID: ")
				fmt.Scan(&id)
				fmt.Print("Name: ")
				fmt.Scan(&name)
				gym.AddMember(id, name)
			} else {
				gym.Info()
			}

		case 3:
			fmt.Println("1. Deposit")
			fmt.Println("2. Withdraw")
			fmt.Println("3. Balance")
			var c int
			fmt.Scan(&c)

			var amount float64
			switch c {
			case 1:
				fmt.Scan(&amount)
				wallet.Deposit(amount)
			case 2:
				fmt.Scan(&amount)
				wallet.Withdraw(amount)
			case 3:
				wallet.Info()
			}

		case 4:
			for _, emp := range employees {
				emp.Info()
				fmt.Println("------")
			}

		case 0:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
