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

func (u *User) Hello(name string) {
	fmt.Println("Hello", name, " my name is ", u.Name)
}

func (u User) HelloNo(name string) {
	fmt.Println("Hello", name, " my name is ", u.Name)
}

/**
 * 修改属性值
 * @param {[type]} in interface{}
 */
func Set(in interface{}) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("can not set value")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("not found this filed")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("new name")
	}
}

type Manager struct {
	User
	title string
}

func main11() {
	fmt.Println("--------------------------", reflect.Array)
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n%#v\n", t.Field(0), t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0}))

	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(100)
	fmt.Println(x)

	u := User{1, "chenjianghao", 27}
	Set(&u)
	fmt.Println(u)

	u.Hello("fantasy")

	av := reflect.ValueOf(u)
	mv := av.MethodByName("HelloNo")
	args := []reflect.Value{reflect.ValueOf("Joe")}
	fmt.Println(args)
	mv.Call(args)
}
