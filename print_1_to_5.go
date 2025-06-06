package main

import (
	"fmt"
	"sync"
)

func printNum(c chan int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Printing %d\n",<-c)
}
func main(){

	var wg sync.WaitGroup

	ch := make(chan int)
	n:= 5
	for i:= 0; i < n ;i++{
		wg.Add(1)
		go printNum(ch,&wg)
	}
	for i:= 0; i < n ;i++{
		ch<-i+1
	}
	wg.Wait()
}