package main

import(
	"RT-Prog/Project/ex5/driver"
	"fmt"
	"time"
)

func main() {
	err, floor := driver.Elev_init()
	if err != 1 {
		fmt.Println("Could not initialize")
		return
	}
	
	External_call(floor)
	
	
}

func External_call(curr_floor int) {
	var press driver.Event
	//var p_btn driver.Event.Type
	
	for {
		press = driver.Elev_get_button()
		
		fmt.Printf("Floor : %d 		Type : %d\n", press.Floor, press.Type)
		
		if (press.Floor >= 0) && (press.Floor < driver.N_floors) {
		
			if press.Floor == curr_floor {
				// only need to open door
				fmt.Println("press=curr")
				driver.Elev_door_light(1)
				time.Sleep(1*time.Second)
				driver.Elev_door_light(0)
			} else if press.Floor > curr_floor {
				//call from below
				fmt.Println("press>curr")
				
				// Start elevator
				driver.Elev_motor(100)
				fmt.Println("elevator going up")
				
				// Poll sensors until the elevator is on the right floor
				done := false
				for !done {
					for i := curr_floor; i < driver.N_floors; i++ {
						time.Sleep(1*time.Millisecond)
						if driver.Io_read_bit(driver.Sensors[i]) == 1 {
							curr_floor = i
							fmt.Printf("up: %d\n", i)
							if i == press.Floor || i == driver.N_floors - 1 {
								done = true
								break
							}
						}
					}
				}
				driver.Elev_motor(0)
				
				driver.Elev_door_light(1)
				time.Sleep(1*time.Second)
				driver.Elev_door_light(0)
			} else if press.Floor < curr_floor {
				// call from above
				fmt.Println("press<curr")
				// Poll sensors until the elevator is on the right floor
				
				// Start elevator
				driver.Elev_motor(-100)
				fmt.Println("elevator going down")
				
				// Poll sensors until the elevator is on the right floor
				done := false
				for !done {
					for i := 0; i < curr_floor; i++ {
						time.Sleep(1*time.Millisecond)
						if driver.Io_read_bit(driver.Sensors[i]) == 1 {
							curr_floor = i
							fmt.Printf("down: %d\n", i)
							if i == press.Floor || i == driver.N_floors - 1 {
								done = true
								break
							}
						}
					}
				}
				driver.Elev_motor(0)
				
				driver.Elev_door_light(1)
				time.Sleep(1*time.Second)
				driver.Elev_door_light(0)
			}
		}
	}
}
