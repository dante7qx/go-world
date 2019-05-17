package interfaces

/**
参考：https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/

1. 一个类型实现了一个 interface 中所有方法，该类型实现了该 interface
2. interface 变量存储的是实现者的值
*/
type Phone interface {
	Call(string) string
	Set(string, string)
}

type IPhone struct {
	Name string
	From string
}

func (iphone *IPhone) Set(name string, from string) {
	iphone.Name = name
	iphone.From = from
}

func (iphone IPhone) Call(phoneNumber string) string {
	return iphone.Name + "来自" + iphone.From + ", 拨叫 " + phoneNumber
}
