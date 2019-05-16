package _reflect

import (
	"reflect"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Say(msg string) (r string) {
	return u.Name + "对你说" + msg
}

type Manager struct {
	User
	Title string
}

/**
	反射
 */
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("传入对象必须是 struct 类型！")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("当前对象的类型:", t)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(field.Type, field.Name, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Type, m.Name, m.Func)
	}
	fmt.Println("-------------------------------")
}

func Info2(o interface{}) {
	cv := reflect.ValueOf(o)
	if cv.Kind() == reflect.Ptr && !cv.Elem().CanSet() {
		fmt.Println("传入对象无法被修改！")
		return
	} else {
		cv = cv.Elem()
	}
	if cf := cv.FieldByName("Title"); cf.Kind() == reflect.String {
		cf.SetString("公司首席执行官")
	}
	fmt.Println("Manager", cv)

	if cf := cv.FieldByName("User"); cf.Kind() == reflect.Struct {
		for i := 0; i < cf.NumField(); i++ {
			f := cf.Field(i)
			switch f.Kind() {
			case reflect.String:
				f.SetString("YouNa")
				break
			case reflect.Int:
				f.SetInt(2002)
				break
			}
		}

		callReturn := cf.MethodByName("Say").Call([]reflect.Value{reflect.ValueOf("Go反射机制")})
		fmt.Println(callReturn)

		cf.FieldByName("Name").SetString("最终幻想10")

	}
	fmt.Println("Manager", cv)
	fmt.Println("-------------------------------")
}
