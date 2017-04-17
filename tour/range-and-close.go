package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1  // 初始化 x,y 的值

	for i := 0; i < n; i++ { // 这里循环运行 10 次
		c <- x               // 将 x 的值传入 c
		x, y = y, x+y        // 将 y 赋值给 x， 将 x + y 赋值给 y
	}

	close(c)
}

func main() {
	c := make(chan int, 10) // 创建 10 个 channel ,类型为 int

	// 运行 goroutine, 调用 fibonacci 函数, 运行 fibonacci 函数，并将中间过程存入 channel
	go fibonacci(cap(c), c)


	for i := range c {
		fmt.Println(i)      // 打印 channel 中收到的数值
	}
}
