package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fanout()
}

func fanout() {
	emps := 20

	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee : send signal : ", emp)
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
