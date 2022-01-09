package main

import "fmt"

// Default
//type error interface{
//	Error() string
//}

type AppError struct{
	Err error
	Custom string
	Field int
}

func (e *AppError) Error() string{
	return e.Err.Error()
}

func main(){
	err := myFunc()
	if err != nil{
		fmt.Println(err.Error())
	}
}

func myFunc() error {
	return &AppError{
		Err: fmt.Errorf("my error"),
		Custom: "value here",
		Field: 89,
	}
}