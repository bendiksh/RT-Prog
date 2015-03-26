package driver

import . "time"
import . "math"
import . "fmt"
	
//
// Matrices of buttons, lights and sensor to loop through easier
//
var Button_matrix = [N_floors][N_buttons]int{
	{BUTTON_UP1,BUTTON_DOWN1,BUTTON_COMMAND1},
	{BUTTON_UP2,BUTTON_DOWN2,BUTTON_COMMAND2},
	{BUTTON_UP3,BUTTON_DOWN3,BUTTON_COMMAND3},
	{BUTTON_UP4,BUTTON_DOWN4,BUTTON_COMMAND4},
}

var Light_matrix = [N_floors][N_buttons]int{
	{LIGHT_UP1,LIGHT_DOWN1,LIGHT_COMMAND1},
	{LIGHT_UP2,LIGHT_DOWN2,LIGHT_COMMAND2},
	{LIGHT_UP3,LIGHT_DOWN3,LIGHT_COMMAND3},
	{LIGHT_UP4,LIGHT_DOWN4,LIGHT_COMMAND4},
}

var button=[N_floors][N_buttons]int{
	{0,0,0},
	{0,0,0},
	{0,0,0},
	{0,0,0},
}

var Sensors = [N_floors]int{SENSOR_FLOOR1,SENSOR_FLOOR2,SENSOR_FLOOR3,SENSOR_FLOOR4}

//
// Functions
//
func Elev_init() (err, floor int){ // returns a status int and a floor int
	if (int(Io_init()) != 1) {
		return 1, -1
	}
	
	// turn off all lights
	for i := 0; i < N_floors; i++ {
		if i != 0 {
			Elev_set_btn_light(i, Button_down, 0)
		}
		if i != (N_floors - 1) {
			Elev_set_btn_light(i, Button_up, 0)
		}
		Elev_set_btn_light(i, Button_command, 0)
	}

	Elev_stop_light(0)
	Elev_door_light(0)
	
	// find current floor
	floor = 0
	for i := 0; i < N_floors; i++{
		if(Io_read_bit(Sensors[i]) == 1) {
			floor = i
		}
	}
	Printf("Init floor : %d\n", floor)
	Elev_floor_ind(floor)

	return
}

func Elev_floor_ind(floor int){
	if (floor & 0x02 == 0x02) {
		Io_set_bit(LIGHT_FLOOR_IND1)
	} else{
		Io_clear_bit(LIGHT_FLOOR_IND1)
	}
	if (floor & 0x01 == 0x01) {
		Io_set_bit(LIGHT_FLOOR_IND2)
	} else{
		Io_clear_bit(LIGHT_FLOOR_IND2)
	}
}


func Elev_poll_buttons() Event_t{ 
	for i := 0; i < N_floors; i++ {
		for j := 0; j < N_buttons; j++{
			if (Io_read_bit(Button_matrix[i][j]) == 1 && button[i][j] == 0){
				button[i][j] = 1
				return Event_t{i,j} // returns on which floor (i) the button (j) has been pressed
			}else if (Io_read_bit(Button_matrix[i][j]) == 0){
				button[i][j] = 0
			}
		}
	}
	return Event_t{-1,-1}
}

func Elev_set_btn_light(floor, button, value int){
	if value == 1 {
		Io_set_bit(Light_matrix[floor][button])
	}else {
		Io_clear_bit(Light_matrix[floor][button])
	}
}

/*func Elev_floor_sensor() int{
	if (Io_read_bit(SENSOR_FLOOR1) == 1){
		return 0
	}else if (Io_read_bit(SENSOR_FLOOR2) == 1){
		return 1
	}else if (Io_read_bit(SENSOR_FLOOR3) == 1){
		return 2
	}else if (Io_read_bit(SENSOR_FLOOR4) == 1){
		return 3
	}else {
		return -1
	}
}*/

func Elev_poll_sensors(low, high, goal int) int {
	done := false
	curr := 0
	for !done {
		for i := low; i < high; i++ {
			Sleep(1 * Millisecond)
			if Io_read_bit(Sensors[i]) == 1 {
				if i != curr {
					Printf("Floor sensor: %d\n", i)
				}
				curr = i
				Elev_floor_ind(i)
				if i == goal || i == N_floors - 1 {
					done = true
					break
				}
			}
		}
	}
	return curr
}

func Elev_stop_light(i int){
	if (i == 1) {
		Io_set_bit(LIGHT_STOP)
	}else{
		Io_clear_bit(LIGHT_STOP)
	}
}

func Elev_door_light(i int){
	if (i == 1) {
		Io_set_bit(LIGHT_DOOR_OPEN)
	}else{
		Io_clear_bit(LIGHT_DOOR_OPEN)
	}
}

var prev_speed int

func Elev_motor(speed int){
	if ( speed > 0 ){
		Io_clear_bit(MOTORDIR)
	}else if (speed < 0){
		Io_set_bit(MOTORDIR)
	}else{
		if (prev_speed < 0){
			Io_clear_bit(MOTORDIR)
		}else if(prev_speed > 0){
			Io_set_bit(MOTORDIR)
		}
		Sleep(10*Millisecond) // time.Sleep and time.Millisecond
	}
	prev_speed = speed
	
	// write speed to motor
	Io_write_analog(MOTOR, int(2048 + 4*Abs(float64(speed)))) // math.Abs
}
