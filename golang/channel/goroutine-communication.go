package main

import "fmt"

// 使用两个goroutine按自然数顺序打印0-9
func main() {
	even := make(chan int)
	odd := make(chan int)
	done := make(chan struct{})

	const x = 10

	// odd goroutine
	go func() {
		for num := range odd {
			fmt.Println("odd goroutine", num)
			even <- num + 1
		}
	}()

	// even goroutine
	go func() {
		for num := range even {
			fmt.Println("-----------------even goroutine", num)
			if num+1 >= x {
				done <- struct{}{}
			} else {
				odd <- num + 1
			}
		}
	}()

	odd <- 0
	<-done
	close(even)
	close(odd)
	close(done)
}
