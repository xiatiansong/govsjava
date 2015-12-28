package main

import (
	"fmt"
	"strconv" //字符串转换
)

type (
	//类型别名
	text string
	Byte int64
)

//常量组
const (
	Zeroday = 0
	Monday  = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	MAX_CON  = 100
	pMAX_CON = 200
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main3() {
	var n text = "中文"
	fmt.Println("a ...interface{}", n)
	var a, b, c, d = 1, 2, 3, 65
	fmt.Println(a, b, c, string(d))
	fmt.Println(strconv.Itoa(d))
	fmt.Println(Sunday)
	fmt.Println(MAX_CON)
	fmt.Println(pMAX_CON)
	fmt.Println(KB, ZB, YB)
	//GO中 ++ -- 是作为语句，而不是表达式，即不能放在等号右边
	a++
	fmt.Println(a)
}
