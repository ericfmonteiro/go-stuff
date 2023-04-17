package concurrencyconcepts

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		countWG("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func countWG(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
