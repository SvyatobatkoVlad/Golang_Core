package main

import "fmt"

func main(){
	var a interface{}

	a = "jelly"
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n",a,a)

	a = 42
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n",a,a)

	b := 23
	fmt.Println(b)
	 //var f string

	//f = b.(string)
	//fmt.Println(f)



}