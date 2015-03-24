package main

import(
	"RT-Prog/Project/heis/driver"
	"fmt"
	"time"
	"log"
)

func main() {
	err, floor := driver.Elev_init()
	if err != 0 {
		log.Fatalln("fail")
	}
	
	call_chan := make(chan driver.Event, 2)
	
	go Buttons(call_chan)
	go Elevator(floor, call_chan)
	
	time.Sleep(100*time.Second)
	
}

func Elevator(curr_floor int, call_chan chan driver.Event) {
	var call driver.Event
	
	for {
		call =<- call_chan
		fmt.Printf("Floor : %d, Type: %d\n", call.Floor, call.Type)
		if (call.Floor >= 0) && (call.Floor < driver.N_floors) {
			driver.Elev_set_btn_light(call.Floor, call.Type, 1)
			switch {
				case call.Floor == curr_floor:
					driver.Elev_door_light(1)
					time.Sleep(1*time.Second)
					driver.Elev_door_light(0)
					//break
				case call.Floor > curr_floor:
					driver.Elev_motor(200)
					curr_floor = driver.Elev_poll_sensors(curr_floor, driver.N_floors, call.Floor)
					driver.Elev_motor(0)
					driver.Elev_set_btn_light(call.Floor, call.Type, 0)
					driver.Elev_door_light(1)
					time.Sleep(1*time.Second)
					driver.Elev_door_light(0)
					//break
				case call.Floor < curr_floor:
					driver.Elev_motor(-200)
					curr_floor = driver.Elev_poll_sensors(0, curr_floor, call.Floor)
					driver.Elev_motor(0)
					driver.Elev_set_btn_light(call.Floor, call.Type, 0)
					driver.Elev_door_light(1)
					time.Sleep(1*time.Second)
					driver.Elev_door_light(0)
					//break
			}
		}
	}
	
}

func Buttons(call_chan chan driver.Event){
	//q := make([]driver.Event, 0, 10)
	for {
		press := driver.Elev_poll_buttons()
		if (press.Floor >= 0) && (press.Floor < driver.N_floors) {
			//q = append(q, press)
			fmt.Println(press)
			call_chan <- press
		}
	}
	//fmt.Println(q)
}
