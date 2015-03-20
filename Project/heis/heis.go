package main

import(
	"RT-Prog/Project/heis/driver"
	"fmt"
	//"time"
	"log"
)

func main() {
	err, floor := driver.Elev_init()
	if err != 0 {
		log.Fatalln("fail")
	}
	
	go Buttons()
	go Elevator(floor)
	
}

func Elevator(curr_floor int) {
	
}

func Buttons(){
	q := make([]driver.Event, 0, 10)
	for {
		press := driver.Elev_poll_buttons()
		q = append(q, press)
	}
	fmt.Println(q)
}
