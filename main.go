package main

import (
	"fmt"
	"time"
)

func Ping(ch chan string) {
	for {
		ch <- "ping"
	}
}

func Pong(ch chan string) {
	for {
		msg := <-ch
		time.Sleep(time.Second * 1)
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
		fmt.Println("pong")
	}
}

func main() {
	ch := make(chan string)
	go Ping(ch)
	go Pong(ch)

	select {}
}
