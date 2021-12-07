package channels

func ProduceEvenNums(nums []int, c chan int) {

	for _, num := range nums {
		if num%2 == 0 {
			c <- num
		}
	}
	close(c)
}
