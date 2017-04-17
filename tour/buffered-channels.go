package main

import "fmt"
/*
Channels can be buffered. 提供一个 缓存长度，作为初始化的第二个参数，用于缓存 channel

ch := make(chan int, 100)

发送至 buffered channel 阻塞直到buffer满， 接收阻塞直到 buffer空

*/

func main() {
	ch := make(chan int, 2)
	ch <- 1             // one buffered channel is full
	ch <- 2             // another buffered channel is full , no other channel can be used
	fmt.Println(<-ch)   // first buffered channel is empty
	ch <- 3
	fmt.Println(<-ch)   // another buffered channel is empty
	fmt.Println(<-ch)

	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}