package main

import(
	. "driver"
	"fmt"
	"time"
	"netlib"
	"tcpcomm"
)

type newJob_t struct{
	floor int
	jobType int
	dir int
}

type status_t struct{
	floor int
	dir int
	dest int
	state int
}


var lights [3][N_floors]int

var network_data *netlib.Network_info

func jobFinished(jobType int, floor int, comm_chan chan tcpcomm.Comm_t) {
	if(jobType != 2) {
		var comm tcpcomm.Comm_t
		comm.Type = Elev_done
		comm.Floor = floor
		comm.Dir = jobType
		comm.ElevDest = -1
		comm.IP = network_data.MyIP
		comm_chan <- comm
	}
} 

func elevate(newJob_chan chan newJob_t, comm_chan chan tcpcomm.Comm_t, status_chan chan status_t) {
	var Calls[3][N_floors]int
	floor_chan := make(chan int, 10)
	go floor(floor_chan)
	_, floor := Elev_init()
	state := IDLE 
	dest := floor
	var dir int 
	for {
			status_chan <- status_t{floor,dir,dest,state}
			switch state {
				case IDLE:
					Motor(0)
					fmt.Println("IDLE")
					loop:
						for {
							select {
								case newJob := <- newJob_chan:
									Calls[newJob.jobType][newJob.floor] = 1
									break
								default:
									break loop
							}
						}
					find_job_loop:
						for i:=2;i>-1;i-- {
							for j:=0;j<N_floors;j++ {
								old_j := j
								if(i == 1){
									j = N_floors - j - 1
								} 
								if(Calls[i][j] == 1) {
									if(j != floor){
										state = i
										dest = j
										dir = getDir(floor, dest)
										break find_job_loop
									}else{
										openDoor(dir, floor)
										Calls[i][j] = 0
										jobFinished(i,j,comm_chan)
									}
								}
								j = old_j
							}
						}
					if(state != IDLE) {
						break
					}
					select {
						case newJob := <- newJob_chan:
							if(newJob.floor != floor){
								state = newJob.jobType
								Calls[newJob.jobType][newJob.floor] = 1
								dest = newJob.floor
								dir = getDir(floor, newJob.floor)
							}else{
								Calls[newJob.jobType][newJob.floor] = 0
								jobFinished(newJob.jobType, newJob.floor, comm_chan)
								openDoor(dir, floor)
								dir = 0
								state = IDLE
							}
							break
						case floor = <- floor_chan:
							fmt.Println("Be carefull, I'm fragile! ")
							break	
					}
					break

				case INSIDECALL:
					fmt.Println("INSIDECALL")
					Motor(dir * 200)
					select {
						case newJob := <- newJob_chan:
							Calls[newJob.jobType][newJob.floor] = 1
							if(dir == -1){
								dest = min(dest, newJob.floor)
							}else if(dir == 1){
								dest = max(dest, newJob.floor)
							}
							break
						case floor = <- floor_chan:
								stop := false
								if(Calls[2][floor] == 1){
									Calls[2][floor] = 0
									stop = true
								}
								if(Calls[0][floor] == 1 && dir == 1) {
									Calls[0][floor] = 0
									jobFinished(0,floor,comm_chan)
									stop = true
								}
								if(Calls[1][floor] == 1 && dir == -1) {
									Calls[1][floor] = 0
									jobFinished(1,floor,comm_chan)
									stop = true
								}
								if(floor == dest){
										state = IDLE
								}
								if(stop == true){
									openDoor(dir, floor)
									dir = getDir(floor, dest)
								}
								break	
					}
					break
				case UPCALL:
					fmt.Println("UPCALL")
					Motor(200 * getDir(floor, dest))
					select {
						case newJob := <- newJob_chan:
							Calls[newJob.jobType][newJob.floor] = 1
							if(dest > newJob.floor && newJob.jobType == 0 && dir < 0) {
								dest = newJob.floor
							}
						case floor = <- floor_chan:
							stop := false
							if(Calls[2][floor] == 1) {
								stop = true
								Calls[2][floor] = 0
							}
							if(floor == dest) {
								stop = true
								Calls[0][floor] = 0
								Calls[2][floor] = 0
								jobFinished(0,floor,comm_chan)
								state = IDLE
							}
							if(stop == true){
									openDoor(dir, floor)
									dir = getDir(floor, dest)
								}
							break
					}
				case DOWNCALL:
					fmt.Println("DOWNCALL")
					Motor(200 * getDir(floor, dest))
					select {
						case newJob := <- newJob_chan:
							Calls[newJob.jobType][newJob.floor] = 1
							if(dest < newJob.floor && newJob.jobType == 1 && dir > 0) {
								dest = newJob.floor
							}
							break
						case floor = <- floor_chan:
							stop := false
							if(Calls[2][floor] == 1) {
								stop = true
								Calls[2][floor] = 0
							}
							if(floor == dest) {
								Calls[1][floor] = 0
								jobFinished(1,floor,comm_chan)
								stop = true
								state = IDLE
							}
							if(stop == true){
									openDoor(dir, floor)
									dir = getDir(floor, dest)
								}
							break
					}
			}
		}	
}

func getDir(from int, to int) int {
	if(from > to){
		return -1
	}else if(from < to){
		return 1
	}else{
		return 0
	}
}

func floor(floor_chan chan int) {
	last_floor := Get_floor()
	for {
		floor := Get_floor()
		if(floor != -1) {
			if(floor != last_floor) {
				Floor_ind(floor)
				floor_chan <- floor
			}
			last_floor = floor
		}
	time.Sleep(50*time.Millisecond)
	}
}

func buttons(event_chan chan Event_t){
	for {
		press := Poll_buttons()
		if (press.Floor >= 0) && (press.Type < 3) {
			event_chan <- press
		}
		time.Sleep(10*time.Millisecond)
	}
}

func openDoor(dir int, floor int) {
	if(dir != 0) {
		Motor(2 * dir)
		time.Sleep(150*time.Millisecond)
	}
	Motor(0)
	Button_light(floor, 2, 0)
	Door_light(1)
	time.Sleep(2*time.Second)
	Door_light(0)
}

func eventHandler(comm_chan chan tcpcomm.Comm_t, newJob_chan chan newJob_t) {
	event_chan := make(chan Event_t,50)
	go buttons(event_chan)
	var comm tcpcomm.Comm_t
	var newJob newJob_t
	for{
		event := <- event_chan

		if (event.Floor >= 0) && (event.Floor <  N_floors) {
			switch event.Type {
				case Button_up:
					comm.Type = Button_up
					comm.Floor = event.Floor
					comm_chan <- comm
					break
				case Button_down:
					comm.Type = Button_down
					comm.Floor = event.Floor
					comm_chan <- comm
					break
				case Button_command:
					Button_light(event.Floor, 2, 1)	
					newJob.jobType = 2
					newJob.floor = event.Floor
					newJob_chan <- newJob
					break
			}
		}
		time.Sleep(50*time.Millisecond)
	}
}

func networkHandler(comm_chan chan tcpcomm.Comm_t, newJob_chan chan newJob_t, 
					status_chan chan status_t) {
	recv_comm_chan := make(chan tcpcomm.Comm_t, 10)
	go tcpcomm.TCP_receiver_routine(network_data.MyIP, S_port, recv_comm_chan)
	var newJob newJob_t
	var status status_t
	for {
		select {
			case status = <- status_chan:
				break			
    		case comm := <- recv_comm_chan:
				if(comm.Type == Status) {
					comm.Type = Status
					comm.Floor = status.floor
					comm.Dir = status.dir
					comm.ElevDest = status.dest
					comm.IP = network_data.MyIP
					comm.State = status.state
					err := tcpcomm.TCP_send_msg(network_data.MasterIP, M_port, comm)
					for {
						if err != nil {
							fmt.Println("TCP_send_error_timeout", err)
							time.Sleep(500 * time.Millisecond)
						}else{
							break
						}
						err = tcpcomm.TCP_send_msg(network_data.MasterIP, M_port, comm)
					}
				}else if(comm.Type == Job) {
					newJob.jobType = comm.Dir
					newJob.floor = comm.ElevDest
					newJob_chan <- newJob
				}else if(comm.Type == Turn_off_lights) {
					lights[comm.Dir][comm.ElevDest] = 0
					Button_light(comm.ElevDest, comm.Dir, 0)
				}else if(comm.Type == Turn_on_lights) {
					lights[comm.Dir][comm.ElevDest] = 1
					Button_light(comm.ElevDest, comm.Dir, 1)
				}
			case comm := <- comm_chan:
				err := tcpcomm.TCP_send_msg(network_data.MasterIP, M_port, comm)
					for {
						if err != nil {
							fmt.Println("TCP_send_error_timeout", err)
							time.Sleep(500 * time.Millisecond)
						}else{
							break
						}
						err = tcpcomm.TCP_send_msg(network_data.MasterIP, M_port, comm)
					}
		}
		time.Sleep(50*time.Millisecond)
	}
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
