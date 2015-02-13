package main

import (
	"tcpcomm"
)

const (
	sAddr = "129.241.187.136:34933"
	cAddr = "129.241.187.156:30009"
)

func main () {
	//a:=tcpcomm.Info_struct{123,100}
	//tcpcomm.TCP_struct_sender(cAddr, a)
	go tcpcomm.TCP_struct_receiver(cAddr)
}


