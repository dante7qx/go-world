package main

import (
	"fmt"
	"./_struct"
	"./_pointer"
	"./interfaces"
)

func main() {
	// 结构体
	//structTest()

	// 指针
	//pointerTest()

	// 接口测试
	interfaceTest()
}

/**
	结构体测试
 */
func structTest() {
	author1 := _struct.Author{Name: "但丁", Age: 33, Description: "一个沉迷Code的猎魔人！"}
	address := _struct.Address{Country: "中国", Province: "北京市", Street: "方巷胡同13弄", Code: 1003}

	fmt.Println("你好，Go 的世界！", author1.Book(address,"无字天书"))
}

/**
	指针测试
 */
func pointerTest() {
	// 基础类型指针
	//_pointer.BasicPinter()

	// 结构体类型指针
	obj := _pointer.Obj{}
	obj.Set("我是一个结构体")
	//obj.Name = "I am a struct"
	objName := _pointer.GetObjName(obj)
	fmt.Println(objName, obj)
}

/**
	接口测试
 */
func interfaceTest()  {
	var phone interfaces.Phone	// 声明接口
	iphone := interfaces.IPhone{}	// 接口实现类实例化
	phone = &iphone	// 接口实现类，接口指向引用，不是值
	phone.Set("IPhone 8", "美国")
	str := phone.Call("13932014370")
	fmt.Println(str)
}