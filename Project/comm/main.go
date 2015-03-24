package main

import (

	"mainlib"
	"time"
	"fmt"
	"tcpcomm"
	"queue"
)

const (
	sAddr = "129.241.187.152:34933"
	cAddr = "127.0.0.1:30009"
)

func main () {
	network_data := mainlib.Network_init("129.241.187.141")
	command := &tcpcomm.Command_t{Command : make(map[string]string)}
	btn_queue := queue.Make_btn_queue()
	tcpcomm.TCP_slave(network_data.MyIP, command)
	time.Sleep(1 * time.Second)

	testpress := make(map[string]string)
	testpress["button id"] = "knappabc"	
	testpress["message type"] = "button press"	
	for {
		time.Sleep(1 * time.Second)
		for queue.Empty(btn_queue) == false {
			fmt.Println(queue.Pop(btn_queue) )
		}

		network_data.Lock()
		if(network_data.Master == 1 && network_data.TCP_master_started == 0){
			go tcpcomm.TCP_master_routine(network_data, btn_queue)
		}
		fmt.Println(network_data.IPmap)
		network_data.Unlock()
		time.Sleep(1 * time.Second)
		tcpcomm.TCP_slave_send_msg(testpress, network_data.MyIP)
	}
}


