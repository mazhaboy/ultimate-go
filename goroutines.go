package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      sync.WaitGroup
	counter int
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)

}
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()

	}
}

// func printPrime(prefix string) {
// 	defer wg.Done()

// next:
// 	for outer := 2; outer < 5000; outer++ {
// 		for inner := 2; inner < outer; inner++ {
// 			if outer%inner == 0 {
// 				continue next
// 			}
// 		}
// 		fmt.Printf("%s:%d\n", prefix, outer)
// 	}
// 	fmt.Println("Comleted", prefix)

// }
