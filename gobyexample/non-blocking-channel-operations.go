package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive
	select {
		case msg := <- messages:
			fmt.Println("received message", msg)
		default:
			fmt.Println("no message received")
	}

	// non-blocking send
	msg := "hi"
	select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
	}

	// multiple non-blocking receives. same idea for receives too
	select {
		case msg := <- messages:
			fmt.Println("received message", msg)
		case sig := <- signals:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
	}
}
