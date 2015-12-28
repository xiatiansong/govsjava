package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) Increment(num int) {
	*tz += TZ(num)
}

func main1() {
	var a TZ
	a.Increment(100)
	fmt.Println(a)
}
