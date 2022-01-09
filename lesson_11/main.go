package main

import (
	"fmt"
	"log"
)

type AppError struct {
	Message string
	Err error
}

func (ae *AppError) Error() string {
	return ae.Message
}

func main(){
	//a := []int{1,2,3}
	//c := a[3]
	//fmt.Println(c)
	c := divide(1, 0)
	fmt.Println("after panic: ", c)
}

func divide(a,b int) int{
	defer func() {
		if err := recover(); err != nil {
			log.Println("panc happened: ", err)
		}
	}()

	return a / b
}

//
//func div(a, b int) int {
//	if b == 0{
//		panic(&AppError{
//			Message: "this my custom err",
//			Err: nil,
//		})
//	}
//	return a / b
//}