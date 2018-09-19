package restful

import (
	"net/http"
		"log"
		"io/ioutil"
)

func DoGet(url string)  (r string, e error) {
	resp, err := http.Get(url)
	result := ""
	if resp != nil {
		defer resp.Body.Close()	// go的特殊语法，main函数执行结束前会执行resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				log.Fatal(e)
			}
			result = string(body)
		}
	} else {
		log.Fatal(err)
	}
	return result, err
}
