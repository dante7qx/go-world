package _json

import (
	"encoding/json"
	"log"
	"fmt"
)

const  (
	JSON_STR = `{"servers":[{"serverIndex":1001,"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverIndex":1002,"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	UNCERTAIN_STR = `{"Name":"竹尘","Age":6,"Parents":["雪峰","欣敏"]}`
)

type Server struct {
	ServerIndex int
	ServerName string
	ServerIP string
}

type ServerList struct {
	Servers []Server
}

func JsonToObject() ServerList {
	var servers ServerList
	err := json.Unmarshal([]byte(JSON_STR), &servers)
	if err != nil {
		log.Fatal("Json to object: ", err)
	}
	return servers
}

func ParseUncertainJson() {
	b := []byte(UNCERTAIN_STR)
	var f interface{}

	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal("Uncertain json to object: ", err)
	}

	// f里面存储了一个map类型，他们的key是string，值存储在空的interface{}里
	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}

