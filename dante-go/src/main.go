package main

import (
	"fmt"
	"errors"
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
	//interfaceTest()

	// 错误处理
	//errTest(9, 3)

	// 数组和切片
	//arrTest()
	//sliceTest()

	// map
	mapTest()

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

/**
	错误接口，参考：https://blog.csdn.net/tennysonsky/article/details/78946265
		1. error接口，它是Go语言内建的接口类型
		2. panic
		3. recover
 */
func errTest(a, b float64)  {
	result, err := divide(a, b)
	if err != nil {
		fmt.Println(result, err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		result = 0.0
		err = errors.New("运行时错误：除数不能是0.")
		return
	}
	result = a / b
	err = nil
	return
}

/**
	数组和切片Slice(无界数组)
 */
func arrTest()  {
	var strArr [3]string = [3]string{"Docker","k8s","OpenShift"}
	for i, str := range strArr {
		fmt.Println(i, str)
	}
}

func sliceTest()  {
	var sliceList = make([]string, 0, 6)
	sliceList = append(sliceList, "竹", "尘", "水", "玉")

	i := 0
	for ; ;  {
		if i >= len(sliceList) {
			break
		}
		fmt.Println(i, sliceList[i])
		i++
	}
	fmt.Println("sliceList len", len(sliceList), "cap", cap(sliceList))

	var sublist = sliceList[2:]
	fmt.Println(sublist)

	var copylist []string = make([]string, 4, 6)
	copy(copylist, sliceList)
	fmt.Println(copylist)
}

/**
	Map
		1. var map_variable map[key_data_type]value_data_type
		2. map_variable := make(map[key_data_type]value_data_type)
 */
func mapTest() {
	maps := map[string]string {"Java": "爪哇", "Pythob": "蟒蛇🐍", "Gradle": "大象🐘"}
	for k, v := range maps {
		fmt.Println(k, v)
	}

	val, ok := maps ["Java"]
	if ok {
		fmt.Println("中文 ->", val)
	} else {
		fmt.Println("没有收录 Java")
	}
}