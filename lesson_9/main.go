package main

import (
	"fmt"
	"time"
)

func main(){
	//ch := make(chan string)
	////gorutines
	//go sayHello(ch)
	//fmt.Println("Hello world", <- ch)

	data := make(chan int)
	exit := make(chan int)

	go func(){
		for i := 0; i < 10; i++{
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)
}

//func say(word string){
//	fmt.Println(word)
//}
//
//func sayHello(data, exit chan int){
//	for i := 0; i <5; i++{
//		time.Sleep(100 * time.Millisecond)
//		say("hello")
//	}
//	// close(exit)
//	exit <- "done"
//}

//======================================================================

func selectOne(data, exit chan int){
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}