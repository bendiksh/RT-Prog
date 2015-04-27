package netlib

import(
	"time"
	"udpcomm"
	"sync"
	"net"


)

type Network_info struct{
	sync.Mutex
	Master int
	TCP_master_started int
	IPmap map[string]int
	MyIP string
	MasterIP string
}

func externalIP() (string) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String()
		}
	}
	return ""
}

func Network_init() *Network_info{
	network_data := &Network_info{Master: 0, IPmap: make(map[string]int), 
									MyIP: externalIP(), TCP_master_started : 0}
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
			network_data.MasterIP = lowest_key(network_data.IPmap)
			if network_data.MyIP == network_data.MasterIP {
				if(network_data.Master != 2) {				
					network_data.Master = 1
				}

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

