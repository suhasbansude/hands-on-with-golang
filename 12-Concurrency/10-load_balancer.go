package main

import (
	"math/rand"
	"time"
)

type Request struct {
	fn func() int // The operation to perform
	c  chan int   // The channel to return the result
}

func Sleep(t int) {
	time.Sleep(time.Second * time.Duration(t))
}

func requester(work chan<- Request) {
	c := make(chan int)
	for {
		Sleep(rand.Int63n(nWorker * 2 * Second))
	}
}

func main() {

}
