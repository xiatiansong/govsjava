package main

import "fmt"

type Connecter interface {
	Connect()
}

//内嵌接口
type USB interface {
	Name() string
	Connecter
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("pc connecter : ", pc.name)
}

type TVConnecter struct {
	name string
}

func (pc TVConnecter) Connect() {
	fmt.Println("pc connecter : ", pc.name)
}

func DisconnectUsb(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnect : ", pc.name)
		return
	}
	fmt.Println("Unknown Device")
}

func Disconnect(usb interface{}) {
	//空接口参数使用switch进行类型判断
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("PhoneDisconnect : ", v.name)
	case TVConnecter:
		fmt.Println("TVDisconnect : ", v.name)
	default:
		fmt.Println("Unknown Device")
	}
	//一般的ok pattern模式判断
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnect1 : ", pc.name)
		return
	}

	fmt.Println("Unknown Device1")
}

func main2() {
	var u = PhoneConnecter{"PhoneConnecter"}
	u.Connect()
	Disconnect(u)
	DisconnectUsb(u)
	//类型转换
	var a Connecter = Connecter(u)
	a.Connect()
	tv := TVConnecter{"TVConnecter"}
	tv.Connect()
	fmt.Println("----------------------------------------")
	Disconnect(tv)

	var a interface{}
	fmt.Println(a == nil)
}
