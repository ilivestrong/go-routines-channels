package channels

func ProduceSum(first, second int, c chan int) {
	c <- first + second // send to channel
	close(c)            // not mandatory, but lets receiver know that there are no more values to be received
}
