package driver


const(
	N_floors = 4
	N_buttons = 3 
	M_port = "30009"
	S_port = "20009"
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
	Status = 8
	Job = 10
	Turn_off_lights = 12
	Turn_on_lights = 13
	Kill = 666
	IDLE = -1
	INSIDECALL = 2
	UPCALL = 0
	DOWNCALL = 1
)




