package main
import (
"fmt"
"net"
"log"
"time"
)
const (
sAddr = "129.241.187.136:34933"
cAddr = "129.241.187.156:30009"
)
func main () {
l, err := net.Listen("tcp", cAddr)
if err != nil {
log.Fatalln(err)
}
defer l.Close()
acc, err := l.Accept()
if err != nil {
log.Fatalln(err)
}
for i:= 0; i< 25; i++ {
go handleRequest(acc)
time.Sleep(100 * time.Millisecond)
}
acc.Close()
}
func handleRequest(conn net.Conn) {
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
}
