package _reflect

import (
	"reflect"
	"fmt"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Say(msg string)(r string) {
	return u.Name + "对你说" + msg
}

/**
	反射
 */
func Info(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("当前对象的类型:", t)
}