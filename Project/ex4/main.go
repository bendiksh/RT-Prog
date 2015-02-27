package main

import (

	"mainlib"
	"time"
)

const (
	sAddr = "129.241.187.152:34933"
	cAddr = "129.241.187.152:30009"
)

func main () {
	//a:=tcpcomm.Info_struct{123,100}
	//tcpcomm.TCP_struct_sender(cAddr, a)
	//go tcpcomm.TCP_struct_receiver(cAddr)
	go mainlib.IP_list("YOYO dette er min IP")
	for {
		time.Sleep(4 * time.Second)
	}
}


