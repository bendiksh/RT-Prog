package tcpcomm

import (
	"net"
	"log"
	"fmt"
)

type Info_struct struct {
	Status uint8;
	Etasje uint8;
}

func TCP_send(Addr string,  buf []byte){
	tcpAddr, err := net.ResolveTCPAddr("tcp", Addr)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
			
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Reply: ", string(buf[0:1024]))
}

func TCP_struct_sender(Addr string, data Info_struct){
	buf := make([]byte, 1024)
	buf[0] = data.Status
	buf[1] = data.Etasje
	TCP_send(Addr, buf);
}

func TCP_struct_receiver(Addr string ){
	for i := 1;i < 2; i=1 {
		fmt.Println("AHH")
	} 
}

