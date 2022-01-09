package main

import "fmt"

func main() {
	//arrays()
	slces()
}

func arrays() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Println(a[1])

	numbers := [...]int{1,2,3}
	fmt.Println(numbers)
}

func slces(){
	letters := []string{"a", "b", "c"}
	letters[0] = "1"

	letters = append(letters, "d")
	fmt.Println(letters)
	letters = append(letters, "d", "e", "f")
	fmt.Println(letters)


	createSlice := make([]string, 3)
	createSlice[0] = "0"
	createSlice[1] = "1"
	createSlice[2] = "2"
	createSlice = append(createSlice, "3")
	fmt.Println(createSlice)

	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))

	t := createSlice[:]
	fmt.Println(t)
	tt := createSlice[:2]
	fmt.Println(tt)
	ttt := createSlice[2:]
	fmt.Println(ttt)
}