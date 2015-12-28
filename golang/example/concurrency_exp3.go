package main

import (
	"fmt"
	"time"
)

func main16() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for {
		select {
		case c <- 0:
			fmt.Println("tranform value 0 to c")
		case c <- 1:
			fmt.Println("tranform value 1 to c")
		}
	}

	//空的select可以阻塞main函数，可用于GUI
	//select {}
	//select设置超时

	b := make(chan bool)
	select {
	case v := <-b:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

}
