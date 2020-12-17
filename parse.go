package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//var protocol = flag.String("p","tcp","the protocol u want")
//var cores = flag.Int("c",1,"the number of cores u want")
//var threads = flag.Int("t",0,"the number of threads u want")
var  (
	ip string
	protocol string
	threads int
	cores int
	timeOut int
)
//var timeOut = flag.Int("d",5,"the duration u want")

func(ps *PortScan) parse() error {
	flag.StringVar(&protocol,"p","tcp","the protocol u want")
	flag.IntVar(&cores,"c",1,"the number of cores u want")
	flag.IntVar(&timeOut,"d",500,"the duration u want")
	flag.IntVar(&threads,"t",1,"the number of threads u want")
	flag.StringVar(&ip,"h","","the ip u want to scan")

	flag.Parse()

	if len(flag.Args()) > 1 {
		flag.Usage()
		os.Exit(1)
	}
	ps.protocol = protocol
	ps.cores = cores
	ps.threads = threads
	ps.timeOut = time.Duration(timeOut)*time.Millisecond
	if ip != "" {
		ipSection := strings.Split(ip, "-")
		if len(ipSection) > 1 {
			ps.ipBegin = ipSection[0]
			ps.ipEnd  = ipSection[1]
			ps.ipTable = ps.getIpList(ps.ipBegin,ps.ipEnd)
		}else{
			ipAddr ,err:= net.ResolveIPAddr("ip",ip)
			if err!= nil {
				return err
			}else{
				ps.addr = ipAddr.String()
			}
		}
	}else{
		flag.Usage()
		os.Exit(0)
	}
	if ps.threads>0{
		if ps.cores>0{
			if ps.cores<runtime.NumCPU(){
				runtime.GOMAXPROCS(ps.cores)
			}else{
				runtime.GOMAXPROCS(runtime.NumCPU())
			}
		}else{
			runtime.GOMAXPROCS(1)
		}
	}else{
		ps.threads = 1
	}
	return nil
}

func(ps *PortScan) getIpList(minIp, maxIp string) []string {
	ipArr := make([]string, 0)
	minIpaddress := net.ParseIP(minIp)
	maxIpaddress := net.ParseIP(maxIp)
	if minIpaddress == nil || maxIpaddress == nil {
		fmt.Println("ip地址格式不正确")
	} else {
		minIpSplitArr := strings.Split(minIp, ".")
		maxIpSplitArr := strings.Split(maxIp, ".")

		minIP1, _ := strconv.Atoi(minIpSplitArr[0])
		minIP2, _ := strconv.Atoi(minIpSplitArr[1])
		minIP3, _ := strconv.Atoi(minIpSplitArr[2])
		minIP4, _ := strconv.Atoi(minIpSplitArr[3])

		maxIP1, _ := strconv.Atoi(maxIpSplitArr[0])
		maxIP2, _ := strconv.Atoi(maxIpSplitArr[1])
		maxIP3, _ := strconv.Atoi(maxIpSplitArr[2])
		maxIP4, _ := strconv.Atoi(maxIpSplitArr[3])

		if minIP1 <= maxIP1 {
			for i1 := minIP1; i1 <= maxIP1; i1++ {
				minIP1 = i1
				var i2 int
				var maxi2 int
				if minIP1 == maxIP1 { //如果第一个数相等
					i2 = minIP2
					maxi2 = maxIP2
				} else {
					i2 = 0
					maxi2 = 255
				}
				for ii2 := i2; ii2 <= maxi2; ii2++ {
					minIP2 = ii2
					var i3 int
					var maxi3 int
					if minIP1 == maxIP1 && minIP2 == maxIP2 { //如果第一个数相等 并且 第二个数相等
						i3 = minIP3
						maxi3 = maxIP3
					} else {
						i3 = 0
						maxi3 = 255
					}
					for ii3 := i3; ii3 <= maxi3; ii3++ {
						minIP3 = ii3
						var i4 int
						var maxi4 int
						if minIP1 == maxIP1 && minIP2 == maxIP2 && minIP3 == maxIP3 { //如果第一个数相等 并且 第二个数相等 并且 第三个数相等
							i4 = minIP4
							maxi4 = maxIP4
						} else {
							i4 = minIP4
							maxi4 = 255
						}
						for ii4 := i4; ii4 <= maxi4; ii4++ {
							minIP4 = ii4
							newIP := fmt.Sprintf("%v.%v.%v.%v", minIP1, minIP2, minIP3, minIP4)
							ipArr = append(ipArr, newIP)

						}
						minIP4 = 0

					}
					minIP3 = 0

				}
				minIP2 = 0

			}
		}
	}
	return ipArr
}