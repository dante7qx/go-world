package concurrencys

import "fmt"

/**
Go 协程（Goroutine）是与其他函数同时运行的函数，轻量级的线程
	参考：https://go.dev/tour/concurrency/1

在函数调用前加上 go 关键字，这次调用就会在一个新的 goroutine 中并发执行。当被调用的函数返回时，这个 goroutine 也自动结束。
需要注意的是，如果这个函数有返回值，那么这个返回值会被丢弃。

Go 协程（Goroutine）之间通过信道（channel）进行通信，简单的说就是多个协程之间通信的管道。信道可以防止多个协程访问共享内存时发生资源争抢的问题。

Channel（管道）- https://www.runoob.com/w3cnote/go-channel-intro.html

Channel（管道） 可以被认为是协程之间通信的管道。与水流从管道的一端流向另一端一样，数据可以从信道的一端发送并在另一端接收。
每个 channel 都有一个类型。此类型是允许信道传输的数据类型。channel 是类型相关的，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。
需要通过内置函数 make 来创建一个信道。
	a := make(chan int)
	c := make(chan int, 1024) // 从带缓冲的channel中读数据

箭头的指向就是数据的流向
	ch <- v    // 发送值v到Channel ch中
	v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v

select

select语句选择一组可能的send操作和receive操作去处理。它类似switch,但是只是用来处理通讯(communication)操作。
它的case可以是send语句，也可以是receive语句，亦或者default。

receive语句可以将值赋值给一个或者两个变量。它必须是一个receive操作。

最多允许有一个default case,它可以放在case列表的任何位置，尽管我们大部分会将它放在最后。
*/

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i //写入
		fmt.Println("create :", i)
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue // 读出
		fmt.Println("receive:", v)
	}
}

func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
