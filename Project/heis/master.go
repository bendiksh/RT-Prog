package main

import(
	 . "RT-Prog/Project/heis/driver"
	"fmt"
	"net"
	"RT-Prog/Project/heis/comm/mainlib"
	"math"
)

func main() {
	msgChan
	msgMap = make(map[string]*Message_t)
	job_q := Make_job_queue()
	job := Event_t{}
	
	for {
		msg := <- msgChan		// received message from slave
		t := msg.Type
		switch t {
			case 0: //UpCalls
				Push(job_q, msg.IP, msg.Floor, 1)
				
				TCP_send_msg(IP, PORT, {7,0,1,msg.Floor}) // request status update from slaves, last two fields are used to set lights
				
			case 1://DownCalls
				Push(job_q, msg.IP, msg.Type, msg.Floor, -1)
				
				TCP_send_msg(IP, PORT, {7,0,-1,msg.Floor}) // request status update from slaves, last two fields are used to set lights
			case 7:
				msgMap[msg.IP] = msg
		}
		if len(msgMap) == len(IPmap){
			f, d := Pop(job_q)
			j := Event_t{f,d} // using Type as Dir
			
			findBest(msgMap, j) // run as goroutine?
		}	
	}
}

func findBest(elevMap map[string]*Message_t, job Event_t) {
	var bestIP string
	var distance int
	bestDist := 99
	
	for key, val := range elevMap {
		distance = int(math.Abs(float64(job.Floor) - float64(val.Floor)))
		if distance < bestDist {
			if job.Type == val.Dir {
				bestDist = distance
				bestIP = key
			}else if val.Dir == 0 {
				bestDist = distance
				bestIP = key
			}
			// need something for when no elevator is chosen, ex request status update and run again
		}
	}
	// return something or send message directly from here
}
