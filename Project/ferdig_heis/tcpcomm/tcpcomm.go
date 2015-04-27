package tcpcomm

import (
	"net"
	"fmt"
	"encoding/json"
	"time"
	. "driver"

)

type Comm_t struct {
	Type int
	Floor int
	Dir int
	ElevDest int
	IP string
	JobID string
	State int
}



func TCP_send_msg(IP string, port string, msg Comm_t) error{
	addr, err := net.ResolveTCPAddr("tcp", IP + ":" + port)
	if err != nil {
		fmt.Println("TCP_send_error1", err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	tries := 0
	for {
		if err != nil {
			fmt.Println("TCP_send_error2", err)
			fmt.Println(IP,msg)
			time.Sleep(100 * time.Millisecond)
			tries++
		}else {
			break
		}
		if(tries == 2) {
			return err
		}
		conn, err = net.DialTCP("tcp", nil, addr)
	
	}
	buf, _ := json.Marshal(msg)
	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("TCP_send_error3")
	}
	conn.Close()
	return nil
}



func TCP_receiver_routine(myIP string, port string, comm_chan chan Comm_t){

	l, err := net.Listen("tcp", myIP + ":" + port)
	if err != nil {
		fmt.Println(err)
	}
	for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println(err)
			}
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			var comm Comm_t
			json.Unmarshal(buf[:n], &comm)
			conn.Close()
			comm_chan <- comm
			if(comm.Type == Kill) {
				l.Close()
				if err != nil {
					fmt.Println(err)
				}		
				return
			}
	}

}
