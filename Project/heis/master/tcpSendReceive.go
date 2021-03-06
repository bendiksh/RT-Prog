package main

import (
	"net"
	"fmt"
	"encoding/json"
	. "RT-Prog/Project/heis/driver"

)

/*type Elev_t struct {
	UpCalls []int
	DownCalls []int
	Floor int
	Dir int
	IP string
}*/

const(
	LPORT="20009"
	WPORT="30009"
)

func main(){
	upMsg := Message_t{Button_up, 3, 0, 0, ""}
	doneMsg := Message_t{Elev_done, 2, 0, 0, ""}
	statusMsg := Message_t{Status, 1, 0, 1, ""}
	
	msgChan := make(chan Message_t, 10)
	
	network_data := *mainlib.Network_info
	
	go TCP_receiver_routine(network_data.MyIP, LPORT, msgChan)
	
	TCP_send_msg(IP, WPORT, upMsg)
	
	for {
		msg := <- msgChan
		switch msg.Type {
			case Job:
				fmt.Println("Received job")
				time.Sleep(2*time.Second)
				TCP_send_msg(IP, WPORT, doneMsg)
				fmt.Println("Elev_done sent")
				break
			case Status:
				fmt.Println("Received status request")
				TCP_send_msg(IP, WPORT, statusMsg)
		}
	}
}

func TCP_send_msg(IP string, port string, msg Message_t){
	addr, err := net.ResolveTCPAddr("tcp", IP + ":" + port)
	if err != nil {
		fmt.Println("TCP_send_error1", err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("TCP_send_error2", err)
		return
	}
	buf, _ := json.Marshal(msg)
	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("TCP_send_error3")
	}
	conn.Close()
}

func TCP_receiver(myIP string, port string) Message_t{
	l, err := net.Listen("tcp", myIP + ":" + port)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	l.Close()
	conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	var returnValue Message_t
	json.Unmarshal(buf[:n], &returnValue)
	return returnValue
}


func TCP_receiver_routine(myIP string, port string, comm chan Message_t){
	for {
		comm <- TCP_receiver(myIP, port)
	}

}
