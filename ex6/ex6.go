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
	sAddr, err := net.ResolveUDPAddr("udp", ":20009")
	if err != nil {
		log.Fatalln(err)
	}
	sConn, err := net.ListenUDP("udp", sAddr)
	sConn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	for {
		buf := make([]byte, 1024)
		_, err := sConn.Read(buf)
		if err != nil {
			fmt.Println("I'm the master")
			
			buf2 := []byte(strconv.Itoa(i))
			for {
				fmt.Printf("%d\n", i)
				i += 1
				cConn, err := net.DialUDP("udp", nil, sAddr)
				if err != nil {
					fmt.Println("Master sending error")	
					log.Fatalln(err)
				}
				_, err = cConn.Write(buf2)
				if err != nil {
					fmt.Println("Master sending error2")
					log.Fatalln(err)
				}
				err = cConn.Close()
				if err != nil {
					fmt.Println("Master sending error3")
					log.Fatalln(err)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}else {
			i = int(buf[0])	
			fmt.Println(i)
		}
	}
}
