package driver

import "net"

const(
	N_floors = 4
	N_buttons = 3 
)

type Event_t struct {
	Floor int
	Type int
}

// Event.Type alternatives
const(
	Button_up = 0
	Button_down = 1
	Button_command = 2
	Elev_done = 3
	Lights_up_on = 4
	Lights_down_on = 5
	Lights_off = 6
	Status = 7
)

type Elev_t struct {
	UpCalls [N_floors]int
	DownCalls [N_floors]int
	Floor int
	Dir int
	Conn *net.TCPConn
}

type Message_t struct {
	Type int
	Floor int
	Dir int
	ElevDest int
	IP string
}
