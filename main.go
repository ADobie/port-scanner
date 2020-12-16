package main

import "fmt"

func main(){
	var portScanner *PortScan
	portScanner=new(PortScan)
	if err:=portScanner.scan();err!=nil{
		fmt.Println(err)
	}
}
