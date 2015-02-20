package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func main() {
	sAddr, err := net.ResolveUDPAddr("udp", ":20009")
	if err != nil {
		log.Fatalln(err)
	}
	sConn, err := net.ListenUDP("udp", sAddr)
	sConn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("listening on ", sConn.LocalAddr().String())
	buf := make([]byte, 1024)
	n, err := sConn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server: ", string(buf[0:n]))
	sConn.Close()
}
