package main

import (
	"fmt"
	"net"
	"log"
	"sync"
)

func sender(sAddr *net.UDPAddr, wg *sync.WaitGroup) {
	_, err := net.ResolveUDPAddr("udp", "129.241.187.255:20009")
	if err != nil {
		log.Fatalln(err)
	}
	buf2 := []byte("Dette er en melding")
	cConn, err := net.DialUDP("udp", nil, sAddr)
	//defer cConn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = cConn.Write(buf2)
	if err != nil {
		log.Fatalln(err)
	}
	err = cConn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	wg.Done()
	
}


func main() {
	sAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:20009")
	if err != nil {
		log.Fatalln(err)
	}
	
	sConn, err := net.ListenUDP("udp", sAddr)
	//defer sConn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("listening on ", sConn.LocalAddr().String())
	buf := make([]byte, 1024)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sender(sConn.LocalAddr().(*net.UDPAddr), &wg)
	_, _ = sConn.Read(buf)
	n, err := sConn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server: read:", string(buf[0:n]))
	sConn.Close()
	wg.Wait()
}
