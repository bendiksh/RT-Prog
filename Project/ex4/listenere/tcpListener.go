package main
import (
	"fmt"
	"net"
	"log"
	"time"
)
const (
	sAddr = "129.241.187.152:34933"
	cAddr = "129.241.187.152:30009"
)
func main () {
	l, err := net.Listen("tcp", cAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for i:= 0; i< 25; i++ {
		handleRequest(l)
		time.Sleep(100 * time.Millisecond)
	}
}
func handleRequest(l net.Listener) {
	conn, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	buf:= make([]byte, 1024)
	n2, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Reply: ", buf[0:n2])
	conn.Write([]byte("Message received.\000"))
	if err != nil {
		log.Fatalln(err)
	}
	conn.Close()
}
