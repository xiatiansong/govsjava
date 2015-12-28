package main

import (
	"fmt"
)

func main14() {
	fmt.Println("---------------------------")
	c1, c2 := make(chan int), make(chan string)
	//
	c := make(chan bool, 2)

	go func() {
		//设计select中多个case都关闭会比较麻烦，一般是直接第一个关闭时就退出select
		//a, b := false, false
		for {
			/*
				select {
				case v, ok := <-c1:
					//被关闭
					if !ok {
						if !a {
							c <- true
							a = true
						}
						break
					}
					fmt.Println("--------c1-------", v)
				case v, ok := <-c2:
					//被关闭
					if !ok {
						if !b {
							c <- true
							b = true
						}
						break
					}
					fmt.Println("---------c2-------", v)
				}
			*/

			select {
			case v, ok := <-c1:
				//被关闭
				if !ok {
					c <- true
					break
				}
				fmt.Println("--------c1-------", v)
			case v, ok := <-c2:
				//被关闭
				if !ok {
					c <- true
					break
				}
				fmt.Println("---------c2-------", v)
			}
		}
	}()

	c1 <- 2
	c2 <- "Hello"
	c1 <- 3
	c2 <- "World"
	c1 <- 4
	c2 <- "Hello~~"
	c1 <- 5
	c2 <- "EveryOne"
	close(c1)
	//close(c2)

	for i := 0; i < 2; i++ {
		<-c
	}
}
