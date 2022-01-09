package main

import (
	"fmt"
	"io"
	"strings"
)

func main(){
	r := strings.NewReader("Hello world")

	arr := make([]byte, 8)

	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d err = %v\n", n, err )
		fmt.Printf("arr n bytes: %q", arr[:n])
		if err == io.EOF{
			break
		}
	}
}