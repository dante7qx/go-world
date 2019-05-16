package concurrencys

import "fmt"

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
