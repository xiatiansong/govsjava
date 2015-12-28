package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main13() {
	fmt.Println("------------1---------------")
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!!!")
		c <- true
		//外界有for循环取channel，就必须关闭channel
		close(c)
	}()
	fmt.Println("-----------2----------------")
	//<-c

	//迭代channel，必须要关闭channel
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("----------3-----------------")

	//有缓存的channel
	c2 := make(chan bool, 1)
	go func() {
		fmt.Println("GO GO GO!!!")
		c2 <- true
	}()
	<-c2

	fmt.Println("-----------多核并行-----------")
	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
		  c3 := make(chan bool, 10)
			for i := 0; i < 10; i++ {
				go Go(c3, i)
			}
			for i := 0; i < 10; i++ {
				<-c3
			}
	*/

	//通过同步包来实现多个goroutine来完成打印内容
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go GoWaitGroup(&wg, i)
	}
	fmt.Println("-----------任务提交完，等待完成-----------")
	wg.Wait()
	fmt.Println("-----------任务已完成-----------")
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(a, index)
	c <- true
}

func GoWaitGroup(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(a, index)
	wg.Done()
}
