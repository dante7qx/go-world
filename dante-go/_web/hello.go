package _web

import (
	"net/http"
	"html/template"
	"log"
	"strings"
	"strconv"
	"fmt"
)


func StartWebServer(port int) {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/login", login)
	log.Println("开启服务器，监听端口：", port)
	err := http.ListenAndServe(strings.Join([]string{"localhost", strconv.Itoa(port)}, ":"), nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["age"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprint(w, "Hello, Go!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("/Users/dante/Documents/Technique/Go/project/dante-go/src/_db/login.gtpl")
		if err != nil {
			log.Fatal("Parse html error: ", err)
		} else {
			t.Execute(w, nil)
		}

	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		// 使用 r.FromValue("username")，不需要先调用 r.ParseForm()。因为FromValue() 会自动调用 ParseForm()
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}