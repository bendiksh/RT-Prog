package driver

//import "net"

const(
	N_floors = 4
	N_buttons = 3 
)

type Event struct {
	Floor int
	Type int
}

// Event.Type alternatives
const(
	Button_up = 0
	Button_down = 1
	Button_command = 2
	//JOB_DONE         = 3
	//PASSED_FLOOR     = 4
	//DIRECTION_CHANGE_UP = 5
	//DIRECTION_CHANGE_DOWN = 6
	//DIRECTION_CHANGE_STOP = 7
	//TURN_ON_UP_LIGHT = 8
	//TURN_ON_DOWN_LIGHT = 9
	//TURN_OFF_LIGHTS = 10
)
