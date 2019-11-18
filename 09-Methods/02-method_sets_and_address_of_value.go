package main

import (
	"fmt"
)

type notifier interface {
	notifyByMail()
	notifyByMobile()
}

type user struct {
	name  string
	email string
}

func (u user) notifyByMail() {
	fmt.Printf("Sending User Email to %s <%s>\n", u.name, u.email)
}

func (u *user) notifyByMobile() {
	fmt.Printf("Sending User Message to %s <%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notifyByMail()
	n.notifyByMobile()
}

func main() {
	/* Case 1 : Compilation Fail
	user{"ABC", "abc@gmail.com"}.notify()
	*/

	u := user{"ABC", "abc@mkcl.org"}
	u.notifyByMail()
	u.notifyByMobile()

	// Case 2 : Compilation Fail
	// sendNotification(u)

	// Case 3 :
	sendNotification(&u)
}
