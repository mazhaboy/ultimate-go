package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	wg.Add(2)
	fmt.Println("create go routine")

	// go printPrime("A")
	// go printPrime("B")
	fmt.Println("Waiting to finish")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	wg.Wait()
	fmt.Println("\nTerminating program")

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
