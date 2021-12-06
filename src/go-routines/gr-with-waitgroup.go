package goroutines

import (
	"fmt"
	"sync"
)

func StartWithWG() {
	fmt.Println("Inside main routine")
	var wg sync.WaitGroup
	wg.Add(1) // increments go routine counter
	go func() {
		defer wg.Done() // decrement go routine counter
		fmt.Println("I am a go routine with wait group enabled.")
	}()
	wg.Wait() // waits on main thread for go routine(s) to get executed before terminating the thread
	fmt.Println("Main routine and go routine finished")
}
