package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type PortScan struct{
	timeOut       time.Duration
	protocol      string
	cores         int
	threads       int
	addr          string
	availablePort []string
	wasteTime int64
	rwLocker sync.RWMutex
	ipBegin string
	ipEnd string
	ipTable []string
}

func(ps *PortScan) scan() error {

	err:= ps.parse()
	if err != nil{
		return err
	}
	fmt.Println("开始扫描...")
	if len(ps.ipTable)==0 {
		if isAlive(ps.addr) {
		start := time.Now().Unix()
		channel := make(chan int, ps.threads)
		for port := 1; port < 65535; port++ {
			channel <- 1
			go ps.check(ps.addr, port, &channel)
		}
		end := time.Now().Unix()
		ps.wasteTime = end - start
		fmt.Printf("扫描结束,耗时%d秒,%s:%s共扫描%d个可用端口\r\n", ps.wasteTime, ps.protocol, ps.addr, len(ps.availablePort))
		return nil
	}else{
		fmt.Println(ps.addr,"当前ip无主机存活")
		return nil
		}
}else{
		start:=time.Now().Unix()
		channel:=make(chan int,ps.threads)
		for _,addr:=range ps.ipTable {
			if isAlive(addr) {
				for port := 1; port < 65535; port++ {
					channel <- 1
					go ps.check(addr, port, &channel)
				}
			}else{
				fmt.Println(addr,"当前ip无主机存活")
			}
		}
		end := time.Now().Unix()
		ps.wasteTime = end - start
		fmt.Printf("扫描结束,耗时%d秒,%s:%s到%s共扫描%d个可用端口\r\n", ps.wasteTime, ps.protocol, ps.ipBegin, ps.ipEnd, len(ps.availablePort))
		return nil
	}
}

func(ps *PortScan) check(ip string,port int,channel *chan int){
	//fmt.Println("123")
		ps.connect(ip, port)
		<-*channel
}

func(ps *PortScan) connect(ip string,port int){
	cn:= make(chan string,1)
	go func(){
		conn,err:=net.Dial("tcp",fmt.Sprintf("%s:%d",ip,port))
		if err == nil{
			conn.Close()
			cn<-ip+":"+strconv.Itoa(port)
		}else{
			cn<-"0"
		}
	}()
	select {
	case result:=<-cn:
		if result!="0"{
			ps.appendAvailable(result)
			ps.out(result)
		}
		case <- time.After(ps.timeOut):
	}
	return
}