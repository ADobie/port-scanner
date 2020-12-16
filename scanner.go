package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type PortScan struct{
	timeOut       time.Duration
	protocol      string
	cores         int
	threads       int
	addr          string
	availablePort []int
	wasteTime int64
	rwLocker sync.RWMutex
}

func(ps *PortScan) scan() error {
	start:=time.Now().Unix()
	err:= ps.parse()
	if err != nil{
		return err
	}
	fmt.Println("开始扫描...")
	channel:=make(chan int,ps.threads)
	ports:=65535
	for port:=1;port<ports;port++ {
		channel <- 1
		go ps.check(port, &channel)
	}
		end := time.Now().Unix()
		ps.wasteTime = end - start
		fmt.Printf("扫描结束,耗时%d秒,%s:%s共扫描%d个可用端口\r\n", ps.wasteTime, ps.protocol, ps.addr, len(ps.availablePort))
		return nil
}

func(ps *PortScan) check(port int,channel *chan int){
	//fmt.Println("123")

	ps.connect(port)
	<- *channel
}

func(ps *PortScan) connect(port int){
	cn:= make(chan int,1)
	go func(){
		conn,err:=net.Dial("tcp",fmt.Sprintf("%s:%d",ps.addr,port))
		if err == nil{
			conn.Close()
			cn<-port
		}else{
			cn<-0
		}
	}()
	select {
	case result:=<-cn:
		if result>0{
			ps.appendAvailable(result)
			ps.out(result)
		}
		case <- time.After(ps.timeOut):
	}
	return
}