package main

import(
	 . "RT-Prog/Project/heis/driver"
	"fmt"
	"time"
	"log"
)

var elev = Elev_t{
	[4]int{0,0,0,0},
	[4]int{0,0,0,0},
	0,
	0,
	nil,
}

func main() {
	var err int
	err, elev.Floor = Elev_init()
	if err != 0 {
		log.Fatalln("fail")
	}
	
	event_chan := make(chan Event_t)
	btn_chan := make(chan Event_t)
	elev_chan := make(chan int)
	
	go buttons(btn_chan, event_chan)
	go elevator(elev_chan, event_chan)
	go eventHandler(event_chan)
	
	for {
		select {
			case b := <-btn_chan:
				fmt.Printf("Button pressed on floor: %d, type: %d\n", b.Floor, b.Type)
			case f := <-elev_chan:
				fmt.Printf("Elevator is at floor: %d\n", f)
		}
	}
	
}

func eventHandler(event_chan chan Event_t) {
	for{
		event := <- event_chan
		if (event.Floor >= 0) && (event.Floor <  N_floors) {
			switch event.Type {
				case Button_up:
					elev.UpCalls[event.Floor] = 1
					break
				case Button_down:
					elev.DownCalls[event.Floor] = 1
					break
				case Button_command:
					if elev.Floor < event.Floor{
						elev.UpCalls[event.Floor] = 1
					}else if elev.Floor > event.Floor{
						elev.DownCalls[event.Floor] = 1
					}
					break
				case Elev_done:
					elev.Floor = event.Floor
					elev.DownCalls[event.Floor] = 0
					elev.UpCalls[event.Floor] = 0
					Button_light(event.Floor, 0, 0)
					Button_light(event.Floor, 1, 0)
					break
				case Lights_up_on:
					Button_light(event.Floor, 0, 1)
					break
				case Lights_down_on:
					Button_light(event.Floor, 1, 1)
					break
				//case 6:
				//case 7:
			}
		}
	}
}

func anyJobs(callList [N_floors]int) (bool, int) {
	for i := 0; i < N_floors; i++{
		if callList[i] == 1 {
			return true, i
		}
	}
	return false, -1
}

func elevator(elev_chan chan int, event_chan chan Event_t) {	
	for {
		// any UpCalls above?
		upAbove, _ := anyJobs(elev.UpCalls)
		
		// any UpCalls below?
		upBelow, i := anyJobs(elev.UpCalls)
		
		// any DownCalls below?
		downBelow, _ := anyJobs(elev.DownCalls)
		
		// any DownCalls above?
		
		
		
		
		//fmt.Printf("Floor : %d, Type: %d\n", call.Floor, call.Type)
		if (call.Floor >= 0) && (call.Floor < N_floors) {
			Button_light(call.Floor, call.Type, 1)
			if call.Floor == elev.Floor{
				fmt.Println("Same floor")
				
				
			}else if call.Floor > elev.Floor {
				fmt.Println("Going up")
				Motor(100)
				
				// Polling sensors
				elev.Floor = Poll_sensors((elev.Floor+1), N_floors, call.Floor, (N_floors - 1))
				
				Motor(0)
				
				
			}else if call.Floor < elev.Floor {
				fmt.Println("Going down")
				Motor(-100)
				
				// Polling sensors
				elev.Floor = Poll_sensors(0, elev.Floor, call.Floor, 0)
				
				Motor(0)
				
		
			}
			Button_light(call.Floor, call.Type, 0)
			Door_light(1)
			time.Sleep(500*time.Millisecond)
			Door_light(0)
			elev_chan <- elev.Floor
			event_chan <- Event_t{elev.Floor, Elev_done}
		}
		
		/*
		if (call.Floor >= 0) && (call.Floor <  N_floors) {
			 Elev_set_btn_light(call.Floor, call.Type, 1)
			switch {
			case call.Floor == curr_floor:
				fmt.Println("Same floor")
				 Elev_set_btn_light(call.Floor, call.Type, 0)
				 Elev_door_light(1)
				time.Sleep(1*time.Second)
				 Elev_door_light(0)
				elev_chan <- curr_floor
				break
				
			case call.Floor > curr_floor:
				fmt.Println("Going up")
				 Elev_motor(200)
				curr_floor =  Elev_poll_sensors(curr_floor,  N_floors, call.Floor)
				 Elev_motor(0)
				 Elev_set_btn_light(call.Floor, call.Type, 0)
				 Elev_door_light(1)
				time.Sleep(1*time.Second)
				 Elev_door_light(0)
				elev_chan <- curr_floor
				break
				
			case call.Floor < curr_floor:
				fmt.Println("Going down")
				 Elev_motor(-200)
				curr_floor =  Elev_poll_sensors(0, curr_floor, call.Floor)
				 Elev_motor(0)
				 Elev_set_btn_light(call.Floor, call.Type, 0)
				 Elev_door_light(1)
				time.Sleep(1*time.Second)
				 Elev_door_light(0)
				elev_chan <- curr_floor
				break
		}
		*/
		
	}
	
}

func buttons(btn_chan chan Event_t, event_chan chan Event_t){
	for {
		press := Poll_buttons()
		if (press.Floor >= 0) && (press.Type < 3) && (Io_read_bit(Sensors[press.Floor]) == 0) {
			if press.Type == 2 {
				Button_light(press.Floor, press.Type, 1)
			}
			btn_chan <- press
			event_chan <- press
		}
	}
}
