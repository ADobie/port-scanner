package main

import "fmt"

func(ps *PortScan)out(ipAndPort string){
	fmt.Println(ipAndPort," is opening")
}

func(ps *PortScan) appendAvailable(ipAndPort string){
	ps.rwLocker.Lock()
	defer ps.rwLocker.Unlock()
	ps.availablePort = append(ps.availablePort,ipAndPort)
}
