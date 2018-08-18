package _pointer

type Obj struct {
	Name string
}

/**
	指针的主要作用就是在函数内部改变传递进来变量的值
 */
func (obj *Obj) Set(name string) {
	obj.Name = name
}

func GetObjName(obj Obj) string {
	// obj 内部的 Name 不会修改
	obj.Name = "I am a struct"
	return obj.Name + "!"
}