package Hotel

import "fmt"

type Hotel struct {
	Name  string
	Rooms int
}

func (h *Hotel) BookRoom() {
	if h.Rooms > 0 {
		h.Rooms--
		fmt.Printf("Room booked at %s, remaining: %d\n", h.Name, h.Rooms)
	} else {
		fmt.Println("No rooms available")
	}
}
