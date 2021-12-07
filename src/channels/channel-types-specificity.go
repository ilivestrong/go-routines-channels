package channels

import (
	"fmt"
)

func Ping(message string, ch chan<- string) {
	defer close(ch)
	ch <- message
}

func Pong(ich <-chan string, och chan<- string) {
	defer close(och)
	och <- fmt.Sprintf("You: %s\nResponse: Hey! Thanks for your message!", <-ich)
}
