package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"time"
)

func main() {
	for i := 0; i <= 255; i++ {
		for j := 0; j <= 255; j++ {
			for k := 0; k <= 255; k++ {
				for l := 0; l <= 255; l++ {
					ip := fmt.Sprintf("%d.%d.%d.%d", i, j, k, l)
					pinger(ip)
				}
			}
		}
	}
}

func pinger(addr string) {
	fmt.Println(addr)
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", addr)
	if err != nil {
		fmt.Println(err)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
