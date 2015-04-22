package main

import(
	 . "RT-Prog/Project/heis/driver"
	"fmt"
	"net"
	"RT-Prog/Project/comm/mainlib"
	"math"
)

func main() {
	msgChan
	var msg Message_t
	var jobsReceived []Event_t// ?? send to another function?? {Floor, Dir}
	msgMap = make(map[string]*Message_t)
	
	for {
		msg = <- msgChan		// received message from slave
		switch msg.Type {
			case 0: //UpCalls
				jobsReceived = append(jobsReceived, [msg.ElevDist, 1])
				
				TCP_send_msg(IP, PORT, {7,0,1,msg.Floor}) // request status update from slaves
				
			case 1://DownCalls
				jobsReceived = append(jobsReceived, {msg.ElevDist, -1})
				
				TCP_send_msg(IP, PORT, {7,0,-1,msg.Floor}) // request status update from slaves
			case 7:
				msgMap[msg.IP] = append(msgMap[msg.IP], msg) // can't append maps
		}
		if len(msgMap) == len(IPmap){
			job := jobsReceived[:1]
			findBest(msgMap) // run as goroutine?
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
