package concurrencyconcepts

import (
	"fmt"
	"time"
)

func GoRoutine() {
	go countRoutine("sheep")
	countRoutine("fish")
}

func countRoutine(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
