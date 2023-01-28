package main

import (
	"dante-go/_algorithm"
	"dante-go/_db"
	"dante-go/_file"
	"dante-go/_json"
	"dante-go/_pointer"
	"dante-go/_reflect"
	"dante-go/_struct"
	"dante-go/_web"
	"dante-go/concurrencys"
	"dante-go/interfaces"
	"dante-go/restful"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// 结构体
	//structTest()

	// 指针
	//pointerTest()

	// 接口测试
	//interfaceTest()

	// 错误处理
	//errTest(9, 2)
	//panicTest()

	// 数组和切片
	//arrTest()
	//sliceTest()

	// map
	//mapTest()

	// Http 请求
	//restTest()

	// 反射测试
	//reflectTest()

	// 并发测试
	//concurrencyTest()

	// 数据库测试
	//mysqlTest()

	// web 测试
	//webTest()

	// json 测试
	//jsonTest()

	// 文件操作测试
	//fileTest()

	// 调用Shell测试
	//invokeShell()

	// 算法测试
	algorithmTest(55)

	// 语法测试
	//grammarTest()
}

/**
结构体测试
*/
func structTest() {
	author1 := _struct.Author{Name: "但丁", Age: 33, Description: "一个沉迷Code的猎魔人！"}
	address := _struct.Address{Country: "中国", Province: "北京市", Street: "方巷胡同13弄", Code: 1003}
	fmt.Println("你好，Go 的世界！", author1.SelfBook(address, "无字天书"))
}

/**
指针测试
*/
func pointerTest() {
	// 基础类型指针
	//_pointer.BasicPinter()

	// 结构体类型指针
	obj := _pointer.Obj{Name: "结构体"}
	sVal := obj.Set("我是一个结构体")
	objName := _pointer.GetObjName(obj)
	fmt.Println(sVal, " | ", obj, " | ", objName)

	// 引用传递和值传递
	//obj := _pointer.Obj{Name: "我有一个新名字"}
	//_pointer.PointerPass(&obj)
	//fmt.Println(obj)
	//obj2 := _pointer.Obj{Name: "我有一个新名字"}
	//_pointer.ValPass(obj2)
	//fmt.Println(obj2)
}

/**
接口测试
*/
func interfaceTest() {
	var phone interfaces.Phone    // 声明接口
	iphone := interfaces.IPhone{} // 接口实现类实例化
	phone = &iphone               // 接口实现类，接口指向引用，不是值
	phone.Set("IPhone 8", "美国")
	str := phone.Call("13932014370")
	fmt.Println(str)
	fmt.Println("========================================")
	var x interfaces.Phone = &interfaces.IPhone{}
	x.Set("11", "333")
	fmt.Println(x.Call("xxx"))
}

/**
错误接口，参考：https://blog.csdn.net/tennysonsky/article/details/78946265
	1. error接口，它是Go语言内建的接口类型
	2. panic
	3. recover
*/
func errTest(a, b float64) {
	result, err := divide(a, b)
	if err != nil {
		fmt.Println(result, err)
	} else {
		fmt.Println(result)
	}
}

func panicTest() {
	var p *int
	// 在可能出现 Panic 错误之前加入 recover()
	defer func() {
		err := recover()
		fmt.Println("Panic 错误捕获处理示例【", err, "】")
	}()
	fmt.Println(*p)
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
func arrTest() {
	var strArr [5]string = [5]string{"Docker", "k8s", "OpenShift", "Istio", "FaaS"}
	for i, str := range strArr {
		fmt.Println(i, str)
	}

	input := []int{1, 2, 3}
	//range循环中的x变量是临时变量。range循环只是将值拷贝到x变量中。因此内存地址都是一样的。创建指针数组的时候，不适合用`range`循环。
	var outRange []*int
	for _, v := range input {
		outRange = append(outRange, &v)
	}
	fmt.Println(*outRange[0], *outRange[1], *outRange[2])
	fmt.Println(outRange)

	// for 循环，内存地址是变化的。
	var outFor []*int
	for i := 0; i < len(input); i++ {
		outFor = append(outFor, &input[i])
	}
	fmt.Println(*outFor[0], *outFor[1], *outFor[2])
	fmt.Println(outFor)
}

func sliceTest() {
	var sliceList = make([]string, 0, 6)
	sliceList = append(sliceList, "竹", "尘", "水", "玉")

	i := 0
	for {
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
	maps := map[string]string{"Java": "爪哇", "Pythob": "蟒蛇🐍", "Gradle": "大象🐘"}
	for k, v := range maps {
		fmt.Println(k, v)
	}
	val, ok := maps["Java1"]
	if ok {
		fmt.Println("中文 ->", val)
	} else {
		fmt.Println("没有收录 Java")
	}

	fmt.Println("======================")
	maps2 := make(map[string]int)
	maps2["key1"] = 123
	maps2["key2"] = 456
	fmt.Println(maps2, "，长度：", len(maps2))
	delete(maps2, "key2")
	fmt.Println(maps2, "，长度：", len(maps2))
}

func restTest() {
	result, err := restful.DoGet("https://api.github.com/users/dante7qx")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func reflectTest() {
	usr := _reflect.User{1001, "但丁", 34}
	_reflect.Info(usr)

	manager := _reflect.Manager{User: _reflect.User{1002, "尤娜", 26}, Title: "CEO"}
	_reflect.Info2(&manager)
}

func concurrencyTest() {
	//concurrencys.Count(20)
	concurrencys.AsyncSay()
}

func mysqlTest() {
	_db.MysqlDB("localhost", "springboot", "root", "iamdante", 3306)
}

func webTest() {
	_web.StartWebServer(8899)
}

func jsonTest() {
	// json to object
	serverList := _json.JsonToObject()
	serverArr := serverList.Servers
	fmt.Println(serverArr)
	for _, server := range serverArr {
		fmt.Println("服务器: ", server.ServerIndex, server.ServerName, server.ServerIP)
	}

	// uncertain json to parse
	_json.ParseUncertainJson()
}

func fileTest() {
	_file.Mkdir("go", "aa/bb/xx")
}

/**
Go语言调用Shell与可执行文件
*/
func invokeShell() {
	command := "source /Users/dante/Desktop/go.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

/**
算法测试
*/
func algorithmTest(searchVal int) {
	var sourceData = []int{100, 32, 78, 97, 21, 102, 55, 76}
	fmt.Println(sourceData, searchVal)
	index := _algorithm.BinarySearch(sourceData, searchVal)
	fmt.Printf("%d 所在位置：%d", searchVal, index)
}

/**
语法测试
*/
func grammarTest() {
	const YYYYMMDDHHMISS = "2006-01-02 10:10:10"
	fmt.Println(time.Now().Format(YYYYMMDDHHMISS))

}

/**
init函数，先于main函数执行，实现包级别的一些初始化操作
init函数的主要作用：

初始化不能采用初始化表达式初始化的变量。
程序运行前的注册。
实现sync.Once功能。
其他
init函数的主要特点：

init函数先于main函数自动执行，不能被其他函数调用；
init函数没有输入参数、返回值；
每个包可以有多个init函数；
包的每个源文件也可以有多个init函数，这点比较特殊；
同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序。
不同包的init函数按照包导入的依赖关系决定执行顺序。

*/
/**
func init() {
	var initFuncInfo string = `
===========================================================================================
init函数，先于main函数执行，实现包级别的一些初始化操作
init函数的主要作用：
	初始化不能采用初始化表达式初始化的变量。
	程序运行前的注册。
	实现sync.Once功能。
init函数的主要特点：
	init函数先于main函数自动执行，不能被其他函数调用；
	init函数没有输入参数、返回值；
	每个包可以有多个init函数；
	包的每个源文件也可以有多个init函数，这点比较特殊；
	同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序。
	不同包的init函数按照包导入的依赖关系决定执行顺序。
===========================================================================================
`
	fmt.Printf("%s\n", initFuncInfo)
}
*/
