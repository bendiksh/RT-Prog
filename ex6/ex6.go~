package main

import (
	"fmt"
	"net"
	"log"
	"strconv"
	"time"
	//"encoding/binary"
)


func main() {
	i := 0
	listenAddr, err := net.ResolveUDPAddr("udp", "localhost:20011")
	sendingAddr, err := net.ResolveUDPAddr("udp", "localhost:20011")
	if err != nil {
		log.Fatalln(err)
	}
	listenConn, err := net.ListenUDP("udp", listenAddr)
	listenConn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := listenConn.Read(buf)
		listenConn.Close()
		if err != nil {
			fmt.Println("I'm the master")
			
			for {	
				i += 1
				fmt.Printf("%d\n", i)
				sendingConn, err := net.DialUDP("udp", nil, sendingAddr)
				if err != nil {
					fmt.Println("Master sending error")	
					//log.Fatalln(err)
				}
				buf2 := []byte(strconv.Itoa(i))
				_, err = sendingConn.Write(buf2)
				if err != nil {
					fmt.Println("Master sending error2")
					//log.Fatalln(err)
				}
				sendingConn.Close()
				time.Sleep(500 * time.Millisecond)
			}
		}else {
			s := string(buf[0:n])
			i, _ = strconv.Atoi(s)	
			//fmt.Printf("%d\n", i)
			listenConn, err = net.ListenUDP("udp", listenAddr)
			listenConn.SetReadDeadline(time.Now().Add(2 * time.Second))
		}
	}
}
