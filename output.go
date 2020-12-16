package main

import "fmt"

func(ps *PortScan)out(port int){
	fmt.Println(port)
}

func(ps *PortScan) appendAvailable(port int){
	ps.rwLocker.Lock()
	defer ps.rwLocker.Unlock()
	ps.availablePort = append(ps.availablePort,port)
}
