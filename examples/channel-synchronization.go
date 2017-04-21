package main

import (
	"fmt"
	"time"
)
// don t under stand
func worker(done chan bool)  {
	fmt.Print("working...")
	time.Sleep(time.Second * 1)
	fmt.Println("done")

	done <- true // send message to channel


	close(done)

}

func main()  {
	done := make(chan bool,1)

	go worker(done)

	k:=<-done // get message from the channel

	println(k)
}