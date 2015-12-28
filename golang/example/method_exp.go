package main

import (
	"fmt"
)

type Atype struct {
	Name string
}

type Btype struct {
	Name string
	//私有字段其实是package级别的，就是包级别可访问
	password string
}

func main9() {
	fmt.Println("--------------------------------")
	a := Atype{"234"}
	a.Print()
	fmt.Println(a)

	b := Btype{"234", "123456"}
	b.Print()
	fmt.Println(b)

	//method value 使用类型来调用方法，第一个参数是receiver
	(*Btype).Print(&b)
	(*Btype).transform(&b, 10)

	fmt.Println(b.password)
}

//值传递
func (a Atype) Print() {
	a.Name = "fantasy"
	fmt.Println("inner A Print", a.Name)
}

//指针传递 引用传递
func (b *Btype) Print() {
	b.Name = "fantasy"
	fmt.Println("inner B Print", b.Name)
}

func (b *Btype) transform(age int) {
	b.Name = "fantasy"
	b.password = "654321"
	fmt.Println("inner B Print", b.Name, age)
}
