package channels

func ProduceSum(first, second int, c chan int) {
	defer close(c) // not mandatory, but lets receiver know that there are no more values to be received

	c <- first + second // send to channel
}
