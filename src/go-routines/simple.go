package goroutines

import (
	"fmt"
	"time"
)

func Start() {
	fmt.Println("Printing from main thread!")
	go aSimpleFunc()                   // calls a func() as go routine
	time.Sleep(time.Millisecond * 100) // required to block thread to wait for go routine execution
	fmt.Println("function ended")
}

func aSimpleFunc() {
	fmt.Println("I am a simple goroutine")
}
