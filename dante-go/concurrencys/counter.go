package concurrencys

import (
	"fmt"
	"runtime"
	"time"
)

func counting(ch chan int) {
	ch <- 1
	fmt.Println("计数中...")
}

func Count(n int) {
	chs := make([]chan int, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		go counting(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
}

func AsyncSay() {
	go say("world") //开一个新的Goroutines执行
	say("hello")
}


func say(s string) {
	for i := 0; i < 5; i++  {
		runtime.Gosched()
		time.Sleep(time.Duration(3))
		fmt.Println(s)
	}
}
