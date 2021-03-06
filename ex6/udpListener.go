package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func main() {
	sAddr, err := net.ResolveUDPAddr("udp", "localhost:20011")
	if err != nil {
		log.Fatalln(err)
	}
	sConn, err := net.ListenUDP("udp", sAddr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("listening on ", sConn.LocalAddr().String())
	buf := make([]byte, 1024)
	sConn.SetReadDeadline(time.Now().Add(50 * time.Second))
	for {
	n, err := sConn.Read(buf)
	if err != nil {
		sConn.Close()
		log.Fatalln(err)
	}
	fmt.Println("server: ", string(buf[0:n]))
	}
	sConn.Close()
}



