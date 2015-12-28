package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main10() {
	fmt.Println("--------------------")
	u := User{100, "fantasy", 20}
	Info(u)
}

func (u User) Hello() {
	fmt.Println("Hello World", u.Name)
}

func Info(in interface{}) {
	t := reflect.TypeOf(in)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("not struct type!")
		return
	}
	v := reflect.ValueOf(in)
	fmt.Println("Values:", v)
	//打印属性和值
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s : %v = %v \n", f.Name, f.Type, val)
	}
	//打印方法和签名
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s : %v\n", m.Name, m.Type)
	}
}
