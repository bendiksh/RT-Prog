package tcpcomm

import (
	"net"
	"fmt"
	"sync"
	"encoding/json"
	"queue"
	"strconv"
	"mainlib"
	"time"
)

const TCP_port = "30009"
var slave_conn net.Conn
var connections map[string]connection_t


type connection_t struct {
	IP string
	ADDR *net.TCPAddr
	Conn *net.TCPConn
	err error
}

func TCP_connect(IP string) connection_t{
	connection := connection_t{IP: IP} 
	connection.ADDR, connection.err = net.ResolveTCPAddr("tcp", IP + ":" + TCP_port)
	if connection.err != nil {
		return connection
	}
	connection.Conn, connection.err = net.DialTCP("tcp", nil, connection.ADDR)
	if connection.err != nil {
		return connection
	}
	return connection
}


func TCP_listen(key string, btn_queue *queue.Btn_queue_t, status status_t, elev_num uint8){
	var buf = make([]byte, 1024)
	for {

		read_len, err := connections[key].Conn.Read(buf)
		if err != nil {
			fmt.Println("TCP read error: ", connections[key].IP)
			connections[key].Conn.Close()
			delete(connections, key)
			return
		}
		var temp_buf = make([]byte,read_len)
		temp_buf = 	buf[:read_len]
		
		var msg map[string]string
		if err := json.Unmarshal(temp_buf, &msg); err != nil {
        	fmt.Println("TCP Unmarshal error: ", connections[key].IP, err)
connections[key].Conn.Close()  
delete(connections, key)      	
return
    }
    	if(msg["message type"] == "button press"){
    		queue.Push(btn_queue, msg["button id"], key)
    	}else{
    		if(msg["busy"] == "true"){
    			status.Busy[elev_num] = true
    		}else{
    			status.Busy[elev_num] = false
    		}
    		status.Destination[elev_num], _ = strconv.Atoi(msg["Destination"])
    		status.Up_or_dwn[elev_num], _ = strconv.Atoi(msg["Up_or_dwn"])
    		status.Door_open[elev_num], _ = strconv.Atoi(msg["Door_open"])
    	}
	}
}



type Command_t struct{
	sync.Mutex
	Command map[string]string
}



func TCP_master_routine(network_data *mainlib.Network_info, btn_queue *queue.Btn_queue_t){
	time.Sleep(2000*time.Millisecond)
	fmt.Println("staring tcp master routine")
	network_data.Lock()
	network_data.TCP_master_started = 1
	network_data.Unlock()
	connections = make(map[string]connection_t)
	var status_test status_t
	var i uint8 = 0
	for {
		for key, _ := range network_data.IPmap {
			if _, ok := connections[key]; ok {
			}else{
				connections[key] = TCP_connect(key)
				i = i + 1
					if connections[key].err != nil {
						fmt.Print("Connection error, ABORTABORT")
						return
					}
				go TCP_listen(key, btn_queue,status_test, i  )
			}
		}
		if(network_data.Master != 1){
			for key := range connections{
				connections[key].Conn.Close()	
			}
			network_data.Lock()
			network_data.TCP_master_started = 0
			network_data.Unlock()
			return
		}
		time.Sleep(100*time.Millisecond)
	}

}

/*
func TCP_master_routine(network_data *mainlib.Network_info, btn_queue *queue.Btn_queue_t){ //status *status_t, command command_t
	fmt.Println("staring tcp master routine")
	network_data.Lock()
	network_data.TCP_master_started = 1
	network_data.Unlock()
	var i uint8 = 0
	connections := make(map[string]connection_t)
	var status_test status_t
	for key, _ := range network_data.IPmap {
					connections[key] = *TCP_connect(key)
					if connections[key].err != nil {
						fmt.Print("Connection error, ABORTABORT")
						return
					}
					go TCP_listen(connections[key], btn_queue,status_test, key, i  )
					i += 1
			} 



	for {
		if(network_data.Master != 1){
			for key := range connections{
				connections[key].Conn.Close()	
			}
			network_data.Lock()
			network_data.TCP_master_started = 0
			network_data.Unlock()
			return
		}
		if(i < len(connections))
		{
			for key, _ := range network_data.IPmap {
				if val, ok := connections[key]; ok {
					
				}else {
					connections[key] = *TCP_connect(key)
					if connections[key].err != nil {
						fmt.Print("Connection error, ABORTABORT")
						return
					}
					go TCP_listen(connections[key], btn_queue,status_test, key, i  )
					i += 1
				}

			}
			
    			//do something here
}
		}
		time.Sleep(100*time.Millisecond)
	}
	
}
*/

type status_t struct {
	sync.Mutex
	Num_of_elevators uint8
	Master bool
	Busy []bool
	Destination []int
	Up_or_dwn []int
	Door_open []int
}

func TCP_accept_connection(myAddr string) {
	l, err := net.Listen("tcp", myAddr + ":" + TCP_port)
	if err != nil {
		fmt.Println(err)
	}

	slave_conn, err = l.Accept()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected")
}


func TCP_slave(myAddr string, command *Command_t){
	TCP_accept_connection(myAddr)
	go TCP_slave_listner(command)
}

func TCP_slave_send_msg(msg map[string]string, myAddr string) {
	defer func() {
        if r := recover(); r != nil {
	    slave_conn.Close()
            fmt.Println("Recovered in f", r)
	    TCP_accept_connection(myAddr)
	    //buf, _ := json.Marshal(msg)
	    //slave_conn.Write(buf)
        }
    	}()
	buf, _ := json.Marshal(msg)
	fmt.Println(buf)
	slave_conn.Write(buf)
}

func TCP_slave_listner(command *Command_t) {
	buf := make([]byte, 1024)
	for{
		read_len, err := slave_conn.Read(buf)
		fmt.Println(buf)
		var temp_buf = make([]byte,read_len)
		temp_buf = 	buf[:read_len]
		if err != nil {
			fmt.Println("TCP slave read error: ", err)
	}
		
		if err := json.Unmarshal(temp_buf, &command.Command); err != nil {
        	fmt.Println("TCP Unmarshal slave error: ", err)
    }else{
    fmt.Println(command.Command)
}
	}
}

