package main

import (
	"fmt"
	"sync"
)

func printNumbers(ch1, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
	close(ch1)
	close(ch2)
}

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(1)
	go printNumbers(ch1, ch2, &wg)

	go func() {
		for {
			select {
			case num, ok := <-ch1:
				if ok {
					fmt.Println(num, "even")
				} else {
					ch1 = nil
				}
			case num, ok := <-ch2:
				if ok {
					fmt.Println(num, "odd")
				} else {
					ch2 = nil
				}
			}

			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("Done!")
}
