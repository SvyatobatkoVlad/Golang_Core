package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main(){
	//writeFile()
	//appendToFile()
	//readFile()

	c := exec.Command("top")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()



}

func appendToFile(){
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil{
		fmt.Println(err)
	}
	defer f.Close()
	if _, err = f.WriteString(" Vlad"); err != nil{
		panic(err)
	}

}

func writeFile(){
	data := []byte(" Vlad")
	err := ioutil.WriteFile("test.txt", data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func readFile(){
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}