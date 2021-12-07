package channels

func ProduceEvenNums(nums []int, c chan int) {
	for _, num := range nums {
		defer close(c) // tell receiver, we are done sending the values

		if num%2 == 0 {
			c <- num // send values to receiver
		}
	}
}
