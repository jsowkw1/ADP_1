package Hotel

import "fmt"

type Hotel struct {
	Name   string
	Rooms  int
	Guests []string
}

func NewHotel(name string, rooms int) Hotel {
	return Hotel{
		Name:  name,
		Rooms: rooms,
	}
}

func (h *Hotel) AddGuest(guest string) {
	if len(h.Guests) < h.Rooms {
		h.Guests = append(h.Guests, guest)
		fmt.Println("Guest added to hotel.")
	} else {
		fmt.Println("No free rooms available.")
	}
}

func (h Hotel) Info() {
	fmt.Println("Hotel:", h.Name)
	fmt.Println("Rooms:", h.Rooms)
	fmt.Println("Guests:", h.Guests)
}
