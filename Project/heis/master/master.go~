package main

import(
	 . "RT-Prog/Project/heis/driver"
	"fmt"
	"time"
	"RT-Prog/Project/heis/comm/mainlib"
	"math"
	. "RT-Prog/Project/heis/comm/tcpcommv2"
)

const(
	//IP = "127.0.0.1"
	PORT = "20009"
)

func main() {
	Master(&mainlib.Network_info)
}

func Master(network_data *mainlib.Network_info) {
	msgChan := make(chan Message_t, 10)
	
	msgMap := make(map[string]Message_t)
	sentMap := make(map[Message_t]int64)
	
	job_q := Make_job_queue()
	
	go TCP_receiver_routine(network_data.MyIP, "30009", msgChan)
	go sentJobs(&sentMap)
	
	for {
		msg := <- msgChan		// received message from slave
		t := msg.Type
		switch t {
			case Button_up:
				fmt.Println("Received call")
				Push(job_q, msg.Floor, 1)
				upMsg := Message_t{7,0,0,msg.Floor,""} // Dir and ElevDest used to set light
				
				for ip := range network_data.IPmap {
					TCP_send_msg(ip, PORT, upMsg)
				}
				break
			case Button_down:
				Push(job_q, msg.Floor, -1)
				downMsg := Message_t{7,0,1,msg.Floor,""} // Dir and ElevDest used to set light
				
				for ip := range network_data.IPmap {// Request status update from all elev's
					TCP_send_msg(ip, PORT, downMsg) 
				}
				break
			case Elev_done:
				fmt.Println("Received Elev_done")
				
				/*doneMsg := Message_t{6,msg.Floor,msg.Dir,0,""} // Floor and Dir used to clear light
				
				for ip := range network_data.IPmap {
					TCP_send_msg(ip, PORT, doneMsg) 
				}
				break*/
			case Status:
				fmt.Println("Received status")
				msgMap[msg.IP] = msg
				
				if len(msgMap) == len(network_data.IPmap){
					f, d := Pop(job_q)
					j := Event_t{f,d} // using Type as Dir
			
					bestIP := findBest(msgMap, j)
					
					bestMsg := Message_t{4,0,j.Type,j.Floor,bestIP}
					TCP_send_msg(bestIP,PORT,bestMsg)
					sentMap[bestMsg] := time.Now().UnixNano()
					
					for k := range msgMap{
						delete(msgMap, k)
					}
				}
		}
			
	}
}

func findBest(elevMap map[string]Message_t, job Event_t) string{
	var bestIP, closestIP string
	var distance int
	bestDist, closest := 99, 99
	
	for key, val := range elevMap {
		distance = int(math.Abs(float64(job.Floor) - float64(val.Floor)))
		if distance < bestDist {
			if job.Type == val.Dir {
				bestDist = distance
				bestIP = key
			}else if val.Dir == 0 {
				bestDist = distance
				bestIP = key
			}else if val.ElevDest == job.Floor{
				bestDist = distance
				bestIP = key
			}
		}else if distance < closest {
			closest = distance
			closestIP = key
		}
	}
	if bestDist < 10{
		return bestIP
	}else {
		return closestIP
	}
}

func sentJobs(sentMap *map[Message_t]int64) {
	
	for {
		t := time.Now().UnixNano()
		for k, v := range sentMap {
			if (t - v) > (5 * 1000000000) {
				fmt.Println("took too long")
			}
		}
		time.Sleep(1 * time.Second)
	}
}































