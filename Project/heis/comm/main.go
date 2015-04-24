package main

import (

	"RT_Prog/Project/heis/comm/mainlib"
	"time"
	"fmt"
	"RT_Prog/Project/heis/comm/tcpcommv2"
	"RT_Prog/Project/heis/comm/queue"
)

const (
	myIP = "127.0.0.1"
)

func main () {
	network_data := mainlib.Network_init("127.0.0.1")
	//command := &tcpcomm.Command_t{Command : make(map[string]string)}
	btn_queue := queue.Make_btn_queue()
	//go tcpcomm.TCP_slave(network_data.MyIP, command)
	time.Sleep(1 * time.Second)

	elev := tcpcommv2.Elev_t{UpCalls : make([]int,4), DownCalls : make([]int,4), Floor: 256, Dir: 0, IP : myIP}
	elev.UpCalls = []int{0,0,0,0} 	
	elev.DownCalls =  []int{0,0,0,0}
	elev.Dir = 0
	elev.IP = myIP

	tcp_chan := make(chan tcpcommv2.Elev_t, 10)
	go tcpcommv2.TCP_receiver_routine(myIP, "30009", tcp_chan)			
	for {
		time.Sleep(1 * time.Second)
		for queue.Empty(btn_queue) == false {
			fmt.Println(queue.Pop(btn_queue) )
		}
		
		network_data.Lock()
		if(network_data.Master == 1){
			for key, _ := range network_data.IPmap {
				tcpcommv2.TCP_send_msg(key, "30009", elev)
			}
		}
		fmt.Println(network_data.IPmap)
		network_data.Unlock()

		select {
    		case elev = <- tcp_chan:
        		fmt.Println(elev)
    		default:
        			fmt.Println("no message received")
    	}

		time.Sleep(1 * time.Second)
	}
}


