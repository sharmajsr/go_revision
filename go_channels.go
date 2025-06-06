package main

import "fmt"
// In this surbhi is first sent to channel and recceived by main, which then prints Hello surbhi. after that shubham is 
// sent from main. If the main exists before goGreet function complete, Hey shubham wont be printed, if not, it will be
func goGreet(c chan string){
	fmt.Printf("Hey %s\n",<-c)
}

func goGreet2(c chan string){
	c<-"surbhi"
	fmt.Println("Msg sent\n")
}

func main(){
	ch := make(chan string)
	go goGreet(ch)
	go goGreet2(ch)
	
	fmt.Printf("Hello %s\n",<-ch)
	ch<-"shubham"

}
