package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "8080", "port")

//go run timeclient.go -host time.nist.gov
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	timestamp1 := time.Now().Unix()
	fmt.Println(timestamp1)
	//循环发送50万次请求
	for i := 0; i < 500000; i++ {
		_, err = conn.Write([]byte("QOTM?"))
		if err != nil {
			fmt.Println("failed:", err)
			os.Exit(1)
		}
		data := make([]byte, 4)
		_, err = conn.Read(data)
		if err != nil {
			fmt.Println("failed to read UDP msg because of ", err)
			os.Exit(1)
		}
		_ = binary.BigEndian.Uint32(data)
		//fmt.Println(time.Unix(int64(t), 0).String())
	}
	timestamp2 := time.Now().Unix()
	fmt.Println(timestamp2 - timestamp1)
	fmt.Println(timestamp2)

	os.Exit(0)
}
