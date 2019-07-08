package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fanout2()
}

func fanout2() {
	emps := 20

	ch := make(chan string, emps)

	// 5 go routines run at a time
	const cap = 5
	sem := make(chan bool, cap)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			sem <- true
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee : send signal : ", emp)
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager : recv'd signal : ", emps)
		emps--
	}

	time.Sleep(time.Second)
	fmt.Println("---------------------")
}
