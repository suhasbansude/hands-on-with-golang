package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumWorkers = 10

type Work struct {
	x, y, z int
}

func worker(in chan *Work, out chan *Work) {
	for w := range in {
		w.z = w.x * w.y
		Sleep(w.z)
		out <- w
	}
}

func sendLotsOfWork(ch chan *Work) {
	for i := 0; i < 50; i++ {
		num1 := rand.Intn(10)
		num2 := rand.Intn(10)
		work := Work{num1, num2, 0}
		ch <- &work
	}
}

func receiveLotsOfResult(ch chan *Work) {
	for w := range ch {
		fmt.Println(*w)
	}
}

func Sleep(t int) {
	time.Sleep(time.Second * time.Duration(t))
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < NumWorkers; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResult(out)
}

func main() {
	Run()
}
