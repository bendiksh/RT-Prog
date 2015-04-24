package master

import(
	 . "RT-Prog/Project/heis/driver"
	"fmt"
	//"net"
	"RT-Prog/Project/heis/comm/mainlib"
	"math"
	. "RT-Prog/Project/heis/comm/tcpcommv2"
)

const(
	//IP = "127.0.0.1"
	PORT = "20009"
)

func Master(network_data *mainlib.Network_info) {
	msgChan := make(chan Message_t, 10)
	msgMap := make(map[string]Message_t)
	job_q := Make_job_queue()
	
	go TCP_receiver_routine(network_data.MyIP, "30009", msgChan)
	
	
	for {
		msg := <- msgChan		// received message from slave
		t := msg.Type
		switch t {
			case 0: //UpCalls
				Push(job_q, msg.Floor, 1)
				upMsg := Message_t{7,0,0,msg.Floor,""} // Dir and ElevDest used to set light
				
				for ip := range network_data.IPmap {
					TCP_send_msg(ip, PORT, upMsg)
				}
				break
			case 1: //DownCalls
				Push(job_q, msg.Floor, -1)
				downMsg := Message_t{7,0,1,msg.Floor,""} // Dir and ElevDest used to set light
				
				for ip := range network_data.IPmap {// Request status update from all elev's
					TCP_send_msg(ip, PORT, downMsg) 
				}
				break
			case 3: //Elev_done
				doneMsg := Message_t{6,msg.Floor,msg.Dir,0,""} // Floor and Dir used to clear light
				
				for ip := range network_data.IPmap {
					TCP_send_msg(ip, PORT, doneMsg) 
				}
				break
			case 7: //Status
				msgMap[msg.IP] = msg
				
				if len(msgMap) == len(network_data.IPmap){
					f, d := Pop(job_q)
					j := Event_t{f,d} // using Type as Dir
			
					bestIP := findBest(msgMap, j)
					
					bestMsg := Message_t{4,0,j.Type,j.Floor,""}
					TCP_send_msg(bestIP,PORT,bestMsg)
					
					for k := range msgMap{
						delete(msgMap, k)
					}
				}
				break
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
			// need something for when no elevator is chosen, ex request status update and run again
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
