package main

import(
	"net"
	"fmt"
	"encoding/json"
	"time"
)



func main(){
	l, err := net.Listen("tcp", "localhost:30000")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
	}
	var msg = make(map[string]string)
		msg["btn1"] = "pressed"
		msgB, _ := json.Marshal(msg)

	for{
		_, err = conn.Write(msgB)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(500*time.Millisecond)
	}
	
}