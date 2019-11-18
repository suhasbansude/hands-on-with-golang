package main

import "fmt"

type Customer struct {
	Name     string
	MobileNo string
	EmailId  string
}

func (c Customer) sendOffer() {
	fmt.Printf("Sending offer to %s<%s>\n", c.Name, c.MobileNo)
}

func (c *Customer) changeMobileNo(mobileNo string) {
	c.MobileNo = mobileNo
}

func (c Customer) printDetails() {
	fmt.Printf("Name : %v , Mobile : %v , Email : %v\n", c.Name, c.MobileNo, c.EmailId)
}

func main() {
	c := Customer{"ABC", "1234567890", "abc@gmail.com"}
	c.sendOffer()
	c.printDetails()
	c.changeMobileNo("6666677777")
	c.printDetails()
}
