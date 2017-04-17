package main

import (
	"fmt"
	"time"
)

/*
   如果其他的条件不满足， select 会执行 default 代码块。
   尝试 default 条件， 尝试 发送或者接收数据，而不阻塞。
*/

func main() {
	tick := time.Tick(100 * time.Millisecond)   // 定时执行，循环
	boom := time.After(500 * time.Millisecond)  // 定时执行，执行一次
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return                               // 跳出
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)    // 等待 50 ms ， 重新跳至 select
		}
	}
}
