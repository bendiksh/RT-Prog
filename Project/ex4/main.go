package main

import (
	"github.com/bendiksh/RT-Prog/Project/ex4/tcpcomm"
)

const (
	sAddr = "129.241.187.152:34933"
	cAddr = "129.241.187.152:30009"
)

func main () {
	a:=tcpcomm.Info_struct{123,100}
	tcpcomm.TCP_struct_sender(cAddr, a)
	//go tcpcomm.TCP_struct_receiver(cAddr)
}


