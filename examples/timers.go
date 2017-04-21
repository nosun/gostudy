package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(time.Second * 1)
	sig    := make(chan bool)

	<-timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second * 2)

	go func() { // wait at
		<-timer2.C
		fmt.Println("timer 2 expired")
		sig <- true
	}()

	time.Sleep(time.Second * 3)

	stop2 := timer2.Stop() // stop at main process,it stopped

	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	close(sig)
	<-sig
}