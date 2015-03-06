package driver

import(
	
)

const(
	N_floors = 4
	N_buttons = 3 
)

// Matrices of buttons, lights and sensor to loop through easier
var Button_matrix = [Num_floors][Num_buttons]int{
	{BUTTON_UP1,BUTTON_DOWN1,BUTTON_COMMAND1},
	{BUTTON_UP2,BUTTON_DOWN2,BUTTON_COMMAND2},
	{BUTTON_UP3,BUTTON_DOWN3,BUTTON_COMMAND3},
	{BUTTON_UP4,BUTTON_DOWN4,BUTTON_COMMAND4}
}

var Light_matrix = [Num_floors][Num_buttons]int{
	{LIGHT_UP1,LIGHT_DOWN1,LIGHT_COMMAND1},
	{LIGHT_UP2,LIGHT_DOWN2,LIGHT_COMMAND2},
	{LIGHT_UP3,LIGHT_DOWN3,LIGHT_COMMAND3},
	{LIGHT_UP4,LIGHT_DOWN4,LIGHT_COMMAND4}
}

var Sensors = [Num_floors]int{SENSOR_FLOOR1,SENSOR_FLOOR2,SENSOR_FLOOR3,SENSOR_FLOOR4}

func Elev_init() {
	if !(Io_init()) {
		return 0
	}
	
	for i := 0; i < N_floors; i++ {
		if i != 0 {
			Elev_set_btn_light(Button_down, i, 0)
		}
		if i != (N_floors - 1) {
			Elev_set_btn_light(Button_down, i, 0)
		}
		Elev_set_btn_light(Button_command, i, 0)
	}
}

func Elev_floor_ind(floor int){
	if floor & 0x02 {
		Io_set_bit(LIGHT_FLOOR_IND1)
	} else{
		Io_clear_bit(LIGHT_FLOOR_IND1)
	}
	if floor & 0x01 {
		Io_set_bit(LIGHT_FLOOR_IND2)
	} else{
		Io_clear_bit(LIGHT_FLOOR_IND2)
	}
}

func Elev_(){

}

func Elev_(){

}
