package main

import (
	"fmt"

	"github.com/ilivestrong/go-routines-channels/src/channels"
)

func main() {
	/*
		------------- Go Routines Example --------------------------
	*/
	// goroutines.Start()
	// goroutines.StartWithWG()
	// goroutines.StartWithWGInClosure()

	/*
		------------- Go Channels Example --------------------------
	*/

	// ------- 1. Simple
	// ch := make(chan int)
	// first, second := 10, 20
	// go channels.ProduceSum(first, second, ch)
	// fmt.Printf("The sum of %d, %d is: %d\n", first, second, <-ch)

	// -------- 2. With range
	nums := []int{2, 4, 5, 8, 77, 32, 35, 98, 99, 100}
	ch := make(chan int, len(nums))
	channels.ProduceEvenNums(nums, ch)
	for num := range ch {
		fmt.Println(num)
	}

}
