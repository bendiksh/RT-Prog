package master

import(
	 . "driver"
	"fmt"
	"time"
	"netlib"
	"strconv"
	"math"
	. "tcpcomm"
)


func Master(network_data *netlib.Network_info) {
	fmt.Println("master")
	msgChan := make(chan Comm_t, 20)
	quitSentJobsChan := make(chan bool)
	
	sentMap := make(map[string]int)
	
	jobMap := make(map[string]map[string]Comm_t)
	go TCP_receiver_routine(network_data.MyIP, M_port, msgChan)
	go sentJobsCheckup(sentMap, network_data.IPmap, quitSentJobsChan, network_data.MyIP)
	
	
	
	for {
		msg := <- msgChan
			time.Sleep(100 * time.Millisecond)		// received message from slave
			t := msg.Type

			switch t {
				case Button_up:
					_, ok := jobMap[fmt.Sprintf("%d%d", msg.Type, msg.Floor)]
					if !ok {
						upMsg := Comm_t{Status,0,0,0,"", fmt.Sprintf("%d%d", msg.Type, msg.Floor), -1}
						jobMap[fmt.Sprintf("%d%d", msg.Type, msg.Floor)] = make(map[string]Comm_t)
						for ip := range network_data.IPmap {// Request status update from all elev's
						TCP_send_msg(ip, S_port, upMsg) 
					}
					}else {
						break
					}
					break
				case Button_down:
					_, ok := jobMap[fmt.Sprintf("%d%d", msg.Type, msg.Floor)]
					if !ok {
						downMsg := Comm_t{Status,0,0,0,"", fmt.Sprintf("%d%d", msg.Type, msg.Floor), -1}
						jobMap[fmt.Sprintf("%d%d", msg.Type, msg.Floor)] = make(map[string]Comm_t)
						for ip := range network_data.IPmap {// Request status update from all elev's
						TCP_send_msg(ip, S_port, downMsg) 
					}
					}
					break
				case Elev_done:
					sentMap[msg.IP] = sentMap[msg.IP] - 1
					msg.Type = Turn_off_lights
					msg.ElevDest = msg.Floor
					msg.Floor = msg.Floor
					msg.Dir = msg.Dir
					for ip := range network_data.IPmap {
						TCP_send_msg(ip, S_port, msg)
					}

				case Status:
						jobMap[msg.JobID][msg.IP] = msg
						
					if len(jobMap[msg.JobID]) == len(network_data.IPmap) {
						f, _ := strconv.ParseInt(string(msg.JobID[1:]),0,32)
						d, _ := strconv.ParseInt(string(msg.JobID[0]),0,32)
						j := Event_t{int(f), int(d)}
						
						bestIP := findBest(jobMap[msg.JobID], j)
						bestMsg := Comm_t{Job,0, j.Type, j.Floor, bestIP, msg.JobID, -1}
						
						TCP_send_msg(bestIP, S_port, bestMsg)
						
						msg.Type = Turn_on_lights
						msg.ElevDest = int(f)
						msg.Dir = bestMsg.Dir
						
						time.Sleep(50 * time.Millisecond)
						
						for ip := range network_data.IPmap {
							TCP_send_msg(ip, S_port, msg)
						}
						delete(jobMap, msg.JobID )
						
						val, ok := sentMap[bestMsg.IP]
						if !ok {
						sentMap[bestMsg.IP] = 1
						}else{
							sentMap[bestMsg.IP] = val + 1
						}
					}
				case Kill:
					quitSentJobsChan <- true
					network_data.TCP_master_started = 0
					return
			}
			
	}
}

func findBest(elevMap map[string]Comm_t, job Event_t) string{

	var jobDir = (-2 * job.Type) + 1
	var min_cost int = 100000
	var cost int
	var lowest_cost_ip string
	
	for key, val := range elevMap {
		if(val.State == IDLE) {
			cost = 10 * int(math.Abs(float64(val.Floor - job.Floor)))
		}else if(val.State == INSIDECALL) {
			if(val.Dir < 0) {
				if(val.ElevDest < job.Floor && val.Floor > job.Floor && jobDir == -1){
					cost = 8 * int(math.Abs(float64(val.Floor - job.Floor)))
				}else{
					cost = 21 * int(math.Abs(float64(val.ElevDest - job.Floor)))
				}
			}else {
				if(val.ElevDest > job.Floor && val.Floor < job.Floor && jobDir == 1){
					cost = 8 * int(math.Abs(float64(val.Floor - job.Floor)))
				}else{
					cost = 21 * int(math.Abs(float64(val.ElevDest - job.Floor)))
				}
			}
			
		}else if(val.State == UPCALL) {
			if(job.Type == UPCALL) {
				cost = 4 * int(math.Abs(float64(val.Floor - job.Floor)))
			}else {
				cost = 18 * int(math.Abs(float64(val.ElevDest - job.Floor))) + 30
			}
		}else if(val.State == DOWNCALL) {
			if(job.Type == DOWNCALL) {
				cost = 4 * int(math.Abs(float64(val.Floor - job.Floor)))
			}else {
				cost = 18 * int(math.Abs(float64(val.ElevDest - job.Floor))) + 30
			}
		}else{
			fmt.Println("state not recognized")
		}
		if(cost < min_cost) {
			min_cost = cost
			lowest_cost_ip = key
		}
		
	}
	return lowest_cost_ip
}

func sentJobsCheckup(sentMap map[string]int,  IPmap map[string]int, quitChan chan bool, MyIP string) {
	for {
	
		select {
		case <- quitChan:
			return
		default:
			for key, _ := range sentMap {
				_, ok := IPmap[key]
				if !ok {
					var comm Comm_t
					comm.Type = Kill
					fmt.Println("restarting master")
					TCP_send_msg(MyIP, M_port, comm)
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}



