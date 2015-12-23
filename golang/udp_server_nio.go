package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "8080", "port")

type MsgInfo struct {
	addr *net.UDPAddr
	msg  uint32
}

func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)

	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	runtime.GOMAXPROCS(2)
	msg := make(chan MsgInfo, 100)
	flag := make(chan bool)

	go func() {
		for {
			receiveMsg(conn, msg)
		}
	}()

	go func() {
		for {
			sendMsg(conn, msg)
		}
	}()

	<-flag
}

func receiveMsg(conn *net.UDPConn, msg chan MsgInfo) {
	data := make([]byte, 1024)
	_, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	daytime := time.Now().Unix()
	msg <- MsgInfo{remoteAddr, uint32(daytime)}
}

func sendMsg(conn *net.UDPConn, msg chan MsgInfo) {
	b := make([]byte, 4)
	msgInfo := <-msg
	binary.BigEndian.PutUint32(b, uint32(msgInfo.msg))
	conn.WriteToUDP(b, msgInfo.addr)
}
