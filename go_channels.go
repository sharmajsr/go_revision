package main

import "fmt"

func goGreet(c chan string){
	fmt.Printf("Hey %s\n",<-c)
}

func goGreet2(c chan string){
	c<-"surbhi"
}

func main(){
	ch := make(chan string)
	go goGreet(ch)
	go goGreet2(ch)
	ch<-"shubham"
	fmt.Printf("Hello %s",<-ch)

}
