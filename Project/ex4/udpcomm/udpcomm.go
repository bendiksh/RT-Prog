package udpcomm

import  (
	"fmt"
	"net"
	"log"
	"strings"
	"time"
	"strconv"
)

func UDP_routine(IPAddr string, chann chan string){
	go UDP_listener(chann)
	for {
		UDP_IP_sender(IPAddr)
		time.Sleep(100*time.Millisecond)
	}

}

func UDP_listener(chann chan string){
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
	buf := make([]byte, 4)
	for {
		_, _ = sConn.Read(buf)
		IP := []string{strconv.Itoa(int(buf[0])),strconv.Itoa(int(buf[1])),strconv.Itoa(int(buf[2])),strconv.Itoa(int(buf[3]))}
		chann <- strings.Join(IP, ".")
	}
}


func UDP_IP_sender(IPAddr string){
	cAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:20009")
	if err != nil {
		log.Fatalln(err)
	}
	IPArray := strings.Split(IPAddr, ".")
	buf2 := make([]byte, 4)
	for i := 0;i<4;i++{
		integer, _ := strconv.Atoi(IPArray[i])
		buf2[i] = byte(integer)
	}
	cConn, err := net.DialUDP("udp", nil, cAddr)
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
}	




