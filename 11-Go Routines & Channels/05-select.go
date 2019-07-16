package main

import (
	"fmt"
	"time"
)

func sender1(ch chan string) {
	for {
		ch <- "from 1st"
		time.Sleep(time.Second * 2)
	}
}

func sender2(ch chan string) {
	for {
		ch <- "from 2nd"
		time.Sleep(time.Second * 2)
	}
}

func receiver(c1, c2 chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("from 1st : ", msg1)
		case msg2 := <-c2:
			fmt.Println("from 2nd : ", msg2)
		}
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sender1(c1)
	go sender2(c2)
	go receiver(c1, c2)
	time.Sleep(10 * time.Second)
}
