package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(val int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting work %d\n", val)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Ending work %d\n", val)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers completed\n")
	// time.Sleep(time.Second * 1)

}
