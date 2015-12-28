package main

import (
	"fmt"
)

func main4() {

	switch a := 1; {
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("None")
	}
	//跳转语句 goto  break  continue
Lable1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				fmt.Println("break lable")
				break Lable1
			}
		}
	}
	flag := true
Lable2:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			//continue Lable2
			goto Lable3
		}
	}
Lable3:
	if flag {
		flag = false
		goto Lable2
	}
	fmt.Println("main end")
}
