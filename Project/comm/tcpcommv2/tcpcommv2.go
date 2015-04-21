package tcpcommv2

import (
	"net"
	"fmt"
	"encoding/json"

)

type Elev_t struct {
	UpCalls []int
	DownCalls []int
	Floor int
	Dir int
	IP string
}

func TCP_send_msg(IP string, port string, msg Elev_t){
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

func TCP_receiver(myIP string, port string) Elev_t{
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
	var returnValue Elev_t
	json.Unmarshal(buf[:n], &returnValue)
	return returnValue
}


func TCP_receiver_routine(myIP string, port string, comm chan Elev_t){
	for {
		comm <- TCP_receiver(myIP, port)
	}

}