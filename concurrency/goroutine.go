package concurrency

import "fmt"

func Goroutine() {
	nums := make([]int, 0)
	doneCh := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			nums = append(nums, i)
		}
		fmt.Println("Block 1st goroutine until it finishes the appending task")
		doneCh <- 0
	}()
	fmt.Println("Unblock main goroutine until it receives the signal from 1st goroutine")
	<-doneCh

	go func() {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		fmt.Println("Done the 2nd goroutine to sum the num.")
		doneCh <- sum
	}()
	fmt.Println("Receive the final sum")
	fmt.Println("Sum: ", <-doneCh)
}
