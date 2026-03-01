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
		// odd channel遍历即是在接收数据，在未发送数据时无缓冲channel是阻塞的
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
