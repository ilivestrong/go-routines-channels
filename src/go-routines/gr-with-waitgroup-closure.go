package goroutines

import (
	"fmt"
	"sync"
)

func StartWithWGInClosure() {
	fmt.Println("Entered main routine")

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
		// this line fixes the issue, where each invocation of
		// go routine now captures a local i value at the time
		// of the iteration, not the last computed value.
	}
	wg.Wait()
	fmt.Println("Main routine and all go routines finished executing")

}
