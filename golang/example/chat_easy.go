package main

import (
	"fmt"
)

var c chan string

func PingPong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From PingPong: Hello Main %d", i)
		i++
	}
}

func main() {
	c = make(chan string)
	go PingPong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From Main: Hello PingPong %d", i)
		fmt.Println(<-c)
	}

}
