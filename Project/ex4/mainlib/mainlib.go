package mainlib

import(
	"fmt"
	"time"
	"udpcomm"

)


func IP_list(myIP string){
	var ikkeTid int = 0
	last_cleanup := time.Now()
	IPmap := make(map[string]int)
	channel := make(chan string, 40)
	go udpcomm.UDP_routine(myIP, channel)
	for {
		newIP := <-channel
		IPmap[newIP] = ikkeTid
		if (time.Since(last_cleanup)) > (2* time.Second){
			last_cleanup = time.Now()
			for key, value := range IPmap {
				if value != ikkeTid{
					delete(IPmap, key)
				}
			}
			ikkeTid = ikkeTid + 1
			fmt.Println(IPmap)
		}
	}
}
