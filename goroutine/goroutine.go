package main

import (
	"fmt"
	"runtime"
	"time"
)
// THere is a race condition if we use "go run -race goroutine.go" to check.
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched() // Without this line the program would hung forever.
			}
		}(i)

	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)

}
