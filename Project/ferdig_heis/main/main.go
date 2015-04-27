package main

import(
	. "driver"
	"fmt"
	"time"
	"netlib"
	"master"
	"tcpcomm"
)

func main() {
	network_data = netlib.Network_init()
	
	time.Sleep(2 * time.Second)

	
	comm_chan := make(chan tcpcomm.Comm_t, 10)
	newJob_chan := make(chan newJob_t, 10)
	status_chan := make(chan status_t, 10)
	
	go networkHandler(comm_chan, newJob_chan, status_chan)
	go eventHandler(comm_chan, newJob_chan)
	go elevate(newJob_chan, comm_chan, status_chan)
	for {
		
		if(network_data.TCP_master_started == 0 && network_data.Master == 1) {
			go master.Master(network_data)
			network_data.TCP_master_started = 1
			for i:=0;i<N_floors;i++ {
				for j:=0;j<3;j++ {
					if(lights[j][i] == 1) {
						comm_chan <- 
						tcpcomm.Comm_t{j,i,j,i,network_data.MyIP, "", IDLE}
					}
				}
			}
		}else if(network_data.Master == 0 && network_data.TCP_master_started == 1) {
			tcpcomm.TCP_send_msg(network_data.MyIP, M_port, 
			tcpcomm.Comm_t{Kill, -1, -1, -1, "DIE!", "-1-1", IDLE})
			network_data.TCP_master_started = 0
			fmt.Println("NOT master")
		}
		time.Sleep(1 * time.Second)
		
	}
	
}
