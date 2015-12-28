package main

import (
	"fmt"
)

func main5() {
	var a [2]int
	var b = [3]int{2: 1}
	c := [4]int{1, 2, 3, 4}
	d := [...]int{12, 23, 14, 45, 26, 37, 18, 39, 20}
	fmt.Println(a, b, c, d)
	//数组为值类型，要使用引用传递需使用slice
	var p *[4]int = &c
	fmt.Println(p)

	var e = [2][3]int{
		{1, 2, 3},
		{3, 2, 1}}
	fmt.Println(e)

	num := len(d)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if d[i] < d[j] {
				d[i], d[j] = d[j], d[i]
			}
		}
	}
	fmt.Println(d)
}
