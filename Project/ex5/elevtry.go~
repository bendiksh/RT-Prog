package main

import(
	"RT-Prog/Project/ex5/driver"
	"fmt"
	"time"
)

func main(){
	err, floor := driver.Elev_init()
	if err != 1 {
		fmt.Println("Could not initialize")
		return
	}
	
	curr_floor := floor
	
	for {
		p_floor, p_btn := driver.Elev_get_button()
		if (p_ floor == curr_floor) {
		// only need to open door
			driver.Elev_door_light(1)
			time.Sleep(1*Second)
			driver.Elev_door_light(0)
		}else if(p_floor < curr_floor) {
		//call from below
			for driver.Sensors[p_floor] != 1 {
				driver.Elev_motor(-10)
			}
			driver.Elev_motor(0)
			driver.Elev_door_light(1)
			time.Sleep(1*Second)
			driver.Elev_door_light(0)
		}else if(p_floor > curr_floor) {
		// call from above
			for driver.Sensors[p_floor] != 1 {
				driver.Elev_motor(10)
			}
			driver.Elev_motor(0)
			driver.Elev_door_light(1)
			time.Sleep(1*Second)
			driver.Elev_door_light(0)
		}
	}
}
