package concurrencyconcepts

import (
	"fmt"
	"time"
)

func Channels() {
	c := make(chan string, 50)
	go countChannels("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func countChannels(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
