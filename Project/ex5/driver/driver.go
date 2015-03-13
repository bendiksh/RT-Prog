package driver

import(
	
)

const(
	N_floors = 4
	N_buttons = 3 
)

// Matrices of buttons, lights and sensor to loop through easier
var Button_matrix = [N_floors][N_buttons]int{
	{BUTTON_UP1,BUTTON_DOWN1,BUTTON_COMMAND1},
	{BUTTON_UP2,BUTTON_DOWN2,BUTTON_COMMAND2},
	{BUTTON_UP3,BUTTON_DOWN3,BUTTON_COMMAND3},
	{BUTTON_UP4,BUTTON_DOWN4,BUTTON_COMMAND4}
}

var Light_matrix = [N_floors][N_buttons]int{
	{LIGHT_UP1,LIGHT_DOWN1,LIGHT_COMMAND1},
	{LIGHT_UP2,LIGHT_DOWN2,LIGHT_COMMAND2},
	{LIGHT_UP3,LIGHT_DOWN3,LIGHT_COMMAND3},
	{LIGHT_UP4,LIGHT_DOWN4,LIGHT_COMMAND4}
}

var button=[N_floors][N_buttons]int{
	{0,0,0},
	{0,0,0},
	{0,0,0},
	{0,0,0}
}

var Sensors = [N_floors]int{SENSOR_FLOOR1,SENSOR_FLOOR2,SENSOR_FLOOR3,SENSOR_FLOOR4}

func Elev_init() (int,int){ // returns a status int and a floor int
	if !(Io_init()) {
		return 0, 0
	}
	
	// turn off all lights
	for i := 0; i < N_floors; i++ {
		if i != 0 {
			Elev_set_btn_light(Button_down, i, 0)
		}
		if i != (N_floors - 1) {
			Elev_set_btn_light(Button_down, i, 0)
		}
		Elev_set_btn_light(Button_command, i, 0)
	}

	Elev_stop_light(0)
	Elev_door_light(0)
	Elev_floor_ind(0)
	
	// find current floor
	floor := 0
	for i := 0; i < N_floors; i++{
		if(Io_read_bit(Sensors[i]) == 1) {
			floor = i
		}
	}

	return 1, floor
}

func Elev_floor_ind(floor int){
	if (floor & 0x02 == 0x02) {
		Io_set_bit(LIGHT_FLOOR_IND1)
	} else{
		Io_clear_bit(LIGHT_FLOOR_IND1)
	}
	if (floor & 0x01 == 0x01){
		Io_set_bit(LIGHT_FLOOR_IND2)
	} else{
		Io_clear_bit(LIGHT_FLOOR_IND2)
	}
}

func Elev_get_button(){ // need an Event type
	for i := 0; i < N_floors; i++ {
		for j := 0; j < N_buttons; j++{
			if (Io_read_bit(Button_matrix[i][j]) == 1 && button[i][j] == 0){
				button[i][j] = 1
				return // need an Event type
			}else if (Io_read_bit(Button_matrix[i][j]) == 0){
				button[i][j] = 0
			}
		}
	}
	return // need an Event type
}

func Elev_set_btn_light(button, floor, value int){
	if value == 1 {
		Io_set_bit(Light_matrix[floor][button])
	}else {
		Io_clear_bit(Light_matrix[floor][button])
	}
}

func Elev_floor_sensor() int{
	if (Io_read_bit(SENSOR_FLOOR1)){
		return 0
	}else if (Io_read_bit(SENSOR_FLOOR2)){
		return 1
	}else if (Io_read_bit(SENSOR_FLOOR3)){
		return 2
	}else if (Io_read_bit(SENSOR_FLOOR4)){
		return 3
	}else {
		return -1
	}
}
