package main

import "fmt"
/*

select 语句允许一个 goroutine 等待多个通信操作。
select 阻塞直到某个条件能够执行, 然后执行那个条件下的语句. 如果同时有多个条件具备， 他会随机的选择一个执行。

*/

func fibonacciSelect(c, quit chan int) {    // 传入了 两个 channel
	x, y := 0, 1
	for {
		select {
		case c <- x:                  // 将 x 写入 c channel， 这个刚开始就可以执行
			x, y = y, x+y
		case <-quit:                  // 这个条件刚开始不满足， 直到quit chan 有通信 才执行
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)    // 创建一个 channel
	quit := make(chan int) // 创建一个 channel

	// a new go process
	go func() {                        // 执行 闭包
		for i := 0; i < 10; i++ {      // 循环 10 次
			fmt.Println(<-c)           // 从 channel c 中接收值，打印
		}
		quit <- 0                      // 将 0 传递给 quit channel
	}()

	fibonacciSelect(c, quit)                 // 调用 fibnacci function，传入 c,quit
}
