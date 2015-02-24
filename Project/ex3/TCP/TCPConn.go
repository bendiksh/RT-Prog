package main

import (
	"fmt"
	"net"
	"log"
)

func main () {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "129.241.187.136:34933")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	buf := make([]byte, 1024)
	buf = []byte("heihei\000")
	_, err = conn.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
	
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println("Reply: ", string(reply))
	_, err = conn.Read(reply)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Reply: ", string(reply))
	
	conn.Close()
}
