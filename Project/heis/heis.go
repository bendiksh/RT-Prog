package main

import(
	"RT-Prog/Project/heis/driver"
	"fmt"
	"time"
	"log"
)

var elev = driver.Elev_t{
	[4]int{0,0,0,0},
	[4]int{0,0,0,0},
	0,
	0,
	nil,
}

func main() {
	var err int
	err, elev.Floor = driver.Elev_init()
	if err != 0 {
		log.Fatalln("fail")
	}
	
	call_chan := make(chan driver.Event_t)
	btn_chan := make(chan driver.Event_t)
	elev_chan := make(chan int)
	//Job_done := make(chan int)
	
	go Buttons(btn_chan, call_chan)
	go Elevator(elev_chan, call_chan)
	
	for {
		select {
			case b := <-btn_chan:
				fmt.Printf("Button pressed on floor: %d, type: %d\n", b.Floor, b.Type)
			case f := <-elev_chan:
				fmt.Printf("Elevator is at floor: %d\n", f)
		}
	}
	
}

func HandleJobs() {
	
}

func Elevator(elev_chan chan int, call_chan chan driver.Event_t) {
	//var call driver.Event_t
	
	for {
		call := <- call_chan
		fmt.Printf("Floor : %d, Type: %d\n", call.Floor, call.Type)
		if (call.Floor >= 0) && (call.Floor < driver.N_floors) {
			driver.Elev_set_btn_light(call.Floor, call.Type, 1)
			if call.Floor == elev.Floor{
				fmt.Println("Same floor")
				
				
			}else if call.Floor > elev.Floor {
				fmt.Println("Going up")
				driver.Elev_motor(100)
				
				// Polling sensors
				elev.Floor = driver.Elev_poll_sensors((elev.Floor+1), driver.N_floors, call.Floor, (driver.N_floors -1))
				/*done := false
				//curr := 0
				for !done {
					for i := elev.Floor + 1; i < driver.N_floors; i++ {
						time.Sleep(1 * time.Millisecond)
						if driver.Io_read_bit(driver.Sensors[i]) == 1 {
							//if i != curr {
								fmt.Printf("Floor sensor: %d\n", i)
							//}
							elev.Floor = i
							driver.Elev_floor_ind(i)
							if (i == call.Floor) || (i == driver.N_floors - 1) {
								done = true
								//Job_done <- i
								break
							}
						}
					}
				}*/
				
				//curr_floor =<- Job_done
				driver.Elev_motor(0)
				
				
			}else if call.Floor < elev.Floor {
				fmt.Println("Going down")
				driver.Elev_motor(-100)
				
				// Polling sensors
				elev.Floor = driver.Elev_poll_sensors(0, elev.Floor, call.Floor, 0)
				/*done := false
				//curr := 0
				for !done {
					for i := 0; i < elev.Floor; i++ {
						time.Sleep(1 * time.Millisecond)
						if driver.Io_read_bit(driver.Sensors[i]) == 1 {
							//if i != curr {
								fmt.Printf("Floor sensor: %d\n", i)
							//}
							//curr = i
							driver.Elev_floor_ind(i)
							if (i == call.Floor) || (i == 0) {
								done = true
								//Job_done <- i
								break
							}
						}
					}
				}*/
				
				//curr_floor =<- Job_done
				driver.Elev_motor(0)
				
		
			}
			driver.Elev_set_btn_light(call.Floor, call.Type, 0)
			driver.Elev_door_light(1)
			time.Sleep(1*time.Second)
			driver.Elev_door_light(0)
			elev_chan <- elev.Floor
		}
		
		/*
		if (call.Floor >= 0) && (call.Floor < driver.N_floors) {
			driver.Elev_set_btn_light(call.Floor, call.Type, 1)
			switch {
			case call.Floor == curr_floor:
				fmt.Println("Same floor")
				driver.Elev_set_btn_light(call.Floor, call.Type, 0)
				driver.Elev_door_light(1)
				time.Sleep(1*time.Second)
				driver.Elev_door_light(0)
				elev_chan <- curr_floor
				break
				
			case call.Floor > curr_floor:
				fmt.Println("Going up")
				driver.Elev_motor(200)
				curr_floor = driver.Elev_poll_sensors(curr_floor, driver.N_floors, call.Floor)
				driver.Elev_motor(0)
				driver.Elev_set_btn_light(call.Floor, call.Type, 0)
				driver.Elev_door_light(1)
				time.Sleep(1*time.Second)
				driver.Elev_door_light(0)
				elev_chan <- curr_floor
				break
				
			case call.Floor < curr_floor:
				fmt.Println("Going down")
				driver.Elev_motor(-200)
				curr_floor = driver.Elev_poll_sensors(0, curr_floor, call.Floor)
				driver.Elev_motor(0)
				driver.Elev_set_btn_light(call.Floor, call.Type, 0)
				driver.Elev_door_light(1)
				time.Sleep(1*time.Second)
				driver.Elev_door_light(0)
				elev_chan <- curr_floor
				break
		}
		*/
		
	}
	
}

func Buttons(btn_chan chan driver.Event_t, call_chan chan driver.Event_t){
	for {
		press := driver.Elev_poll_buttons()
		if (press.Floor >= 0) && (press.Type < 3) && (driver.Io_read_bit(driver.Sensors[press.Floor]) == 0) {
			if press.Type == 2 {
				driver.Elev_set_btn_light(press.Floor, press.Type, 1)
			}
			//fmt.Println(press)
			btn_chan <- press
			call_chan <- press
		}
	}
	//fmt.Println(q)
}
