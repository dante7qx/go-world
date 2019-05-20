package _pointer

import "fmt"

func BasicPinter() {

	var num int8 = 100
	fmt.Println("num 值", num, "内存地址", &num)

	// 指针类型，变量指向值的内存地址，访问值使用 *var
	var numPointer *int8
	numPointer = &num
	fmt.Println("numPointer 内存地址", numPointer, "值", *numPointer)
}