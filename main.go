package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var sizeChannel int

func init() {
	if len(os.Args) < 2 {
		fmt.Println("please, add argument channel size")
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sizeChannel = num
}
func main() {
	channelIp := make(chan string, sizeChannel)
	var wg sync.WaitGroup
	go generateIp(channelIp)
	for i := 0; i < sizeChannel; i++ {
		wg.Add(1)
		go pinger(channelIp, &wg)
	}
	wg.Wait()
	fmt.Println("ok")
}

func generateIp(c chan<- string) {
	for i := 0; i <= 255; i++ {
		for j := 0; j <= 255; j++ {
			for k := 0; k <= 255; k++ {
				for l := 0; l <= 255; l++ {
					c <- fmt.Sprintf("%d.%d.%d.%d", i, j, k, l)
				}
			}
		}
	}
	close(c)
}
func pinger(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for addr := range c {
		pingRes(addr)
	}

}

func pingRes(addr string) {
	var arrOutput []string
	arrOutput = append(arrOutput, addr)
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", addr)
	if err != nil {
		arrOutput = append(arrOutput, err.Error())
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		arrOutput = append(arrOutput, fmt.Sprintf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt))
	}
	p.OnIdle = func() {
		arrOutput = append(arrOutput, "finish")
	}
	err = p.Run()
	if err != nil {
		arrOutput = append(arrOutput, err.Error())
	}
	arrOutput = append(arrOutput, "\n")
	for _, v := range arrOutput {
		fmt.Println(v)
	}
}
