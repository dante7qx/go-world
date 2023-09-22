package _web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
*
1. 服务绑定 localhost, 只有容器内可以使用
2. 服务绑定 0.0.0.0, 容器外可以使用

参考: https://www.jianshu.com/p/3e94a1277c11
*/
const HOST = "0.0.0.0"

func StartWebServer(port int) {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/login", login)
	log.Println("开启服务器，监听端口：", port)
	err := http.ListenAndServe(strings.Join([]string{HOST, strconv.Itoa(port)}, ":"), nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	} //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["age"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	_, err2 := fmt.Fprint(w, "Hello, Go!")
	if err2 != nil {
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("/Users/dante/Documents/Project/go-world/src/dante-go/_db/login.gtpl")
		if err != nil {
			log.Fatal("Parse html error: ", err)
		} else {
			err := t.Execute(w, nil)
			if err != nil {
				return
			}
		}

	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		// 使用 r.FromValue("username")，不需要先调用 r.ParseForm()。因为FromValue() 会自动调用 ParseForm()
		err := r.ParseForm()
		if err != nil {
			return
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		template.HTMLEscape(w, []byte(r.Form.Get("username")))                      //输出到客户端
	}
}
