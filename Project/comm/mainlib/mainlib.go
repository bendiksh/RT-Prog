package mainlib

import(
	"time"
	"udpcomm"
	"sync"

)

type Network_info struct{
	sync.Mutex
	Master int
	TCP_master_started int
	IPmap map[string]int
	MyIP string
}

func Network_init(IP string) *Network_info{
	network_data := &Network_info{Master: 0, IPmap: make(map[string]int), MyIP: IP, TCP_master_started : 0}
	go Am_I_master(network_data)
	return network_data
}

func Am_I_master(network_data *Network_info){
	var ikkeTid int = 0
	last_cleanup := time.Now()
	channel := make(chan string, 40)
	go udpcomm.UDP_routine(network_data.MyIP, channel)
	for {
		newIP := <-channel
		network_data.IPmap[newIP] = ikkeTid
		if (time.Since(last_cleanup)) > (1* time.Second){
			last_cleanup = time.Now()
			for key, value := range network_data.IPmap {
				if value != ikkeTid{
					delete(network_data.IPmap, key)
				}
			}
			ikkeTid = ikkeTid + 1
			network_data.Lock()
			if network_data.MyIP == lowest_key(network_data.IPmap) {
				network_data.Master = 1

			}else{
				network_data.Master = 0
			}
			network_data.Unlock()
		}
	}
}

func lowest_key(IPs map[string]int) string{
	lowest_IP := "256.256.256.256"
	for key, _ := range IPs {
				if key < lowest_IP {
					lowest_IP = key
				}
			}
	return lowest_IP
}
