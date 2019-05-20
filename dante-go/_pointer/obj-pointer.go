package _pointer

type Obj struct {
	Name string
}

/**
	1. 指针的主要作用就是在函数内部改变传递进来变量的值，即引用传递

	2. 写在方法前的类型(o struct_type)，说明这个方法被struct_type类型绑定了，
声明该类型的对象即可调用该类型的方法

	3. 引申：在Java中，对象和集合做为参数就是引用传递，不过语法上不需要添加指针标记 *
 */
func (obj *Obj) Set(name string) string {
	obj.Name = name
	return "99999 - 00000"
}

/**
	obj 内部的 Name 不会修改，即值传递
 */
func GetObjName(obj Obj) string {
	obj.Name = "I am a struct"
	return obj.Name + "!"
}