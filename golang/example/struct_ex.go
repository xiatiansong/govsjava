package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	person  *Person
	Contact struct {
		Phone          string
		Province, City string
	}
	Class string
}

func ValueTrans(person Person) {
	person.Age = 13
	fmt.Println("ValueTrans", person)
}

func PointerTrans(person *Person) {
	//操作指针直接使用属性，不需要 *person 这种方式
	person.Age = 59
	person.Name = "universe"
	fmt.Println("PointerTrans", person)
}

func main8() {
	fmt.Println("--------------------")
	p1 := Person{}
	p2 := Person{"xiatiansong", 28}
	p3 := &Person{Name: "xiatiansong", Age: 28}
	p1.Name = "fantasy"
	fmt.Println(p1, p2, p3)
	ValueTrans(p2)
	fmt.Println(p2)
	PointerTrans(p3)
	fmt.Println(p3)

	//匿名struct
	p4 := struct {
		Name string
		Age  int
	}{
		Name: "universe", Age: 29,
	}
	fmt.Println(p4)

	p5 := Student{Class: "三年二班", person: &Person{Name: "universe", Age: 29}}
	fmt.Println(p5)
	//p5.person = p3
	p5.Contact.Phone = "15815532362"
	p5.Contact.Province = "Guangdong"
	p5.Contact.City = "Shenzhen"
	fmt.Println(p5)
}
