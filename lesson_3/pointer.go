package main

import "fmt"

	type Ponter struct {
		name string
		age int
	}

	func main(){
		// pointers() //see how work pointer
		structs()
	}

	func structs(){
		p1 := Ponter{
			name: "Vlad",
			age: 22,
		}

		fmt.Println(p1) //  {Vlad 22}
		p1.age = 23
		fmt.Println(p1.age) //  23

		p2 := Ponter{
			name: "Vlad",
		}
		fmt.Println(p2) //  {Vlad 0}
		g := &p1

		fmt.Println(g) //  &{Vlad 23}
		fmt.Println(*g) //  {Vlad 23}

	}

	func pointers() {
		a := "Hello world"
		b := 42
		fmt.Println(a) // Hello world
		fmt.Println(b) // 42

		p := &a		// put link in p
		fmt.Println(p)	// 0xc00009e210  see link
		fmt.Println(*p) // Hello world see p from link
		*p = "oh my"	// change p for link
		fmt.Println(a)  // oh my
	}


