package main

import (
	"fmt"
)

func main7() {
	origin := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(origin)
	s1 := origin[4:9] //左闭右开
	//s1[0] = 100
	fmt.Println(s1, origin)
	//make
	s2 := make([]int, 3)
	fmt.Println(s2)
	copy(s2, s1)
	fmt.Println(s2)
}
