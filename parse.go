package main

import (
	"flag"
	"net"
	"os"
	"runtime"
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
	ipAddr,err := net.ResolveIPAddr("ip",ip)
	if err != nil {
		return err
	}
	ps.addr = ipAddr.String()
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