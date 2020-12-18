package main

import (
	"github.com/go-ping/ping"
	"time"
)

func isAlive(target string) bool  {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Timeout = 50*time.Millisecond
	pinger.SetPrivileged(true)
	pinger.Run()// blocks until finished
	stats := pinger.Statistics()

	//fmt.Println(stats)
	// 有回包，就是说明IP是可用的
	if stats.PacketsRecv >= 1 {
		return true
	}
	return false
}
