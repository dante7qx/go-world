package _pointer


func PointerPass(obj *Obj)  {
	obj.Name = "我是修改后的名字"
}

func ValPass(obj Obj) {
	obj.Name = "我是修改后的名字"
}


